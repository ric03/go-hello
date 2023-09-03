package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "playground"
	deployments := clientset.AppsV1().Deployments(namespace)
	//list, err := deployments.List(context.TODO(), metav1.ListOptions{LabelSelector: "app.kubernetes.io/instance"})
	list, err := deployments.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d deployments in the %s namespace\n", len(list.Items), namespace)
	for _, el := range list.Items {

		for key, value := range el.Annotations {
			if key == "meta.helm.sh/release-name" {
				fmt.Printf("%s, helm release = %s", el.Name, value)
			}
		}
	}

	//for {
	//	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//
	//	// Examples for error handling:
	//	// - Use helper functions like e.g. errors.IsNotFound()
	//	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	//	namespace := "default"
	//	pod := "example-xxxxx"
	//	_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	//	if errors.IsNotFound(err) {
	//		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	//	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	//		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	//			pod, namespace, statusError.ErrStatus.Message)
	//	} else if err != nil {
	//		panic(err.Error())
	//	} else {
	//		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	//	}
	//
	//	time.Sleep(10 * time.Second)
	//}
}

func GetDeploymentStatus(clientset *kubernetes.Clientset) {
	/**
	Options:
	1. check if ready pods are 'equal'

		➜  k8s-nginx k get deploy
		NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
		nginx-healthy-k8s-nginx   1/1     1            1           3s
		nginx-k8s-nginx           0/5     3            0           45m

	2. check status

		HEALTHY DEPLOYMENT:

			status:
			  availableReplicas: 1
			  conditions:
			  - lastTransitionTime: "2023-09-03T22:13:00Z"
				lastUpdateTime: "2023-09-03T22:13:00Z"
				message: Deployment has minimum availability.
				reason: MinimumReplicasAvailable
				status: "True"
				type: Available
			  - lastTransitionTime: "2023-09-03T22:12:59Z"
				lastUpdateTime: "2023-09-03T22:13:00Z"
				message: ReplicaSet "nginx-healthy-k8s-nginx-6bcdf86479" has successfully progressed.
				reason: NewReplicaSetAvailable
				status: "True"
				type: Progressing
			  observedGeneration: 1
			  readyReplicas: 1
			  replicas: 1
			  updatedReplicas: 1



		UNHEALTHY DEPLOYMENT

			status:
			  conditions:
			  - lastTransitionTime: "2023-09-03T22:08:53Z"
				lastUpdateTime: "2023-09-03T22:08:53Z"
				message: Deployment does not have minimum availability.
				reason: MinimumReplicasUnavailable
				status: "False"
				type: Available
			  - lastTransitionTime: "2023-09-03T21:27:04Z"
				lastUpdateTime: "2023-09-03T22:13:32Z"
				message: ReplicaSet "nginx-k8s-nginx-5999446c89" is progressing.
				reason: ReplicaSetUpdated
				status: "True"
				type: Progressing
			  observedGeneration: 7
			  replicas: 7
			  unavailableReplicas: 7
			  updatedReplicas: 3

	*/
}
