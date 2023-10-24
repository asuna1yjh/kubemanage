package kubecontroller

import (
	"gin_demo/routes"

	"github.com/gin-gonic/gin"
)

func InitKubeRouter(deployment *Deployment) routes.Option {
	return func(r *gin.RouterGroup) {
		user := r.Group("/k8s")
		{
			user.GET("/deployment/list", deployment.GetList)
		}
	}
}
