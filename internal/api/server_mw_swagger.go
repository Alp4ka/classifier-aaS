package api

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func (s *HTTPServer) mwSwagger() fiber.Handler {
	return swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.yaml",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	})
}
