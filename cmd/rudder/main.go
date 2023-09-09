package main

import (
	rudder "github.com/KommodoreX/dp-rudder/internal/rudder"
	"github.com/KommodoreX/dp-rudder/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.LoggerDragonFly.Base().Sync()
	logger.LoggerDragonFly.Base().Info("Starting dragonfly....")
	rudder.Init()
}
