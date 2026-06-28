package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
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
	}

	app := fiber.New(appConfig)

	app.Use(logger.New())

	return app
}