package logs

import "go.uber.org/fx"

var Module = fx.Module("logger",
	fx.Provide(NewLogger),
)
