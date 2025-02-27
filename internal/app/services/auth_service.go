package services

import (
	"context"
	"errors"
	"time"

	"GoWorld/internal/app/config"
	"GoWorld/internal/app/models"
	"GoWorld/internal/app/repositories"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.AuthResponse, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.AuthResponse, error)
}

type AuthServiceImpl struct {
	userRepo  repositories.UserRepository
	logger    *zap.Logger
	jwtSecret string
}

func NewAuthService(
	userRepo repositories.UserRepository,
	logger *zap.Logger,
	cfg *config.Config,
) AuthService {
	return &AuthServiceImpl{
		userRepo:  userRepo,
		logger:    logger,
		jwtSecret: cfg.Security.JWTSecret,
	}
}

var (
	ErrUserExists         = errors.New("邮箱已被注册")
	ErrInvalidCredentials = errors.New("无效的登录凭证")
)

// Register 实现注册逻辑
func (s *AuthServiceImpl) Register(ctx context.Context, req *models.RegisterRequest) (*models.AuthResponse, error) {
	// 检查邮箱是否已存在
	existing, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
	if existing != nil {
		return nil, ErrUserExists
	}

	// 密码哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("密码哈希失败", zap.Error(err))
		return nil, errors.New("注册失败")
	}

	// 创建用户对象
	user := &models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  true,
	}

	// 保存到数据库
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		s.logger.Error("创建用户失败", zap.Error(err))
		return nil, errors.New("注册失败")
	}

	// 生成JWT
	token, err := s.generateJWT(user)
	if err != nil {
		s.logger.Error("JWT生成失败", zap.Error(err))
		return nil, errors.New("注册失败")
	}

	return &models.AuthResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}, nil
}

// Login 实现登录逻辑
func (s *AuthServiceImpl) Login(ctx context.Context, req *models.LoginRequest) (*models.AuthResponse, error) {
	// 查找用户
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.logger.Warn("用户不存在", zap.String("email", req.Email))
		return nil, ErrInvalidCredentials
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		s.logger.Warn("密码验证失败", zap.String("email", req.Email))
		return nil, ErrInvalidCredentials
	}

	// 生成JWT
	token, err := s.generateJWT(user)
	if err != nil {
		s.logger.Error("JWT生成失败", zap.Error(err))
		return nil, errors.New("登录失败")
	}

	return &models.AuthResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}, nil
}

// generateJWT 生成JWT令牌
func (s *AuthServiceImpl) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
