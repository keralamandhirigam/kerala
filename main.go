package main

import (
	"fmt"
	"kerala/witchcraft/routes"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := routes.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello World")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
