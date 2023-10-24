package user

import (
	"gin_demo/logic"
	"gin_demo/middlewares"
	"gin_demo/routes"

	"github.com/google/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	us *logic.UserUseCase
}

func NewUserRouter(us *logic.UserUseCase) routes.Option {
	ur := &UserRouter{
		us: us,
	}
	return func(r *gin.RouterGroup) {
		r.POST("/register", ur.RegisterHandler)
		r.POST("/login", ur.LoginHandler)
		user := r.Group("/user").Use(middlewares.JWTAuthMiddleware())
		{
			user.GET("/info", ur.InfoHandler)
		}
	}
}

var ProviderSet = wire.NewSet(NewUserRouter)
