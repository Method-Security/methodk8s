package node

import (
	"context"

	methodk8s "github.com/method-security/methodk8s/generated/go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func EnumerateNodes(ctx context.Context, k8config *rest.Config) (*methodk8s.NodeReport, error) {
	resources := methodk8s.NodeReport{}
	errors := []string{}
	config := k8config

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		errors = append(errors, err.Error())
		return &methodk8s.NodeReport{Errors: errors}, err
	}

	nodesList, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return &methodk8s.NodeReport{Errors: errors}, err
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

	resources = methodk8s.NodeReport{
		Nodes:      nodes,
		ClusterUrl: &config.Host,
		Errors:     errors,
	}

	return &resources, nil
}

func isNodeReady(node *corev1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == corev1.NodeReady {
			return condition.Status == corev1.ConditionTrue
		}
	}
	return false
}
