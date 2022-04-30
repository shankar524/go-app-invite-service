package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewRequestHandler),
	fx.Provide(NewEnv),
	fx.Provide(NewTextGenerator),
	fx.Provide(NewCache),
)
