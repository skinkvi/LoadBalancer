package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	var err error
	Log, err = zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       true,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableCaller:     true,
		DisableStacktrace: true,
	}.Build()
	if err != nil {
		panic(err)
	}
}
