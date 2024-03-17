package api

import (
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
)

func (s *HTTPServer) mwLogging(c *fiber.Ctx) error {
	rid := c.Locals(requestid.ConfigDefault.ContextKey).(string)

	start := time.Now()
	mlogger.L(c.UserContext()).Info(
		"Request",
		field.String("rid", rid),
		field.String("method", c.Method()),
		field.String("path", c.Path()),
		field.String("req", string(c.Body())),
		field.String("start", start.String()),
	)

	err := c.Next()
	if err != nil {
		if err = c.App().ErrorHandler(c, err); err != nil {
			_ = c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	end := time.Now()
	mlogger.L(c.UserContext()).Info(
		"Response",
		field.String("rid", rid),
		field.String("resp", string(c.Response().Body())),
		field.Int("status", c.Response().StatusCode()),
		field.String("start", start.String()),
		field.String("end", end.String()),
		field.String("latency", end.Sub(start).String()),
	)

	return err
}
