package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	SAVEDIR     = ".xkcd"
	APIEndpoint = "https://xkcd.com"
)

type Comic struct {
	Num        int
	Title      string
	Transcript string
}

func main() {
	savedir := os.Getenv("HOME") + "/" + SAVEDIR

	// 保存用のディレクトリを用意する
	if err := os.Mkdir(savedir, 0755); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}

	f, err := os.Open(strings.Join([]string{savedir, ".index.text"}, "/"))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	} else if err != nil {
		f, err = os.Create(strings.Join([]string{savedir, ".index.text"}, "/"))
		if err != nil {
			panic(err)
		}
	}

	// 既にダウンロード済みのファイルはindexにあるので読み込む
	scanner := bufio.NewScanner(f)
	var indicesSet = make(map[string]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		indicesSet[strings.TrimSpace(line)] = struct{}{}
	}
	f.Close()

	// 入力されたIDを読み込む
	comicID := os.Args[1]
	if _, err := strconv.Atoi(comicID); err != nil {
		log.Fatal("commicID should be a number.")
	}

	// 既にダウンロードされていればそれを読み込む
	if _, ok := indicesSet[comicID]; ok {
		f, err = os.Open(strings.Join([]string{savedir, comicID, "info.0.json"}, "/"))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		comic := Comic{}
		if err := json.NewDecoder(f).Decode(&comic); err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", comic.Transcript)
	} else {
		// 未ダウンロードならばダウンロードする
		fmt.Println("Downloading...")
		resp, err := http.Get(strings.Join([]string{APIEndpoint, comicID, "info.0.json"}, "/"))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		comic := Comic{}
		if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", comic.Transcript)

		// 保存する
		if err := os.Mkdir(strings.Join([]string{savedir, comicID}, "/"), 0755); err != nil {
			if !os.IsExist(err) {
				panic(err)
			}
		}
		f, err = os.Create(strings.Join([]string{savedir, comicID, "info.0.json"}, "/"))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := json.NewEncoder(f).Encode(comic); err != nil {
			panic(err)
		}

		// indexに追加する
		f, err = os.OpenFile(strings.Join([]string{savedir, ".index.text"}, "/"), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := f.WriteString(fmt.Sprintf("%s\n", comicID)); err != nil {
			panic(err)
		}
	}
}
