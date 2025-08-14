package zap

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestNewProductionBySugar(t *testing.T) {
	// 创建一个日志记录器（生产配置）
	logger, _ := zap.NewProduction()
	defer logger.Sync() // 刷新缓冲区，如果有

	// 获取一个 SugaredLogger
	sugar := logger.Sugar()

	// 示例 1：使用松散类型的键值对进行结构化日志记录
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	// 示例 2：Printf 风格的日志记录
	sugar.Infof("Failed to fetch URL: %s", "http://example.com")
}

func TestNewProductionByLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 使用强类型的 Field 值进行结构化日志记录
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestNewDevelopment(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		// 处理错误
	}
	defer logger.Sync()
	logger.Debug("Development logging initialized")
}
