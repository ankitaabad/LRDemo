package main

import "github.com/gofiber/fiber/v2"
func ApiKeyMissing(ctx *fiber.Ctx) error{
	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Message":"apikey is missing"})
}

func ApiKeyOrSecretKeyMissing(ctx *fiber.Ctx) error{
	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Message":"apikey or apisecret is missing"})
}

func Unauthorized(ctx *fiber.Ctx) error{
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message":"authorization token is missing"})
}

func okResponse(ctx *fiber.Ctx, data []byte) error{
	ctx.Set("content-type", "application/json")
	return ctx.Status(fiber.StatusOK).Send(data)
}

func missingBody(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message":"Empty Body"})
}

func requestMalformed(ctx *fiber.Ctx,err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message":err.Error()})
}