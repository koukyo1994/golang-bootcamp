package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"  // GIFデコーダを登録する
	"image/jpeg" // JPEGデコーダを登録する
	"image/png"  // PNGデコーダを登録する
	"io"
	"os"
)

var imgFormat = flag.String("format", "jpeg", "specify image format to convert to.")

func toImgFormat(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	if format == "jpeg" || format == "jpg" {
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	} else if format == "png" {
		return png.Encode(out, img)
	} else if format == "gif" {
		return gif.Encode(out, img, &gif.Options{})
	} else {
		return fmt.Errorf("format %s not supported", format)
	}
}

func main() {
	flag.Parse()
	if err := toImgFormat(os.Stdin, os.Stdout, *imgFormat); err != nil {
		fmt.Fprintf(os.Stderr, "convert format: %v\n", err)
		os.Exit(1)
	}
}
