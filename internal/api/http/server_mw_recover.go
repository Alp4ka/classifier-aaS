package http

import (
	"github.com/gofiber/fiber/v2"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *Server) mwRecoverer() fiber.Handler {
	return fiberrecover.New(fiberrecover.Config{EnableStackTrace: true})
}
