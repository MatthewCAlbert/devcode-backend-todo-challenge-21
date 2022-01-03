package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ActivityGroupController struct {
	Controller
	DB *gorm.DB
}

func (c *ActivityGroupController) All(context *fiber.Ctx) error {
	var res = map[string]interface{}{
		"budiman": "saya",
	}
	return c.Controller.HandleSendOK(context, res)
}

func (c *ActivityGroupController) GetOne(id uint) {

}

func (c *ActivityGroupController) Create(id uint) {

}

func (c *ActivityGroupController) Update(id uint) {

}

func (c *ActivityGroupController) Delete(id uint) {

}
