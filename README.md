# k8s client

linux
```
GOOS=linux;GOARCH=amd64
./go_build_main_go_linux  -ns kube-system
```
mac
```
./package  -n istio-system
```

```
k8s version:
  Client Version: v1.11.0+d4cacc0
  Server Version: v1.11.0+d4cacc0

  Client Version: v1.14.7
  Server Version: v1.14.7
```
```
ca.crt token获取路径: sa in pod:/var/run/secrets/kubernetes.io/serviceaccount
```
