package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/html")

}

func main() {
	http.HandleFunc("/", handleRoot)

	fmt.Println("Server running on localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server starting failure", err)
	}
}
