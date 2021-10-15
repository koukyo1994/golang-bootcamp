package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const APIEndpoint = "https://api.github.com/repos"

var (
	repo   = flag.String("repo", "", "Github repo to fetch, format: {owner}/{repo}")
	action = flag.String("action", "list", "Action to take, options: list, get, create, update, close")
	number = flag.Int("number", 0, "Issue number to fetch")
	editor = flag.String("editor", "", "Editor to use for editing")
	title  = flag.String("title", "", "Title of issue")
)

type User struct {
	Login string `json:"login"`
	URL   string `json:"html_url"`
}

type Issue struct {
	ID        int    `json:"id"`
	URL       string `json:"html_url"`
	Number    int    `json:"number"`
	User      *User  `json:"user"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	State     string `json:"state"`
	Assignee  *User  `json:"assignee"`
	Comments  int    `json:"comments"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func listIssues(repo string) {
	url := strings.Join([]string{APIEndpoint, repo, "issues"}, "/")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}
	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	for _, issue := range result {
		log.Printf("#%d %s %s", issue.Number, issue.User.Login, issue.Title)
	}
}

func getIssue(repo string, number int) {
	url := strings.Join([]string{APIEndpoint, repo, "issues", strconv.Itoa(number)}, "/")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	log.Printf("#%d %s %s", result.Number, result.User.Login, result.Title)
	log.Printf("Body: %s", result.Body)
	log.Printf("Created At: %s Updated At: %s", result.CreatedAt, result.UpdatedAt)
}

func createIssue(repo string, title string, body string) {
	url := strings.Join([]string{APIEndpoint, repo, "issues"}, "/")
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Error: %s", resp.Status)
	} else {
		log.Printf("Issue created")
	}
}

func updateIssue(repo string, number int, body string) {
	url := strings.Join([]string{APIEndpoint, repo, "issues", strconv.Itoa(number)}, "/")
	req, err := http.NewRequest("PATCH", url, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	} else {
		log.Printf("Issue updated")
	}
}

func closeIssue(repo string, number int) {
	url := strings.Join([]string{APIEndpoint, repo, "issues", strconv.Itoa(number)}, "/")
	body := `{"state": "closed"}`
	req, err := http.NewRequest("PATCH", url, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	} else {
		log.Printf("Issue closed")
	}
}

func main() {
	flag.Parse()
	if *repo == "" {
		flag.PrintDefaults()
		return
	}
	switch *action {
	case "list":
		listIssues(*repo)
	case "get":
		getIssue(*repo, *number)
	case "create":
		if *title == "" {
			fmt.Printf("Title: ")
			fmt.Scanf("%s", title)
		}

		var body string
		if *editor != "" {
			body = launchEditor(*editor)
			fmt.Println(body)
		} else {
			fmt.Printf("Body: ")
			fmt.Scanf("%s", &body)
		}
		data := strings.Join([]string{"{\"title\":\"", *title, "\",\"body\":\"", body, "\"}"}, "")
		createIssue(*repo, *title, data)
	case "update":
		if *number == 0 {
			log.Fatal("Number is required")
		}
		if *title == "" {
			fmt.Printf("Title: ")
			fmt.Scanf("%s", title)
		}

		var body string
		if *editor != "" {
			body = launchEditor(*editor)
			fmt.Println(body)
		} else {
			fmt.Printf("Body: ")
			fmt.Scanf("%s", &body)
		}
		data := strings.Join([]string{"{\"title\":\"", *title, "\",\"body\":\"", body, "\"}"}, "")
		updateIssue(*repo, *number, data)
	case "close":
		if *number == 0 {
			log.Fatal("Number is required")
		}
		closeIssue(*repo, *number)
	default:
		flag.PrintDefaults()
	}
}

func launchEditor(e string) string {
	filePath := os.Getenv("HOME") + "/.tmp"
	if err := makeFile(filePath); err != nil {
		return ""
	}
	defer deleteFile(filePath)

	if err := openEditor(e, filePath); err != nil {
		return ""
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(b)
}

func makeFile(filePath string) (err error) {
	if !isFileExists(filePath) {
		if err = ioutil.WriteFile(filePath, []byte(""), 0644); err != nil {
			return
		}
	}
	return
}

func isFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || !os.IsNotExist(err)
}

func deleteFile(filePath string) (err error) {
	if isFileExists(filePath) {
		if err = os.Remove(filePath); err != nil {
			return
		}
	}
	return
}

func openEditor(program string, args ...string) error {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
