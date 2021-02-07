package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main(){

	app := fiber.New()
	
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("API up and running")
	})

	const PORT = ":5000"
	log.Fatal(app.Listen(PORT))
}