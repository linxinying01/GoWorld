package handlers

import (
	"GoWorld/internal/app/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger *zap.Logger
	config *config.Config
}

func NewUserHandler(logger *zap.Logger, cfg *config.Config) *UserHandler {
	return &UserHandler{
		logger: logger,
		config: cfg,
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
