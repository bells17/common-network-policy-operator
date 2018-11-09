// MIT LICENSE

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CommonNetworkPolicySpec defines the desired state of CommonNetworkPolicy
type CommonNetworkPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// CommonNetworkPolicyStatus defines the observed state of CommonNetworkPolicy
type CommonNetworkPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// CommonNetworkPolicy is the Schema for the commonnetworkpolicies API
// +k8s:openapi-gen=true
type CommonNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CommonNetworkPolicySpec   `json:"spec,omitempty"`
	Status CommonNetworkPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// CommonNetworkPolicyList contains a list of CommonNetworkPolicy
type CommonNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CommonNetworkPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CommonNetworkPolicy{}, &CommonNetworkPolicyList{})
}
