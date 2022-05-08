package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
)

func (r *FiberRouter) GetOrganization(path string, handler func(ctx service.OrganizationContext)) {
	r.App.Get(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) CreateOrganization(path string, handler func(service.OrganizationContext)) {
	r.App.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PatchOrganization(path string, handler func(service.OrganizationContext)) {
	r.App.Patch(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) DeleteOrganization(path string, handler func(service.OrganizationContext)) {
	r.App.Delete(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}
