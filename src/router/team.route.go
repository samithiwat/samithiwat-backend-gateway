package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
)

func (r *FiberRouter) GetTeam(path string, handler func(handler.TeamContext)) {
	r.App.Get(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) CreateTeam(path string, handler func(handler.TeamContext)) {
	r.App.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PatchTeam(path string, handler func(handler.TeamContext)) {
	r.App.Patch(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) DeleteTeam(path string, handler func(handler.TeamContext)) {
	r.App.Delete(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}
