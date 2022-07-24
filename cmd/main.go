package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "501 Not Implemented", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
