package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting daemon")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	port := "8080"

	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
