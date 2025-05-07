package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gizemcetincaglar/todo_app/routes"
	"github.com/gizemcetincaglar/todo_app/services"
)

func main() {
	
	err := services.LoadUsersFromFile()
	if err != nil {
		log.Fatalf("Kullanıcılar yüklenemedi: %v", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("TodoApp API is running ✅")
	})

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Sunucu %s portunda dinleniyor...", port)
	log.Fatal(app.Listen(":" + port))
}
