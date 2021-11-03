package main

import (
	"io"
	"os"
)

type WriteCounter struct {
	w io.Writer
	n *int64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	*wc.n += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var n int64
	return &WriteCounter{w, &n}, &n
}

func main() {
	var w io.Writer
	w = os.Stdout
	w, n := CountingWriter(w)
	w.Write([]byte("hello\n"))
	println(*n)
	w.Write([]byte("world\n"))
	println(*n)
}
