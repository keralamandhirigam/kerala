package main

import (
	"kerala/witchcraft/routes"
	"log"
)

func main() {
	err := routes.Init()
	if err != nil {
		log.Fatal(err)
	}
}
