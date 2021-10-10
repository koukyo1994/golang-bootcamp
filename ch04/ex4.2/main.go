package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var (
	algorithm = flag.String("algorithm", "sha256", "algorithm to use")
)

func main() {
	flag.Parse()
	var input string
	fmt.Scan(&input)
	switch *algorithm {
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
	case "sha1":
		fmt.Printf("%x\n", sha1.Sum([]byte(input)))
	case "md5":
		fmt.Printf("%x\n", md5.Sum([]byte(input)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
	}
}
