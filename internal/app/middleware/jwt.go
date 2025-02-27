package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"GoWorld/internal/app/config"
	"GoWorld/internal/app/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

const (
	UserIDKey    = "userID"
	UserEmailKey = "userEmail"
)

// JWTMiddleware 封装中间件需要的依赖
type JWTMiddleware struct {
	jwtSecret string
	logger    *zap.Logger
}

// NewJWTMiddleware 构造函数（通过Wire注入）
func NewJWTMiddleware(config *config.Config, logger *zap.Logger) *JWTMiddleware {
	return &JWTMiddleware{
		jwtSecret: config.Security.JWTSecret,
		logger:    logger.With(zap.String("module", "jwt_middleware")),
	}
}

// Handler 中间件处理器
func (m *JWTMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractToken(c)
		if err != nil {
			m.sendError(c, http.StatusUnauthorized, err)
			return
		}

		claims, err := m.validateToken(tokenString)
		if err != nil {
			m.sendError(c, http.StatusUnauthorized, err)
			return
		}

		m.setUserContext(c, claims)
		c.Next()
	}
}

// extractToken 从请求中提取JWT令牌
func extractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("缺少认证令牌")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("无效的令牌格式")
	}

	return parts[1], nil
}

// validateToken 验证JWT令牌有效性
func (m *JWTMiddleware) validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return []byte(m.jwtSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("令牌验证失败: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的令牌声明")
}

// setUserContext 将用户信息存入上下文
func (m *JWTMiddleware) setUserContext(c *gin.Context, claims jwt.MapClaims) {
	if userID, ok := claims["sub"].(float64); ok {
		c.Set(UserIDKey, uint(userID))
	}

	if email, ok := claims["email"].(string); ok {
		c.Set(UserEmailKey, email)
	}
}

// sendError 统一错误响应
func (m *JWTMiddleware) sendError(c *gin.Context, code int, err error) {
	m.logger.Warn("JWT认证失败",
		zap.Error(err),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
	)

	c.AbortWithStatusJSON(code, models.ErrorResponse{
		Error: err.Error(),
	})
}
