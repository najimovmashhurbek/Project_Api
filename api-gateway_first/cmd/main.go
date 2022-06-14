package main

import (
	"github.com/najimovmashhurbek/project-api/api-gateway.ozim/api"
	"github.com/najimovmashhurbek/project-api/api-gateway.ozim/config"
	"github.com/najimovmashhurbek/project-api/api-gateway.ozim/pkg/logger"
	"github.com/najimovmashhurbek/project-api/api-gateway.ozim/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
