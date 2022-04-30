package routes

import (
	"github.com/shankar524/go-app-invite-service/api/controller"
	"github.com/shankar524/go-app-invite-service/lib"
)

type PublicRoutes struct {
	handler         lib.RequestHandler
	tokenController controller.TokenController
}

func (s PublicRoutes) Setup() {
	publicRoutes := s.handler.Gin.Group("/api/v1/public")
	{
		publicRoutes.POST("/token/validate", s.tokenController.ValidateToken)
	}
}

func NewPublicRoutes(handler lib.RequestHandler, tokenController controller.TokenController) PublicRoutes {
	return PublicRoutes{
		handler:         handler,
		tokenController: tokenController,
	}
}
