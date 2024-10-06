package router

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-docker-example/internal/handler"
	"go-fiber-docker-example/internal/model"
	"go-fiber-docker-example/internal/repository"
	"go-fiber-docker-example/internal/service"
)

func initScheduleDI(dbconn *model.PrismaClient) *handler.ScheduleHandler {
	repository := repository.NewScheduleRepository(dbconn)
	service := service.NewScheduleService(repository)
	handler := handler.NewScheduleHandler(service)
	return handler
}

func initScheduleRouter(router fiber.Router, handler *handler.ScheduleHandler) {
	scheduleRouter := router.Group("/schedule")
	scheduleRouter.Post("/", handler.CreateSchedule)
	scheduleRouter.Get("/", handler.ListSchedule)
	scheduleRouter.Get("/:id", handler.GetScheduleByID)
	scheduleRouter.Patch("/:id", handler.UpdateSchedule)
	scheduleRouter.Delete("/:id", handler.DeleteSchedule)
}
