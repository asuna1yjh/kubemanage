package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cross() func(c *gin.Context) {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost", "http://127.0.0.1"},    // 允许跨域发来请求的网站
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool { // 自定义过滤源站的方法
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		})
	}
}
