package ingress

import (
	"context"

	methodk8s "github.com/method-security/methodk8s/generated/go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	gatewayclientset "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

type K8Resources struct {
	Ingress *methodk8s.Ingress `json:"ingress" yaml:"ingress"`
}

type K8ResourceReport struct {
	Resources K8Resources `json:"resources" yaml:"resources"`
	Errors    []string    `json:"errors" yaml:"errors"`
}

func EnumerateIngresses(k8config *rest.Config, onlyGateway bool) (*K8ResourceReport, error) {
	resources := K8Resources{}
	errors := []string{}
	config := k8config

	clientset, err := gatewayclientset.NewForConfig(config)
	if err != nil {
		errors = append(errors, err.Error())
		return &K8ResourceReport{Errors: errors}, err
	}

	gatewayList, err := clientset.GatewayV1beta1().Gateways("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		errors = append(errors, err.Error())
		return &K8ResourceReport{Errors: errors}, err
	}

	ingress := &methodk8s.Ingress{}
	gateways := []*methodk8s.GatewayInfo{}
	for _, gateway := range gatewayList.Items {
		listeners := []*methodk8s.Listener{}
		for _, listener := range gateway.Spec.Listeners {
			protocol, err := methodk8s.NewProtocolTypesFromString(string(listener.Protocol))

			if err != nil {
				errors = append(errors, err.Error())
				protocol, _ = methodk8s.NewProtocolTypesFromString("UNKNOWN")
			}
			listenerInfo := methodk8s.Listener{
				Name:     string(listener.Name),
				Port:     int(listener.Port),
				Protocol: protocol,
			}
			listeners = append(listeners, &listenerInfo)
		}

		gatewayInfo := methodk8s.GatewayInfo{
			Name:      gateway.GetName(),
			Namespace: gateway.GetNamespace(),
			Listeners: listeners,
		}
		gateways = append(gateways, &gatewayInfo)
	}

	if gateways != nil {
		ingress.Gateways = gateways
	}

	// '--gateway' flag not set
	if !onlyGateway {
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			errors = append(errors, err.Error())
			return &K8ResourceReport{Errors: errors}, err
		}

		ingressList, err := clientset.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errors = append(errors, err.Error())
			return &K8ResourceReport{Errors: errors}, err
		}

		ingresses := []*methodk8s.IngressInfo{}
		for _, ingress := range ingressList.Items {
			rules := []*methodk8s.Rule{}
			for _, rule := range ingress.Spec.Rules {
				for _, path := range rule.HTTP.Paths {
					ruleInfo := methodk8s.Rule{
						Host:        rule.Host,
						Path:        path.Path,
						ServiceName: path.Backend.Service.Name,
						ServicePort: int(path.Backend.Service.Port.Number),
					}
					rules = append(rules, &ruleInfo)
				}
			}

			ingressInfo := methodk8s.IngressInfo{
				Name:      ingress.GetName(),
				Namespace: ingress.GetNamespace(),
				Rules:     rules,
			}
			ingresses = append(ingresses, &ingressInfo)
		}
		if ingresses != nil {
			ingress.Ingresses = ingresses
		}
	}

	resources.Ingress = ingress

	k8ResourceReport := K8ResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &k8ResourceReport, nil
}
