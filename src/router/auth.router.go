package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
)

func (r *FiberRouter) GetAuth(path string, handler func(ctx handler.AuthContext)) {
	r.auth.Get(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PostAuth(path string, handler func(handler.AuthContext)) {
	r.auth.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}
