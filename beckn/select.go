package beckn

import "github.com/gofiber/fiber/v2"

func Select(c *fiber.Ctx) error {
	return c.JSON("SELECT CALLED")
}
