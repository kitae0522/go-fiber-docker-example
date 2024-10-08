package handler

import (
	"go-fiber-docker-example/internal/dto"
	"go-fiber-docker-example/internal/service"
	"go-fiber-docker-example/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	newUser := new(dto.UserCreateProfileReq)
	if err := utils.Bind(ctx, newUser); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 프로필 설정 실패. Body Binding 과정에서 문제 발생", err)
	}

	err := h.userService.CreateProfile(newUser)
	if err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 프로필 실패. Repository에서 문제 발생", err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.UserCreateProfileRes{
		IsError:    false,
		StatusCode: fiber.StatusCreated,
		Message:    "✅ 프로필 설정 완료",
	})
}

func (h *UserHandler) GetProfileByTag(ctx *fiber.Ctx) error {
	profileTag := ctx.Params("userTag")
	targetUser := new(dto.UserGetProfileTagReq)

	if len(profileTag) == 0 {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 유저 조회 실패. Binding 과정에서 문제 발생", nil)
	}
	targetUser.UserTag = profileTag

	findUser, err := h.userService.GetProfileByTag(targetUser)
	if err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 유저 조회 실패. Repository에서 문제 발생", err)
	}

	if findUser == nil {
		statusCode := fiber.StatusNotFound
		errMessage := "❌ 유저 조회 실패. 해당하는 유저가 존재하지 않습니다."
		return ctx.Status(statusCode).JSON(dto.UserGetProfileRes{
			IsError:    false,
			StatusCode: statusCode,
			Message:    errMessage,
			User:       nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.UserGetProfileRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 유저 조회 완료",
		User:       findUser,
	})
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	profileTag := ctx.Params("userTag")
	updateValue := new(dto.UserUpdateProfileReq)

	if len(profileTag) == 0 {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 프로필 수정 실패. Binding 과정에서 문제 발생", nil)
	}
	if err := utils.Bind(ctx, updateValue); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 프로필 수정 실패. Body Binding 과정에서 문제 발생", err)
	}

	updatedUser, err := h.userService.UpdateProfile(updateValue)
	if err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 프로필 수정 실패. Repository에서 문제 발생", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.UserUpdateProfileRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 프로필 수정 완료",
		User:       updatedUser,
	})
}

func (h *UserHandler) DeleteProfile(ctx *fiber.Ctx) {}
