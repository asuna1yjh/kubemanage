package kubecontroller

import (
	"gin_demo/controllers/common"
	"gin_demo/controllers/types"
	"gin_demo/logic/kubeusecase"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Deployment struct {
	uc *kubeusecase.DeploymentUseCase
}

func NewDeployment(uc *kubeusecase.DeploymentUseCase) *Deployment {
	return &Deployment{
		uc: uc,
	}
}

// CreateDeployment 创建deployment
// @Summary 创建deployment
// @Description 创建deployment
// @Tags k8s
// @Accept json
// @Produce json
// @Param body body string true "body"
// @Success 200 {object} string "{"code": 200, msg="", "data": "新增成功"}"
// @Router /k8s/deployment/create [post]
func (d *Deployment) CreateDeployment(c *gin.Context) {
	// 1. 参数校验
	p := new(types.CreateDeploymentRequest)
	if err := common.Parameter(c, p); err != nil {
		zap.L().Debug("CreateDeployment", zap.Any("p", p))
		return
	}
	// 2. 业务逻辑
	if err := d.uc.CreateDeployment(p); err != nil {
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	common.ResponseSuccess(c, "新增成功")

}

// DeleteDeployment 删除deployment
// @Summary 删除deployment
// @Description 删除deployment
// @Tags k8s
// @Accept json
// @Produce json
// @Param name query string true "Deployment名称"
// @Param namespace query string true "命名空间"
// @Success 200 {object} string "{"code": 200, msg="", "data": "删除成功"}"
// @Router /k8s/deployment/delete [delete]
func (d *Deployment) DeleteDeployment(c *gin.Context) {
	// 1. 参数校验
	p := new(types.DeleteDeploymentRequest)
	if err := common.Parameter(c, p); err != nil {
		zap.L().Debug("DeleteDeployment", zap.Any("p", p))
		return
	}
	// 2. 业务逻辑
	if err := d.uc.DeleteDeployment(p); err != nil {
		if err == common.ErrorNotExist {
			common.ResponseError(c, common.CodeNotExist)
			return
		}
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	common.ResponseSuccess(c, "删除成功")
}

// DetailDeployment 获取deployment详情
// @Summary 获取deployment详情
// @Description 获取deployment详情
// @Tags k8s
// @Accept json
// @Produce json
// @Param name query string true "Deployment名称"
// @Param namespace query string true "命名空间"
// @Success 200 {object} string "{"code": 200, msg="", "data": "获取成功"}"
// @Router /k8s/deployment/detail [get]
func (d *Deployment) DetailDeployment(c *gin.Context) {
	// 1. 参数校验
	p := new(types.DetailDeploymentRequest)
	if err := common.Parameter(c, p); err != nil {
		zap.L().Debug("DetailDeployment", zap.Any("p", p))
		return
	}
	// 2. 业务逻辑
	deployment, err := d.uc.DetailDeployment(p)
	if err != nil {
		if err == common.ErrorNotExist {
			common.ResponseError(c, common.CodeNotExist)
			return
		}
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	common.ResponseSuccess(c, deployment)
}

// GetList 获取deployment列表
// @Summary 获取deployment列表
// @Description 获取deployment列表
// @Tags k8s
// @Accept json
// @Produce json
// @Param namespace query string true "命名空间"
// @Success 200 {object} string "{"code": 200, msg="", "data": "获取成功"}"
// @Router /k8s/deployment/list [get]
func (d *Deployment) GetList(c *gin.Context) {
	// 1. 参数校验
	p := new(types.GetDeploymentListRequest)
	if err := common.Parameter(c, p); err != nil {
		//zap.L().Debug("GetList", zap.Any("p", p))
		return
	}
	list, err := d.uc.GetList(p.Namespace)
	if err != nil {
		return
	}
	common.ResponseSuccess(c, list)
}

// UpdateDeployment 更新deployment
// @Summary 更新deployment
// @Description 更新deployment
// @Tags k8s
// @Accept json
// @Produce json
// @Param body body string true "body"
// @Success 200 {object} string "{"code": 200, msg="", "data": "更新成功"}"
// @Router /k8s/deployment/update [post]
func (d *Deployment) UpdateDeployment(c *gin.Context) {
	// 1. 参数校验
	p := new(types.UpdateDeploymentRequest)
	if err := common.Parameter(c, p); err != nil {
		zap.L().Debug("GetList", zap.Any("p", p))
		return
	}
	// 2. 业务逻辑
	data, err := d.uc.UpdateDeployment(p)
	if err != nil {
		if err == common.ErrInvalidParams {
			common.ResponseError(c, common.CodeInvalidParams)
			return
		}
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	common.ResponseSuccess(c, data)
}

// RestartDeployment 重启deployment
// @Summary 重启deployment
// @Description 重启deployment
// @Tags k8s
// @Accept json
// @Produce json
// @Param body body string true "body"
// @Success 200 {object} string "{"code": 200, msg="", "data": "重启成功"}"
// @Router /k8s/deployment/restart [post]
func (d *Deployment) RestartDeployment(c *gin.Context) {
	// 1. 参数校验
	p := new(types.RestartDeploymentRequest)
	if err := common.Parameter(c, p); err != nil {
		zap.L().Debug("GetList", zap.Any("p", p))
		return
	}
	// 2. 业务逻辑
	data, err := d.uc.RestartDeployment(p)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerBusy, "重启失败")
		return
	}
	common.ResponseSuccessWithMsg(c, "重启成功", data)
}

// ScaleDeployment 扩容deployment
// @Summary 扩容deployment
// @Description 扩容deployment
// @Tags k8s
// @Accept json
// @Produce json
// @Param body body string true "body"
// @Success 200 {object} string "{"code": 200, msg="", "data": "扩容成功"}"
// @Router /k8s/deployment/scale [put]
func (d *Deployment) ScaleDeployment(c *gin.Context) {
	// 1. 参数校验
	p := new(types.ScaleDeploymentRequest)
	if err := common.Parameter(c, p); err != nil {
		zap.L().Debug("GetList", zap.Any("p", p))
		return
	}
	// 2. 业务逻辑
	data, err := d.uc.ScaleDeployment(p)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerBusy, "扩容失败")
		return
	}
	common.ResponseSuccessWithMsg(c, "扩容成功", gin.H{"Replicas": data.Spec.Replicas})
}

// CountDeployment 获取deployment数量
// @Summary 获取deployment数量
// @Description 获取deployment数量
// @Tags k8s
// @Accept json
// @Produce json
// @Param body body string true "body"
// @Success 200 {object} string "{"code": 200, msg="", "data": "扩容成功"}"
// @Router /k8s/deployment/scale [get]
func (d *Deployment) CountDeployment(c *gin.Context) {
	// 1. 参数校验
	p := new(types.CountDeploymentRequest)
	if err := common.Parameter(c, p); err != nil {
		zap.L().Debug("GetList", zap.Any("p", p))
		return
	}
	// 2. 业务逻辑
	data, err := d.uc.CountDeployment(p)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerBusy, "获取数量失败")
		return
	}
	common.ResponseSuccessWithMsg(c, "获取数量成功", types.CountDeploymentResponse{Count: data})
}
