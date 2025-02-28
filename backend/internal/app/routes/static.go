package routes

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func RegisterStatic(r *gin.Engine) {
	// 开发环境：禁用静态缓存
	if gin.Mode() == gin.DebugMode {
		r.NoRoute(func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173"+c.Request.URL.Path)
		})
		return
	}

	// 生产环境：托管静态文件
	staticDir := filepath.Join("web", "dist")

	// 静态资源
	r.Static("/assets", filepath.Join(staticDir, "assets"))
	r.StaticFile("/favicon.ico", filepath.Join(staticDir, "favicon.ico"))

	// 前端路由处理
	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(staticDir, "index.html"))
	})
}
