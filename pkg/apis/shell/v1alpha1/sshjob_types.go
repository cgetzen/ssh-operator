package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SSHJobSpec defines the desired state of SSHJob
type SSHJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
  Namespace string `json:"namespace"`
  Pod       string `json:"pod"`
  Container string `json:"container"`
}

// SSHJobStatus defines the observed state of SSHJob
type SSHJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
  Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SSHJob is the Schema for the sshjobs API
// +k8s:openapi-gen=true
type SSHJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SSHJobSpec   `json:"spec,omitempty"`
	Status SSHJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SSHJobList contains a list of SSHJob
type SSHJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SSHJob{}, &SSHJobList{})
}
