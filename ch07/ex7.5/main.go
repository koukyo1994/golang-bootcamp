package main

import (
	"bytes"
	"fmt"
	"io"
)

type Reader struct {
	r      io.Reader
	readed int64
	n      int64
}

func (r *Reader) Read(p []byte) (int, error) {
	if r.readed >= r.n {
		return 0, io.EOF
	}
	b, err := r.r.Read(p)
	if err != nil {
		return b, err
	}

	b = min(b, int(r.n-r.readed))
	r.readed += int64(b)
	return b, err
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &Reader{r, 0, n}
}

func main() {
	sentence := "This is a sample sentence."
	reader := bytes.NewReader([]byte(sentence))
	limitreader := LimitReader(reader, 10)
	data, err := io.ReadAll(limitreader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", len(data))
	fmt.Printf("%s\n", data)
}
