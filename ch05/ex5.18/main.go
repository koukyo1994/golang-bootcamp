package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	defer func() {
		if closeErr := f.Close(); err != nil {
			err = closeErr
		}
	}()
	return local, n, err
}

func main() {
	url := os.Args[1]
	filename, _, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v", err)
		os.Exit(1)
	}
	fmt.Println(filename)
}
