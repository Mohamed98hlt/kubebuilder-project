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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MaPageWebSpec defines the desired state of MaPageWeb
type MaPageWebSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
    Application string `json:"application,omitempy"`
    Contenu string `json:"content,omitempty"` 
	// +optional
    Pref  Case `json:"pref,omitempty"`
	

}


// +kubebuilder:validation:Enum=Critique;Important;Normal
type Case string

const (

    CriticalCase Case = "Critique"

    ImportantCase Case = "Important"

    NormalCase Case = "Normal"
)

// MaPageWebStatus defines the observed state of MaPageWeb
type MaPageWebStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MaPageWeb is the Schema for the mapagewebs API
type MaPageWeb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MaPageWebSpec   `json:"spec,omitempty"`
	Status MaPageWebStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MaPageWebList contains a list of MaPageWeb
type MaPageWebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaPageWeb `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MaPageWeb{}, &MaPageWebList{})
}
