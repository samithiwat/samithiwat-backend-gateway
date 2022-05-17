package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
)

func (r *FiberRouter) GetUser(path string, handler func(ctx handler.UserContext)) {
	r.user.Get(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) CreateUser(path string, handler func(handler.UserContext)) {
	r.user.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PatchUser(path string, handler func(handler.UserContext)) {
	r.user.Patch(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) DeleteUser(path string, handler func(handler.UserContext)) {
	r.user.Delete(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}
