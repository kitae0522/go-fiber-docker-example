package routers

import (
	"github.com/gofiber/fiber/v2"
)

func EnrollRouter(app *fiber.App) {
	apiRouter := app.Group("/")
	initCalenderRouter(apiRouter)

	apiRouter.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})
}
