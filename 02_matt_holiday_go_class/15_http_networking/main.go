package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

// The go standard libr has many packages for making web servers:
// 	client & server sockets
// 	route multiplexing
// 	HTTP and HTML, including HTML templates
// 	JSON and other data formats
// 	cryptographic security
// 	SQL database access
// 	compression utilities
// 	image generation
// There are also lots of 3rd-party packages with improvements

// An HTTP handler func is an instance of an interface
// 	type Handler interface { ServerHTTP(ResponseWriter, *Request ) }
//  type HandlerFunc func(ResponseWriter, *Request)
// 	func (f HandlerFunc) ServerHTTP(w ResponseWriter, r *Request) {
// 	f(w, r)}
//  The HTTP framework can call a method on a function
// 	func handler(w http.ResponseWriter, r *http.Request) {
//  fmt.Fprintf(w, "Hello, world! from %s\n", r.URL.Path[1:])}

func handler(w http.ResponseWriter, r *http.Request) {

	const base = "https://jsonplaceholder.typicode.com"

	resp, err := http.Get(base + r.URL.Path[:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var item todo

	err = json.Unmarshal(body, &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("mine")
	tmpl.Parse(form)
	tmpl.Execute(w, item)

}

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>
`

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// https://jsonplaceholder.typicode.com/todos/2
	// localhost:8080/todos/2

}
