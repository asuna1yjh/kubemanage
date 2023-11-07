package kubecontroller

import (
	"fmt"
	"gin_demo/controllers/common"
	"gin_demo/controllers/types"
	"gin_demo/logic/kubeusecase"

	"github.com/gin-gonic/gin"
)

// 管理pod的调用
type Pod struct {
	uc *kubeusecase.PodUseCase
}

func NewPod(pod *kubeusecase.PodUseCase) *Pod {
	return &Pod{uc: pod}
}

// GetPods 查询pod列表
// @Summary 查询pod列表
// @Description 查询pod列表
// @Tags k8s
// @Accept json
// @Produce json
// @Param namespace query string true "命名空间"
// @Param limit query int true "分页限制"
// @Param page query int true "页码"
// @Success 200 {object} string "{"code": 200, msg="success", "data": "pod对象列表"}"
func (p *Pod) GetListPods(c *gin.Context) {
	// 1. 参数校验
	param := new(types.GetListPodsRequest)
	err := common.Parameter(c, param)
	if err != nil {
		return
	}
	// 2. 业务逻辑
	total, pods, err := p.uc.GetListPods(param)
	if err != nil {
		common.ResponseError(c, common.CodeQueryError)
	}
	// 3. 返回结果
	common.ResponseSuccess(c, gin.H{
		"total": total,
		"pods":  pods,
	})
}

// GetPodLog 查询pod日志
// @Summary 查询pod日志
// @Description 查询pod日志
// @Tags k8s
// @Accept Message Text
// @Produce json
// @Param namespace query string true "命名空间"
// @Param podName query string true "pod名称"
// @Param containerName query string true "容器名称"
func (p *Pod) GetPodLog(c *gin.Context) {
	// 1. 参数校验
	param := new(types.GetPodLogRequest)
	err := common.Parameter(c, param)
	if err != nil {
		return
	}
	// 2. 业务逻辑
	log, err := p.uc.GetPodLog(param, c)
	if err != nil {
		common.ResponseError(c, common.CodeQueryError)
	}
	// 3. 返回结果
	common.ResponseSuccess(c, gin.H{
		"log": log,
	})

}

// GetPodDetail 查询pod详情
// @Summary 查询pod详情
// @Description 查询pod详情
// @Tags k8s
// @Accept json
// @Produce json
// @Param namespace query string true "命名空间"
// @Param podName query string true "pod名称"
// @Success 200 {object} string "{"code": 200, msg="success", "data": "pod对象"}"
// @Router /api/k8s/pod/detail [get]
func (p *Pod) GetPodDetail(c *gin.Context) {
	// 1. 参数校验
	param := new(types.GetPodDetailRequest)
	err := common.Parameter(c, param)
	if err != nil {
		return
	}
	// 2. 业务逻辑
	pod, err := p.uc.GetPodDetail(param)
	if err != nil {
		common.ResponseError(c, common.CodeQueryError)
	}
	// 3. 返回结果
	common.ResponseSuccess(c, pod)
}

// DeletePod 删除pod
// @Summary 删除pod
// @Description 删除pod
// @Tags k8s
// @Accept json
// @Produce json
// @Param namespace query string true "命名空间"
// @Param podName query string true "pod名称"
// @Success 200 {object} string "{"code": 200, msg="success", "data": "pod对象"}"
// @Router /api/k8s/pod/del [delete]
func (p *Pod) DeletePod(c *gin.Context) {
	// 1. 参数校验
	param := new(types.GetPodDetailRequest)
	err := common.Parameter(c, param)
	if err != nil {
		return
	}
	// 2. 业务逻辑
	err = p.uc.DeletePod(param)
	if err != nil {
		if err == common.ErrorNotExist {
			common.ResponseError(c, common.CodeNotExist)
			return
		}
		common.ResponseError(c, common.CodeDeleteError)
		return
	}
	// 3. 返回结果
	common.ResponseSuccess(c, nil)
}

// WebShell 连接pod的shell界面
// @Summary 连接pod的shell界面
// @Description 连接pod的shell界面
// @Tags k8s
// @Accept Message Text
// @Produce Message Text
// @Param namespace query string true "命名空间"
// @Param podName query string true "pod名称"
// @Param containerName query string true "容器名称"
// @Success 200 {object} string "{"code": 200, msg="success", "data": "pod对象"}"
// @Router /api/k8s/pod/webshell [get]
func (p *Pod) WebShell(c *gin.Context) {
	// 1. 参数校验
	param := new(types.GetPodRequest)
	err := common.Parameter(c, param)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 2. 业务逻辑
	err = p.uc.WebShell(param, c)
	if err != nil {
		common.ResponseError(c, common.CodeQueryError)
	}
	// 3. 返回结果
	common.ResponseSuccess(c, nil)
}

// CountPod 查询pod数量
// @Summary 查询pod数量
// @Description 查询pod数量
// @Tags k8s
// @Accept json
// @Produce json
// @Param namespace query string true "命名空间"
// @Success 200 {object} string "{"code": 200, msg="success", "data": "pod数量"}"
// @Router /api/k8s/pod/count [get]
func (p *Pod) CountPod(c *gin.Context) {
	// 1. 参数校验
	param := new(types.CountPodRequest)
	err := common.Parameter(c, param)
	if err != nil {
		return
	}
	// 2. 业务逻辑
	count, err := p.uc.CountPod(param)
	if err != nil {
		common.ResponseError(c, common.CodeQueryError)
	}
	// 3. 返回结果
	common.ResponseSuccess(c, count)
}
