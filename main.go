package main

import (
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
