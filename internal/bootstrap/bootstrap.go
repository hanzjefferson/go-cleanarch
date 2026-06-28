package bootstrap

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/hanzjefferson/go-cleanarch/internal/delivery/http_fiber"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	Viper     *viper.Viper
	Logrus    *logrus.Entry
	Validator *validator.Validate
	Fiber     *fiber.App
}

func (b *Bootstrap) Boot() {
	http_fiber.RegisterMiddlewares(b.Fiber)
	http_fiber.RegisterHandlers(b.Fiber)

	addr := fmt.Sprintf("%s:%d",
		b.Viper.GetString("server.host"),
		b.Viper.GetInt("server.port"),
	)
	b.Fiber.Listen(addr)
}
