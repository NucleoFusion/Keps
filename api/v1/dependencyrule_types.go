/*
Copyright 2026.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SubjectRef struct {
	// APIGroup of the subject resource (e.g. "apps").
	// Empty means core API group.
	// +optional
	APIGroup string `json:"apiGroup,omitempty"`

	// Kind of the subject resource (e.g. Deployment).
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`

	// Name selects a single object.
	// Mutually exclusive with Selector.
	// +optional
	Name string `json:"name,omitempty"`

	// Selector selects multiple objects.
	// Mutually exclusive with Name.
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
}

type DependencyRef struct {
	// APIGroup of the dependency resource.
	// Empty means core API group.
	// +optional
	APIGroup string `json:"apiGroup,omitempty"`

	// Kind of the dependency resource.
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`

	// Name of the dependency resource.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

type EnforcementMode string

const (
	EnforcementStrict EnforcementMode = "Strict"
	EnforcementWarn   EnforcementMode = "Warn"
)

// DependencyRuleSpec defines the desired state of DependencyRule
type DependencyRuleSpec struct {
	// Subject defines the resource(s) this rule applies to.
	// +kubebuilder:validation:Required
	Subject SubjectRef `json:"subject"`

	// DependsOn lists resources that must exist before the subject is created or updated.
	// +kubebuilder:validation:MinItems=1
	DependsOn []DependencyRef `json:"dependsOn"`

	// Enforcement determines how violations are handled.
	// +kubebuilder:validation:Enum=Strict;Warn
	// +kubebuilder:default=Strict
	Enforcement EnforcementMode `json:"enforcement,omitempty"`
}

// DependencyRuleStatus defines the observed state of DependencyRule.
type DependencyRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// For Kubernetes API conventions, see:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	// conditions represent the current state of the DependencyRule resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DependencyRule is the Schema for the dependencyrules API
type DependencyRule struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of DependencyRule
	// +required
	Spec DependencyRuleSpec `json:"spec"`

	// status defines the observed state of DependencyRule
	// +optional
	Status DependencyRuleStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// DependencyRuleList contains a list of DependencyRule
type DependencyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []DependencyRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DependencyRule{}, &DependencyRuleList{})
}
