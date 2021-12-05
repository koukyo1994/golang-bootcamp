package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var done = make(chan struct{})

func fetch(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("fetch: getting %s: %s", url, resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("fetch: reading %s: %v", url, err)
	}
	fmt.Printf("%s", b)
	close(done)
	return nil
}

func main() {
	for _, url := range os.Args[1:] {
		go fetch(url)
	}
	<-done
}
