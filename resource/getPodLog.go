package resource

import (
	"fmt"
	"goclient/client"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

//todo get stream log failed
func GetPodLog(ns string, podName string, containerName string) {

	var (
		clientset *kubernetes.Clientset
		tailLines int64
		req       *rest.Request
		res       rest.Result
		logs      []byte
		err       error
	)

	// 初始化k8s客户端
	clientset = client.GetClient()

	// 生成获取POD日志请求
	req = clientset.CoreV1().Pods(ns).GetLogs(podName, &core_v1.PodLogOptions{Container: containerName, TailLines: &tailLines})

	// req.Stream()也可以实现Do的效果

	// 发送请求
	if res = req.Do(); res.Error() != nil {
		err = res.Error()
		goto FAIL
	}

	// 获取结果
	if logs, err = res.Raw(); err != nil {
		goto FAIL
	}

	fmt.Println("容器输出:", string(logs))
	return

FAIL:
	fmt.Println(err)
	return
}
