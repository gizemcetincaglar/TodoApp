package controllers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gizemcetincaglar/todo_app/services"
	"github.com/gizemcetincaglar/todo_app/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz istek formatı",
		})
	}

	user := services.FindUserByUsername(req.Username)
	if user == nil || user.Password != req.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Geçersiz kullanıcı adı veya şifre",
		})
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Token oluşturulamadı",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Giriş başarılı",
		"token":   token,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
