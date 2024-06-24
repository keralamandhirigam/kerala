package routes

import (
	_ "kerala/witchcraft/docs"
	"kerala/witchcraft/handler"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func RouteGroup(app *fiber.App, client *mongo.Client) {
	h := handler.NewHandler(client)
	app.Post("/create", h.CreateNewWitchcraft)

}

// func Init() error {
// 	app := fiber.New()
// 	r := RouteGroup(app)
// 	db, err := database.MongoConnection()
// 	if err != nil {
// 		return err
// 	}
// 	handler := handler.NewHandler(db)

// 	app.Get("/swagger/*", fiberSwagger.WrapHandler)
// 	r.Post("/", handler.CreateNewWitchcraft)
// 	r.Put("/:id", handler.UpdateWitchcraft)
// 	r.Get("/", handler.GetAllWitchcrafts)
// 	r.Get("/:id", handler.GetWitchcraftById)
// 	r.Delete("/:id", handler.DeleteWitchcraftById)

// 	err = app.Listen(":8080")
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
