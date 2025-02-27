//go:build wireinject
// +build wireinject

package wire

import (
	"GoWorld/internal/app/handlers"

	"github.com/google/wire"
)

// 初始化应用程序依赖（入口函数）
func InitializeApp() (*handlers.UserHandler, error) {
	wire.Build(SuperSet)
	return &handlers.UserHandler{}, nil
}
