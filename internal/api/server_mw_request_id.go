package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func (s *HTTPServer) mwGetRequestIDer() fiber.Handler {
	return requestid.New()
}
