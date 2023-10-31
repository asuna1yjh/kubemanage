package types

type GetDeploymentListRequest struct {
	Namespace string `json:"namespace" binding:"required" form:"namespace"`
}

type GetDeploymentListResponse struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// 创建deployment的请求参数
type CreateDeploymentRequest struct {
	Name          string            `json:"name" binding:"required"`
	Namespace     string            `json:"namespace" binding:"required" `
	Replicas      int32             `json:"replicas" binding:"required"`
	Image         string            `json:"image" binding:"required"`
	Labels        map[string]string `json:"label" binding:"required" comment:"标签"`
	Cpu           string            `json:"cpu"  comment:"Cpu限制"`
	Memory        string            `json:"memory"  comment:"内存限制"`
	ContainerPort int32             `json:"container_port" binding:"required"  comment:"容器端口"`
	HealthCheck   bool              `json:"health_check"   comment:"健康检查开关"`
	HealthPath    string            `json:"health_path"  comment:"Http健康检查路径"`
}

// 删除deployment的请求参数
type DeleteDeploymentRequest struct {
	Name      string `json:"name" binding:"required"`
	Namespace string `json:"namespace" binding:"required"`
}

// 获取deployment详情的请求参数
type DetailDeploymentRequest struct {
	Name      string `json:"name" binding:"required" form:"name"`
	Namespace string `json:"namespace" binding:"required" form:"namespace"`
}

// UpdateDeploymentRequest 更新deployment的请求参数
type UpdateDeploymentRequest struct {
	Namespace string `json:"namespace" binding:"required" comment:"命名空间" form:"namespace"`
	Content   string `json:"content" binding:"required" comment:"更新内容" form:"content"`
}

// RestartDeploymentRequest 重启deployment的请求参数
type RestartDeploymentRequest struct {
	Name      string `json:"name" binding:"required" form:"name"`
	Namespace string `json:"namespace" binding:"required" form:"namespace"`
}

// ScaleDeploymentRequest 扩容deployment的请求参数
type ScaleDeploymentRequest struct {
	Name string `json:"name" binding:"required" form:"name"`
	// 命名空间
	Namespace string `json:"namespace" binding:"required" form:"namespace"`
	// 期望副本数
	ScaleNum int32 `json:"scale_num" binding:"required" form:"scale_num"`
}

// CountDeploymentRequest 获取deployment数量的请求参数
type CountDeploymentRequest struct {
	Namespace string `json:"namespace" binding:"required" form:"namespace"`
}

// CountDeploymentResponse 获取deployment列表的请求参数
type CountDeploymentResponse struct {
	Count int32 `json:"count"`
}
