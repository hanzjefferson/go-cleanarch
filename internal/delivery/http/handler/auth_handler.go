package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hanzjefferson/go-cleanarch/internal/model"
	"github.com/hanzjefferson/go-cleanarch/internal/usecase"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	Log         *logrus.Logger
	AuthUseCase *usecase.AuthUseCase
}

func NewAuthHandler(
	log *logrus.Logger,
	authUseCase *usecase.AuthUseCase,
) *AuthHandler {
	return &AuthHandler{
		Log:         log,
		AuthUseCase: authUseCase,
	}
}

func (h *AuthHandler) Login(ctx fiber.Ctx) error {
	req := new(model.LoginRequest)
	if err := ctx.Bind().Body(req); err != nil {
		h.Log.Warnf("failed to bind request: %+v", err)
		return fiber.ErrBadRequest
	}

	data, err := h.AuthUseCase.Login(req)
	if err != nil {
		h.Log.Warnf("failed to login user: %+v", err)
		return err
	}

	return ctx.JSON(
		&model.HTTPResponse{
			Message: "login succesfully",
			Data: data,
		},
	)
}

func (h *AuthHandler) Register(ctx fiber.Ctx) error {
	req := new(model.RegisterRequest)
	if err := ctx.Bind().Body(req); err != nil {
		h.Log.Warnf("failed to destruct req body:%+v", err)
		return fiber.ErrBadRequest
	}

	_, err := h.AuthUseCase.Create(req)
	if err != nil {
		return err
	}

	return ctx.JSON(
		model.HTTPResponse{
			Message: "register succesfully",
		},
	)
}
