package middlewares

import (
	"fibreApi/constants"
	"fibreApi/mysession"
	"fibreApi/types"
	"log"

	"github.com/gofiber/fiber/v2"
)

// ProtectMe ...
func ProtectMe(ctx *fiber.Ctx) error{
	session, err := mysession.SessionStore.Get(ctx)

	if err != nil{
		log.Fatal(err)
	}

	user := session.Get(constants.KLogin)

	if user == nil{
		return ctx.Status(401).JSON(types.Status{Success: false, Message: "Not Authenticated"})
	}

	return ctx.Next()
}

// IsFree ...
func IsFree(ctx *fiber.Ctx) error{
	session, err := mysession.SessionStore.Get(ctx)

	if err != nil{
		log.Fatal(err)
	}

	user := session.Get(constants.KLogin)

	if user != nil{
		return ctx.Status(401).JSON(types.Status{Success: false, Message: "Not Authorized"})
	}

	return ctx.Next()
}