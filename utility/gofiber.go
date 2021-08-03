package utility

import "github.com/gofiber/fiber/v2"

func FiberError(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{"message": message})
}

func FiberSuccess(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{"message": message})
}
