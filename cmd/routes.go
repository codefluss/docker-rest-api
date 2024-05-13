package main

import (
	"github.com/codefluss/docker-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/facts", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}