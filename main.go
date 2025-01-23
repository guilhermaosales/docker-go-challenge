package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Full Cycle Rocks!</h1>")
}

func main() {
	log.Print("Starting server...")
	http.HandleFunc("/", handler)

	port := "8080"
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
