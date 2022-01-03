package routes

import (
	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/controller"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodoListRoutes struct {
	Router fiber.Router
	DB     *gorm.DB
}

func (a *TodoListRoutes) Init() {
	todoListController := controller.TodoListController{DB: a.DB}

	// Activity Group
	// TodoList
	a.Router.Get("/", todoListController.All)
	a.Router.Get("/:id", func(c *fiber.Ctx) error {
		return todoListController.HandleSendPlaceholder(c)
	})
	a.Router.Post("/", func(c *fiber.Ctx) error {
		return todoListController.HandleSendPlaceholder(c)
	})
	a.Router.Delete("/:id", func(c *fiber.Ctx) error {
		return todoListController.HandleSendPlaceholder(c)
	})
	a.Router.Patch("/:id", func(c *fiber.Ctx) error {
		return todoListController.HandleSendPlaceholder(c)
	})
}
