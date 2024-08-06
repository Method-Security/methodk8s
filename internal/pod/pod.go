package pod

import (
	"context"
	"strings"

	methodk8s "github.com/method-security/methodk8s/generated/go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func EnumeratePods(ctx context.Context, k8config *rest.Config) (*methodk8s.PodReport, error) {
	resources := methodk8s.PodReport{}
	errors := []string{}
	config := k8config

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		errors = append(errors, err.Error())
		return &methodk8s.PodReport{Errors: errors}, err
	}

	podsList, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return &methodk8s.PodReport{Errors: errors}, err
	}

	pods := []*methodk8s.Pod{}
	for _, pod := range podsList.Items {
		containers := []*methodk8s.Container{}
		for _, container := range pod.Spec.Containers {
			ports := []*methodk8s.ContainerPort{}
			for _, port := range container.Ports {
				protocol, err := methodk8s.NewProtocolTypesFromString(string(port.Protocol))
				if err != nil {
					errors = append(errors, err.Error())
					protocol, _ = methodk8s.NewProtocolTypesFromString("UNDEFINED")
				}

				portInfo := methodk8s.ContainerPort{
					Port:     int(port.ContainerPort),
					Protocol: protocol,
				}
				ports = append(ports, &portInfo)
			}

			securityContext := methodk8s.SecurityContext{
				RunAsRoot:                container.SecurityContext != nil && container.SecurityContext.RunAsUser != nil && *container.SecurityContext.RunAsUser == 0,
				AllowPrivilegeEscalation: container.SecurityContext != nil && container.SecurityContext.AllowPrivilegeEscalation != nil && *container.SecurityContext.AllowPrivilegeEscalation,
				ReadOnlyRootFilesystem:   container.SecurityContext != nil && container.SecurityContext.ReadOnlyRootFilesystem != nil && *container.SecurityContext.ReadOnlyRootFilesystem,
			}

			containerInfo := methodk8s.Container{
				Name:            container.Name,
				Image:           container.Image,
				Ports:           ports,
				SecurityContext: &securityContext,
			}
			containers = append(containers, &containerInfo)
		}

		status, err := methodk8s.NewStatusTypesFromString(strings.ToUpper(string(pod.Status.Phase)))
		if err != nil {
			errors = append(errors, err.Error())
			status, _ = methodk8s.NewStatusTypesFromString("UNKNOWN")
		}
		statusInfo := methodk8s.Status{
			Status: status,
			PodIp:  pod.Status.PodIP,
			HostIp: pod.Status.HostIP,
		}

		podInfo := methodk8s.Pod{
			Name:       pod.GetName(),
			Namespace:  pod.GetNamespace(),
			Version:    pod.GetResourceVersion(),
			Status:     &statusInfo,
			Node:       pod.Spec.NodeName,
			Containers: containers,
		}
		pods = append(pods, &podInfo)
	}

	resources = methodk8s.PodReport{
		Pods:       pods,
		ClusterUrl: &config.Host,
		Errors:     errors,
	}

	return &resources, nil
}
