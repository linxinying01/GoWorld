// internal/app/models/user.go
package models

import (
    "time"
    "gorm.io/gorm"
)

// User 系统用户模型
type User struct {
    gorm.Model           // 内嵌 gorm.Model 包含 ID, CreatedAt, UpdatedAt, DeletedAt
    
    // 身份验证信息
    Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"` // 唯一邮箱
    Password  string    `gorm:"type:varchar(100);not null" json:"-"`               // 密码（不序列化到JSON）
    
    // 用户基本信息
    FirstName string    `gorm:"type:varchar(50)" json:"first_name"`                // 名字
    LastName  string    `gorm:"type:varchar(50)" json:"last_name"`                 // 姓氏
    AvatarURL string    `gorm:"type:varchar(255)" json:"avatar_url"`               // 头像地址
    
    // 状态信息
    IsActive  bool      `gorm:"default:true" json:"is_active"`                     // 是否激活
    LastLogin time.Time `json:"last_login"`                                        // 最后登录时间
    LastIP    string    `gorm:"type:varchar(45)" json:"last_ip"`                   // 最后登录IP
    
    // 权限控制（示例）
    Roles     string    `gorm:"type:varchar(255);default:'user'" json:"roles"`     // 用户角色
    
    // 时间戳（gorm.Model 已包含，这里为自定义扩展示例）
    VerifiedAt *time.Time `json:"verified_at"`                                     // 邮箱验证时间
}

// TableName 自定义表名（可选）
func (User) TableName() string {
    return "users"
}