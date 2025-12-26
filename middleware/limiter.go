package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type limiter struct {
	mu       sync.Mutex
	requests map[string]int
	resetAt  time.Time
	limit    int
	window   time.Duration
}

func NewLimiter(limit int, window time.Duration) gin.HandlerFunc {
	l := &limiter{
		requests: make(map[string]int),
		resetAt:  time.Now().Add(window),
		limit:    limit,
		window:   window,
	}

	return func(c *gin.Context) {
		l.mu.Lock()
		defer l.mu.Unlock()

		if time.Now().After(l.resetAt) {
			l.requests = make(map[string]int)
			l.resetAt = time.Now().Add(l.window)
		}

		ip := c.ClientIP()
		l.requests[ip]++

		if l.requests[ip] > l.limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
