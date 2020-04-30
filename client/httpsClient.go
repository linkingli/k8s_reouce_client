package client

import (
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/cert"
	"log"
)

func GetClient(tokenPath, caPath, masterUrl string) *kubernetes.Clientset {
	token, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		fmt.Println(err)
	}
	tlsClientConfig := rest.TLSClientConfig{}
	if _, err := cert.NewPool(caPath); err != nil {
		log.Printf("Expected to load root CA config from %s, but got err: %v", caPath, err)
	} else {
		tlsClientConfig.CAFile = caPath
	}
	config := rest.Config{
		Host:            masterUrl,
		TLSClientConfig: tlsClientConfig,
		BearerToken:     string(token),
		BearerTokenFile: tokenPath,
	}
	client, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Printf("faild to create kuber client %s", err)
	}
	return client
}
