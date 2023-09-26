/*
Copyright 2019, 2021, 2022, 2023 The Multi-Cluster App Dispatcher Authors.

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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
)

// PendingPodSpecApplyConfiguration represents an declarative configuration of the PendingPodSpec type for use
// with apply.
type PendingPodSpecApplyConfiguration struct {
	PodName    *string           `json:"podname,omitempty"`
	Conditions []v1.PodCondition `json:"conditions,omitempty"`
}

// PendingPodSpecApplyConfiguration constructs an declarative configuration of the PendingPodSpec type for use with
// apply.
func PendingPodSpec() *PendingPodSpecApplyConfiguration {
	return &PendingPodSpecApplyConfiguration{}
}

// WithPodName sets the PodName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodName field is set to the value of the last call.
func (b *PendingPodSpecApplyConfiguration) WithPodName(value string) *PendingPodSpecApplyConfiguration {
	b.PodName = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *PendingPodSpecApplyConfiguration) WithConditions(values ...v1.PodCondition) *PendingPodSpecApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}
