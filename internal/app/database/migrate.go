// internal/app/database/migrate.go
package database

import (
    "gorm.io/gorm"
    "GoWorld/internal/app/models"
)

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &models.User{},
        // 添加其他模型...
    )
}