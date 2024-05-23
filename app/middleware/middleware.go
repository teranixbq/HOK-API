package middleware

import (
	"sync"
	"github.com/gofiber/fiber/v2"
)

var RequestCount int
var mutex sync.Mutex

func RequestLimitMiddleware(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()

	if RequestCount > 1000 {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"message": "Too many requests",
		})
	}
	RequestCount++
	return c.Next()
}