package service

import (
	"context"

	methodk8s "github.com/method-security/methodk8s/generated/go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func EnumerateServices(ctx context.Context, k8config *rest.Config) (*methodk8s.ServiceReport, error) {
	resources := methodk8s.ServiceReport{}
	errors := []string{}
	config := k8config

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		errors = append(errors, err.Error())
		return &methodk8s.ServiceReport{Errors: errors}, err
	}

	servicesList, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return &methodk8s.ServiceReport{Errors: errors}, err
	}

	services := []*methodk8s.Service{}
	for _, service := range servicesList.Items {
		ports := []*methodk8s.ServicePort{}
		for _, port := range service.Spec.Ports {
			protocol, err := methodk8s.NewProtocolTypesFromString(string(port.Protocol))
			if err != nil {
				errors = append(errors, err.Error())
				protocol, _ = methodk8s.NewProtocolTypesFromString("UNDEFINED")
			}

			portInfo := methodk8s.ServicePort{
				Name:       port.Name,
				Protocol:   protocol,
				Port:       int(port.Port),
				TargetPort: port.TargetPort.String(),
			}
			ports = append(ports, &portInfo)
		}

		serviceInfo := methodk8s.Service{
			Name:      service.GetName(),
			Namespace: service.GetNamespace(),
			Type:      string(service.Spec.Type),
			ManagedBy: service.GetLabels()["app.kubernetes.io/managed-by"],
			Ports:     ports,
		}
		services = append(services, &serviceInfo)
	}

	resources = methodk8s.ServiceReport{
		Services: services,
		Errors:   errors,
	}

	return &resources, nil
}
