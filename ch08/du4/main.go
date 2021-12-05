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
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// walkDirはdirをルートとするファイルツリーをたどり、
// 見つかったファイルのそれぞれの大きさをfileSizesに送ります
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
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
	select {
	case sema <- struct{}{}: //トークンの取得
	case <-done:
		return nil // キャンセルされた
	}
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

	// 入力を検出するとキャンセルできる
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// ファイルツリーを走査する
	filesizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, filesizes)
	}
	go func() {
		n.Wait()
		close(filesizes)
	}()

	// 定期的に結果を表示する
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// 結果を表示する
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// 既存のゴルーチンが完了できるようにfileSizesを空にする
			for range filesizes {
				// 何もしない
			}
			return
		case size, ok := <-filesizes:
			if !ok {
				break loop // filesizesが閉じられた
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
