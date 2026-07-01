package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hanzjefferson/go-cleanarch/internal/delivery/http/handler"
)

type RouterConfig struct {
	Router fiber.Router
	AuthHandler *handler.AuthHandler
}

func (c *RouterConfig) SetupRouter(r fiber.Router){
	r.Post("/login", c.AuthHandler.Login)
	r.Post("/register", c.AuthHandler.Register)
}