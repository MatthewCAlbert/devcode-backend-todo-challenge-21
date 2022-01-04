package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Routes struct {
	Router *fiber.App
	DB     *gorm.DB
}

func (a *Routes) Init() {

	// Middleware
	a.Router.Use("/", func(c *fiber.Ctx) error {
		c.Accepts(fiber.MIMEApplicationJSON)
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.Next()
	})

	// Base Routes
	a.Router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ini aplikasi gan")
	})

	a.Router.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "Baduts not found")
	})

	// Init the Rest of Routes
	activityGroupRouter := ActivityGroupRoutes{
		DB: a.DB, Router: a.Router.Group("/activity-groups"),
	}
	activityGroupRouter.Init()

	todoListRouter := TodoListRoutes{
		DB: a.DB, Router: a.Router.Group("/todo-items"),
	}
	todoListRouter.Init()

}
