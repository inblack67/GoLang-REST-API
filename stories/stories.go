package stories

import (
	"fibreApi/db"
	"fibreApi/models"

	"github.com/gofiber/fiber/v2"
)

// Hello ...
type Hello struct{
	success bool
	msg string
}

// GetAllStories ...
func GetAllStories(ctx *fiber.Ctx) error{
	var stories []models.Story
	db.PgConn.Find(&stories)
	return ctx.JSON(stories)
}

// GetSingleStory ...
func GetSingleStory(ctx *fiber.Ctx) error{
	return ctx.SendString("Get Single Story")
}

// CreateStory ...
func CreateStory(ctx *fiber.Ctx) error{
	return ctx.SendString("Add new story")
}

// DeleteStory ...
func DeleteStory(ctx *fiber.Ctx) error{
	return ctx.SendString("Delete story")
}