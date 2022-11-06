package server

import (
	"github.com/co-codin/model"
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

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type",
	}))

	router.Get("/goly", getAllRedirects)

	router.Listen(":3000")
}
