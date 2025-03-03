/*
Copyright 2025.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StandaloneSpec defines the desired state of Standalone.
type StandaloneSpec struct {
	Auth          *AuthSettings        `json:"auth"`
	Metrics       *MetricsSettings     `json:"metrics"`
	Image         string               `json:"image"`
	Volumes       []Volume             `json:"volumes"`
	Replicas      int32                `json:"replicas"`
	Config        string               `json:"config"`
	Command       []string             `json:"command,omitempty"`
	Args          []string             `json:"args,omitempty"`
	Port          int32                `json:"port"`
	Env           []corev1.EnvVar      `json:"env,omitempty"`
	VolumeMounts  []corev1.VolumeMount `json:"volumeMounts,omitempty"`
	OtherSettings `json:",inline"`
}

// StandaloneStatus defines the observed state of Standalone.
type StandaloneStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Standalone is the Schema for the standalones API.
type Standalone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StandaloneSpec   `json:"spec,omitempty"`
	Status StandaloneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StandaloneList contains a list of Standalone.
type StandaloneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Standalone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Standalone{}, &StandaloneList{})
}
