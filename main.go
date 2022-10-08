package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

type podMetric struct {
	Metadata struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Containers []struct {
		Name  string `json:"name"`
		Usage struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"usage"`
	} `json:"containers"`
}

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("", "C:/Users/Rig/.kube/config")
	clientset, _ := kubernetes.NewForConfig(config)
	pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	GetMetrics()
}

func GetMetrics() {
	config, _ := clientcmd.BuildConfigFromFlags("", "C:/Users/Rig/.kube/config")

	clientset, err := metricsv.NewForConfig(config)
	if err != nil {
		fmt.Printf(err.Error())
		panic("ERRO 1")
	}

	for i := 0; i != -1; i++ {
		data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/namespaces/app/pods/nginx-deployment-69cb66965-jwfmf").DoRaw(context.TODO())
		if err != nil {
			fmt.Printf(err.Error())
			panic("Error Collecting Metrics for the PODs.")
		}

		formated_data := podMetric{}
		fmt.Println(string(data))
		json.Unmarshal([]byte(string(data)), &formated_data)
		fmt.Printf("Pod: %s Usage: %s\n", formated_data.Metadata, formated_data.Containers)
		time.Sleep(1 * time.Second)
	}
}
