package kubecontroller

import (
	"gin_demo/routes"

	"github.com/gin-gonic/gin"
)

func InitKubeRouter(deployment *Deployment) routes.Option {
	return func(r *gin.RouterGroup) {
		k8sRouter := r.Group("/k8s")
		{
			k8sRouter.POST("/deployment/create", deployment.CreateDeployment)
			k8sRouter.GET("/deployment/list", deployment.GetList)
			k8sRouter.DELETE("/deployment/delete", deployment.DeleteDeployment)
			k8sRouter.GET("/deployment/detail", deployment.DetailDeployment)
			k8sRouter.PUT("/deployment/update", deployment.UpdateDeployment)
			k8sRouter.POST("/deployment/restart", deployment.RestartDeployment)
			k8sRouter.PUT("/deployment/scale", deployment.ScaleDeployment)
			k8sRouter.GET("/deployment/count", deployment.CountDeployment)
		}
	}
}
