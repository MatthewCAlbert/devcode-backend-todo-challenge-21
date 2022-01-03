package routes

import (
	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/controller"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ActivityGroupRoutes struct {
	Router fiber.Router
	DB     *gorm.DB
}

func (a *ActivityGroupRoutes) Init() {
	activityController := controller.ActivityGroupController{DB: a.DB}

	// Activity Group
	a.Router.Get("/", activityController.All)
	a.Router.Get("/:id", func(c *fiber.Ctx) error {
		return activityController.HandleSendPlaceholder(c)
	})
	a.Router.Post("/", func(c *fiber.Ctx) error {
		return activityController.HandleSendPlaceholder(c)
	})
	a.Router.Delete("/:id", func(c *fiber.Ctx) error {
		return activityController.HandleSendPlaceholder(c)
	})
	a.Router.Patch("/:id", func(c *fiber.Ctx) error {
		return activityController.HandleSendPlaceholder(c)
	})
}
