package main

import (
	"flag"
	"goclient/client"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {

	//get pods in ns by args
	var namespaceUser string
	flag.StringVar(&namespaceUser,"ns","default","namespace")
	flag.Parse()

	clientset := client.GetClient("conf/token", "conf/ca.crt", "10.125.31.171:8443")
	pods, err := clientset.CoreV1().Pods(namespaceUser).List(v1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}


}
