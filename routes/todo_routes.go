package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gizemcetincaglar/todo_app/controllers"
	"github.com/gizemcetincaglar/todo_app/middleware"
)

func RegisterTodoRoutes(api fiber.Router) {
	todo := api.Group("/todos", middleware.JWTProtected())

	todo.Get("/mine", controllers.GetMyTodos)
	todo.Get("/all", controllers.GetAllTodos)
	todo.Post("/", controllers.CreateTodo)
	todo.Delete("/:id", controllers.DeleteTodo)
}
