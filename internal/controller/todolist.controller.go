package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodoListController struct {
	Controller
	DB *gorm.DB
}

func (c *TodoListController) All(context *fiber.Ctx) error {
	var res = map[string]interface{}{
		"budiman": "saya",
	}
	return c.Controller.HandleSendOK(context, res)
}

func (c *TodoListController) GetOne(id uint) {

}

func (c *TodoListController) Create(id uint) {

}

func (c *TodoListController) Update(id uint) {

}

func (c *TodoListController) Delete(id uint) {

}
