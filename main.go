package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tigerza117/eb_api/router"
)

func main()  {
	app := fiber.New()

	router.New(app)

	if err := app.Listen(":8888"); err != nil {
		log.Fatal(err)
	}
}