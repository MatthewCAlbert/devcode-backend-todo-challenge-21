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

	// TodoList
	a.Router.Get("/", todoListController.All)
	a.Router.Get("/:id", todoListController.GetOne)
	a.Router.Post("/", todoListController.Create)
	a.Router.Delete("/:id", todoListController.Delete)
	a.Router.Patch("/:id", todoListController.Update)
}
