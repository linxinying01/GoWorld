package main

import (
	"GoWorld/internal/app/config"
	"GoWorld/internal/app/logger"
	"GoWorld/internal/app/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// 初始化日志（根据配置环境）
	logger.InitGlobalLogger(cfg.App.Env)

	// 注入 Logger 和 Config 到业务层（推荐使用依赖注入框架，如 Wire）
	userHandler := handlers.NewUserHandler(logger.Logger, &cfg)

	// 启动服务器（示例使用 Gin）
	r := gin.Default()
	r.GET("/users", userHandler.GetUsers)
	r.Run(cfg.Server.Port)
}