package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

var tmpl = template.Must(template.New("table").Parse(`
<html lang="en">
  <head>
    <meta charset="utf-8">
	<title>Item Table</title>
  </head>
  <body>
    <table>
	  <tr style='text-align: left'>
	    <th>Item</th>
		<th>Price</th>
	  </tr>
    {{range .DB}}
	<tr>
	  <td>{{.Item}}</td>
	  <td>{{.Price}}</td>
	</tr>
	{{end}}
	</table>
  </body>
</html>
`))

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	type DBItem struct {
		Item  string
		Price dollars
	}

	var Items []DBItem
	for item, price := range db {
		Items = append(Items, DBItem{item, price})
	}
	if err := tmpl.Execute(w, struct {
		DB []DBItem
	}{Items}); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "something went wrong\n")
		return
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	http.ListenAndServe("localhost:8000", mux)
}
