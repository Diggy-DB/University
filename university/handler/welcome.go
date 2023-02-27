package handler

import "github.com/gofiber/fiber/v2"

func Welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to my University API")
}
