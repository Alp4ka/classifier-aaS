package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func (s *HTTPServer) mwRequestID() fiber.Handler {
	return requestid.New()
}
