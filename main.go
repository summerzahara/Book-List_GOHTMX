package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		books := map[string][]Book{
			"Books": {
				{Title: "Pride & Prejudice", Author: "Jane Austen"},
				{Title: "The Trial", Author: "Franz Kafka"},
				{Title: "Kafka by the Shore", Author: "Haruki Marukami"},
			},
		}
		tmpl.Execute(w, books)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) //Simulate loading
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "book-list-element", Book{Title: title, Author: author})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-book/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
