/*

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
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CommonNetworkPolicySpec defines the desired state of CommonNetworkPolicy
type CommonNetworkPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	NamePrefix        string                         `json:"namePrefix,omitempty"`
	ExcludeNamespaces []string                       `json:"excludeNamespaces,omitempty"`
	PolicySpec        networkingv1.NetworkPolicySpec `json:"policySpec"`
}

// CommonNetworkPolicyStatus defines the observed state of CommonNetworkPolicy
type CommonNetworkPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// CommonNetworkPolicy is the Schema for the commonnetworkpolicies API
type CommonNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CommonNetworkPolicySpec   `json:"spec,omitempty"`
	Status CommonNetworkPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CommonNetworkPolicyList contains a list of CommonNetworkPolicy
type CommonNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CommonNetworkPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CommonNetworkPolicy{}, &CommonNetworkPolicyList{})
}
