package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tmpl = template.Must(template.New("table").Parse(`
<html lang="en">
  <head>
    <meta charset="utf-8">
	<title>Music Table</title>
  </head>
  <body>
    <table>
	<tr style='text-align: left'>
	  <th><a href="?sortBy=Title">Title</a></th>
	  <th><a href="?sortBy=Artist">Artist</a></th>
	  <th><a href="?sortBy=Album">Album</a></th>
	  <th><a href="?sortBy=Year">Year</a></th>
	  <th><a href="?sortBy=Length">Length</a></th>
	</tr>
	{{range .Tracks}}
	<tr>
	  <td>{{.Title}}</td>
	  <td>{{.Artist}}</td>
	  <td>{{.Album}}</td>
	  <td>{{.Year}}</td>
	  <td>{{.Length}}</td>
	</tr>
	{{end}}
	</table>
  </body>
</html>
`))

var defaultTracks = []*Track{
	{"Go", "Delilah", "from the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Soveig", "Smash", 2011, length("4m24s")},
}

type arbitraryKeySort struct {
	t   []*Track
	key string
}

func (x arbitraryKeySort) Len() int      { return len(x.t) }
func (x arbitraryKeySort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x arbitraryKeySort) Less(i, j int) bool {
	switch x.key {
	case "Title":
		return x.t[i].Title < x.t[j].Title
	case "Artist":
		return x.t[i].Artist < x.t[j].Artist
	case "Album":
		return x.t[i].Album < x.t[j].Album
	case "Year":
		return x.t[i].Year < x.t[j].Year
	case "Length":
		return x.t[i].Length < x.t[j].Length
	}
	return false
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		sortBy = ""
	)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "sortBy" {
			sortBy = v[0]
		}
	}

	tracks := append([]*Track{}, defaultTracks...)

	sort.Sort(arbitraryKeySort{tracks, sortBy})
	if err := tmpl.Execute(w, struct {
		Tracks []*Track
	}{tracks}); err != nil {
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
