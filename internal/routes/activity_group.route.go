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
	a.Router.Get("/:id", activityController.GetOne)
	a.Router.Post("/", activityController.Create)
	a.Router.Delete("/:id", activityController.Delete)
	a.Router.Patch("/:id", activityController.Update)
}
