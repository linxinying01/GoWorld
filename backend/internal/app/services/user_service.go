package services

import (
	"GoWorld/backend/internal/app/repositories"

	"go.uber.org/zap"
)

type UserService struct {
	logger *zap.Logger
	repo   repositories.UserRepository
}

// 通过 Wire 自动注入依赖
func NewUserService(repo repositories.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

// func (s *UserService) GetUser(id string) {
// 	s.logger.Info("查询用户", zap.String("user_id", id))

// 	// 调用 repository
// }

func (s *UserService) GetUserByID(id string) {
	s.logger.Info("查询用户", zap.String("user_id", id))

	// 调用 repository
}
