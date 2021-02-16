package stories

import (
	"errors"
	"fibreApi/constants"
	"fibreApi/db"
	"fibreApi/models"
	"fibreApi/mysession"
	"fibreApi/types"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Hello ...
type Hello struct{
	success bool
	msg string
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

// GetMyStories ...
func GetMyStories(ctx *fiber.Ctx) error{

	session, sessionErr := mysession.SessionStore.Get(ctx)

	if sessionErr != nil{
		log.Fatal(sessionErr)
	}

	data := session.Get(constants.KLogin)

	user, ok := data.(types.SSession)

	if !ok{
		log.Fatal("session typecast err")
	}

	var stories []models.Story
	dbc := db.PgConn.Find(&stories, models.Story{UserID: user.User})
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
			return ctx.Status(404).JSON(types.Status{Success: false, Message: "Story does not exist"})
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

	myuuid, errUUID := uuid.NewV4()

	if errUUID != nil{
		log.Fatal("errUUID",errUUID)
	}

	session, sessionErr := mysession.SessionStore.Get(ctx)

	if sessionErr != nil{
		log.Fatal(sessionErr)
	}

	data := session.Get(constants.KLogin)

	user, ok := data.(types.SSession)

	if !ok{
		log.Fatal("session typecast err")
	}

	newStory.UUID = myuuid
	newStory.UserID = user.User

	err2 := db.PgConn.Create(&newStory).Error

	if(err2 != nil){
		return ctx.Status(401).JSON(err2)
	}

	return ctx.Status(201).JSON(newStory)

}

// DeleteStory ...
func DeleteStory(ctx *fiber.Ctx) error{

	session, sessionErr := mysession.SessionStore.Get(ctx)

	if sessionErr != nil{
		log.Fatal(sessionErr)
	}

	data := session.Get(constants.KLogin)

	user, ok := data.(types.SSession)

	if !ok{
		log.Fatal("session typecast err")
	}

	id := ctx.Params("id")

	var story models.Story

	err := db.PgConn.Find(&story, id).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if(notFoundErr || models.Story{} == story){
			return ctx.Status(404).JSON(types.Status{Success: false, Message: "Story does not exist"})
		}

	if story.UserID != user.User{
		return ctx.Status(404).JSON(types.Status{Success: false, Message: "Not Authorized"})
	}

	dbc := db.PgConn.Delete(&story, id)
	if(dbc.Error != nil){
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(types.Status{Success: true, Message: "Story deleted successfully"})
}