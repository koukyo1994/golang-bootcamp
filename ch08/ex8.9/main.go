package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{}, 20)

type fileInfo struct {
	root   string
	nfiles int64
	nbytes int64
}

// wa;lDirはdirをルートとするファイルツリーをたどり、
// 見つかったファイルのそれぞれの大きさをfileSizesに送ります
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// direntsはディレクトリdirの項目を返します
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	// 最初のディレクトリを決める
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// ファイルツリーを走査する
	fileSizeCounters := make([]chan int64, len(roots))
	waitGroups := make([]*sync.WaitGroup, len(roots))
	for i := range roots {
		fileSizeCounters[i] = make(chan int64)
		waitGroups[i] = &sync.WaitGroup{}
	}

	for i, root := range roots {
		waitGroups[i].Add(1)
		go walkDir(root, waitGroups[i], fileSizeCounters[i])
	}
	for i := range roots {
		go func(i int) {
			waitGroups[i].Wait()
			close(fileSizeCounters[i])
		}(i)
	}

	// 定期的に結果を表示する
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// 結果を表示する
	var fileInfos = make([]fileInfo, len(roots))
	for i, root := range roots {
		fileInfos[i] = fileInfo{root, 0, 0}
	}
loop:
	for {
		select {
		case <-tick:
			printDiskUsage(fileInfos)
		default:
			running := false
			for i := range roots {
				size, ok := <-fileSizeCounters[i]
				running = running || ok
				if ok {
					fileInfos[i].nfiles++
					fileInfos[i].nbytes += size
				}
			}
			if !running {
				printDiskUsage(fileInfos)
				break loop
			}
		}
	}
	printDiskUsage(fileInfos)
}

func printDiskUsage(fileInfos []fileInfo) {
	fmt.Println("========================")
	for _, value := range fileInfos {
		fmt.Printf("%s: %d files  %.1f GB\n", value.root, value.nfiles, float64(value.nbytes)/1e9)
	}
}
