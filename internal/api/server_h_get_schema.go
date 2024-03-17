package api

import (
	"github.com/gofiber/fiber/v2"
)

func (s *HTTPServer) hGetSchema(ctx *fiber.Ctx) error {
	return nil
}

var _ fiber.Handler = (*HTTPServer)(nil).hCreateSchema
