package middleware

import (
	"strings"
	"time"

	"github.com/MQEnergy/mqshop/pkg/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
)

// CacheMiddleware http cache middleware
func CacheMiddleware() fiber.Handler {
	return cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		KeyGenerator: func(ctx *fiber.Ctx) string {
			keyG := ctx.IP() + ":" + ctx.Path()
			switch ctx.Method() {
			case fiber.MethodGet:
				params := make([]string, 0)
				for k, v := range ctx.Queries() {
					params = append(params, k+"="+v)
				}
				keyG += ":" + strings.Join(params, "&")
			case fiber.MethodPost:
				keyG += ":" + string(ctx.BodyRaw())
			}
			return utils.CopyString(helper.GenerateHash(keyG))
		},
		Expiration:   30 * time.Second,
		CacheControl: true, // enables client side caching if set to true
		Methods:      []string{fiber.MethodGet, fiber.MethodHead, fiber.MethodPost},
	})
}
