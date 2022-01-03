package middleware

import (
	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandler struct {
}

func (e ErrorHandler) Default(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError
	message := ""
	status := ""

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
		message = e.Message
	}

	switch code {
	case 400:
		status = constant.STATUS_BADREQUEST
	case 404:
		status = constant.STATUS_NOT_FOUND
	default:
		status = "false"
	}

	var res = map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    make(map[string]string),
	}

	// Return statuscode with error message
	return c.Status(code).JSON(res)
}
