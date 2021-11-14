package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
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

var tracks = []*Track{
	{"Go", "Delilah", "from the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Soveig", "Smash", 2011, length("4m24s")},
}

type multiKeysSort struct {
	t    []*Track
	keys []string
}

func (x multiKeysSort) Len() int      { return len(x.t) }
func (x multiKeysSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x multiKeysSort) Less(i, j int) bool {
	var equal = func(i, j int, key string) bool {
		switch key {
		case "Title":
			return x.t[i].Title == x.t[j].Title
		case "Artist":
			return x.t[i].Artist == x.t[j].Artist
		case "Album":
			return x.t[i].Album == x.t[j].Album
		case "Year":
			return x.t[i].Year == x.t[j].Year
		case "Length":
			return x.t[i].Length == x.t[j].Length
		}
		return false
	}

	var less = func(i, j int, key string) bool {
		switch key {
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
	for _, key := range x.keys {
		if equal(i, j, key) {
			continue
		}
		return less(i, j, key)
	}
	return false
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // 列幅を計算して表を印字する
}

func main() {
	printTracks(tracks)
	sort.Sort(multiKeysSort{tracks, []string{"Title", "Year"}})
	printTracks(tracks)
}
