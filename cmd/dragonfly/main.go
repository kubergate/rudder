package main

import (
	"github.com/NomadXD/dragonfly/internal/dragonfly"
	"github.com/NomadXD/dragonfly/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.LoggerDragonFly.Sync()
	logger.LoggerDragonFly.Info("Starting dragonfly....")
	dragonfly.Init()
}