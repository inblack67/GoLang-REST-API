package stories

import (
	"context"
	"encoding/json"
	"errors"
	"fibreApi/cache"
	"fibreApi/constants"
	"fibreApi/db"
	"fibreApi/models"
	"fibreApi/mysession"
	"fibreApi/types"
	"fmt"
	"log"
	"time"

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

	cachedMarshalledStories, getErr := cache.RedisClient.Get(context.Background(), constants.KGetStories).Result()

	if getErr != nil{

		fmt.Println("Get All Stories DB Query")

		var stories []models.Story

		dbc := db.PgConn.Find(&stories)

		if dbc.Error != nil{
			return ctx.Status(401).JSON(dbc.Error)
		}

		marshalledStories, marshallErr  := json.Marshal(stories)

		if marshallErr != nil{
			log.Fatal("marshallErr", marshallErr)
		}

		setErr := cache.RedisClient.Set(context.Background(), constants.KGetStories, marshalledStories, time.Hour * 24).Err()

		if setErr != nil{
			log.Fatal("setErr", setErr)
		}

		return ctx.Status(200).JSON(stories)
	}

	var cachedStories []models.Story

	unmarshalErr := json.Unmarshal([]byte(cachedMarshalledStories), &cachedStories)

	if unmarshalErr != nil{
		log.Fatal("unmarshalErr", unmarshalErr)
	}

	fmt.Println("Get All Stories Redis Query")

	return ctx.Status(200).JSON(cachedStories)
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
	dbc := db.PgConn.Find(&stories, models.Story{UserID: user.User.ID})
	if dbc.Error != nil{
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
		if notFoundErr || (models.Story{} == story) {
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
	newStory.UserID = user.User.ID

	err2 := db.PgConn.Create(&newStory).Error

	if err2 != nil{
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
		if notFoundErr || (models.Story{} == story){
			return ctx.Status(404).JSON(types.Status{Success: false, Message: "Story does not exist"})
		}

	if story.UserID != user.User.ID{
		return ctx.Status(404).JSON(types.Status{Success: false, Message: "Not Authorized"})
	}

	dbc := db.PgConn.Delete(&story, id)
	if dbc.Error != nil{
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(types.Status{Success: true, Message: "Story deleted successfully"})
}