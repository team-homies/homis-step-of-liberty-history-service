package middleware

import "github.com/gofiber/fiber/v2"

func JWTMiddleware(c *fiber.Ctx) error {
	// tokenString := c.Get("X-Authorization")

	c.Locals("user_id", 1)

	return c.Next()
}
