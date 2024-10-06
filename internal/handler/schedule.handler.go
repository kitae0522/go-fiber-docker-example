package handler

import (
	"log"

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
		return utils.CreateErrorRes(ctx, statusCode, errMessage, err)
	}

	err := h.scheduleService.CreateSchedule(newSchedule)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		errMessage := "❌ 스케쥴 생성 실패. Repository에서 문제 발생"
		log.Printf("%v: %v", errMessage, err)
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError:    true,
			StatusCode: statusCode,
			Message:    errMessage,
			Error:      err,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.NewScheduleItemRes{
		IsError:    false,
		StatusCode: fiber.StatusCreated,
		Message:    "✅ 스케쥴 생성 완료",
	})
}

func (h *ScheduleHandler) ListSchedule(ctx *fiber.Ctx) error {
	scheduleList, err := h.scheduleService.ListSchedule()
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		errMessage := "❌ 스케쥴 조회 실패. Repository에서 문제 발생"
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError:    true,
			StatusCode: statusCode,
			Message:    errMessage,
			Error:      err,
		})
	}

	if len(scheduleList) == 0 {
		statusCode := fiber.StatusNotFound
		errMessage := "❌ 스케쥴 조회 실패. 스케쥴이 없습니다."
		return ctx.Status(statusCode).JSON(dto.ListScheduleItemRes{
			IsError:       false,
			StatusCode:    statusCode,
			Message:       errMessage,
			ScheduleCount: 0,
			ScheduleList:  nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ListScheduleItemRes{
		IsError:       false,
		StatusCode:    fiber.StatusOK,
		Message:       "✅ 스케쥴 조회 완료",
		ScheduleCount: len(scheduleList),
		ScheduleList:  scheduleList,
	})
}

func (h *ScheduleHandler) GetScheduleByID(ctx *fiber.Ctx) error {
	scheduleID := ctx.Params("id")
	targetSchedule := new(dto.GetScheduleItemReq)

	if len(scheduleID) == 0 {
		statusCode := fiber.StatusBadRequest
		errMessage := "❌ 스케쥴 조회 실패. Binding 과정에서 문제 발생"
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError:    true,
			StatusCode: statusCode,
			Message:    errMessage,
		})
	}
	targetSchedule.ID = scheduleID

	findSchedule, err := h.scheduleService.GetScheduleByID(targetSchedule)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		errMessage := "❌ 스케쥴 조회 실패. Repository에서 문제 발생"
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError:    true,
			StatusCode: statusCode,
			Message:    errMessage,
			Error:      err,
		})
	}

	if findSchedule == nil {
		statusCode := fiber.StatusNotFound
		errMessage := "❌ 스케쥴 조회 실패. 해당하는 스케쥴이 없습니다."
		return ctx.Status(statusCode).JSON(dto.GetScheduleItemRes{
			IsError:    false,
			StatusCode: statusCode,
			Message:    errMessage,
			Schedule:   nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.GetScheduleItemRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 스케쥴 조회 완료",
		Schedule:   findSchedule,
	})
}

func (h *ScheduleHandler) UpdateSchedule(ctx *fiber.Ctx) error {
	scheduleID := ctx.Params("id")
	targetSchedule := new(dto.UpdateScheduleItemReq)

	if len(scheduleID) == 0 {
		statusCode := fiber.StatusBadRequest
		errMessage := "❌ 스케쥴 조회 실패. Binding 과정에서 문제 발생"
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError:    true,
			StatusCode: statusCode,
			Message:    errMessage,
		})
	}
	targetSchedule.ID = scheduleID

	if err := utils.Bind(ctx, targetSchedule); err != nil {
		statusCode := fiber.StatusBadRequest
		errMessage := "❌ 스케쥴 생성 실패. Body Binding 과정에서 문제 발생"
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError:    true,
			StatusCode: statusCode,
			Message:    errMessage,
			Error:      err,
		})
	}

	updateSchedule, err := h.scheduleService.UpdateSchedule(targetSchedule)

	if updateSchedule == nil || err != nil {
		statusCode := fiber.StatusNotFound
		errMessage := "❌ 스케쥴 수정 실패. 해당하는 스케쥴이 없습니다."
		return ctx.Status(statusCode).JSON(dto.GetScheduleItemRes{
			IsError:    false,
			StatusCode: statusCode,
			Message:    errMessage,
			Schedule:   nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.GetScheduleItemRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 스케쥴 수정 완료",
		Schedule:   updateSchedule,
	})
}

func (h *ScheduleHandler) DeleteSchedule(ctx *fiber.Ctx) error {
	scheduleID := ctx.Params("id")
	targetSchedule := new(dto.DeleteScheduleItemReq)

	if len(scheduleID) == 0 {
		statusCode := fiber.StatusBadRequest
		errMessage := "❌ 스케쥴 조회 실패. Binding 과정에서 문제 발생"
		return ctx.Status(statusCode).JSON(utils.ErrorRes{
			IsError:    true,
			StatusCode: statusCode,
			Message:    errMessage,
		})
	}
	targetSchedule.ID = scheduleID

	success, err := h.scheduleService.DeleteSchedule(targetSchedule)

	if !success || err != nil {
		statusCode := fiber.StatusNotFound
		errMessage := "❌ 스케쥴 삭제 실패. 해당하는 스케쥴이 없습니다."
		return ctx.Status(statusCode).JSON(dto.DeleteScheduleItemRes{
			IsError:    false,
			StatusCode: statusCode,
			Message:    errMessage,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.DeleteScheduleItemRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 스케쥴 삭제 완료",
	})
}
