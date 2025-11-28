package cache

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type Item struct {
	data      []byte
	timestamp time.Time
}

type CacheMiddleware struct {
	cache map[string]*Item
	mu    sync.Mutex
	ttl   time.Duration
}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func NewCacheMiddleware(ttl time.Duration) *CacheMiddleware {
	return &CacheMiddleware{cache: make(map[string]*Item), ttl: ttl}
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (m *CacheMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.RequestURI()

		m.mu.Lock()
		item, exists := m.cache[key]
		m.mu.Unlock()
		if exists && time.Since(item.timestamp) < m.ttl {
			c.Data(http.StatusOK, "application/json", item.data)
			c.Abort()
			return
		}
		writer := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		c.Next()
		m.mu.Lock()
		m.cache[key] = &Item{
			data:      writer.body.Bytes(),
			timestamp: time.Now().UTC(),
		}
		m.mu.Unlock()
	}

}
