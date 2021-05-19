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

package v1

import (
	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RhodsQuickStartSpec defines the desired state of RhodsQuickStart
type RhodsQuickStartSpec struct {
	// displayName is the display name of the Quick Start.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +required
	DisplayName string `json:"displayName"`
	// icon is a base64 encoded image that will be displayed beside the Quick Start display name.
	// The icon should be an vector image for easy scaling. The size of the icon should be 40x40.
	// +optional
	Icon string `json:"icon,omitempty"`
	// tags is a list of strings that describe the Quick Start.
	// +optional
	Tags []string `json:"tags,omitempty"`
	// durationMinutes describes approximately how many minutes it will take to complete the Quick Start.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +required
	DurationMinutes int `json:"durationMinutes"`
	// description is the description of the Quick Start. (includes markdown)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=256
	// +required
	Description string `json:"description"`
	// prerequisites contains all prerequisites that need to be met before taking a Quick Start. (includes markdown)
	// +optional
	Prerequisites []string `json:"prerequisites,omitempty"`
	// introduction describes the purpose of the Quick Start. (includes markdown)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +required
	Introduction string `json:"introduction"`
	// tasks is the list of steps the user has to perform to complete the Quick Start.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	// +required
	Tasks []RhodsQuickStartTask `json:"tasks"`
	// conclusion sums up the Quick Start and suggests the possible next steps. (includes markdown)
	// +optional
	Conclusion string `json:"conclusion,omitempty"`
	// nextQuickStart is a list of the following Quick Starts, suggested for the user to try.
	// +optional
	NextQuickStart []string `json:"nextQuickStart,omitempty"`
	// accessReviewResources contains a list of resources that the user's access
	// will be reviewed against in order for the user to complete the Quick Start.
	// The Quick Start will be hidden if any of the access reviews fail.
	// +optional
	AccessReviewResources []authorizationv1.ResourceAttributes `json:"accessReviewResources,omitempty"`
}

// RhodsQuickStartTask is a single step in a Quick Start.
type RhodsQuickStartTask struct {
	// title describes the task and is displayed as a step heading.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +required
	Title string `json:"title"`
	// description describes the steps needed to complete the task. (includes markdown)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +required
	Description string `json:"description"`
	// review contains instructions to validate the task is complete. The user will select 'Yes' or 'No'.
	// using a radio button, which indicates whether the step was completed successfully.
	// +optional
	Review *RhodsQuickStartTaskReview `json:"review,omitempty"`
	// summary contains information about the passed step.
	// +optional
	Summary *RhodsQuickStartTaskSummary `json:"summary,omitempty"`
}

// RhodsQuickStartTaskReview contains instructions that validate a task was completed successfully.
type RhodsQuickStartTaskReview struct {
	// instructions contains steps that user needs to take in order
	// to validate his work after going through a task. (includes markdown)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +required
	Instructions string `json:"instructions"`
	// failedTaskHelp contains suggestions for a failed task review and is shown at the end of task. (includes markdown)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +required
	FailedTaskHelp string `json:"failedTaskHelp"`
}

// RhodsQuickStartTaskSummary contains information about a passed step.
type RhodsQuickStartTaskSummary struct {
	// success describes the succesfully passed task.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +required
	Success string `json:"success"`
	// failed briefly describes the unsuccessfully passed task. (includes markdown)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=128
	// +required
	Failed string `json:"failed"`
}

// RhodsQuickStartStatus defines the observed state of RhodsQuickStart
type RhodsQuickStartStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RhodsQuickStart is the Schema for the rhodsquickstarts API
type RhodsQuickStart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RhodsQuickStartSpec   `json:"spec,omitempty"`
	Status RhodsQuickStartStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RhodsQuickStartList contains a list of RhodsQuickStart
type RhodsQuickStartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RhodsQuickStart `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RhodsQuickStart{}, &RhodsQuickStartList{})
}
