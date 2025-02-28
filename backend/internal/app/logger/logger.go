package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(env string) *zap.Logger {
	 var config zap.Config

	 if env == "production" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	 } else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	 }

	 logger, _ := config.Build()
	 return logger
}

// 全局 Logger 示例 (推荐通过依赖注入传递)
var Logger *zap.Logger

func InitGlobalLogger(env string) {
	Logger = NewLogger(env)
}