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

type Enable struct {
	Title               string            `json:"title,omitempty"`
	ActionLabel         string            `json:"actionLabel,omitempty"`
	Description         string            `json:"description,omitempty"`
	Variables           map[string]string `json:"variables,omitempty"`
	VariableDisplayText map[string]string `json:"variableDisplayText,omitempty"`
	VariableHelpText    map[string]string `json:"variableHelpText,omitempty"`
	ValidationSecret    string            `json:"validationSecret,omitempty"`
	ValidationJob       string            `json:"validationJob,omitempty"`
	ValidationConfigMap string            `json:"validationConfigMap,omitempty"`
}

// ConsoleApplicationSpec defines the desired state of ConsoleApplication
type ConsoleApplicationSpec struct {
	DisplayName       string   `json:"displayName,omitempty"`
	Provider          string   `json:"provider,omitempty"`
	Description       string   `json:"description,omitempty"`
	Route             string   `json:"route,omitempty"`
	RouteNamespace    string   `json:"routeNamespace,omitempty"`
	RouteSuffix       string   `json:"routeSuffix,omitempty"`
	ServiceName       string   `json:"serviceName,omitempty"`
	EndPoint          string   `json:"endPoint,omitempty"`
	Link              string   `json:"link,omitempty"`
	Img               string   `json:"img,omitempty"`
	DocsLink          string   `json:"docsLink,omitempty"`
	GetStartedLink    string   `json:"getStartedLink,omitempty"`
	Category          string   `json:"category,omitempty"`
	Support           string   `json:"support,omitempty"`
	QuickStart        string   `json:"quickStart,omitempty"`
	ComingSoon        string   `json:"comingSoon,omitempty"`
	IsEnabled         *bool    `json:"isEnabled,omitempty"`
	KfdefApplications []string `json:"kfdefApplications"`
	CsvName           string   `json:"csvName,omitempty"`
	Enable            Enable   `json:"enable,omitempty"`
	FeatureFlag       string   `json:"featureFlag,omitempty"`
}

// ConsoleApplicationStatus defines the observed state of ConsoleApplication
type ConsoleApplicationStatus struct {
	Enabled *bool `json:"enabled,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConsoleApplication is the Schema for the consoleapplications API
type ConsoleApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsoleApplicationSpec   `json:"spec,omitempty"`
	Status ConsoleApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConsoleApplicationList contains a list of ConsoleApplication
type ConsoleApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsoleApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConsoleApplication{}, &ConsoleApplicationList{})
}
