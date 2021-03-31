package user

import (
	"github.com/gofiber/fiber/v2"
)

func (control controller) Register(ctx *fiber.Ctx) error {
	return ctx.SendString("REGISTER!")
}
