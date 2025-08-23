package middleware

import (
	"sync"
	
	"bosh-admin/core/ctx"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
}

func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}

func RateLimiter(rps int, burst int) gin.HandlerFunc {
	rateLimiter := NewIPRateLimiter(rate.Limit(rps), burst)
	return ctx.Handler(func(c *ctx.Context) {
		limiter := rateLimiter.GetLimiter(c.ClientIP())
		if !limiter.Allow() {
			c.TooManyRequests()
			c.Abort()
			return
		}
		c.Next()
	})
}
