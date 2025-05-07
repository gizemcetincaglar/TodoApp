package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gizemcetincaglar/todo_app/controllers"
)

func RegisterAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
}
