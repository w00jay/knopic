/*
Linstor Operator
Copyright 2019 LINBIT USA, LLC.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LinstorNodeSetSpec defines the desired state of LinstorNodeSet
type LinstorNodeSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// StoragePools is a list of StoragePools for LinstorNodeSet to manage.
	StoragePools *StoragePools `json:"storagePools"`

	//DisableDRBDKernelModuleInjection turns off automatic injection of the DRBD
	// kernel module on the host system when set to true.
	DisableDRBDKernelModuleInjection bool `json:"disableDRBDKernelModuleInjection"`

	//DrbdRepoCred is the name of the k8s secret with the repo credential
	DrbdRepoCred string `json:"drbdRepoCred"`

	//LinstorSatImage & Version  is the LINSTOR Satellite image location and version/tag
	LinstorSatImage   string `json:"linstorSatImage"`
	LinstorSatVersion string `json:"linstorSatVersion"`

	//LinstorKernelModImage & Version  is the DRBD Kernel injection image location and version/tag
	LinstorKernelModImage   string `json:"linstorKernelModImage"`
	LinstorKernelModVersion string `json:"linstorKernelModVersion"`
}

// LinstorNodeSetStatus defines the observed state of LinstorNodeSet
type LinstorNodeSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Keep the first letters in the json for Errors lowercase and all other
	// statuses uppercase. This causes `kubectl get TYPE NAME -oyaml` to sort
	// errors to below the other potentially very long Statuses which is the preferred UX.

	// Errors remaining that will trigger reconciliations.
	Errors []string `json:"errors"`
	// SatelliteStatuses by hostname.
	SatelliteStatuses map[string]*SatelliteStatus `json:"SatelliteStatuses"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinstorNodeSet is the Schema for the linstornodesets API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=linstornodesets,scope=Namespaced
type LinstorNodeSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LinstorNodeSetSpec   `json:"spec,omitempty"`
	Status LinstorNodeSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinstorNodeSetList contains a list of LinstorNodeSet
type LinstorNodeSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinstorNodeSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LinstorNodeSet{}, &LinstorNodeSetList{})
}
