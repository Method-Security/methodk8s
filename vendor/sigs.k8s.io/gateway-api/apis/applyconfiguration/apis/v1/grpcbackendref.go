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
	apisv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// GRPCBackendRefApplyConfiguration represents an declarative configuration of the GRPCBackendRef type for use
// with apply.
type GRPCBackendRefApplyConfiguration struct {
	BackendRefApplyConfiguration `json:",inline"`
	Filters                      []GRPCRouteFilterApplyConfiguration `json:"filters,omitempty"`
}

// GRPCBackendRefApplyConfiguration constructs an declarative configuration of the GRPCBackendRef type for use with
// apply.
func GRPCBackendRef() *GRPCBackendRefApplyConfiguration {
	return &GRPCBackendRefApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *GRPCBackendRefApplyConfiguration) WithGroup(value apisv1.Group) *GRPCBackendRefApplyConfiguration {
	b.Group = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *GRPCBackendRefApplyConfiguration) WithKind(value apisv1.Kind) *GRPCBackendRefApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *GRPCBackendRefApplyConfiguration) WithName(value apisv1.ObjectName) *GRPCBackendRefApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *GRPCBackendRefApplyConfiguration) WithNamespace(value apisv1.Namespace) *GRPCBackendRefApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *GRPCBackendRefApplyConfiguration) WithPort(value apisv1.PortNumber) *GRPCBackendRefApplyConfiguration {
	b.Port = &value
	return b
}

// WithWeight sets the Weight field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Weight field is set to the value of the last call.
func (b *GRPCBackendRefApplyConfiguration) WithWeight(value int32) *GRPCBackendRefApplyConfiguration {
	b.Weight = &value
	return b
}

// WithFilters adds the given value to the Filters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Filters field.
func (b *GRPCBackendRefApplyConfiguration) WithFilters(values ...*GRPCRouteFilterApplyConfiguration) *GRPCBackendRefApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFilters")
		}
		b.Filters = append(b.Filters, *values[i])
	}
	return b
}
