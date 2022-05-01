package middlewares

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/shankar524/go-app-invite-service/lib"

	"go.uber.org/ratelimit"
)

type APIRateLimiterMiddleware struct {
	Rate int
}

func (a APIRateLimiterMiddleware) RateLimit(c *gin.Context) {
	var clientRecord sync.Map
	client := c.ClientIP()
	lif, ok := clientRecord.Load(client)
	if !ok {
		lif = ratelimit.New(a.Rate)
	}

	lm, ok := lif.(ratelimit.Limiter)
	if !ok {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "rate limit exceeded"})
		return
	}
	lm.Take()
	clientRecord.Store(client, lm)
}

func NewAPIRateLimiterMiddleware(env lib.Env) APIRateLimiterMiddleware {
	return APIRateLimiterMiddleware{env.ApiRate}
}
