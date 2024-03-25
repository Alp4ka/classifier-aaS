package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func (s *Server) mwRateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        s.rateLimit,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimiterMiddleware: limiter.SlidingWindow{},
	},
	)
}
