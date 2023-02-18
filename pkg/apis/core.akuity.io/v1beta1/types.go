package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	NamespaceClassResourceName          = "namespaceclass"
	NamespaceClassResourcePlural        = "namespaceclasses"
	NamespaceClassShortName             = "nsc"
	GroupName                    string = "core.akuity.io"
	Name                         string = NamespaceClassResourcePlural + "." + GroupName
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespaceClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespaceClassSpec `json:"spec,omitempty"`
}

type NamespaceClassSpec struct {
	Resources string `json:"resources"`
}

type NamespaceClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceClass `json:"items"`
}
