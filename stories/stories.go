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

// Status ...
type Status struct{
	success bool
	msg string
}

// GetAllStories ...
func GetAllStories(ctx *fiber.Ctx) error{
	var stories []models.Story
	db.PgConn.Find(&stories)
	return ctx.Status(200).JSON(stories)
}

// GetSingleStory ...
func GetSingleStory(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var story models.Story
	db.PgConn.Find(&story, id)
	if story.Title == ""{
		return ctx.Status(404).JSON(Status{success: false, msg: "Story does not exists"})
	}
	return ctx.Status(200).JSON(story)
}

// CreateStory ...
func CreateStory(ctx *fiber.Ctx) error{
	newBook := new(models.Story)
	if err := ctx.BodyParser(newBook); err != nil {
            return err
    }
	db.PgConn.Create(&newBook)
	return ctx.Status(201).JSON(newBook)
}

// DeleteStory ...
func DeleteStory(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var story models.Story
	db.PgConn.Find(&story, id)
	if story.Title == ""{
		return ctx.Status(404).JSON(Status{success: false, msg: "Story does not exists"})
	}
	db.PgConn.Delete(&story, id)
	return ctx.Status(200).JSON(Status{success: true, msg: "Story deleted successfully"})
}