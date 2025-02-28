package handlers

import (
	"GoWorld/backend/internal/app/config"
	"GoWorld/backend/internal/app/middleware"
	"GoWorld/backend/internal/app/models"
	"GoWorld/backend/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger      *zap.Logger
	config      *config.Config
	userService *services.UserService
}

func NewUserHandler(logger *zap.Logger, cfg *config.Config, userService *services.UserService) *UserHandler {
	return &UserHandler{
		logger:      logger,
		config:      cfg,
		userService: userService,
	}
}

// Response 定义通用的响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	h.logger.Info("Handling GetUser request",
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
	)

	response := Response{
		Code: 200,
		Msg:  "success",
		Data: struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{
			Name: "test user",
			Age:  25,
		},
	}

	c.JSON(200, response)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get(middleware.UserIDKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "用户未认证"})
		return
	}

	// 类型断言确保安全
	id, ok := userID.(uint)
	if !ok {
		h.logger.Error("无效的用户ID类型")
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "服务器错误"})
		return
	}

	h.logger.Debug("获取用户信息", zap.Uint("user_id", id))

	// 调用服务层获取用户信息
	// h.userService.GetUserByID(strconv.Itoa(id))

	// ... 处理逻辑

	// h.logger.Info("Handling GetProfile request", zap.String("user_id", fmt.Sprintf("%s", user)))
}
