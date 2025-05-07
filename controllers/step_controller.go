package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gizemcetincaglar/todo_app/services"
)

type CreateStepRequest struct {
	Content string `json:"content"`
}

func GetSteps(c *fiber.Ctx) error {
	todoID := c.Params("todoId")
	steps := services.GetStepsByTodoID(todoID)
	return c.JSON(steps)
}

func CreateStep(c *fiber.Ctx) error {
	todoID := c.Params("todoId")

	var req CreateStepRequest
	if err := c.BodyParser(&req); err != nil || req.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "İçerik boş olamaz",
		})
	}

	step := services.CreateStep(todoID, req.Content)
	return c.Status(fiber.StatusCreated).JSON(step)
}

func CompleteStep(c *fiber.Ctx) error {
	stepID := c.Params("stepId")
	if services.MarkStepCompleted(stepID) {
		return c.JSON(fiber.Map{"message": "Adım tamamlandı"})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Adım bulunamadı"})
}

func DeleteStep(c *fiber.Ctx) error {
	stepID := c.Params("stepId")
	if services.DeleteStep(stepID) {
		return c.JSON(fiber.Map{"message": "Adım silindi"})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Adım bulunamadı"})
}
