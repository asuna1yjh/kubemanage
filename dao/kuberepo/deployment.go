package kuberepo

import (
	"context"

	appv1 "k8s.io/api/apps/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	KubeFactory *KubeFactory
}

func NewDeployment(kubeFactory *KubeFactory) *Deployment {
	return &Deployment{
		KubeFactory: kubeFactory,
	}
}

func (d *Deployment) GetDeploymentList(namespce string) (list *appv1.DeploymentList, err error) {
	list, err = d.KubeFactory.Client.AppsV1().Deployments(namespce).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return
	}
	return
}

func (d *Deployment) CreateDeployment(deploy *appv1.Deployment) (err error) {
	_, err = d.KubeFactory.Client.AppsV1().Deployments(deploy.Namespace).Create(context.Background(), deploy, v1.CreateOptions{})
	if err != nil {
		return err
	}
	return
}
