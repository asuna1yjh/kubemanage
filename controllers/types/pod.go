package types

type GetPodLogRequest struct {
	Namespace     string `json:"namespace" binding:"required" form:"namespace"`
	PodName       string `json:"pod_name" binding:"required" form:"pod_name"`
	ContainerName string `json:"container_name" binding:"required" form:"container_name"`
}

// GetPodDetailRequest 查询pod详情
type GetPodDetailRequest struct {
	Namespace string `json:"namespace" binding:"required" form:"namespace"`
	PodName   string `json:"pod_name" binding:"required" form:"pod_name"`
}

// GetListPodsRequest 获取pod列表的请求参数
type GetListPodsRequest struct {
	Namespace string `json:"namespace" binding:"required" form:"namespace"`
	Size      int    `json:"size" binding:"required,gt=1" form:"size"`
	Page      int    `json:"page" binding:"required,gt=0" form:"page"`
}

// GetPodRequest 查询pod详情
type GetPodRequest struct {
	Namespace     string `json:"namespace" binding:"required" form:"namespace"`
	PodName       string `json:"pod_name" binding:"required" form:"pod_name"`
	ContainerName string `json:"container_name" binding:"required" form:"container_name"`
}

// GetPodsRequest 查询pod列表
type CountPodRequest struct {
	Namespace string `json:"namespace" binding:"required" form:"namespace"`
}
