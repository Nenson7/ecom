package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type PageData struct {
	Title string
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		data := PageData{
			Title: "My title",
		}
		tmpl.Execute(w, data)
	})


	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<p>you have clicked</p>"))
    })

	fmt.Println("Server running on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server starting failure", err)
	}
}
