package router

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-docker-example/internal/handler"
	"go-fiber-docker-example/internal/model"
	"go-fiber-docker-example/internal/repository"
	"go-fiber-docker-example/internal/service"
)

func initUserDI(dbconn *model.PrismaClient) *handler.UserHandler {
	repository := repository.NewUserRepository(dbconn)
	service := service.NewUserService(repository)
	handler := handler.NewUserHandler(service)
	return handler
}

func initUserRouter(router fiber.Router, handler *handler.UserHandler) {
	userRouter := router.Group("/user")
	userRouter.Get("/:id", handler.GetProfileByTag)
	userRouter.Patch("/:id", handler.UpdateProfile)
	userRouter.Delete("/:id", handler.DeleteProfile)
}
