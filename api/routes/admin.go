package routes

import (
	"log"

	"github.com/shankar524/go-app-invite-service/api/controller"
	"github.com/shankar524/go-app-invite-service/api/repository"
	"github.com/shankar524/go-app-invite-service/cron"
	"github.com/shankar524/go-app-invite-service/lib"
	"github.com/shankar524/go-app-invite-service/middlewares"
)

type AdminRoutes struct {
	handler         lib.RequestHandler
	tokenController controller.TokenController
	tokenRepository repository.TokenRepository
	validator       middlewares.APIValidationMiddleware
	scheduler       cron.Cron
}

func (s AdminRoutes) Setup() {

	if err := s.tokenRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}

	s.scheduler.Job.Start()

	adminRoutes := s.handler.Gin.Group("/api/v1/admin")
	adminRoutes.Use(s.validator.Validate)
	{
		adminRoutes.GET("/token", s.tokenController.GetAll)
		adminRoutes.POST("/token", s.tokenController.Create)
		adminRoutes.GET("/token/:id", s.tokenController.GetByID)
		adminRoutes.PUT("/token/:id/disable", s.tokenController.DisableTokenByID)
	}
}

func NewAdminRoutes(handler lib.RequestHandler, tc controller.TokenController, tr repository.TokenRepository, validator middlewares.APIValidationMiddleware, scheduler cron.Cron) AdminRoutes {
	return AdminRoutes{
		handler:         handler,
		tokenController: tc,
		tokenRepository: tr,
		validator:       validator,
		scheduler:       scheduler,
	}
}
