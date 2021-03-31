package user

import (
	"github.com/gofiber/fiber/v2"
)

func (control controller) Login(ctx *fiber.Ctx) error {
	return ctx.SendString("LOGIN!")
}