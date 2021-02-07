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
	id := ctx.Params("id")
	var story models.Story
	db.PgConn.Find(&story, id)
	return ctx.JSON(story)
}

// CreateStory ...
func CreateStory(ctx *fiber.Ctx) error{
	newBook := new(models.Story)
	if err := ctx.BodyParser(newBook); err != nil {
            return err
    }
	db.PgConn.Create(&newBook)
	return ctx.JSON(newBook)
}

// DeleteStory ...
func DeleteStory(ctx *fiber.Ctx) error{
	return ctx.SendString("Delete story")
}