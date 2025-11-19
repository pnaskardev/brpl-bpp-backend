package beckn

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Select(c *fiber.Ctx) error {

	var body = c.Body()

	fmt.Println("INCOMING SELECT REQUEST", body)

	return c.Status(200).JSON("SELECT CALLED")
}
