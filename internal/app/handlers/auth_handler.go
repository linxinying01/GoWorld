package handlers

import (
	"GoWorld/internal/app/models"
	"GoWorld/internal/app/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authService services.AuthService
	logger      *zap.Logger
}

func NewAuthHandler(authService services.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

// Register 用户注册
// @Summary 用户注册
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "注册信息"
// @Success 201 {object} models.AuthResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("无效的注册请求", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "无效的请求格式"})
		return
	}

	authRes, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error("注册失败", zap.Error(err))
		c.JSON(errorToStatusCode(err), models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, authRes)
}

// Login 用户登录
// @Summary 用户登录
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "登录信息"
// @Success 200 {object} models.AuthResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("无效的登录请求", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "无效的请求格式"})
		return
	}

	authRes, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error("登录失败", zap.Error(err))
		c.JSON(errorToStatusCode(err), models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, authRes)
}

// errorToStatusCode 将业务错误转换为HTTP状态码
func errorToStatusCode(err error) int {
	switch {
	case errors.Is(err, services.ErrUserExists):
		return http.StatusConflict
	case errors.Is(err, services.ErrInvalidCredentials):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
