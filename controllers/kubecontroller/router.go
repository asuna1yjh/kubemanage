package kubecontroller

import (
	"gin_demo/routes"

	"github.com/gin-gonic/gin"
)

func InitKubeRouter(deployment *Deployment, pod *Pod) routes.Option {
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
		// 管理所有pod的请求
		{
			k8sRouter.GET("/pod/list", pod.GetListPods)
			k8sRouter.GET("/pod/logs", pod.GetPodLog)
			k8sRouter.GET("/pod/detail", pod.GetPodDetail)
			k8sRouter.DELETE("/pod/delete", pod.DeletePod)
			k8sRouter.GET("/pod/webshell", pod.WebShell)
			k8sRouter.GET("/pod/count", pod.CountPod)
		}
	}
}
