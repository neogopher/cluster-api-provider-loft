/*
Copyright 2023.

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
	clusterv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VirtualClusterInstanceSpec defines the desired state of VirtualClusterInstance
type VirtualClusterInstanceSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.
	// +optional
	ControlPlaneEndpoint clusterv1beta1.APIEndpoint `json:"controlPlaneEndpoint"`

	// Project is the loft project under which the virtualclusterinstance will be created.
	// +optional
	Project string `json:"project"`
}

// VirtualClusterInstanceStatus defines the observed state of VirtualClusterInstance
type VirtualClusterInstanceStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Ready defines if the virtual cluster instance is ready.
	// +optional
	Ready bool `json:"ready"`

	// Phase describes the current phase the virtual cluster instance is in
	// +optional
	Phase VirtualClusterInstancePhase `json:"phase,omitempty"`

	// Reason describes the reason in machine readable form why the cluster is in the current
	// phase
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message describes the reason in human readable form why the cluster is in the currrent
	// phase
	// +optional
	Message string `json:"message,omitempty"`
}

// VirtualClusterInstancePhase describes the phase of a virtual cluster
type VirtualClusterInstancePhase string

// These are the valid admin account types
const (
	VirtualClusterUnknown  VirtualClusterInstancePhase = ""
	VirtualClusterPending  VirtualClusterInstancePhase = "Pending"
	VirtualClusterDeployed VirtualClusterInstancePhase = "Deployed"
	VirtualClusterFailed   VirtualClusterInstancePhase = "Failed"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VirtualClusterInstance is the Schema for the virtualclusterinstances API
type VirtualClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterInstanceSpec   `json:"spec,omitempty"`
	Status VirtualClusterInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VirtualClusterInstanceList contains a list of VirtualClusterInstance
type VirtualClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualClusterInstance{}, &VirtualClusterInstanceList{})
}
