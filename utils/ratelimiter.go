package utils

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	rateLimit  = 30            // max requests per window per IP
	rateWindow = time.Minute  // sliding window size
	cleanupTTL = 10 * time.Minute
)

type ipEntry struct {
	mu       sync.Mutex
	count    int
	windowAt time.Time
	lastSeen time.Time
}

var (
	ipTable     sync.Map
	cleanupOnce sync.Once
)

func getEntry(ip string) *ipEntry {
	v, loaded := ipTable.LoadOrStore(ip, &ipEntry{windowAt: time.Now(), lastSeen: time.Now()})
	if loaded {
		e := v.(*ipEntry)
		e.lastSeen = time.Now()
		return e
	}
	return v.(*ipEntry)
}

func startCleanup() {
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			cutoff := time.Now().Add(-cleanupTTL)
			ipTable.Range(func(key, val any) bool {
				e := val.(*ipEntry)
				e.mu.Lock()
				stale := e.lastSeen.Before(cutoff)
				e.mu.Unlock()
				if stale {
					ipTable.Delete(key)
				}
				return true
			})
		}
	}()
}

func RateLimitMiddleware() gin.HandlerFunc {
	cleanupOnce.Do(startCleanup)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		e := getEntry(ip)

		e.mu.Lock()
		now := time.Now()
		if now.Sub(e.windowAt) >= rateWindow {
			e.count = 0
			e.windowAt = now
		}
		e.count++
		allowed := e.count <= rateLimit
		e.mu.Unlock()

		if !allowed {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"status": "Failed",
				"desc":   "too many requests — please slow down",
			})
			return
		}
		c.Next()
	}
}
