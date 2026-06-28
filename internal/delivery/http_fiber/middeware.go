package http_fiber

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func RegisterMiddlewares(r fiber.Router){
	r.Use(logger.New())
}