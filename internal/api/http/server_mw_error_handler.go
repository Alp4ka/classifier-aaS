package http

import "github.com/gofiber/fiber/v2"

type HandlerResp struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message,omitempty"`
}

func (s *Server) mwErrorHandler(c *fiber.Ctx, err error) error {
	// TODO(Gorkovets Roman): User friendly error handling.

	return c.Status(fiber.StatusBadRequest).JSON(HandlerResp{
		Success: false,
		Message: err.Error(),
	})
}

var _ fiber.ErrorHandler = (*Server)(nil).mwErrorHandler
