package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tigerza117/eb_api/router/user"
)

func New(app *fiber.App)  {
	user.New(app)
}
