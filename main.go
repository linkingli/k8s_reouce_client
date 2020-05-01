package main

import (
	"flag"
	"goclient/client"
	"goclient/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {

	//get pods in ns by args by https
	var namespaceUser string
	flag.StringVar(&namespaceUser, "ns", "default", "namespace")
	flag.Parse()

	clientset := client.GetClient()
	pods, err := clientset.CoreV1().Pods(namespaceUser).List(v1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}

	//apply deployment by yaml
	resource.ApplyDeployment()

	//change deployment image by yaml
	resource.ChangeDeploymentImage()

	//rollupdate Image
	resource.RollUpdateDeployment()

	resource.GetPodLog("istio-system", "istiod-6c45d49679-sps4q", "discovery")
}
