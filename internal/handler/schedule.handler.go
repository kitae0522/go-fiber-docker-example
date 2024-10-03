package handler

import (
	"log"
	// "reflect"

	"github.com/gofiber/fiber/v2"

	"go-fiber-docker-example/internal/dto"
    "go-fiber-docker-example/internal/service"
	"go-fiber-docker-example/pkg/utils"
)

type ScheduleHandler struct {
	scheduleService *service.ScheduleService
}

func NewScheduleHandler(service *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{scheduleService: service}
}

func (h *ScheduleHandler) CreateSchedule(ctx *fiber.Ctx) error {
	newSchedule := new(dto.NewScheduleItemReq)
	if err := utils.Bind(ctx, newSchedule); err != nil {
		statusCode := fiber.StatusBadRequest
		errMessage := "❌ 스케쥴 생성 실패. Body Binding 과정에서 문제 발생"
		log.Printf("%v: %v", errMessage, err)
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError: true,
			StatusCode: statusCode,
			Error: err,
		})
	}

	err := h.scheduleService.CreateSchedule(newSchedule)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		errMessage := "❌ 스케쥴 생성 실패. Repository에서 문제 발생"
		log.Printf("%v: %v", errMessage, err)
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError: true,
			StatusCode: statusCode,
			Error: err,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.NewScheduleItemRes{
		IsError: false,
		Message: "✅ 스케쥴 생성 완료",
	})
}

func (h *ScheduleHandler) GetScheduleByID(ctx *fiber.Ctx) error {
	scheduleID := dto.GetScheduleItemReq{ID: ctx.Params("id")}
	
	findSchedule, err := h.scheduleService.GetScheduleByID(&scheduleID)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		errMessage := "❌ 스케쥴 조회 실패. Repository에서 문제 발생"
		log.Printf("%v: %v", errMessage, err)
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError: true,
			StatusCode: statusCode,
			Error: err,
		})
	}

	// TODO: 아이템 없을 시, 없다고 리턴

	return ctx.Status(fiber.StatusCreated).JSON(dto.GetScheduleItemRes{
		IsError: false,
		Message: "✅ 스케쥴 조회 완료",
		Schedule: findSchedule,
	})
}