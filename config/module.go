// lib/config/fx.go
package config

import "go.uber.org/fx"

var Module = fx.Module("config", fx.Provide(LoadConfigFromYAML))
