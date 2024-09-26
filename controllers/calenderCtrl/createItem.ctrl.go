package calenderCtrl

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-docker-example/dto/calender"
	"go-fiber-docker-example/utils"
)

func CreateCalender(ctx *fiber.Ctx) error {
	newItem := new(dto.NewCalenderItemReq)
	if err := utils.Bind(ctx, newItem); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorRes{
			IsError:    true,
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
		})
	}

	return ctx.JSON(dto.NewCalenderItemRes{
		IsError: false,
		ID:      fmt.Sprintf("Make %s, %s, %s, %s", newItem.Title, newItem.StartDate, newItem.EndDate, newItem.Content),
	})
}
