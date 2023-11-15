package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	// "github.com/gofiber/storage/redis"
)

// type Auth interface {
// 	Limiter() func(fiberCtx *fiber.Ctx) error
// }

// type auth struct {
// }

// func NewLimiter() Auth {
// 	return &auth{}
// }

var Limiter = limiter.New(limiter.Config{
	Next: func(c *fiber.Ctx) bool {
		return c.Get("is_qa", "0") == "1"
	},
	KeyGenerator: func(c *fiber.Ctx) string {
		return fmt.Sprintf("%s#%s#%s", c.Get("X-Real-Ip"), c.Route().Method, c.Route().Path)
	},
	Max:                60,
	Expiration:         1 * time.Minute,
	SkipFailedRequests: true,
	LimitReached: func(c *fiber.Ctx) error {
		return fiber.ErrTooManyRequests
	},
	// Storage: redis.New(),
})
