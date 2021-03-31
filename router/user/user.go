package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tigerza117/eb_api/controller/user"
)

func New(app *fiber.App)  {
	controller := user.New()
	auth := app.Group("/auth")
	{
		auth.Post("/register", controller.Register)
		auth.Post("/login", controller.Login)
	}
}
