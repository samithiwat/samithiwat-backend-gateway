package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/samithiwat/samithiwat-backend-gateway/src/model"
)

type FiberRouter struct {
	*fiber.App
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New()

	r.Use(cors.New())
	r.Use(logger.New())

	r.Get("/docs/*", swagger.HandlerDefault)

	return &FiberRouter{r}
}

type FiberCtx struct {
	*fiber.Ctx
}

func NewFiberCtx(c *fiber.Ctx) *FiberCtx {
	return &FiberCtx{c}
}

func (c *FiberCtx) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberCtx) JSON(statusCode int, v interface{}) {
	c.Ctx.Status(statusCode).JSON(v)
}

func (c *FiberCtx) ID(id *int32) error {
	v, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	*id = int32(v)
	
	return nil
}

func (c *FiberCtx) PaginationQueryParam(query *model.PaginationQueryParams) error {
	if err := c.QueryParser(query); err != nil {
		return err
	}
	return nil
}
