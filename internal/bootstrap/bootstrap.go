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

func (b *Bootstrap) Booting() error {
	userRepo := repository.NewUserRepo(b.SQL)

	authUc := usecase.NewAuthUseCase(b.Log, b.JWT, userRepo)

	httpRouter := &http.RouterConfig{
		AuthHandler: httpHandler.NewAuthHandler(b.Log, authUc),
	}

	b.Fiber.Route("api", httpRouter.SetupRouter)
	return nil
}

func (b *Bootstrap) Boot() {
	debug := b.Config.GetBool("app.debug")

	for k, v := range b.Config.GetStringMapString("app.secret") {
		if v == "" {
			if debug {
				b.Log.Warnf("secret with key 'app.secret.%s' is empty!", k)
			} else {
				b.Log.Fatalf("secret with key 'app.secret.%s' is empty!", k)
			}
		}
	}

	if err := b.Booting(); err != nil {
		b.Log.Fatalf("boot failed: %+v", err)
	}

	addr := fmt.Sprintf("%s:%d",
		b.Config.GetString("server.host"),
		b.Config.GetInt("server.port"),
	)
	b.Fiber.Listen(addr)
}
