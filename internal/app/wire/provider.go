package wire

import (
	"GoWorld/internal/app/config"
	"GoWorld/internal/app/database"
	"GoWorld/internal/app/handlers"
	"GoWorld/internal/app/logger"
	"GoWorld/internal/app/middleware"
	"GoWorld/internal/app/repositories"
	"GoWorld/internal/app/services"

	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	AuthHandler   *handlers.AuthHandler
	UserHandler   *handlers.UserHandler
	JWTMiddleware *middleware.JWTMiddleware
}

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
func ProvideDB(cfg *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	return database.NewMySQLDB(cfg, logger)
}

// 依赖集合分组
var (
	// 基础设施层
	InfraSet = wire.NewSet(
		ProvideConfig,
		ProvideLogger,
		ProvideDB,
	)

	// Repository层
	RepositorySet = wire.NewSet(
		// wire.Bind(new(repositories.UserRepository), new(*repositories.UserRepo)),
		repositories.NewUserRepository,
	)

	// Service 提供者
	ServiceSet = wire.NewSet(
		// wire.Bind(new(services.AuthService), new(*services.AuthServiceImpl)),
		services.NewAuthService,
		services.NewUserService,
	)

	// Handler 提供者
	HandlerSet = wire.NewSet(
		handlers.NewAuthHandler,
		handlers.NewUserHandler,
		middleware.NewJWTMiddleware,

		wire.Struct(new(App), "*"),
	)
)

// 完整依赖集合
var SuperSet = wire.NewSet(
	InfraSet,
	RepositorySet,
	ServiceSet,
	HandlerSet,
)
