package bootstrap

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	Config    *viper.Viper
	Log       *logrus.Logger
	Validator *validator.Validate
	Fiber     *fiber.App
	SQL       *sqlx.DB
}

func (b *Bootstrap) Boot() {

	addr := fmt.Sprintf("%s:%d",
		b.Config.GetString("server.host"),
		b.Config.GetInt("server.port"),
	)
	b.Fiber.Listen(addr)
}
