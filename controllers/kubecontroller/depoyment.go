package kubecontroller

import (
	"gin_demo/controllers/common"
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

func (s *Deployment) GetList(c *gin.Context) {
	ns, ok := c.GetQuery("namespace")
	if !ok {
		zap.L().Error("参数错误")
		common.ResponseError(c, common.CodeInvalidParams)
	}
	list, err := s.uc.GetList(ns)
	if err != nil {
		return
	}
	common.ResponseSuccess(c, list)
}
