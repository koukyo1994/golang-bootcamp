package main

import (
	"archive/tar"
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

var filename = flag.String("file", "", "archive file to open")
var saveDir = flag.String("save_dir", "./", "directory to save the extracted files")

func main() {
	flag.Parse()
	if strings.HasSuffix(*filename, ".zip") {
		if err := extractZip(*filename, *saveDir); err != nil {
			fmt.Fprintf(os.Stderr, "cannot open zip file: %v\n", err)
			os.Exit(1)
		}
	} else if strings.HasSuffix(*filename, ".tar") {
		if err := extractTar(*filename, *saveDir); err != nil {
			fmt.Fprintf(os.Stderr, "cannot open tar file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Format of the file `%s` is not supported\n", *filename)
		os.Exit(1)
	}
}

func extractTar(filename string, saveDir string) error {
	file, _ := os.Open(filename)
	defer file.Close()
	reader := tar.NewReader(file)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if strings.HasSuffix(header.Name, "/") {
			fmt.Fprintf(os.Stderr, "Creating directory: %s\n", header.Name)
			dirPath := path.Join(saveDir, header.Name)
			os.MkdirAll(dirPath, 0777)
		} else {
			fmt.Fprintf(os.Stderr, "Extracting file: %s\n", header.Name)
			path := path.Join(saveDir, header.Name)
			writeTo, _ := os.Create(path)
			defer writeTo.Close()
			for {
				_, err := io.Copy(writeTo, reader)
				if err != nil {
					return err
				} else {
					break
				}
			}
		}
	}
	return nil
}

func extractZip(filename string, saveDir string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if strings.HasSuffix(file.Name, "/") {
			fmt.Fprintf(os.Stderr, "Creating directory: %s\n", file.Name)
			dirPath := path.Join(saveDir, file.Name)
			os.MkdirAll(dirPath, 0777)
		} else {
			fmt.Fprintln(os.Stderr, "Extracting file:", file.Name)
			readFrom, err := file.Open()
			if err != nil {
				return err
			}
			defer readFrom.Close()
			path := path.Join(saveDir, file.Name)
			writeTo, _ := os.Create(path)
			defer writeTo.Close()
			for {
				_, err := io.Copy(writeTo, readFrom)
				if err != nil {
					return err
				} else {
					break
				}
			}
		}
	}
	return nil
}
