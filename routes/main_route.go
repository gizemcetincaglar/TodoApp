package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Her alt modül için route'ları yüklemek için fonksiyonlar
	RegisterAuthRoutes(api)
	RegisterTodoRoutes(api)
	RegisterStepRoutes(api)
}
