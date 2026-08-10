package routes

import "github.com/gofiber/fiber/v3"

func CreateSetsRoutes(app *fiber.App) {
	sets := app.Group("/api/sets")

	sets.Get("/:id", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello, world"})
	})
}
