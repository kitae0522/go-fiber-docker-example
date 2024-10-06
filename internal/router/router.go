package router

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-docker-example/internal/model"
)

func EnrollRouter(app *fiber.App, dbconn *model.PrismaClient) {
	apiRouter := app.Group("/")
	initScheduleRouter(apiRouter, initScheduleDI(dbconn))
	initAuthRouter(apiRouter, initAuthDI(dbconn))

	apiRouter.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "pong",
		})
	})
}
