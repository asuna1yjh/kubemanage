package kubeusecase

import (
	"bufio"
	"context"
	"fmt"
	"gin_demo/controllers/common"
	"gin_demo/controllers/types"
	"gin_demo/dao/kuberepo"
	"gin_demo/pkg/terminal"
	"io"
	"net/http"

	"k8s.io/client-go/kubernetes/scheme"

	"k8s.io/client-go/tools/remotecommand"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"

	corev1 "k8s.io/api/core/v1"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodUseCase struct {
	KubeFactory *kuberepo.KubeFactory
}

// resultResp 是返回pod列表的结构体
type resultResp struct {
	Total int `json:"total"`
	data  []*corev1.Pod
}

func newResultResp(data []*corev1.Pod) *resultResp {
	return &resultResp{data: data}
}

func NewPodUseCase(kubeFactory *kuberepo.KubeFactory) *PodUseCase {
	return &PodUseCase{KubeFactory: kubeFactory}
}

func (p *PodUseCase) GetListPods(param *types.GetListPodsRequest) (total int, pods []*corev1.Pod, err error) {
	// 1. 查询pod列表
	podList, err := p.KubeFactory.Client.CoreV1().Pods(param.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("查询pod列表失败", zap.Error(err))
		return 0, nil, err
	}
	data := make([]*corev1.Pod, len(podList.Items))
	for k, _ := range podList.Items {
		data[k] = &podList.Items[k]
	}
	resp := newResultResp(data)
	resp.QueryPods(param.Page, param.Size)
	// 2. 返回结果
	return resp.Total, resp.data, nil
}

func (r *resultResp) QueryPods(page, size int) {
	fmt.Printf("page: %d, size: %d\n", page, size)
	//计算要返回的数据的起始索引和结束索引
	startIndex := (page - 1) * size
	endIndex := page * size
	// 1. 判断是否超出了索引
	// 1.1 startIndex超出了索引
	if startIndex > len(r.data) {
		startIndex = len(r.data)
	}
	if endIndex > len(r.data) {
		endIndex = len(r.data)
	}
	r.Total = len(r.data)
	fmt.Printf("startIndex: %d, endIndex: %d\n", startIndex, endIndex)
	fmt.Printf("len(r.data): %d\n", len(r.data))
	r.data = r.data[startIndex:endIndex]

}

func (p *PodUseCase) GetPodLog(param *types.GetPodLogRequest, c *gin.Context) (log string, err error) {
	// http升级为websocket
	// 1. 升级为websocket
	UP := websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 2. 升级
	conn, err := UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Error("升级为websocket失败", zap.Error(err))
		return "", err
	}
	var logNum = int64(100)
	options := &corev1.PodLogOptions{
		Container: param.ContainerName, // 容器名
		Follow:    true,                // 是否跟随
		TailLines: &logNum,             // 读取日志的行数
	}
	logs := p.KubeFactory.Client.CoreV1().Pods(param.Namespace).GetLogs(param.PodName, options)
	stream, err := logs.Stream(context.Background())
	if err != nil {
		zap.L().Error("获取pod日志失败", zap.Error(err))
		return "", err
	}
	var stopChan = make(chan struct{})
	go func() {
		select {
		case <-stopChan:
			err := stream.Close()
			zap.L().Debug("关闭日志stream")
			if err != nil {
				zap.L().Error("关闭stream失败", zap.Error(err))
				return
			}
			err = conn.Close()
			zap.L().Debug("关闭websocket连接")
			if err != nil {
				zap.L().Error("关闭conn失败", zap.Error(err))
				return
			}
		}
	}()
	// 1. 关闭stream

	//	defer stream.Close()
	//defer func(stream io.ReadCloser) {
	//	fmt.Printf("关闭stream\n")
	//	err := stream.Close()
	//	if err != nil {
	//		zap.L().Error("关闭stream失败", zap.Error(err))
	//	}
	//}(stream)
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				zap.L().Error("读取日志失败", zap.Error(err))
				return
			}
			if string(p) == "exit" {
				close(stopChan)
				return
			}
		}
	}()
	// 2. 读取日志
	reader := bufio.NewReader(stream)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			zap.L().Error("reader.ReadString to failed", zap.Error(err))
			return "", err
		}
		// 3. 发送日志
		err = conn.WriteMessage(websocket.TextMessage, []byte(line))
		if err != nil {
			zap.L().Error("conn.WriteMessage to failed", zap.Error(err))
			return "", err
		}
	}
	// 3. 返回日志
	return "", nil
}

func (p *PodUseCase) GetPodDetail(param *types.GetPodDetailRequest) (pod *corev1.Pod, err error) {
	pod, err = p.KubeFactory.Client.CoreV1().Pods(param.Namespace).Get(context.Background(), param.PodName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("获取pod详情失败", zap.Error(err))
		return
	}
	// 清空pod的managedFields字段
	pod.ManagedFields = nil
	return
}

func (p *PodUseCase) DeletePod(param *types.GetPodDetailRequest) (err error) {
	pod, err := p.GetPodDetail(param)
	if err != nil {
		if pod.Name == "" {
			zap.L().Info("pod不存在", zap.Error(err))
			return common.ErrorNotExist
		}
		return err
	}
	if err := p.KubeFactory.Client.CoreV1().Pods(param.Namespace).Delete(context.Background(), param.PodName, metav1.DeleteOptions{}); err != nil {
		zap.L().Error("删除pod失败", zap.Error(err))
		return err
	}
	return
}

func (p *PodUseCase) WebShell(param *types.GetPodRequest, c *gin.Context) (err error) {
	// 1. 查询container是否存在
	pod, err := p.GetPodDetail(&types.GetPodDetailRequest{
		Namespace: param.Namespace,
		PodName:   param.PodName,
	})
	if err != nil {
		if pod.Name == "" {
			zap.L().Info("pod不存在", zap.Error(err))
			return common.ErrorNotExist
		}
		zap.L().Error("查询pod失败", zap.Error(err))
		return common.ErrorServerBusy
	}
	for k, _ := range pod.Spec.Containers {
		if pod.Spec.Containers[k].Name == param.ContainerName {
			break
		}
		if k == len(pod.Spec.Containers)-1 {
			zap.L().Info("container不存在", zap.Error(err))
			return common.ErrorNotExist
		}
	}
	// 2. 升级为websocket
	UP := websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 2. 升级
	conn, err := UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Error("升级为websocket失败", zap.Error(err))
		return err
	}
	// 3. 升级为流
	session := terminal.NewTerminalSession(conn)
	// 构造请求
	req := p.KubeFactory.Client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(param.PodName).
		Namespace(param.Namespace).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: param.ContainerName,
			Command:   []string{"/bin/sh"},
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)
	// ParameterCodec
	// 执行请求
	exec, err := remotecommand.NewSPDYExecutor(p.KubeFactory.Config, "POST", req.URL())
	if err != nil {
		zap.L().Error("创建SPDYExecutor失败", zap.Error(err))
		return err
	}
	// 3. 升级为流
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  session,
		Stdout: session,
		Stderr: session,
		Tty:    true,
	})
	if err != nil {
		zap.L().Error("执行请求失败", zap.Error(err))
		session.Close()
		return err
	}
	return nil
}

func (p *PodUseCase) CountPod(param *types.CountPodRequest) (ret int64, err error) {
	podList, err := p.KubeFactory.Client.CoreV1().Pods(param.Namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("查询pod列表失败", zap.Error(err))
		return 0, err
	}
	return int64(len(podList.Items)), nil
}
