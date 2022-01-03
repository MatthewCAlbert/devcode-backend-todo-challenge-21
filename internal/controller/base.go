package controller

import (
	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/constant"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (c *Controller) GetJsonResponse(data interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"status":  constant.STATUS_SUCCESS,
		"message": constant.STATUS_SUCCESS,
		"data":    data,
	}

	return res
}

func (c *Controller) HandleSendOK(context *fiber.Ctx, data interface{}) error {
	context.Status(fiber.StatusOK)
	return context.JSON(c.GetJsonResponse(data))
}

func (c *Controller) HandleSendCreated(context *fiber.Ctx, data interface{}) error {
	context.Status(fiber.StatusCreated)
	return context.JSON(c.GetJsonResponse(data))
}

func (c *Controller) HandleSendPlaceholder(context *fiber.Ctx) error {
	context.Status(fiber.StatusOK)
	return context.JSON(c.GetJsonResponse(make(map[string]string)))
}
