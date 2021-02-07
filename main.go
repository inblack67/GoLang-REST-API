package main

import (
	"fibreApi/stories"
	"log"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/stories", stories.GetAllStories)
	app.Get("/api/stories/:id", stories.GetSingleStory)
	app.Post("/api/stories", stories.CreateStory)
	app.Delete("/api/stories/:id", stories.DeleteStory)
}

func main(){

	app := fiber.New()
	
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("API up and running")
	})

	setupRoutes(app)

	const PORT = ":5000"
	log.Fatal(app.Listen(PORT))
}