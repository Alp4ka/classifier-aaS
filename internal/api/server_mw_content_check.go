package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (s *HTTPServer) mwContentChecker() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodGet || c.Is("json") {
			return c.Next()
		}
		return errors.New("only JSON requests are allowed")
	}
}
