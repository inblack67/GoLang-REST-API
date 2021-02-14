package auth

import (
	"errors"
	"fibreApi/db"
	"fibreApi/models"
	"log"
	"time"

	"github.com/alexedwards/argon2id"
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

// GetAllUsers ...
func GetAllUsers(ctx *fiber.Ctx) error{
	var users []models.User
	dbc := db.PgConn.Find(&users)
	if(dbc.Error != nil){
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(users)
}

// GetSingleUser ...
func GetSingleUser(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var user models.User

	err := db.PgConn.Find(&user, id).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if(notFoundErr || models.User{} == user){
			return ctx.Status(404).JSON(Status{Success: false, Message: "User does not exist"})
		}

	return ctx.Status(200).JSON(user)
}

// RegisterUser ...
func RegisterUser(ctx *fiber.Ctx) error{
	newUser := new(models.User)
	if err := ctx.BodyParser(newUser); err != nil {
            return err
    }

	hashedPassword, err := argon2id.CreateHash(newUser.Password, argon2id.DefaultParams)

	if err != nil{
		log.Fatal(err)
	}

	newUser.Password = hashedPassword
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	err2 := db.PgConn.Create(&newUser).Error

	if(err2 != nil){
		return ctx.Status(401).JSON(err2)
	}

	return ctx.Status(201).JSON(newUser)

}

// DeleteUser ...
func DeleteUser(ctx *fiber.Ctx) error{
	id := ctx.Params("id")

	var user models.User

	err := db.PgConn.Find(&user, id).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if(notFoundErr || models.User{} == user){
			return ctx.Status(404).JSON(Status{Success: false, Message: "user does not exist"})
		}

	dbc := db.PgConn.Delete(&user, id)
	if(dbc.Error != nil){
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(Status{Success: true, Message: "user deleted successfully"})
}