package bootstrap

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/hanzjefferson/go-cleanarch/internal/model"
	"github.com/spf13/viper"
)

type structValidator struct {
    validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
    return v.validate.Struct(out)
}

func NewFiber(validate *validator.Validate, config *viper.Viper) *fiber.App {
	appConfig := fiber.Config{
		AppName: config.GetString("app.name"),
		ServerHeader: config.GetString("app.name"),
		StructValidator: &structValidator{
			validate: validate,
		},
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var fiberErr *fiber.Error
			if errors.As(err, &fiberErr) {
				code = fiberErr.Code
			}

			return ctx.Status(code).JSON(model.HTTPResponse{
				Message: err.Error(),
			})
		},
	}

	app := fiber.New(appConfig)

	app.Use(logger.New())
	app.Use(recoverer.New())

	return app
}