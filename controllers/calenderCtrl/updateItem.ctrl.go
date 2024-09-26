package calenderCtrl

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-docker-example/dto/calender"
	"go-fiber-docker-example/utils"
)

func UpdateCalender(ctx *fiber.Ctx) error {
	updateItem := new(dto.UpdateCalenderItemReq)
	if err := utils.Bind(ctx, updateItem); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorRes{
			IsError:    true,
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.UpdateCalenderItemRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		ID:         updateItem.ID,
	})
}
