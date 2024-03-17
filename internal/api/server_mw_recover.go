package api

import (
	"github.com/gofiber/fiber/v2"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *HTTPServer) mwGetRecoverer() fiber.Handler {
	return fiberrecover.New()
}
