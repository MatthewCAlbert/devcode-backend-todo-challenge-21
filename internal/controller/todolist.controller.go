package controller

import (
	"fmt"

	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// var querySelectCast = []string{
// 	"CAST(activity_group_id as CHAR(21)) as activity_group_id",
// 	"id", "title", "is_active", "priority", "created_at", "updated_at", "deleted_at",
// }

type TodoListController struct {
	Controller
	DB *gorm.DB
}

type QueryTodo struct {
	ActivityGroupId string `query:"activity_group_id"`
}

func (c *TodoListController) All(context *fiber.Ctx) error {

	// Parse Query
	q := new(QueryTodo)
	context.QueryParser(q)

	var tx *gorm.DB
	res := []map[string]interface{}{}

	if q.ActivityGroupId != "" {
		tx = c.DB.Model(&model.Todo{}).Find(&res, "activity_group_id = ? and is_active = true", q.ActivityGroupId)
	} else {
		tx = c.DB.Model(&model.Todo{}).Find(&res, "is_active = true")
	}
	if tx.Error != nil {
		return tx.Error
	}

	return c.HandleSendOK(context, res, false)
}

func (c *TodoListController) GetOne(context *fiber.Ctx) error {
	id := context.Params("id")

	res := map[string]interface{}{}
	tx := c.DB.Model(&model.Todo{}).Take(&res, "id = ?", id)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Todo with ID %s Not Found", id))
	}

	return c.HandleSendOK(context, res, false)
}

type CreateTodo struct {
	ActivityGroupID *uint   `json:"activity_group_id" validate:"required"`
	Title           string  `json:"title" validate:"required,min=1"`
	IsActive        *bool   `json:"is_active"`
	Priority        *string `json:"priority" validate:"min=1,max=15"`
}

func (c *TodoListController) Create(context *fiber.Ctx) error {
	todo := new(CreateTodo)

	// Parse In
	context.BodyParser(todo)
	if todo.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "title cannot be null")
	}
	if todo.ActivityGroupID == nil {
		return fiber.NewError(fiber.StatusBadRequest, "activity_group_id cannot be null")
	}

	// Get
	ag := model.ActivityGroup{}
	tx := c.DB.Model(&model.ActivityGroup{}).First(&ag, "id = ?", *todo.ActivityGroupID)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Activity with ID %d Not Found", *todo.ActivityGroupID))
	}

	// Create
	newTodo := model.Todo{
		Title: todo.Title, ActivityGroup: ag,
		IsActive: true, Priority: "very-high",
	}
	res := c.DB.Create(&newTodo)
	if res.Error != nil {
		return res.Error
	}

	return c.HandleSendCreated(context, newTodo)
}

type UpdateTodo struct {
	Title    *string `json:"title" validate:"min=1"`
	IsActive *bool   `json:"is_active"`
	Priority *string `json:"priority" validate:"min=1,max=15"`
}

func (c *TodoListController) Update(context *fiber.Ctx) error {
	id := context.Params("id")

	// Get Body
	updatedTodo := new(UpdateTodo)

	context.BodyParser(updatedTodo)

	// Process
	todo := model.Todo{}
	tx := c.DB.Model(&model.Todo{}).First(&todo, "id = ?", id)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Todo with ID %s Not Found", id))
	}

	if updatedTodo.Title != nil && *updatedTodo.Title != "" {
		todo.Title = *updatedTodo.Title
	}
	if updatedTodo.IsActive != nil {
		todo.IsActive = *updatedTodo.IsActive
	}
	if updatedTodo.Priority != nil {
		todo.Priority = *updatedTodo.Priority
	}

	tx = c.DB.Save(&todo)
	if tx.Error != nil {
		return tx.Error
	}

	return c.HandleSendOK(context, todo, true)
}

func (c *TodoListController) Delete(context *fiber.Ctx) error {
	id := context.Params("id")

	ag := model.Todo{}
	tx := c.DB.Model(&model.Todo{}).First(&ag, "id = ?", id)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Todo with ID %s Not Found", id))
	}

	tx = c.DB.Delete(&ag)
	if tx.Error != nil {
		return tx.Error
	}

	return c.HandleSendPlaceholder(context)
}
