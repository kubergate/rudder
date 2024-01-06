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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RudderSpec defines the desired state of Helm
type RudderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	GatewayController *GatewayController `json:"gatewayController,omitempty"`

	XdsServerConfig *XdsServerConfig `json:"xdsServerConfig,omitempty"`

	KubernetesWatchConfig *KubernetesWatchConfig `json:"kubernetesWatchConfig,omitempty"`

	DataStoreConfig *DataStoreConfig `json:"dataStoreConfig,omitempty"`
}

// RudderStatus defines the observed state of Helm
type RudderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Rudder is the Schema for the helms API
type Rudder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RudderSpec   `json:"spec,omitempty"`
	Status RudderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RudderList contains a list of Helm
type RudderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rudder `json:"items"`
}

type GatewayController struct {
	ControllerNames []string `json:"controllerNames,omitempty"`
}

type XdsServerConfig struct {
	XdsMode string `json:"xdsMode,omitempty"`
}

type KubernetesWatchConfig struct {
	Namespaces []string `json:"namespaces,omitempty"`
}

type DataStoreConfig struct {
	DBPath  string `json:"dbPath,omitempty"`
	Timeout int32  `json:"timeout,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Rudder{}, &RudderList{})
}
