package ingress

import (
	"context"

	methodk8s "github.com/method-security/methodk8s/generated/go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	gatewayclientset "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

func EnumerateIngresses(k8config *rest.Config, onlyGateway bool) (*methodk8s.IngressReport, error) {
	resources := methodk8s.IngressReport{}
	errors := []string{}
	config := k8config

	clientset, err := gatewayclientset.NewForConfig(config)
	if err != nil {
		errors = append(errors, err.Error())
		return &methodk8s.IngressReport{Errors: errors}, err
	}

	gatewayList, err := clientset.GatewayV1beta1().Gateways("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return &methodk8s.IngressReport{Errors: errors}, err
	}

	gateways := []*methodk8s.Gateway{}
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

		gatewayInfo := methodk8s.Gateway{
			Name:      gateway.GetName(),
			Namespace: gateway.GetNamespace(),
			Listeners: listeners,
		}
		gateways = append(gateways, &gatewayInfo)
	}

	// '--gateway' flag not set
	ingresses := []*methodk8s.Ingress{}
	if !onlyGateway {
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			errors = append(errors, err.Error())
			return &methodk8s.IngressReport{Errors: errors}, err
		}

		ingressList, err := clientset.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errors = append(errors, err.Error())
			return &methodk8s.IngressReport{Errors: errors}, err
		}

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

			ingressInfo := methodk8s.Ingress{
				Name:      ingress.GetName(),
				Namespace: ingress.GetNamespace(),
				Rules:     rules,
			}
			ingresses = append(ingresses, &ingressInfo)
		}
	}

	resources = methodk8s.IngressReport{
		Gateways:  gateways,
		Ingresses: ingresses,
		Errors:    errors,
	}

	return &resources, nil
}
