package zap

import (
	"errors"
	"testing"

	"go.uber.org/zap"
)

func TestWith(t *testing.T) {
	logger, _ := zap.NewProduction()

	logger.Error("出现问题",
		zap.Error(errors.New("x/0")),
		zap.Stack("stacktrace"),
	)

	// 跳过一些帧
	logger.Error("出现问题",
		zap.StackSkip("stacktrace", 2),
	)
}
