# k8s resouce client

linux
```
GOOS=linux;GOARCH=amd64
./go_build_main_go_linux  -ns kube-system
```
mac
```
./package  -ns istio-system
```

version
```
k8s version:
  Client Version: v1.11.0+d4cacc0
  Server Version: v1.11.0+d4cacc0

  Client Version: v1.14.7
  Server Version: v1.14.7
```
ca and token 
```
kubectl exec -it  poName -n nsName  /bin/sh
TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
CA=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)

```
```
kubeconfig 连接
https://www.cnblogs.com/jokerjason/p/11898428.html
```
```
golang文件追加覆盖
https://www.cnblogs.com/kumata/p/10161754.html
```

```
ssh传输文件
https://www.cnblogs.com/-xuzhankun/p/11056576.html
```
