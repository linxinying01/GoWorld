package database

import (
	"GoWorld/backend/internal/app/config"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化 MySQL 连接
func NewMySQLDB(cfg *config.Config, zapLogger *zap.Logger) (*gorm.DB, error) {
	// 初始化 GORM 配置
	gormConfig := &gorm.Config{
		Logger: NewGormLogger(zapLogger),
	}

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(cfg.Database.MySQL.DSN), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// 获取通用数据库对象
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}

	// 连接池配置
	sqlDB.SetMaxOpenConns(cfg.Database.MySQL.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MySQL.MaxIdleConns)

	if lifetime, err := time.ParseDuration(cfg.Database.MySQL.ConnMaxLifetime); err == nil {
		sqlDB.SetConnMaxLifetime(lifetime)
	}

	if idleTime, err := time.ParseDuration(cfg.Database.MySQL.ConnMaxIdleTime); err == nil {
		sqlDB.SetConnMaxIdleTime(idleTime)
	}

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	// 数据库迁移
	if err := AutoMigrate(db); err != nil {
		panic("数据库迁移失败：" + err.Error())
	}

	return db, nil
}

// GORM 日志适配器
func NewGormLogger(zapLogger *zap.Logger) logger.Interface {
	return logger.New(
		&gormLogWriter{zapLogger: zapLogger},
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
}

type gormLogWriter struct {
	zapLogger *zap.Logger
}

func (w *gormLogWriter) Printf(format string, args ...interface{}) {
	w.zapLogger.Sugar().Debugf(format, args...)
}
