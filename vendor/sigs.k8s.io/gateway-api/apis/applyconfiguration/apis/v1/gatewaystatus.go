/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// GatewayStatusApplyConfiguration represents an declarative configuration of the GatewayStatus type for use
// with apply.
type GatewayStatusApplyConfiguration struct {
	Addresses  []GatewayStatusAddressApplyConfiguration `json:"addresses,omitempty"`
	Conditions []metav1.ConditionApplyConfiguration     `json:"conditions,omitempty"`
	Listeners  []ListenerStatusApplyConfiguration       `json:"listeners,omitempty"`
}

// GatewayStatusApplyConfiguration constructs an declarative configuration of the GatewayStatus type for use with
// apply.
func GatewayStatus() *GatewayStatusApplyConfiguration {
	return &GatewayStatusApplyConfiguration{}
}

// WithAddresses adds the given value to the Addresses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Addresses field.
func (b *GatewayStatusApplyConfiguration) WithAddresses(values ...*GatewayStatusAddressApplyConfiguration) *GatewayStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAddresses")
		}
		b.Addresses = append(b.Addresses, *values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *GatewayStatusApplyConfiguration) WithConditions(values ...*metav1.ConditionApplyConfiguration) *GatewayStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithListeners adds the given value to the Listeners field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Listeners field.
func (b *GatewayStatusApplyConfiguration) WithListeners(values ...*ListenerStatusApplyConfiguration) *GatewayStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithListeners")
		}
		b.Listeners = append(b.Listeners, *values[i])
	}
	return b
}