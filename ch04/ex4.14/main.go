package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"
)

const APIEndpoint = "https://api.github.com"

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // マークダウン(Markdown)形式
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type MileStone struct {
	Title       string
	URL         string
	State       string
	Creator     *User
	Description string
}

func fetchIssues(owner, repo string) (*IssuesSearchResult, error) {
	result := IssuesSearchResult{}
	issueURL := APIEndpoint + "/search/issues"

	q := url.QueryEscape("repo:" + owner + "/" + repo + " bug")
	resp, err := http.Get(issueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func fetchMilestone(owner, repo string) ([]*MileStone, error) {
	var result []*MileStone
	url := APIEndpoint + "/repos/" + owner + "/" + repo + "/milestones"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

var tmpl = template.Must(template.New("bugrepots").Parse(`
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Bug and Milestone Reports</title>
  </head>
  <body>
    <h1>{{.Issues.TotalCount}} bugs found</h1>
    <table>
      <tr style='text-align: left'>
	    <th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	  </tr>
	  {{range .Issues.Items}}
	  <tr>
	    <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	  </tr>
	  {{end}}
	</table>
	<h1>milestones</h1>
	<table>
	  <tr style='text-align: left'>
		<th>Title</th>
		<th>State</th>
		<th>Creator</th>
		<th>Description</th>
	  </tr>
	  {{range .Milestone}}
	  <tr>
		<td><a href='{{.URL}}'>{{.Title}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.Creator.HTMLURL}}'>{{.Creator.Login}}</a></td>
		<td>{{.Description}}</td>
	  </tr>
	  {{end}}
	</table>
  </body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		owner = "golang"
		repo  = "go"
	)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		switch k {
		case "owner":
			owner = v[0]
		case "repo":
			repo = v[0]
		}
	}

	issues, err := fetchIssues(owner, repo)
	if err != nil {
		log.Print(err)
	}

	milestone, err := fetchMilestone(owner, repo)
	if err != nil {
		log.Print(err)
	}

	if err := tmpl.Execute(w, struct {
		Issues    *IssuesSearchResult
		Milestone []*MileStone
	}{issues, milestone}); err != nil {
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
