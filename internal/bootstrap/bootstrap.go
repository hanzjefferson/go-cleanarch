package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	Config    *viper.Viper
	Log       *logrus.Entry
	Validator *validator.Validate
	Fiber     *fiber.App
	SQL       *sql.DB
}

func (b *Bootstrap) Boot() {

	addr := fmt.Sprintf("%s:%d",
		b.Config.GetString("server.host"),
		b.Config.GetInt("server.port"),
	)
	b.Fiber.Listen(addr)
}
