package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	// Вывод одновременно в консоль и в файл app.log
	cfg.OutputPaths = []string{"stdout", "app.log"}
	cfg.ErrorOutputPaths = []string{"stderr"}

	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	return cfg.Build()
}