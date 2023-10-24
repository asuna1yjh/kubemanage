package kuberepo

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 这是对kubenetes的操作，封装了一些常用的操作
// Path: pkg/kuberepo/kubemanage/v1/init.go

type KubeFactory struct {
	Client kubernetes.Interface
}

func NewKubeFactory() *KubeFactory {
	// 从本地的kubeconfig文件加载配置, 此处为在k8s集群外部使用
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		zap.L().Error("获取k8s config 失败")
		return nil
	}
	// 创建一个clientSet
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		zap.L().Error("获取k8s clientSet 失败")
		return nil
	}
	zap.L().Info("获取k8s clientSet 成功")
	return &KubeFactory{Client: clientset}
}
