package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	var elements, attributes string
	attrs := make(map[string]string)

	fmt.Printf("element names: \n")
	fmt.Scanln(&elements)
	names := strings.Split(elements, ",")

	fmt.Printf("attributes: \n")
	fmt.Scanln(&attributes)
	splittedAttributes := strings.Split(attributes, ",")
	for _, attr := range splittedAttributes {
		keyValue := strings.Split(attr, ":")
		if len(keyValue) != 2 {
			fmt.Printf("invalid attribute format: %v\n", attr)
		} else {
			attrs[keyValue[0]] = keyValue[1]
		}
	}

	dec := xml.NewDecoder(strings.NewReader(fetch(url)))
	var stack []string
	var attrsMap map[string]string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			attrsMap = make(map[string]string)    // 新しい開始要素が来たら初期化
			for _, a := range tok.Attr {
				attrsMap[a.Name.Local] = a.Value
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			// fmt.Printf("%v %v\n", stack, attrsMap)
			if containsAll(stack, names) && attributesMatch(attrsMap, attrs) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func attributesMatch(x, y map[string]string) bool {
	for k, v := range y {
		if x[k] != v {
			return false
		}
	}
	return true
}

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return string(b)
}
