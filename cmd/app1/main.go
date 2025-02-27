package main

import (
	"GoWorld/internal/app/wire"

	"github.com/gin-gonic/gin"
)

func main() {
	// 通过 Wire 初始化所有配置
	userHandler, err := wire.InitializeApp()
	if err != nil {
		panic("依赖注入失败：" + err.Error())
	}

	// 初始化路由
	r := gin.Default()
	r.GET("/users", userHandler.GetUsers)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		panic("服务器启动失败：" + err.Error())
	}
}