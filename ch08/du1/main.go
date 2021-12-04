package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// wa;lDirはdirをルートとするファイルツリーをたどり、
// 見つかったファイルのそれぞれの大きさをfileSizesに送ります
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// direntsはディレクトリdirの項目を返します
func dirents(dir string) []os.FileInfo {
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
	filesizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, filesizes)
		}
		close(filesizes)
	}()

	// 結果を表示する
	var nfiles, nbytes int64
	for size := range filesizes {
		nfiles++
		nbytes += size
	}
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
