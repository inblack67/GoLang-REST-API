package auth

import (
	"errors"
	"fibreApi/constants"
	"fibreApi/db"
	"fibreApi/models"
	"fibreApi/mysession"
	"fibreApi/structs"
	"fibreApi/types"
	"log"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Hello ...
type Hello struct{
	success bool
	msg string
}

// GetAllUsers ...
func GetAllUsers(ctx *fiber.Ctx) error{
	var users []models.User
	dbc := db.PgConn.Find(&users)
	if dbc.Error != nil {
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(users)
}

// GetSingleUser ...
func GetSingleUser(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var user models.User

	err := db.PgConn.Preload("Stories").Find(&user, id).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if notFoundErr {
			return ctx.Status(404).JSON(types.Status{Success: false, Message: "User does not exist"})
		}

	return ctx.Status(200).JSON(user)
}

// RegisterUser ...
func RegisterUser(ctx *fiber.Ctx) error{
	newUser := new(models.User)
	if err := ctx.BodyParser(newUser); err != nil {
            return err
    }

	validationError := newUser.ValidateMe()

	if validationError != nil{
		return ctx.Status(400).JSON(validationError)
	}

	hashedPassword, err := argon2id.CreateHash(newUser.Password, argon2id.DefaultParams)

	if err != nil{
		log.Fatal(err)
	}

	newUser.Password = hashedPassword

	myuuid, errUUID := uuid.NewV4()

	if errUUID != nil{
		log.Fatal("errUUID",errUUID)
	}

	newUser.UUID = myuuid

	err2 := db.PgConn.Create(&newUser).Error

	if err2 != nil {
		return ctx.Status(401).JSON(err2)
	}

	return ctx.Status(201).JSON(newUser)

}

// LoginUser ...
func LoginUser(ctx *fiber.Ctx) error{

	session, sessionErr := mysession.SessionStore.Get(ctx)

	if sessionErr != nil{
		log.Fatal(sessionErr)
	}

	credentials := new(structs.SLogin)

	if err := ctx.BodyParser(credentials); err != nil {
            return err
    }

	var user models.User

	err := db.PgConn.Find(&user, models.User{Username: credentials.Username}).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if notFoundErr {
			return ctx.Status(404).JSON(types.Status{Success: false, Message: "Invalid Credentials"})
		}

	isValidPassword, argonErr := argon2id.ComparePasswordAndHash(credentials.Password, user.Password)

	if argonErr != nil{
		log.Fatal(err)
	}

	if !isValidPassword{
		return ctx.Status(404).JSON(types.Status{Success: false, Message: "Invalid Credentials"})
	}

	session.Set(constants.KLogin, types.SSession{User: user})

	defer session.Save()

	return ctx.Status(200).JSON(types.Status{Success: true, Message: "Logged In"})
}

// GetMe ...
func GetMe(ctx *fiber.Ctx) error{

	session, sessionErr := mysession.SessionStore.Get(ctx)

	if sessionErr != nil{
		log.Fatal(sessionErr)
	}

	data := session.Get(constants.KLogin)

	return ctx.Status(200).JSON(data)
}

// LogoutUser ...
func LogoutUser(ctx *fiber.Ctx) error{

	session, sessionErr := mysession.SessionStore.Get(ctx)

	if sessionErr != nil{
		log.Fatal(sessionErr)
	}

	err := session.Destroy()

	if err != nil{
		log.Fatal(err)
	}

	return ctx.Status(200).JSON(types.Status{Success: true, Message: "Logged Out"})
}

// DeleteUser ...
func DeleteUser(ctx *fiber.Ctx) error{
	id := ctx.Params("id")

	var user models.User

	err := db.PgConn.Find(&user, id).Error

	notFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
		if notFoundErr {
			return ctx.Status(404).JSON(types.Status{Success: false, Message: "user does not exist"})
		}

	dbc := db.PgConn.Delete(&user, id)
	if dbc.Error != nil {
		return ctx.Status(401).JSON(dbc.Error)
	}
	return ctx.Status(200).JSON(types.Status{Success: true, Message: "user deleted successfully"})
}