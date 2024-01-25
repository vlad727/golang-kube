package main

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	config, _    = clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	clientset, _ = kubernetes.NewForConfig(config)
)

func main() {
	list, _ := clientset.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})

	for _, item := range list.Items {
		log.Info(item.Name)
	}
}
