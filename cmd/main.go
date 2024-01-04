package main

import (
	"github.com/AnarShia/FillabApi/database"
	"github.com/AnarShia/FillabApi/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDbSqlite()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// get all users
	app.Get("/users", handlers.ListUsers)

	// create a user
	app.Post("/users", handlers.CreateUser)

	// get a user
	app.Get("/users/:id", handlers.GetUser)

	// update a user
	app.Put("/users/:id", handlers.UpdateUser)

	// delete a user
	app.Delete("/users/:id", handlers.DeleteUser)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":4000")
}
