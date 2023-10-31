package kubeusecase

import (
	"context"
	"gin_demo/controllers/common"
	"gin_demo/controllers/types"
	"gin_demo/dao/kuberepo"
	"strconv"
	"time"

	"k8s.io/apimachinery/pkg/util/json"

	"k8s.io/apimachinery/pkg/util/intstr"

	corev1 "k8s.io/api/core/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"go.uber.org/zap"

	appv1 "k8s.io/api/apps/v1"
)

type DeploymentUseCase struct {
	Factor *kuberepo.Deployment
}

func NewDeploymentUseCase(d *kuberepo.Deployment) *DeploymentUseCase {
	return &DeploymentUseCase{Factor: d}
}

func (c *DeploymentUseCase) GetList(namespce string) (list *appv1.DeploymentList, err error) {
	list, err = c.Factor.GetDeploymentList(namespce)
	if err != nil {
		zap.L().Error("获取deployment列表失败", zap.Any("err", err))
		return
	}
	for k, _ := range list.Items {
		zap.L().Info("获取deployment列表成功", zap.Any("list", list.Items[k]))

	}
	return
}

func (c *DeploymentUseCase) CreateDeployment(p *types.CreateDeploymentRequest) (err error) {
	// 创建deployment
	deploy := appv1.Deployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      p.Name,
			Labels:    p.Labels,
			Namespace: p.Namespace,
		},
		Spec: appv1.DeploymentSpec{
			Replicas: &p.Replicas,
			Selector: &v1.LabelSelector{
				MatchLabels: p.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name:   p.Name,
					Labels: p.Labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            p.Name,
							Image:           p.Image,
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: p.ContainerPort,
								},
							},
						},
					},
				},
			},
		},
	}
	// 如果需要健康检查
	if p.HealthCheck {
		deploy.Spec.Template.Spec.Containers[0].ReadinessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: p.HealthPath,
					Port: intstr.IntOrString{
						IntVal: p.ContainerPort,
					},
					Scheme: "HTTP",
				},
			},
			InitialDelaySeconds: 10,
			TimeoutSeconds:      2,
			PeriodSeconds:       5,
			SuccessThreshold:    1,
			FailureThreshold:    3,
		}
		deploy.Spec.Template.Spec.Containers[0].LivenessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: p.HealthPath,
					Port: intstr.IntOrString{
						IntVal: p.ContainerPort,
					},
					Scheme: "HTTP",
				},
			},
			InitialDelaySeconds: 10,
			TimeoutSeconds:      2,
			PeriodSeconds:       5,
			SuccessThreshold:    1,
			FailureThreshold:    3,
		}
	}
	// 创建deployment
	if err := c.Factor.CreateDeployment(&deploy); err != nil {
		zap.L().Error("创建deployment失败", zap.Any("err", err))
		return err
	}
	return err
}

func (c *DeploymentUseCase) DeleteDeployment(p *types.DeleteDeploymentRequest) (err error) {
	// 1. 查看deployment是否存在
	if data, err := c.Factor.KubeFactory.Client.AppsV1().Deployments(p.Namespace).Get(context.Background(), p.Name, v1.GetOptions{}); err != nil {
		if data.Name == "" {
			zap.L().Info("deployment不存在", zap.Any("err", err))
			return common.ErrorNotExist
		}
		zap.L().Error("查询deployment失败", zap.Any("err", err))
		return err
	}
	// 2. 删除deployment
	if err := c.Factor.KubeFactory.Client.AppsV1().Deployments(p.Namespace).Delete(context.Background(), p.Name, v1.DeleteOptions{}); err != nil {
		zap.L().Error("删除deployment失败", zap.Any("err", err))
		return err
	}
	return
}

func (c *DeploymentUseCase) DetailDeployment(p *types.DetailDeploymentRequest) (*appv1.Deployment, error) {
	// 1. 查看deployment是否存在
	data, err := c.Factor.KubeFactory.Client.AppsV1().Deployments(p.Namespace).Get(context.Background(), p.Name, v1.GetOptions{})
	if err != nil {
		if data == nil {
			zap.L().Info("deployment不存在", zap.Any("err", err))
			return nil, common.ErrorNotExist
		}
		zap.L().Error("查询deployment失败", zap.Any("err", err))
		return nil, common.ErrorServerBusy
	}
	data.ObjectMeta.ManagedFields = nil
	return data, nil
}

func (c *DeploymentUseCase) UpdateDeployment(p *types.UpdateDeploymentRequest) (*appv1.Deployment, error) {
	deploy := &appv1.Deployment{}
	err := json.Unmarshal([]byte(p.Content), deploy)
	if err != nil {
		zap.L().Error("转换deploy失败", zap.Any("err", err))
		return nil, common.ErrInvalidParams
	}
	zap.L().Debug("转换deploy成功", zap.Any("deploy", deploy))
	// 调用sdk去更新deployment
	data, err := c.Factor.KubeFactory.Client.AppsV1().Deployments(p.Namespace).Update(context.Background(), deploy, v1.UpdateOptions{})
	if err != nil {
		zap.L().Error("更新deployment失败", zap.Any("err", err))
		return nil, err
	}
	data.ObjectMeta.ManagedFields = nil
	return data, nil
}

func (c *DeploymentUseCase) RestartDeployment(p *types.RestartDeploymentRequest) (data *appv1.Deployment, err error) {
	// 1. 查看deployment是否存在
	data, err = c.Factor.KubeFactory.Client.AppsV1().Deployments(p.Namespace).Get(context.Background(), p.Name, v1.GetOptions{})
	if err != nil {
		if data.Name == "" {
			zap.L().Info("deployment不存在", zap.Any("err", err))
			return
		}
		zap.L().Error("查询deployment失败", zap.Any("err", err))
		return
	}
	// 2. 重启deployment
	// 随便改一个无关的值,就会触发重启
	//data.ObjectMeta.Labels["timestamp"] = "restart-" + time.Now().Format("20060102150405")
	data.Spec.Template.ObjectMeta.Labels["timestamp"] = strconv.Itoa(int(time.Now().Unix()))
	// 直接更新deployment
	data, err = c.Factor.KubeFactory.Client.AppsV1().Deployments(p.Namespace).Update(context.Background(), data, v1.UpdateOptions{})
	if err != nil {
		zap.L().Error("重启deployment失败", zap.Any("err", err))
	}
	return
}

func (c *DeploymentUseCase) ScaleDeployment(p *types.ScaleDeploymentRequest) (data *appv1.Deployment, err error) {
	// 1. 查看deployment是否存在
	data, err = c.Factor.KubeFactory.Client.AppsV1().Deployments(p.Namespace).Get(context.Background(), p.Name, v1.GetOptions{})
	if err != nil {
		if data.Name == "" {
			zap.L().Info("deployment不存在", zap.Any("err", err))
			return
		}
		zap.L().Error("查询deployment失败", zap.Any("err", err))
		return
	}
	// 2. 扩容deployment
	// 2.1 获取当前副本数
	// 2.2 修改副本数
	data.Spec.Replicas = &p.ScaleNum
	// 2.3 更新deployment
	data, err = c.Factor.KubeFactory.Client.AppsV1().Deployments(p.Namespace).Update(context.Background(), data, v1.UpdateOptions{})
	if err != nil {
		zap.L().Error("扩容deployment失败", zap.Any("err", err))
		return
	}
	return
}

func (c *DeploymentUseCase) CountDeployment(p *types.CountDeploymentRequest) (count int32, err error) {
	// 1. 获取deployment列表
	list, err := c.Factor.GetDeploymentList(p.Namespace)
	if err != nil {
		zap.L().Error("获取deployment列表失败", zap.Any("err", err))
		return
	}
	count = int32(len(list.Items))
	return
}
