package container

import (
	"Crash-Billing-service/internal/config"
	repo "Crash-Billing-service/internal/repository/pgsql"
	"Crash-Billing-service/internal/service"
	"Crash-Billing-service/internal/transport/rpc/handlers"
	"Crash-Billing-service/internal/transport/server"
	"Crash-Billing-service/pkg/db"
	"Crash-Billing-service/pkg/logger"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		config.Module,
		db.Module,
		logger.Module,
		repo.Module,
		service.Module,
		handlers.Module,
		server.Module,
	)
}
