package main

import (
	"github.com/kubergate/rudder/internal/config"
	rudder "github.com/kubergate/rudder/internal/rudder"
	"github.com/kubergate/rudder/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.LoggerRudder.Base().Sync()
	defer logger.LoggerRudder.Sugar().Sync()

	config, err := config.ReadConfigs("deployments/resources/config/rudder.yaml")
	if err != nil {
		logger.LoggerRudder.Base().Error(err.Error())
	}
	logger.LoggerRudder.Base().Info("Starting rudder....")
	rudder.Init(config)
}
