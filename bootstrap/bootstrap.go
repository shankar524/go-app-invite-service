package bootstrap

import (
	"context"

	"github.com/shankar524/go-app-invite-service/api/controller"
	"github.com/shankar524/go-app-invite-service/api/repository"
	"github.com/shankar524/go-app-invite-service/api/routes"
	"github.com/shankar524/go-app-invite-service/api/service"
	"github.com/shankar524/go-app-invite-service/lib"
	"github.com/shankar524/go-app-invite-service/middlewares"

	"go.uber.org/fx"
)

var Module = fx.Options(
	lib.Module,
	middlewares.Module,
	controller.Module,
	routes.Module,
	service.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	database lib.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			conn.SetMaxOpenConns(10)
			go func() {
				routes.Setup()
				handler.Gin.Run(":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			conn.Close()
			return nil
		},
	})
}
