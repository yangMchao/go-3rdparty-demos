package zap

import (
	"errors"
	"testing"

	"go.uber.org/zap"
)

func TestLogLevel(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Debug("Debugging information") // 在生产级别不会被记录
	logger.Info("Information message")    // 将被记录
	logger.Warn("Warning message")        // 将被记录
	logger.Error("Error message",         // 将被记录，并包含堆栈跟踪
		zap.Error(errors.New("sample error")),
	)
	// logger.Fatal("Fatal error") // 将记录并终止程序
}
