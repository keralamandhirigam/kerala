package handler

import (
	"kerala/witchcraft/routes"
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// Create a global Fiber app instance
var app *fiber.App

func init() {
	// Initialize the Fiber app
	app = fiber.New()

	// Initialize routes
	routes.RouteGroup(app)
}

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Use the Fiber adaptor to handle the request
	adaptor.FiberApp(app)(w, r)
}

func main() {
	// For local testing
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
