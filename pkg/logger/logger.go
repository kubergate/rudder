package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// type Logger interface {
//     *zap.Logger
//     Sugar() *zap.SugaredLogger
// }

type Logger struct {
	baseLogger *zap.Logger
}

var LoggerDragonFly Logger

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
	LoggerDragonFly.baseLogger = zap.Must(config.Build())
}

func (l *Logger) Sugar() *zap.SugaredLogger {
	return l.Sugar()
}

func (l *Logger) Base() *zap.Logger {
	return l.baseLogger
}
