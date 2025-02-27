package main

import (
	"GoWorld/internal/app/handlers"
	"GoWorld/internal/app/middleware"
	"GoWorld/internal/app/wire"

	"github.com/gin-gonic/gin"
)

func setupRouter(
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	jwtMiddleware *middleware.JWTMiddleware,
) *gin.Engine {
	r := gin.Default()

	// 认证路由组
	authGroup := r.Group("auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// 需要认证的路由组
	apiGroup := r.Group("/api")
	apiGroup.Use(jwtMiddleware.Handler()) // JWT中间件需要实现
	{
		apiGroup.GET("/users", userHandler.GetUsers)
	}

	return r
}

func main() {
	// 通过 Wire 初始化所有配置
	authHandler, err := wire.InitializeAuth()
	if err != nil {
		panic("依赖注入失败1：" + err.Error())
	}

	userHandler, err := wire.InitializeUser()
	if err != nil {
		panic("依赖注入失败2：" + err.Error())
	}

	jwtMiddleware, err := wire.InitializeJwt()
	if err != nil {
		panic("依赖注入失败3：" + err.Error())
	}

	// 初始化路由
	r := setupRouter(authHandler, userHandler, jwtMiddleware)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		panic("服务器启动失败：" + err.Error())
	}
}
