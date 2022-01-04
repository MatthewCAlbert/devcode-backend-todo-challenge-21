package controller

import (
	"fmt"

	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ActivityGroupController struct {
	Controller
	DB *gorm.DB
}

func (c *ActivityGroupController) All(context *fiber.Ctx) error {
	res := []map[string]interface{}{}
	tx := c.DB.Model(&model.ActivityGroup{}).Find(&res)
	if tx.Error != nil {
		return tx.Error
	}

	return c.HandleSendOK(context, res, false)
}

func (c *ActivityGroupController) GetOne(context *fiber.Ctx) error {
	id := context.Params("id")

	res := map[string]interface{}{}
	tx := c.DB.Model(&model.ActivityGroup{}).Take(&res, "id = ?", id)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Activity with ID %s Not Found", id))
	}

	return c.HandleSendOK(context, res, false)
}

type ActivityGroup struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

func (c *ActivityGroupController) Create(context *fiber.Ctx) error {
	ag := new(ActivityGroup)

	context.BodyParser(ag)

	if ag.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "title cannot be null")
	}
	if ag.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email cannot be null")
	}

	newAg := model.ActivityGroup{Title: ag.Title, Email: ag.Email}
	res := c.DB.Create(&newAg)
	if res.Error != nil {
		return res.Error
	}

	return c.HandleSendCreated(context, newAg)
}

func (c *ActivityGroupController) Update(context *fiber.Ctx) error {
	id := context.Params("id")

	// Get Body
	updatedAg := new(ActivityGroup)

	context.BodyParser(updatedAg)

	if updatedAg.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "title cannot be null")
	}

	// Process
	ag := model.ActivityGroup{}
	tx := c.DB.Model(&model.ActivityGroup{}).First(&ag, "id = ?", id)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Activity with ID %s Not Found", id))
	}

	ag.Title = updatedAg.Title

	tx = c.DB.Save(&ag)
	if tx.Error != nil {
		return tx.Error
	}

	return c.HandleSendOK(context, ag, true)
}

func (c *ActivityGroupController) Delete(context *fiber.Ctx) error {
	id := context.Params("id")

	ag := model.ActivityGroup{}
	tx := c.DB.Model(&model.ActivityGroup{}).First(&ag, "id = ?", id)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Activity with ID %s Not Found", id))
	}

	tx = c.DB.Delete(&ag)
	if tx.Error != nil {
		return tx.Error
	}

	return c.HandleSendPlaceholder(context)
}
