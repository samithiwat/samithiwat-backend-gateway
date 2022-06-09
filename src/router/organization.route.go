package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
)

func (r *FiberRouter) GetOrganization(path string, handler func(ctx handler.OrganizationContext)) {
	r.org.Get(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) CreateOrganization(path string, handler func(handler.OrganizationContext)) {
	r.org.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PatchOrganization(path string, handler func(handler.OrganizationContext)) {
	r.org.Patch(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) DeleteOrganization(path string, handler func(handler.OrganizationContext)) {
	r.org.Delete(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}
