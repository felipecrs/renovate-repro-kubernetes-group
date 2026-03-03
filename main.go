package main

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/yaml"
)

func newClientset() (*kubernetes.Clientset, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, nil).ClientConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func main() {
	pod := corev1.Pod{}
	pod.Name = "example"
	pod.Spec.TerminationGracePeriodSeconds = ptr.To(int64(30))

	out, err := yaml.Marshal(pod)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
