package main

import (
	"fmt"
	// "kerala/witchcraft/routes"
	"log"
	"net/http"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	err := routes.Init()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello World")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
