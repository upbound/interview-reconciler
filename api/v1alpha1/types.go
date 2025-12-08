package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PostgresConnectionSpec defines the desired state of PostgresConnection.
type PostgresConnectionSpec struct {
	// Host is the PostgreSQL server hostname.
	// +kubebuilder:validation:Required
	Host string `json:"host"`

	// Port is the PostgreSQL server port.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port"`

	// Database is the name of the database to connect to.
	// +kubebuilder:validation:Required
	Database string `json:"database"`

	// Credentials contains the reference to the secret containing database
	// credentials.
	// +kubebuilder:validation:Required
	Credentials CredentialsSpec `json:"credentials"`
}

// CredentialsSpec defines the credentials source.
type CredentialsSpec struct {
	// SecretRef is a reference to a Secret containing the credentials.
	// +kubebuilder:validation:Required
	SecretRef SecretReference `json:"secretRef"`
}

// SecretReference contains the reference to a Secret.
type SecretReference struct {
	// Name is the name of the secret.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

// PostgresConnectionStatus defines the observed state of PostgresConnection.
type PostgresConnectionStatus struct {
	// Conditions represent the latest available observations of the
	// PostgresConnection's state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// SecretRef contains the reference to the Secret created by the controller.
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty"`
}

// PostgresConnection is the Schema for the postgresconnections API.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=pgconn
// +kubebuilder:printcolumn:name="Host",type=string,JSONPath=`.spec.host`
// +kubebuilder:printcolumn:name="Database",type=string,JSONPath=`.spec.database`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
type PostgresConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgresConnectionSpec   `json:"spec"`
	Status PostgresConnectionStatus `json:"status,omitempty"`
}

// PostgresConnectionList contains a list of PostgresConnection.
// +kubebuilder:object:root=true
type PostgresConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PostgresConnection{}, &PostgresConnectionList{})
}
