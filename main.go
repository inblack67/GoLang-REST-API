package main

import (
	"fibreApi/auth"
	"fibreApi/db"
	"fibreApi/stories"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App){
	// stories
	app.Get("/api/stories", stories.GetAllStories)
	app.Get("/api/stories/:id", stories.GetSingleStory)
	app.Post("/api/stories", stories.CreateStory)
	app.Delete("/api/stories/:id", stories.DeleteStory)
	
	// auth
	app.Get("/api/users", auth.GetAllUsers)
	app.Get("/api/users/:id", auth.GetSingleUser )
	app.Post("/api/register", auth.RegisterUser )
	app.Delete("/api/users/:id", auth.DeleteUser)
	app.Get("/api/me", auth.GetMe)
	// app.Post("/api/login", auth.RegisterUser)
}

func main(){
	
	db.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4000, http://localhost:3000",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
		AllowMethods: "GET,POST,DELETE",
	}))
	
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("API up and running")
	})

	setupRoutes(app)

	const PORT = ":5000"
	fmt.Println("Server starting on port", PORT)
	log.Fatal(app.Listen(PORT))
}