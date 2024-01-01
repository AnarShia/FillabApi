package main

import (
	"github.com/AnarShia/FillabApi/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/facts", handlers.CreateFact)

}
