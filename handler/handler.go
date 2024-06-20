package handler

import (
	"context"
	"encoding/json"
	"kerala/witchcraft/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	Db *mongo.Client
}

func NewHandler(db *mongo.Client) *Handler {
	return &Handler{
		Db: db,
	}
}

// @Summary Create New Requirment
// @Description Create New Requirment
// @Tags Witchcraft
// @Accept json
// @Produce json
// @Param user body model.CreateWitchcraft true "Create Request"
// @Success 200 {string} string  "Success"
// @Failure 400 {string} string "Bad Request"
// @Router /api/v1/witchcraft [post]
func (h *Handler) CreateNewWitchcraft(c *fiber.Ctx) error {
	coll := h.Db.Database("keralaWitchcraft").Collection("keralaWitchcraft")
	var req model.CreateWitchcraft

	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	dto := model.WitchcraftDAO{
		WitchcraftId:     req.WhichcraftId,
		Name:             req.Name,
		Address:          req.Address,
		Place:            req.Place,
		Remarks:          req.Remarks,
		Payment:          req.Payment,
		BalanceAmount:    req.BalanceAmount,
		VisitorsDuration: req.VisitorsDuration,
	}

	err = CreateOrUpdate(c.Context(), dto, coll)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

// @Summary Update Requirment
// @Description Update Requirment
// @Tags Witchcraft
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param user body model.UpdateWitchcraft true "Create Request"
// @Success 200 {string} string  "Success"
// @Failure 400 {string} string "Bad Request"
// @Router /api/v1/witchcraft/{id} [put]
func (h *Handler) UpdateWitchcraft(c *fiber.Ctx) error {
	coll := h.Db.Database("keralaWitchcraft").Collection("keralaWitchcraft")
	var req model.UpdateWitchcraft

	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id := c.Params("id")
	dto := model.WitchcraftDAO{
		WitchcraftId:     id,
		Name:             req.Name,
		Address:          req.Address,
		Place:            req.Place,
		Remarks:          req.Remarks,
		Payment:          req.Payment,
		BalanceAmount:    req.BalanceAmount,
		VisitorsDuration: req.VisitorsDuration,
	}

	err = CreateOrUpdate(c.Context(), dto, coll)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func CreateOrUpdate(ctx context.Context, req model.WitchcraftDAO, coll *mongo.Collection) error {
	filter := bson.M{
		"witchcraftId": req.WitchcraftId,
	}
	opts := options.Update().SetUpsert(true)
	_, err := coll.UpdateOne(ctx, filter, bson.M{"$set": req}, opts)
	if err != nil {
		return err
	}
	return nil
}

// @Summary Get All Requirments
// @Description Get All Requirments
// @Tags Witchcraft
// @Accept json
// @Produce json
// @Success 200 {object} []model.Witchcraft
// @Failure 400 {string} string "Bad Request"
// @Router /api/v1/witchcraft  [get]
func (h *Handler) GetAllWitchcrafts(c *fiber.Ctx) error {
	coll := h.Db.Database("keralaWitchcraft").Collection("keralaWitchcraft")

	var requirments []model.Witchcraft

	curs, err := coll.Find(c.Context(), bson.M{}, nil)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = curs.All(c.Context(), &requirments)
	if err != nil {
		return err
	}

	if len(requirments) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": []interface{}{}})

	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": requirments})

}

// @Summary Get  Requirment
// @Description Get Requirment
// @Tags Witchcraft
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} []model.Witchcraft
// @Failure 400 {string} string "Bad Request"
// @Router /api/v1/witchcraft/{id}  [get]
func (h *Handler) GetWitchcraftById(c *fiber.Ctx) error {
	coll := h.Db.Database("keralaWitchcraft").Collection("keralaWitchcraft")

	var requirment model.Witchcraft

	err := coll.FindOne(c.Context(), bson.M{"witchcraftId": c.Params("id")}, nil).Decode(&requirment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": requirment})

}

// @Summary Delete Record By ID
// @Description Delete Record By ID
// @Tags Witchcraft
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {string} string  "Success"
// @Failure 400 {string} string "Bad Request"
// @Router /api/v1/witchcraft/{id}  [delete]
func (h *Handler) DeleteWitchcraftById(c *fiber.Ctx) error {
	coll := h.Db.Database("keralaWitchcraft").Collection("keralaWitchcraft")

	_, err := coll.DeleteOne(c.Context(), bson.M{"witchcraftId": c.Params("id")}, nil)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "success"})

}
