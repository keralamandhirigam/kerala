package routes

import (
	"kerala/witchcraft/database"
	_ "kerala/witchcraft/docs"
	"kerala/witchcraft/handler"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func RouteGroup(app *fiber.App) fiber.Router {
	group := app.Group("/api/v1/witchcraft")
	return group
}

func Init() error {
	app := fiber.New()
	r := RouteGroup(app)
	db, err := database.MongoConnection()
	if err != nil {
		return err
	}
	handler := handler.NewHandler(db)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	r.Post("/", handler.CreateNewWitchcraft)
	r.Put("/:id", handler.UpdateWitchcraft)
	r.Get("/", handler.GetAllWitchcrafts)
	r.Get("/:id", handler.GetWitchcraftById)
	r.Delete("/:id", handler.DeleteWitchcraftById)

	// Remove the app.Listen line
	// err = app.Listen(":8080")
	// if err != nil {
	// 	return err
	// }
	return nil
}
