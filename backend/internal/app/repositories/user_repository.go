// internal/app/repositories/user_repository.go
package repositories

import (
	"GoWorld/backend/internal/app/models"
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 接口定义（面向接口编程，方便测试和替换实现）
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, page, pageSize int) ([]*models.User, error)
}

// 具体实现（GORM 版本）
type UserRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

// 确保实现接口
// var _ UserRepository = (*UserRepo)(nil)

func NewUserRepository(db *gorm.DB, logger *zap.Logger) UserRepository {
	return &UserRepo{
		db:     db,
		logger: logger.With(zap.String("module", "user_repository")),
	}
}

// CreateUser 创建用户
func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		r.logger.Error("创建用户失败",
			zap.Error(err),
			zap.Any("user", user),
		)
		return wrapDBError(err, "创建用户失败")
	}
	return nil
}

// GetUserByID 通过ID获取用户
func (r *UserRepo) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		r.logger.Error("查询用户失败",
			zap.Error(err),
			zap.Uint("user_id", id),
		)
		return nil, wrapDBError(err, "查询用户失败")
	}
	return &user, nil
}

// GetUserByEmail 通过邮箱获取用户
func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		r.logger.Error("通过邮箱查询用户失败",
			zap.Error(err),
			zap.String("email", email),
		)
		return nil, wrapDBError(err, "通过邮箱查询用户失败")
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (r *UserRepo) UpdateUser(ctx context.Context, user *models.User) error {
	result := r.db.WithContext(ctx).
		Model(user).
		Updates(user)
	if result.Error != nil {
		r.logger.Error("更新用户失败",
			zap.Error(result.Error),
			zap.Any("user", user),
		)
		return wrapDBError(result.Error, "更新用户失败")
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// DeleteUser 删除用户（软删除）
func (r *UserRepo) DeleteUser(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Delete(&models.User{})
	if result.Error != nil {
		r.logger.Error("删除用户失败",
			zap.Error(result.Error),
			zap.Uint("user_id", id),
		)
		return wrapDBError(result.Error, "删除用户失败")
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// ListUsers 分页列出用户
func (r *UserRepo) ListUsers(ctx context.Context, page, pageSize int) ([]*models.User, error) {
	var users []*models.User
	offset := (page - 1) * pageSize

	if err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		r.logger.Error("获取用户列表失败",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
		)
		return nil, wrapDBError(err, "获取用户列表失败")
	}
	return users, nil
}

// 错误处理辅助函数
func wrapDBError(err error, message string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return errors.New(message + ": " + err.Error())
}

// 定义公共错误
var (
	ErrNotFound = errors.New("记录不存在")
)
