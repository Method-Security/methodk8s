package node

import (
	"context"

	methodk8s "github.com/method-security/methodk8s/generated/go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8Resources struct {
	Nodes []*methodk8s.Node `json:"nodes" yaml:"nodes"`
}

type K8ResourceReport struct {
	Resources K8Resources `json:"resources" yaml:"resources"`
	Errors    []string    `json:"errors" yaml:"errors"`
}

func EnumerateNodes(k8config *rest.Config) (*K8ResourceReport, error) {
	resources := K8Resources{}
	errors := []string{}
	config := k8config

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		errors = append(errors, err.Error())
		return &K8ResourceReport{Errors: errors}, err
	}

	nodesList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return &K8ResourceReport{Errors: errors}, err
	}

	nodes := []*methodk8s.Node{}
	for _, node := range nodesList.Items {
		addresses := []*methodk8s.Address{}
		for _, addr := range node.Status.Addresses {
			address := methodk8s.Address{
				Type:    string(addr.Type),
				Address: addr.Address,
			}
			addresses = append(addresses, &address)
		}

		nodeInfo := methodk8s.Node{
			Name:         node.GetName(),
			Arch:         node.Status.NodeInfo.Architecture,
			Os:           node.Status.NodeInfo.OperatingSystem,
			Instancetype: node.Labels["beta.kubernetes.io/instance-type"],
			Status:       isNodeReady(&node),
			Addresses:    addresses,
		}
		nodes = append(nodes, &nodeInfo)
	}

	if nodes != nil {
		resources.Nodes = nodes
	}

	k8ResourceReport := K8ResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &k8ResourceReport, nil
}

func isNodeReady(node *corev1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == corev1.NodeReady {
			return condition.Status == corev1.ConditionTrue
		}
	}
	return false
}
