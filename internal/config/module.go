package config

import (
	"Crash-Billing-service/internal/transport/server"
	"Crash-Billing-service/pkg/db"
	"Crash-Billing-service/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(
		LoadConfig,
		func(cfg *Config) *db.DBConfig { return cfg.DBConfig },
		func(cfg *Config) *server.GrpcConfig { return cfg.GrpcConfig },
		func(cfg *Config) *logger.Config { return cfg.LoggerConfig },
	),
)
