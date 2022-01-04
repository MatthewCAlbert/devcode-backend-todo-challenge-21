package controller

import (
	"reflect"

	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/constant"
	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/pkg/utils"
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

func (c *Controller) HandleSendOK(context *fiber.Ctx, data interface{}, needLower bool) error {
	context.Status(fiber.StatusOK)
	if !needLower {
		return context.JSON(c.GetJsonResponse(data))
	}
	return context.JSON(c.GetJsonResponse(c.structToValidDataResponse(data, []string{})))
}

func (c *Controller) HandleSendCreated(context *fiber.Ctx, data interface{}) error {
	context.Status(fiber.StatusCreated)
	return context.JSON(c.GetJsonResponse(c.structToValidDataResponse(data, []string{"deleted_at"})))
}

func (c *Controller) HandleSendPlaceholder(context *fiber.Ctx) error {
	context.Status(fiber.StatusOK)
	return context.JSON(c.GetJsonResponse(make(map[string]string)))
}

func (c *Controller) structToValidDataResponse(in interface{}, exclude []string) map[string]interface{} {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	v := reflect.ValueOf(in)
	vType := v.Type()

	result := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		name := vType.Field(i).Name
		name = utils.ToSnakeCase(name)

		// Kick relational field (cmn kepake di create sih)
		if utils.StringInSlice(name, []string{"todos", "activity_group"}) {
			continue
		}

		if utils.StringInSlice(name, exclude) {
			continue
		}

		result[name] = v.Field(i).Interface()
	}

	delete(result, "model")

	return result

}
