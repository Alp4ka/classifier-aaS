package api

import "github.com/gofiber/fiber/v2"

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *HTTPServer) mwErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
		Success: false,
		Message: err.Error(),
	})
}

var _ fiber.ErrorHandler = (*HTTPServer)(nil).mwErrorHandler
