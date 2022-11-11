package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(ctx *fiber.Ctx, err error) error {
	fmt.Println("inside error handler ", err.Error())
	switch err.(type) {
	case *validator.InvalidValidationError:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error()})
	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error()})
	}
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler})
	app.Post("/login", LoginHandler)
	app.Post("/account", CreateAccountHandler)
	app.Put("/account", UpdateProfileHandler)
	app.Listen(":3000")
}
