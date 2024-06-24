// package main

// import (
// 	"kerala/witchcraft/routes"
// 	"log"
// )

// func main() {
// 	err := routes.Init()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
//}

package handler

import (
	"context"
	"kerala/witchcraft/routes"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app *fiber.App
var client *mongo.Client

func init() {
	// Initialize MongoDB client
	var err error
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Fiber app
	app = fiber.New()

	// Initialize routes
	routes.RouteGroup(app, client)
}

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(app)(w, r)
}

func main() {
	// For local testing
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
