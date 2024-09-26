package calenderCtrl

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-docker-example/dto/calender"
	"go-fiber-docker-example/utils"
)

func DeleteCalender(ctx *fiber.Ctx) error {
	deleteItem := new(dto.DeleteCalenderItemReq)
	if err := utils.Bind(ctx, deleteItem); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorRes{
			IsError:    true,
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.DeleteCalenderItemRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		ID:         deleteItem.ID,
	})
}
