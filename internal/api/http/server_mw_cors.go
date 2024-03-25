package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *Server) mwCors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
	})
}
