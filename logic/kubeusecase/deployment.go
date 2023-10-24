package kubeusecase

import (
	"gin_demo/dao/kuberepo"

	appv1 "k8s.io/api/apps/v1"
)

type DeploymentUseCase struct {
	Factor *kuberepo.Deployment
}

func (c *DeploymentUseCase) GetList(namespce string) (list *appv1.DeploymentList, err error) {
	list, err = c.Factor.GetDeploymentList(namespce)
	if err != nil {
		return
	}
	return
}

func NewDeploymentUseCase(d *kuberepo.Deployment) *DeploymentUseCase {
	return &DeploymentUseCase{Factor: d}
}
