package stories

import (
	"errors"
	"fibreApi/db"
	"fibreApi/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Hello ...
type Hello struct{
	success bool
	msg string
}

// Status ...
type Status struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}

// GetAllStories ...
func GetAllStories(ctx *fiber.Ctx) error{
	var stories []models.Story
	dbc := db.PgConn.Find(&stories)
	if(dbc.Error != nil){
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(stories)
}

// GetSingleStory ...
func GetSingleStory(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var story models.Story

	err := db.PgConn.Find(&story, id).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if(notFoundErr || models.Story{} == story){
			return ctx.Status(404).JSON(Status{Success: false, Message: "Story does not exists"})
		}

	return ctx.Status(200).JSON(story)
}

// CreateStory ...
func CreateStory(ctx *fiber.Ctx) error{
	newBook := new(models.Story)
	if err := ctx.BodyParser(newBook); err != nil {
            return err
    }
	err2 := db.PgConn.Create(&newBook).Error

	if(err2 != nil){
		return ctx.Status(401).JSON(err2)
	}

	return ctx.Status(201).JSON(newBook)

}

// DeleteStory ...
func DeleteStory(ctx *fiber.Ctx) error{
	id := ctx.Params("id")

	var story models.Story

	err := db.PgConn.Find(&story, id).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if(notFoundErr || models.Story{} == story){
			return ctx.Status(404).JSON(Status{Success: false, Message: "Story does not exists"})
		}

	dbc := db.PgConn.Delete(&story, id)
	if(dbc.Error != nil){
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(Status{Success: true, Message: "Story deleted successfully"})
}