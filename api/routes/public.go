package routes

import (
	"github.com/shankar524/go-app-invite-service/api/controller"
	"github.com/shankar524/go-app-invite-service/lib"
	"github.com/shankar524/go-app-invite-service/middlewares"
)

type PublicRoutes struct {
	handler         lib.RequestHandler
	tokenController controller.ITokenController
	rateLimiter     middlewares.IRateLimiter
}

func (s PublicRoutes) Setup() {
	publicRoutes := s.handler.Gin.Group("/api/v1/public")
	publicRoutes.Use(s.rateLimiter.RateLimit)
	{
		publicRoutes.POST("/token/validate", s.tokenController.ValidateToken)
	}
}

func NewPublicRoutes(handler lib.RequestHandler, tokenController controller.TokenController, rateLimiter middlewares.APIRateLimiterMiddleware) PublicRoutes {
	return PublicRoutes{
		handler:         handler,
		tokenController: &tokenController,
		rateLimiter:     &rateLimiter,
	}
}
