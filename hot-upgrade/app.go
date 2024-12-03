package hot_upgrade

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kinddemo/tools"
	"log"
	"os"
)

func Run() {
	podName := "my-go-app-6cc69756cf-4j6rs"
	oldImage := "my-go-app:v3"
	newImage := "my-go-app:v5"

	clientset, err := tools.NewSet()
	if err != nil {
		panic(err.Error())
	}

	// readinessGates 所有的容器是否都read
	//pod.Spec.ReadinessGates
	pod, err := clientset.CoreV1().Pods("default").Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// readinessGates 所有的容器是否都read
	//pod.Spec.ReadinessGates
	for _, gate := range pod.Spec.ReadinessGates {
		if gate.String() != "" {
			fmt.Println("pod not read")
			os.Exit(0)
		}
	}

	// spec.containers[x] 更新pod 镜像
	//pod.Spec.Containers[].Image

	for i, container := range pod.Spec.Containers {
		if container.Image == oldImage {
			pod.Spec.Containers[i].Image = newImage
		}
	}

	_, err = clientset.CoreV1().Pods("default").Update(context.Background(), pod, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 监听pod变化
	watcher, err := clientset.CoreV1().Pods("default").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error creating watcher: %v", err)
	}

	ch := watcher.ResultChan()
	for event := range ch {
		pod, ok := event.Object.(*v1.Pod)
		if !ok {
			log.Fatalf("Unexpected type")
		}

		fmt.Printf("Event Type: %s, Pod Name: %s, Phase: %s\n", event.Type, pod.Name, pod.Status.Phase)
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
