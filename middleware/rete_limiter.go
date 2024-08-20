package middleware

import (
	"93HW/config"
	"93HW/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.ClientIP()

		rateLimitKey := "rate_limit" + userID

		count, err := config.RedisClient.Get(utils.Ctx, rateLimitKey).Int()
		if err != nil && err.Error() != "redis: nil" {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if count >= 10 {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		config.RedisClient.Incr(utils.Ctx, rateLimitKey)
		config.RedisClient.Expire(utils.Ctx, rateLimitKey, 60)
		c.Next()
	}
}
