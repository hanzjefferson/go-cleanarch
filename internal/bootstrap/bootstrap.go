package bootstrap

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/hanzjefferson/go-cleanarch/internal/delivery/http"
	httpHandler "github.com/hanzjefferson/go-cleanarch/internal/delivery/http/handler"
	"github.com/hanzjefferson/go-cleanarch/internal/repository"
	"github.com/hanzjefferson/go-cleanarch/internal/usecase"
	"github.com/hanzjefferson/go-cleanarch/pkg/jwt"
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
	JWT       *jwt.Provider
}

func (b *Bootstrap) Boot() {
	userRepo := repository.NewUserRepo(b.SQL)

	authUc := usecase.NewAuthUseCase(b.Log, b.JWT, userRepo)

	httpRouter := &http.RouterConfig{
		AuthHandler: httpHandler.NewAuthHandler(b.Log, authUc),
	}

	b.Fiber.Route("api", httpRouter.SetupRouter)

	addr := fmt.Sprintf("%s:%d",
		b.Config.GetString("server.host"),
		b.Config.GetInt("server.port"),
	)
	b.Fiber.Listen(addr)
}
