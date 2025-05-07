package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gizemcetincaglar/todo_app/services"
)

type CreateTodoRequest struct {
	Name string `json:"name"`
}

func GetMyTodos(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)
	todos := services.GetTodosByUser(userID)
	return c.JSON(todos)
}

func GetAllTodos(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Yetkisiz erişim",
		})
	}
	todos := services.GetAllTodos()
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	var req CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz istek",
		})
	}
	userID := c.Locals("userId").(string)
	todo := services.CreateTodo(userID, req.Name)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	todoID := c.Params("id")
	if todoID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz ID",
		})
	}
	ok := services.DeleteTodo(todoID)
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "TO-DO bulunamadı veya zaten silinmiş",
		})
	}
	return c.JSON(fiber.Map{"message": "TO-DO başarıyla silindi"})
}
