//go:build wireinject
// +build wireinject

package wire

import (
	"GoWorld/internal/app/handlers"
	"GoWorld/internal/app/middleware"

	"github.com/google/wire"
)

// 初始化应用程序依赖（入口函数）
func InitializeAuth() (*handlers.AuthHandler, error) {
	wire.Build(SuperSet)
	return &handlers.AuthHandler{}, nil
}

func InitializeUser() (*handlers.UserHandler, error) {
	wire.Build(SuperSet)
	return &handlers.UserHandler{}, nil
}

func InitializeJwt() (*middleware.JWTMiddleware, error) {
	wire.Build(SuperSet)
	return &middleware.JWTMiddleware{}, nil
}