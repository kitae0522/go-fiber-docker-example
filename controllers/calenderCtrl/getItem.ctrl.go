package calenderCtrl

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-docker-example/dto/calender"
	"go-fiber-docker-example/utils"
)

func GetCalender(ctx *fiber.Ctx) error {
	calenderItem := new(dto.GetCalenderItemReq)
	if err := utils.Bind(ctx, calenderItem); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorRes{
			IsError:    true,
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
		})
	}
	return ctx.Status(200).JSON(dto.GetCalenderItemRes{
		IsError: false,
		ID:      ctx.Params("id"),
	})
}
