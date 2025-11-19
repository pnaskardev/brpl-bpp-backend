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


	
	// 404 Handler
    app.Use(func(c *fiber.Ctx) error {
        return c.SendStatus(404) // => 404 "Not Found"
    })

	log.Fatal(app.Listen(":3000"))

}
