package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
)

func (r *FiberRouter) GetUser(path string, handler func(ctx service.UserContext)) {
	r.App.Get(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) CreateUser(path string, handler func(service.UserContext)) {
	r.App.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PatchUser(path string, handler func(service.UserContext)) {
	r.App.Patch(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) DeleteUser(path string, handler func(service.UserContext)) {
	r.App.Delete(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}
