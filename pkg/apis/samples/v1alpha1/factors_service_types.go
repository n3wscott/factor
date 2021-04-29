/*
Copyright 2019 The Knative Authors

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

// Factor is a Knative abstraction that encapsulates the interface by which Knative
// components express a desire to have a particular image cached.
//
// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Factor struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the Factor (from the client).
	// +optional
	Spec FactorSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the Factor (from the controller).
	// +optional
	Status FactorStatus `json:"status,omitempty"`
}

var (
	// Check that Factor can be validated and defaulted.
	_ apis.Validatable   = (*Factor)(nil)
	_ apis.Defaultable   = (*Factor)(nil)
	_ kmeta.OwnerRefable = (*Factor)(nil)
	// Check that the type conforms to the duck Knative Resource shape.
	_ duckv1.KRShaped = (*Factor)(nil)
)

// FactorSpec holds the desired state of the Factor (from the client).
type FactorSpec struct {
	Base int `json:"base"` // !{base}   !5 = 5+4+3+2+1
}

const (
	// FactorConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	FactorConditionReady = apis.ConditionReady
)

// FactorStatus communicates the observed state of the Factor (from the controller).
type FactorStatus struct {
	duckv1.Status `json:",inline"`

	Expression string `json:"expression"`
	Result int `json:"result"`
}

// FactorList is a list of Factor resources
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type FactorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Factor `json:"items"`
}

// GetStatus retrieves the status of the resource. Implements the KRShaped interface.
func (as *Factor) GetStatus() *duckv1.Status {
	return &as.Status.Status
}
