package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/pnaskardev/brpl-bpp/beckn"
)

func main() {
	fmt.Println("Hello World")

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("OK!!")
	})

	api := app.Group("/bpp")

	api.Get("/select", beckn.Select)

	log.Fatal(app.Listen(":3000"))

}
