package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *HTTPServer) mwCors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
	})
}
