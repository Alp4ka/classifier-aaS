package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func (s *Server) mwRequestID() fiber.Handler {
	return requestid.New()
}
