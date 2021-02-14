package stories

import (
	"errors"
	"fibreApi/db"
	"fibreApi/models"
	"time"

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
			return ctx.Status(404).JSON(Status{Success: false, Message: "Story does not exist"})
		}

	return ctx.Status(200).JSON(story)
}

// CreateStory ...
func CreateStory(ctx *fiber.Ctx) error{
	newStory := new(models.Story)
	if err := ctx.BodyParser(newStory); err != nil {
            return err
    }

	validationError := newStory.ValidateStory()

	if validationError != nil{
		return ctx.Status(400).JSON(validationError)
	}


	newStory.CreatedAt = time.Now()
	newStory.UpdatedAt = time.Now()

	err2 := db.PgConn.Create(&newStory).Error

	if(err2 != nil){
		return ctx.Status(401).JSON(err2)
	}

	return ctx.Status(201).JSON(newStory)

}

// DeleteStory ...
func DeleteStory(ctx *fiber.Ctx) error{
	id := ctx.Params("id")

	var story models.Story

	err := db.PgConn.Find(&story, id).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if(notFoundErr || models.Story{} == story){
			return ctx.Status(404).JSON(Status{Success: false, Message: "Story does not exist"})
		}

	dbc := db.PgConn.Delete(&story, id)
	if(dbc.Error != nil){
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(Status{Success: true, Message: "Story deleted successfully"})
}