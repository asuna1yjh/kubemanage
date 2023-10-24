package routes

import (
	"gin_demo/cmd/option"
	"gin_demo/logger"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.RouterGroup)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

func Setup(opt *option.Option, opts ...Option) *gin.Engine {
	if opt.Config.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置gin的模式为发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), gin.Recovery())
	api := r.Group("/api")
	Include(opts...)
	for _, opt := range options {
		opt(api)
	}
	return r
}
