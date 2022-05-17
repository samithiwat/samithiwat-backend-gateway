package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/middleware"
	"strconv"
)

type FiberRouter struct {
	*fiber.App
	auth fiber.Router
	user fiber.Router
	team fiber.Router
	org  fiber.Router
}

func NewFiberRouter(authGuard middleware.AuthGuard) *FiberRouter {
	r := fiber.New(fiber.Config{
		StrictRouting: true,
		AppName:       "Samithiwat.dev API",
	})

	r.Use(cors.New())
	r.Use(logger.New())

	r.Get("/docs/*", swagger.HandlerDefault)

	auth := NewGroupRoute(r, "/auth", authGuard.Validate)
	user := NewGroupRoute(r, "/user", authGuard.Validate)
	team := NewGroupRoute(r, "/team", authGuard.Validate)
	org := NewGroupRoute(r, "/organization", authGuard.Validate)

	return &FiberRouter{r, auth, user, team, org}
}

func NewGroupRoute(r *fiber.App, path string, middleware func(ctx middleware.AuthContext)) fiber.Router {
	return r.Group(path, func(c *fiber.Ctx) error {
		middleware(NewFiberCtx(c))
		return nil
	})
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

func (c *FiberCtx) ID() (id int32, err error) {
	v, err := c.ParamsInt("id")

	return int32(v), err
}

func (c *FiberCtx) UserID() int32 {
	id := c.Ctx.Get("UserId")

	result, err := strconv.Atoi(id)
	if err != nil {
		result = -1
	}

	return int32(result)
}

func (c *FiberCtx) PaginationQueryParam(query *dto.PaginationQueryParams) error {
	if err := c.QueryParser(query); err != nil {
		return err
	}

	return nil
}

func (c *FiberCtx) Token() string {
	return c.Ctx.Get(fiber.HeaderAuthorization, "")
}

func (c *FiberCtx) Path() string {
	return c.Ctx.Path()
}

func (c *FiberCtx) SetHeader(k string, v string) {
	c.Ctx.Set(k, v)
}

func (c *FiberCtx) Next() {
	c.Ctx.Next()
}
