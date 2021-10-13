package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	SAVEDIR     = ".poster"
	APIEndpoint = "https://omdbapi.com"
)

var (
	APIKey = flag.String("apikey", "", "API key for OMDB")
	Title  = flag.String("title", "", "Title of the movie")
)

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
}

func searchMovie(title string) (Movie, error) {
	var movie Movie
	url := APIEndpoint + "/?apikey=" + *APIKey + "&t=" + title
	resp, err := http.Get(url)
	if err != nil {
		return movie, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return movie, err
	}
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func retrieveImage(url string, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return err
	}
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	savedir := os.Getenv("HOME") + "/" + SAVEDIR

	// 保存用のディレクトリを用意する
	if err := os.Mkdir(savedir, 0755); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}

	// コマンドライン引数を解析する
	flag.Parse()

	// APIキーが指定されていない場合はエラーを出力する
	if *APIKey == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// タイトルが指定されていない場合はエラーを出力する
	if *Title == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := os.Open(strings.Join([]string{savedir, ".index.txt"}, "/"))
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(strings.Join([]string{savedir, ".index.txt"}, "/"))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
	defer f.Close()

	// 既に一度検索したことがあるタイトルはindexに書き込んであるので読み込む
	// true: 検索したことがあり、omdbに存在した false: 検索したがomdbに存在しなかった
	index := make(map[string]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, ":")
		if splitted[1] == "true" {
			index[splitted[0]] = true
		} else {
			index[splitted[0]] = false
		}
	}

	if value, ok := index[*Title]; !ok {
		// 検索したことがないタイトルなので検索する
		movie, err := searchMovie(*Title)
		if err != nil {
			panic(err)
		}
		f, err = os.OpenFile(strings.Join([]string{savedir, ".index.txt"}, "/"), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if movie.Poster == "" {
			fmt.Println("Specified title does not exist on OMDB.")
			if _, err := f.WriteString(strings.Join([]string{*Title, ":false\n"}, ":")); err != nil {
				panic(err)
			}
			os.Exit(0)
		} else if movie.Poster == "N/A" {
			fmt.Println("Specified title does not exist on OMDB.")
			if _, err := f.WriteString(strings.Join([]string{*Title, ":false\n"}, ":")); err != nil {
				panic(err)
			}
			os.Exit(0)
		}

		url := movie.Poster
		fileName := strings.Join([]string{savedir, movie.Title + ".jpg"}, "/")

		if err := retrieveImage(url, fileName); err != nil {
			if _, err_ := f.WriteString(fmt.Sprintf("%s:%s\n", *Title, "false")); err_ != nil {
				panic(err_)
			}
			panic(err)
		} else {
			if _, err_ := f.WriteString(fmt.Sprintf("%s:%s\n", *Title, "true")); err_ != nil {
				panic(err_)
			}
			fmt.Printf("Download succeded! File path: %s\n", fileName)
		}
	} else if value {
		fmt.Printf("Already donwloaded. File path: %s\n", strings.Join([]string{savedir, *Title + ".jpg"}, "/"))
	} else {
		fmt.Println("Specified title does not exist on OMDB.")
	}
}
