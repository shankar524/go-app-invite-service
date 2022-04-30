package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() RequestHandler {
	engine := gin.New()
	engine.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API up and running"})
	})
	return RequestHandler{Gin: engine}
}
