package server

import (
	"strconv"

	"github.com/co-codin/model"
	"github.com/co-codin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllRedirects(ctx *fiber.Ctx) error {
	golies, err := model.GetAllGolies()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all golies",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(golies)
}


func getGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not parse id " + err.Error(),
		})
	}

	goly, err := model.GetGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not retreive goly from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

func createGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly

	err := c.BodyParser(&goly)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = utils.RandomUrl(8)
	}

	err = model.CreateGoly(goly)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not create goly in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(c *fiber.Ctx) error {

}

func deleteGoly(c *fiber.Ctx) error {
	
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type",
	}))

	router.Get("/goly", getAllRedirects)
	router.Get("/goly/:id", getGoly))
	router.Post("/goly", createGoly)
	router.Patch("/goly", updateGoly)
	router.Delete("/goly/:id", deleteGoly)

	router.Listen(":3000")
}
