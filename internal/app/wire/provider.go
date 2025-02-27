package wire

import (
	"GoWorld/internal/app/config"
	"GoWorld/internal/app/handlers"
	"GoWorld/internal/app/logger"
	"GoWorld/internal/app/repositories"
	"GoWorld/internal/app/services"

	"github.com/google/wire"
	"go.uber.org/zap"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 配置提供者
func ProvideConfig() (*config.Config, error) {
	config, err := config.LoadConfig("./configs")
	return &config, err
}

// 日志提供者
func ProvideLogger(cfg *config.Config) (*zap.Logger, error) {
	return logger.NewLogger(cfg.App.Env), nil
}

// 数据库提供者
func ProvideDB(cfg *config.Config) (*gorm.DB, error) {
	// db, err := gorm.Open(postgres.Open(cfg.Database.DSN))
	// // 其它配置
	// return db, err
	return nil, nil
}

// Repository 提供者
var RepositorySet = wire.NewSet(
	// wire.Bind(new(repositories.UserRepository), new(*repositories.UserRepo)),
	repositories.NewUserRepository,
)

// Service 提供者
var ServiceSet = wire.NewSet(
	services.NewUserService,
)

// Handler 提供者
var HandlerSet = wire.NewSet(
	handlers.NewUserHandler,
)

// 应用程序完整依赖集合
var SuperSet = wire.NewSet(
	ProvideConfig,
	ProvideLogger,
	ProvideDB,
	RepositorySet,
	ServiceSet,
	HandlerSet,
	// config.NewConfig,
	// logger.NewLogger,
)
