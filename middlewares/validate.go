package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shankar524/go-app-invite-service/lib"
)

type APIValidationMiddleware struct {
	HeaderField string
	env         lib.Env
}

type IValidator interface {
	Validate(c *gin.Context)
}

func (a APIValidationMiddleware) Validate(c *gin.Context) {
	providedKey := c.Request.Header.Get(a.HeaderField)
	if providedKey != a.env.ApiKey {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid api key"})
		return
	}
}

func NewAPIValidationMiddleware(env lib.Env) IValidator {
	return &APIValidationMiddleware{HeaderField: "api-key", env: env}
}
