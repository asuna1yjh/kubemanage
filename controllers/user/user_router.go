package user

import (
	"gin_demo/logic"
	"gin_demo/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	us logic.UserInterface
}

func NewUserRouter(us logic.UserInterface) func(r *gin.RouterGroup) {
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
