package handler

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-docker-example/internal/dto"
	"go-fiber-docker-example/internal/model"
	"go-fiber-docker-example/internal/service"
	"go-fiber-docker-example/pkg/utils"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	user := new(dto.AuthRegisterReq)
	if err := utils.Bind(ctx, user); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 회원가입 실패. Body Binding 과정에서 문제 발생", err)
	}

	err := h.authService.Register(user)
	if err != nil {
		if _, uniqueErr := model.IsErrUniqueConstraint(err); uniqueErr {
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 회원가입 실패. 중복된 유저가 존재합니다.", err)
		}
		return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 회원가입 실패. Repository에서 문제 발생", err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.AuthRegisterRes{
		IsError:    false,
		StatusCode: fiber.StatusCreated,
		Message:    "✅ 회원가입 완료",
	})
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	user := new(dto.AuthLoginReq)
	if err := utils.Bind(ctx, user); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 로그인 실패. Body Binding 과정에서 문제 발생", err)
	}

	findUser, err := h.authService.Login(user)
	if err != nil {
		switch err {
		case model.ErrNotFound:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. 존재하지 않는 사용자입니다.", err)
		default:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. Repository에서 문제 발생", err)
		}
	}
	if findUser == nil {
		return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. 패스워드가 일치하지 않습니다.", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.AuthLoginRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 로그인 완료",
		User:       findUser,
	})
}
