/*
Copyright 2021.

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

type ConsoleDocumentType string

const (
	Documentation ConsoleDocumentType = "documentation"
	HowTo         ConsoleDocumentType = "how-to"
	QuickStart    ConsoleDocumentType = "quickstart"
	Tutorial      ConsoleDocumentType = "tutorial"
)

// ConsoleDocumentSpec defines the desired state of ConsoleDocument
type ConsoleDocumentSpec struct {
	Type            ConsoleDocumentType `json:"type,omitempty"`
	DisplayName     string              `json:"displayName,omitempty"`
	AppName         string              `json:"appName,omitempty"`
	Provider        string              `json:"provider,omitempty"`
	Description     string              `json:"description,omitempty"`
	Url             string              `json:"url,omitempty"`
	Img             string              `json:"img,omitempty"`
	Icon            string              `json:"icon,omitempty"`
	DurationMinutes int                 `json:"durationMinutes"`
	MarkDown        string              `json:"markDown,omitempty"`
	FeatureFlag     string              `json:"featureFlag,omitempty"`
}

// ConsoleDocumentStatus defines the observed state of ConsoleDocument
type ConsoleDocumentStatus struct {
	Enabled *bool `json:"enabled,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConsoleDocument is the Schema for the consoledocuments API
type ConsoleDocument struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsoleDocumentSpec   `json:"spec,omitempty"`
	Status ConsoleDocumentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConsoleDocumentList contains a list of ConsoleDocument
type ConsoleDocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsoleDocument `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConsoleDocument{}, &ConsoleDocumentList{})
}
