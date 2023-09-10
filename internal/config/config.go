package config

import (
	v1alpha1 "github.com/KommodoreX/dp-rudder/api/v1alpha1/config"
	"github.com/KommodoreX/dp-rudder/pkg/config"
	"github.com/KommodoreX/dp-rudder/pkg/logger"
)

func ReadConfigs(configFilePath string) (v1alpha1.Rudder, error) {
	cfg, err := config.ParseYAML(configFilePath)
	if err != nil {
		logger.LoggerRudder.Base().Error(err.Error())
		return v1alpha1.Rudder{}, err
	}
	logger.LoggerRudder.Base().Sugar().Info(cfg.AllSettings())
	var config v1alpha1.Rudder
	if err := cfg.Unmarshal(&config); err != nil {
		logger.LoggerRudder.Base().Error(err.Error())
		return v1alpha1.Rudder{}, err
	}
	logger.LoggerRudder.Sugar().Infof("Configuration loaded: %+v", config)
	return config, nil
}
