package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type patchStringValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

var pods = []string{"diffusion-0", "diffusion-1", "diffusion-2"}
var podsNamespace = "scoreboards"
var returnBytes []byte
var putLabels = map[string]string{
	"client-type": "internal",
	"sandeep":     "lamba",
}

func main() {
	// Global Variables
	var kubeconfig = flag.String("kubeconfig", "/Users/slamba/.kube/config", "clutser-kubeconfig file")

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("The kubeconfig can not be loaded %v", err)
		os.Exit(1)
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error is %v", err)
		os.Exit(1)
	}
	// Watch the Pods for Namespace
	watcher, err := clientset.CoreV1().Pods(podsNamespace).Watch(metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	// Create Channel for Interface
	ch := watcher.ResultChan()
	for event := range ch {
		isLabel := false
		podEvent, ok := event.Object.(*v1.Pod)
		if !ok {
			log.Fatal("Unexpected type")
		}
		if event.Type == watch.Added || event.Type == watch.Modified {
			for key, value := range putLabels {
				returnBytes = createPayload(key, value)
				// make variables from events
				getName := podEvent.GetName()
				getNamespace := podEvent.GetNamespace()
				getLabels := podEvent.GetLabels()
				// loop for all pods
				for allPods := range pods {
					if pods[allPods] == getName && podsNamespace == getNamespace {
						log.Printf("\tEvent Type : %s\n", event.Type)
						for j := range getLabels {
							if j == key && getLabels[j] == value {
								isLabel = true
							}
						}
						if !isLabel {
							log.Printf("\tPod:%s dont have labels %s=%s\n", pods[allPods], key, value)
							_, updateErr := clientset.CoreV1().Pods(podsNamespace).Patch(pods[allPods], types.JSONPatchType, returnBytes)
							if updateErr != nil {
								log.Printf("\tThere is error in fetching pod details %v\n", updateErr)
							}
							log.Printf("\tPod %s labelled successfully with label %s=%s\n", pods[allPods], key, value)
						} else {
							log.Printf("\tLabel %s=%s already present for pod %s\n", key, value, pods[allPods])
						}
					}
				}
				// isLabelPresent(getName, getNamespace, getLabels)
			}
		}
	}
}

// function to create payload to make labels for pod
func createPayload(podKey string, podValue string) []byte {
	payload := []patchStringValue{{
		Op:    "replace",
		Path:  "/metadata/labels/" + podKey,
		Value: podValue,
	}}
	payloadBytes, _ := json.Marshal(payload)
	return payloadBytes
}

// function to check if labels is present of or not
func isLabelPresent(getName string, getNamespace string, getLabels map[string]string) bool {
	isLabel := false
	for allPods := range pods {
		if pods[allPods] == getName && podsNamespace == getNamespace {
			for j := range getLabels {
				if j == "client-type" && getLabels[j] == "internal" {
					isLabel = true
				}
			}
		}
	}
	return isLabel
}
