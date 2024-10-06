package router

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-docker-example/internal/handler"
	"go-fiber-docker-example/internal/model"
	"go-fiber-docker-example/internal/repository"
	"go-fiber-docker-example/internal/service"
)

func initAuthDI(dbconn *model.PrismaClient) *handler.AuthHandler {
	repository := repository.NewAuthRepository(dbconn)
	service := service.NewAuthService(repository)
	handler := handler.NewAuthHandler(service)
	return handler
}

func initAuthRouter(router fiber.Router, handler *handler.AuthHandler) {
	authRouter := router.Group("/auth")
	authRouter.Post("/register", handler.Register)
	authRouter.Post("/login", handler.Login)
}
