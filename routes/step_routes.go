package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gizemcetincaglar/todo_app/controllers"
	"github.com/gizemcetincaglar/todo_app/middleware"
)

func RegisterStepRoutes(api fiber.Router) {
	step := api.Group("/steps", middleware.JWTProtected())

	step.Get("/:todoId", controllers.GetSteps)
	step.Post("/:todoId", controllers.CreateStep)
	step.Put("/complete/:stepId", controllers.CompleteStep)
	step.Delete("/:stepId", controllers.DeleteStep)
}
