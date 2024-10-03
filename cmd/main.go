package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"go-fiber-docker-example/internal/handler"
	"go-fiber-docker-example/internal/service"
	"go-fiber-docker-example/internal/repository"
	"go-fiber-docker-example/internal/model"
)

const port = ":8080"

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))
	app.Use(logger.New())

	dbconn := model.NewClient()
	if err := dbconn.Prisma.Connect(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer func() {
		if err := dbconn.Prisma.Disconnect(); err != nil {
			// 빠른 실패를 위한 panic은 좋다!
			panic(err)
		}
	}()

	repository := repository.NewScheduleRepository(dbconn)
	service := service.NewScheduleService(repository)
	handler := handler.NewScheduleHandler(service)

	app.Post("/schedule", handler.CreateSchedule)
	app.Get("/schedule/:id", handler.GetScheduleByID)
	
	// routers.EnrollRouter(app)
	log.Fatal(app.Listen(port))
}
