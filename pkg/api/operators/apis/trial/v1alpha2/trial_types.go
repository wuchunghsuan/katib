/*

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

package v1alpha2

import (
	common "github.com/kubeflow/katib/pkg/api/operators/apis/common/v1alpha2"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TrialSpec struct {
	// Describes the objective of the experiment.
	Objective *common.ObjectiveSpec `json:"objective,omitempty"`

	// Key-value pairs for hyperparameters and assignment values.
	ParameterAssignments []common.ParameterAssignment `json:"parameterAssignments"`

	// Raw text for the trial run spec. This can be any generic Kubernetes
	// runtime object. The trial operator should create the resource as written,
	// and let the corresponding resource controller (e.g. tf-operator) handle
	// the rest.
	RunSpec string `json:"runSpec,omitempty"`
	// Whether to retain the trial run object after completed.
	RetainRun bool `json:"retainRun,omitempty"`

	// Raw text for the metrics collector spec. This must be a CronJob object.
	// Deprecated
	MetricsCollectorSpec string `json:"metricsCollectorSpec,omitempty"`
	// Whether to retain the trial metrics collector CronJob object after completed.
	// Deprecated
	RetainMetricsCollector bool `json:"retainMetricsCollector,omitempty"`
}

type TrialStatus struct {
	// Represents time when the Trial was acknowledged by the Trial controller.
	// It is not guaranteed to be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// Represents time when the Trial was completed. It is not guaranteed to
	// be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`

	// Represents last time when the Trial was reconciled. It is not guaranteed to
	// be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	LastReconcileTime *metav1.Time `json:"lastReconcileTime,omitempty"`

	// List of observed runtime conditions for this Trial.
	Conditions []TrialCondition `json:"conditions,omitempty"`

	// Results of the Trial - objectives and other metrics values.
	Observation *common.Observation `json:"observation,omitempty"`
}

// +k8s:deepcopy-gen=true
// TrialCondition describes the state of the trial at a certain point.
type TrialCondition struct {
	// Type of trial condition.
	Type TrialConditionType `json:"type"`

	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`

	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`

	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// TrialConditionType describes the various conditions a Trial can be in.
type TrialConditionType string

const (
	TrialCreated   TrialConditionType = "Created"
	TrialRunning   TrialConditionType = "Running"
	TrialSucceeded TrialConditionType = "Succeeded"
	TrialKilled    TrialConditionType = "Killed"
	TrialFailed    TrialConditionType = "Failed"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Represents the structure of a Trial resource.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Trial struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrialSpec   `json:"spec,omitempty"`
	Status TrialStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrialList contains a list of Trials
type TrialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trial `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Trial{}, &TrialList{})
}
