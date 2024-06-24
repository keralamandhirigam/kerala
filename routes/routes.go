package routes

import (
	"kerala/witchcraft/database"
	_ "kerala/witchcraft/docs"
	"kerala/witchcraft/handler"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func RouteGroup(app *fiber.App) {
	group := app.Group("/api/v1/witchcraft")

	db, err := database.MongoConnection()
	if err != nil {
		panic(err) // For simplicity, panic on error
	}
	h := handler.NewHandler(db)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	group.Post("/", h.CreateNewWitchcraft)
	group.Put("/:id", h.UpdateWitchcraft)
	group.Get("/", h.GetAllWitchcrafts)
	group.Get("/:id", h.GetWitchcraftById)
	group.Delete("/:id", h.DeleteWitchcraftById)
}
