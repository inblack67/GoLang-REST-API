package main

import (
	"fibreApi/auth"
	"fibreApi/cache"
	"fibreApi/db"
	"fibreApi/middlewares"
	"fibreApi/mysession"
	"fibreApi/stories"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App){
	// stories
	app.Get("/api/stories", middlewares.ProtectMe, stories.GetAllStories)
	app.Get("/api/stories/:id", middlewares.ProtectMe, stories.GetSingleStory)
	app.Post("/api/stories", middlewares.ProtectMe, stories.CreateStory)
	app.Delete("/api/stories/:id", middlewares.ProtectMe, stories.DeleteStory)
	
	// auth
	app.Get("/api/users", middlewares.ProtectMe, auth.GetAllUsers)
	app.Get("/api/users/:id", middlewares.ProtectMe, auth.GetSingleUser )
	app.Post("/api/register", middlewares.IsFree, auth.RegisterUser )
	app.Delete("/api/users/:id", middlewares.ProtectMe, auth.DeleteUser)
	app.Get("/api/me", middlewares.ProtectMe, auth.GetMe)
	app.Post("/api/login", middlewares.IsFree, auth.LoginUser)
	app.Post("/api/logout", middlewares.ProtectMe, auth.LogoutUser)
}

func main(){
	
	cache.StartRedis()
	db.ConnectDB()

	mysession.CreateStore()

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