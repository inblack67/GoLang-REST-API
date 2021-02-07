package stories

import (
	"github.com/gofiber/fiber/v2"
)

// GetAllStories ...
func GetAllStories(ctx *fiber.Ctx) error{
	return ctx.SendString("Get All Stories")
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