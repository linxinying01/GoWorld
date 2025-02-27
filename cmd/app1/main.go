package main

import (
	"GoWorld/internal/app/wire"

	"github.com/gin-gonic/gin"
)

func setupRouter(appHandlerSet *wire.App) *gin.Engine {
	r := gin.Default()

	// 认证路由组
	authGroup := r.Group("auth")
	{
		authGroup.POST("/register", appHandlerSet.AuthHandler.Register)
		authGroup.POST("/login", appHandlerSet.AuthHandler.Login)
	}

	// 需要认证的路由组
	apiGroup := r.Group("/api")
	apiGroup.Use(appHandlerSet.JWTMiddleware.Handler()) // JWT中间件需要实现
	{
		apiGroup.GET("/users", appHandlerSet.UserHandler.GetUsers)
	}

	return r
}

func main() {
	// 通过 Wire 初始化所有配置
	appHandlerSet, err := wire.InitializeApp()
	if err != nil {
		panic("依赖注入失败：" + err.Error())
	}

	// 初始化路由
	r := setupRouter(appHandlerSet)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		panic("服务器启动失败：" + err.Error())
	}
}
