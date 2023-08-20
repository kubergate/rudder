package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LoggerDragonFly *zap.Logger

func InitLogger() {
	encoderCfg := zap.NewProductionEncoderConfig()
    encoderCfg.TimeKey = "timestamp"
    encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

    config := zap.Config{
        Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
        Development:       false,
        DisableCaller:     false,
        DisableStacktrace: false,
        Sampling:          nil,
        Encoding:          "json",
        EncoderConfig:     encoderCfg,
        OutputPaths: []string{
            "stderr",
        },
        ErrorOutputPaths: []string{
            "stderr",
        },
        InitialFields: map[string]interface{}{
            "pid": os.Getpid(),
        },
    }
	LoggerDragonFly = zap.Must(config.Build())
}