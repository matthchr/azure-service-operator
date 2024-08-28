// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180301

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type MetricAlert_Spec_ARM struct {
	// Location: Resource location
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The alert rule properties of the resource.
	Properties *MetricAlertProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &MetricAlert_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-03-01"
func (alert MetricAlert_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (alert *MetricAlert_Spec_ARM) GetName() string {
	return alert.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/metricAlerts"
func (alert *MetricAlert_Spec_ARM) GetType() string {
	return "Microsoft.Insights/metricAlerts"
}

// An alert rule.
type MetricAlertProperties_ARM struct {
	// Actions: the array of actions that are performed when the alert rule becomes active, and when an alert condition is
	// resolved.
	Actions []MetricAlertAction_ARM `json:"actions,omitempty"`

	// AutoMitigate: the flag that indicates whether the alert should be auto resolved or not. The default is true.
	AutoMitigate *bool `json:"autoMitigate,omitempty"`

	// Criteria: defines the specific alert criteria information.
	Criteria *MetricAlertCriteria_ARM `json:"criteria,omitempty"`

	// Description: the description of the metric alert that will be included in the alert email.
	Description *string `json:"description,omitempty"`

	// Enabled: the flag that indicates whether the metric alert is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// EvaluationFrequency: how often the metric alert is evaluated represented in ISO 8601 duration format.
	EvaluationFrequency *string  `json:"evaluationFrequency,omitempty"`
	Scopes              []string `json:"scopes,omitempty"`

	// Severity: Alert severity {0, 1, 2, 3, 4}
	Severity *int `json:"severity,omitempty"`

	// TargetResourceRegion: the region of the target resource(s) on which the alert is created/updated. Mandatory if the scope
	// contains a subscription, resource group, or more than one resource.
	TargetResourceRegion *string `json:"targetResourceRegion,omitempty"`

	// TargetResourceType: the resource type of the target resource(s) on which the alert is created/updated. Mandatory if the
	// scope contains a subscription, resource group, or more than one resource.
	TargetResourceType *string `json:"targetResourceType,omitempty"`

	// WindowSize: the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the
	// threshold.
	WindowSize *string `json:"windowSize,omitempty"`
}

// An alert action.
type MetricAlertAction_ARM struct {
	// ActionGroupId: the id of the action group to use.
	ActionGroupId *string `json:"actionGroupId,omitempty"`

	// WebHookProperties: This field allows specifying custom properties, which would be appended to the alert payload sent as
	// input to the webhook.
	WebHookProperties map[string]string `json:"webHookProperties,omitempty"`
}

type MetricAlertCriteria_ARM struct {
	// MicrosoftAzureMonitorMultipleResourceMultipleMetric: Mutually exclusive with all other properties
	MicrosoftAzureMonitorMultipleResourceMultipleMetric *MetricAlertMultipleResourceMultipleMetricCriteria_ARM `json:"microsoftAzureMonitorMultipleResourceMultipleMetricCriteria,omitempty"`

	// MicrosoftAzureMonitorSingleResourceMultipleMetric: Mutually exclusive with all other properties
	MicrosoftAzureMonitorSingleResourceMultipleMetric *MetricAlertSingleResourceMultipleMetricCriteria_ARM `json:"microsoftAzureMonitorSingleResourceMultipleMetricCriteria,omitempty"`

	// MicrosoftAzureMonitorWebtestLocationAvailability: Mutually exclusive with all other properties
	MicrosoftAzureMonitorWebtestLocationAvailability *WebtestLocationAvailabilityCriteria_ARM `json:"microsoftAzureMonitorWebtestLocationAvailabilityCriteria,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because MetricAlertCriteria_ARM represents a discriminated union (JSON OneOf)
func (criteria MetricAlertCriteria_ARM) MarshalJSON() ([]byte, error) {
	if criteria.MicrosoftAzureMonitorMultipleResourceMultipleMetric != nil {
		return json.Marshal(criteria.MicrosoftAzureMonitorMultipleResourceMultipleMetric)
	}
	if criteria.MicrosoftAzureMonitorSingleResourceMultipleMetric != nil {
		return json.Marshal(criteria.MicrosoftAzureMonitorSingleResourceMultipleMetric)
	}
	if criteria.MicrosoftAzureMonitorWebtestLocationAvailability != nil {
		return json.Marshal(criteria.MicrosoftAzureMonitorWebtestLocationAvailability)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the MetricAlertCriteria_ARM
func (criteria *MetricAlertCriteria_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["odata.type"]
	if discriminator == "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria" {
		criteria.MicrosoftAzureMonitorMultipleResourceMultipleMetric = &MetricAlertMultipleResourceMultipleMetricCriteria_ARM{}
		return json.Unmarshal(data, criteria.MicrosoftAzureMonitorMultipleResourceMultipleMetric)
	}
	if discriminator == "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria" {
		criteria.MicrosoftAzureMonitorSingleResourceMultipleMetric = &MetricAlertSingleResourceMultipleMetricCriteria_ARM{}
		return json.Unmarshal(data, criteria.MicrosoftAzureMonitorSingleResourceMultipleMetric)
	}
	if discriminator == "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria" {
		criteria.MicrosoftAzureMonitorWebtestLocationAvailability = &WebtestLocationAvailabilityCriteria_ARM{}
		return json.Unmarshal(data, criteria.MicrosoftAzureMonitorWebtestLocationAvailability)
	}

	// No error
	return nil
}

type MetricAlertMultipleResourceMultipleMetricCriteria_ARM struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// AllOf: the list of multiple metric criteria for this 'all of' operation.
	AllOf []MultiMetricCriteria_ARM `json:"allOf,omitempty"`

	// OdataType: specifies the type of the alert criteria.
	OdataType MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_ARM `json:"odata.type,omitempty"`
}

type MetricAlertSingleResourceMultipleMetricCriteria_ARM struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// AllOf: The list of metric criteria for this 'all of' operation.
	AllOf []MetricCriteria_ARM `json:"allOf,omitempty"`

	// OdataType: specifies the type of the alert criteria.
	OdataType MetricAlertSingleResourceMultipleMetricCriteria_OdataType_ARM `json:"odata.type,omitempty"`
}

type WebtestLocationAvailabilityCriteria_ARM struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`
	ComponentId          *string            `json:"componentId,omitempty"`

	// FailedLocationCount: The number of failed locations.
	FailedLocationCount *float64 `json:"failedLocationCount,omitempty"`

	// OdataType: specifies the type of the alert criteria.
	OdataType WebtestLocationAvailabilityCriteria_OdataType_ARM `json:"odata.type,omitempty"`

	// WebTestId: The Application Insights web test Id.
	WebTestId *string `json:"webTestId,omitempty"`
}

// +kubebuilder:validation:Enum={"Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria"}
type MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_ARM string

const MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_ARM_MicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria = MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_ARM("Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria")

// Mapping from string to MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_ARM
var metricAlertMultipleResourceMultipleMetricCriteria_OdataType_ARM_Values = map[string]MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_ARM{
	"microsoft.azure.monitor.multipleresourcemultiplemetriccriteria": MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_ARM_MicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria,
}

// +kubebuilder:validation:Enum={"Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria"}
type MetricAlertSingleResourceMultipleMetricCriteria_OdataType_ARM string

const MetricAlertSingleResourceMultipleMetricCriteria_OdataType_ARM_MicrosoftAzureMonitorSingleResourceMultipleMetricCriteria = MetricAlertSingleResourceMultipleMetricCriteria_OdataType_ARM("Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria")

// Mapping from string to MetricAlertSingleResourceMultipleMetricCriteria_OdataType_ARM
var metricAlertSingleResourceMultipleMetricCriteria_OdataType_ARM_Values = map[string]MetricAlertSingleResourceMultipleMetricCriteria_OdataType_ARM{
	"microsoft.azure.monitor.singleresourcemultiplemetriccriteria": MetricAlertSingleResourceMultipleMetricCriteria_OdataType_ARM_MicrosoftAzureMonitorSingleResourceMultipleMetricCriteria,
}

type MetricCriteria_ARM struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// CriterionType: Specifies the type of threshold criteria
	CriterionType MetricCriteria_CriterionType_ARM `json:"criterionType,omitempty"`

	// Dimensions: List of dimension conditions.
	Dimensions []MetricDimension_ARM `json:"dimensions,omitempty"`

	// MetricName: Name of the metric.
	MetricName *string `json:"metricName,omitempty"`

	// MetricNamespace: Namespace of the metric.
	MetricNamespace *string `json:"metricNamespace,omitempty"`

	// Name: Name of the criteria.
	Name *string `json:"name,omitempty"`

	// Operator: the criteria operator.
	Operator *MetricCriteria_Operator_ARM `json:"operator,omitempty"`

	// SkipMetricValidation: Allows creating an alert rule on a custom metric that isn't yet emitted, by causing the metric
	// validation to be skipped.
	SkipMetricValidation *bool `json:"skipMetricValidation,omitempty"`

	// Threshold: the criteria threshold value that activates the alert.
	Threshold *float64 `json:"threshold,omitempty"`

	// TimeAggregation: the criteria time aggregation types.
	TimeAggregation *MetricCriteria_TimeAggregation_ARM `json:"timeAggregation,omitempty"`
}

type MultiMetricCriteria_ARM struct {
	// Dynamic: Mutually exclusive with all other properties
	Dynamic *DynamicMetricCriteria_ARM `json:"dynamicThresholdCriterion,omitempty"`

	// Static: Mutually exclusive with all other properties
	Static *MetricCriteria_ARM `json:"staticThresholdCriterion,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because MultiMetricCriteria_ARM represents a discriminated union (JSON OneOf)
func (criteria MultiMetricCriteria_ARM) MarshalJSON() ([]byte, error) {
	if criteria.Dynamic != nil {
		return json.Marshal(criteria.Dynamic)
	}
	if criteria.Static != nil {
		return json.Marshal(criteria.Static)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the MultiMetricCriteria_ARM
func (criteria *MultiMetricCriteria_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["criterionType"]
	if discriminator == "DynamicThresholdCriterion" {
		criteria.Dynamic = &DynamicMetricCriteria_ARM{}
		return json.Unmarshal(data, criteria.Dynamic)
	}
	if discriminator == "StaticThresholdCriterion" {
		criteria.Static = &MetricCriteria_ARM{}
		return json.Unmarshal(data, criteria.Static)
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria"}
type WebtestLocationAvailabilityCriteria_OdataType_ARM string

const WebtestLocationAvailabilityCriteria_OdataType_ARM_MicrosoftAzureMonitorWebtestLocationAvailabilityCriteria = WebtestLocationAvailabilityCriteria_OdataType_ARM("Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria")

// Mapping from string to WebtestLocationAvailabilityCriteria_OdataType_ARM
var webtestLocationAvailabilityCriteria_OdataType_ARM_Values = map[string]WebtestLocationAvailabilityCriteria_OdataType_ARM{
	"microsoft.azure.monitor.webtestlocationavailabilitycriteria": WebtestLocationAvailabilityCriteria_OdataType_ARM_MicrosoftAzureMonitorWebtestLocationAvailabilityCriteria,
}

type DynamicMetricCriteria_ARM struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// AlertSensitivity: The extent of deviation required to trigger an alert. This will affect how tight the threshold is to
	// the metric series pattern.
	AlertSensitivity *DynamicMetricCriteria_AlertSensitivity_ARM `json:"alertSensitivity,omitempty"`

	// CriterionType: Specifies the type of threshold criteria
	CriterionType DynamicMetricCriteria_CriterionType_ARM `json:"criterionType,omitempty"`

	// Dimensions: List of dimension conditions.
	Dimensions []MetricDimension_ARM `json:"dimensions,omitempty"`

	// FailingPeriods: The minimum number of violations required within the selected lookback time window required to raise an
	// alert.
	FailingPeriods *DynamicThresholdFailingPeriods_ARM `json:"failingPeriods,omitempty"`

	// IgnoreDataBefore: Use this option to set the date from which to start learning the metric historical data and calculate
	// the dynamic thresholds (in ISO8601 format)
	IgnoreDataBefore *string `json:"ignoreDataBefore,omitempty"`

	// MetricName: Name of the metric.
	MetricName *string `json:"metricName,omitempty"`

	// MetricNamespace: Namespace of the metric.
	MetricNamespace *string `json:"metricNamespace,omitempty"`

	// Name: Name of the criteria.
	Name *string `json:"name,omitempty"`

	// Operator: The operator used to compare the metric value against the threshold.
	Operator *DynamicMetricCriteria_Operator_ARM `json:"operator,omitempty"`

	// SkipMetricValidation: Allows creating an alert rule on a custom metric that isn't yet emitted, by causing the metric
	// validation to be skipped.
	SkipMetricValidation *bool `json:"skipMetricValidation,omitempty"`

	// TimeAggregation: the criteria time aggregation types.
	TimeAggregation *DynamicMetricCriteria_TimeAggregation_ARM `json:"timeAggregation,omitempty"`
}

// +kubebuilder:validation:Enum={"StaticThresholdCriterion"}
type MetricCriteria_CriterionType_ARM string

const MetricCriteria_CriterionType_ARM_StaticThresholdCriterion = MetricCriteria_CriterionType_ARM("StaticThresholdCriterion")

// Mapping from string to MetricCriteria_CriterionType_ARM
var metricCriteria_CriterionType_ARM_Values = map[string]MetricCriteria_CriterionType_ARM{
	"staticthresholdcriterion": MetricCriteria_CriterionType_ARM_StaticThresholdCriterion,
}

// +kubebuilder:validation:Enum={"Equals","GreaterThan","GreaterThanOrEqual","LessThan","LessThanOrEqual"}
type MetricCriteria_Operator_ARM string

const (
	MetricCriteria_Operator_ARM_Equals             = MetricCriteria_Operator_ARM("Equals")
	MetricCriteria_Operator_ARM_GreaterThan        = MetricCriteria_Operator_ARM("GreaterThan")
	MetricCriteria_Operator_ARM_GreaterThanOrEqual = MetricCriteria_Operator_ARM("GreaterThanOrEqual")
	MetricCriteria_Operator_ARM_LessThan           = MetricCriteria_Operator_ARM("LessThan")
	MetricCriteria_Operator_ARM_LessThanOrEqual    = MetricCriteria_Operator_ARM("LessThanOrEqual")
)

// Mapping from string to MetricCriteria_Operator_ARM
var metricCriteria_Operator_ARM_Values = map[string]MetricCriteria_Operator_ARM{
	"equals":             MetricCriteria_Operator_ARM_Equals,
	"greaterthan":        MetricCriteria_Operator_ARM_GreaterThan,
	"greaterthanorequal": MetricCriteria_Operator_ARM_GreaterThanOrEqual,
	"lessthan":           MetricCriteria_Operator_ARM_LessThan,
	"lessthanorequal":    MetricCriteria_Operator_ARM_LessThanOrEqual,
}

// +kubebuilder:validation:Enum={"Average","Count","Maximum","Minimum","Total"}
type MetricCriteria_TimeAggregation_ARM string

const (
	MetricCriteria_TimeAggregation_ARM_Average = MetricCriteria_TimeAggregation_ARM("Average")
	MetricCriteria_TimeAggregation_ARM_Count   = MetricCriteria_TimeAggregation_ARM("Count")
	MetricCriteria_TimeAggregation_ARM_Maximum = MetricCriteria_TimeAggregation_ARM("Maximum")
	MetricCriteria_TimeAggregation_ARM_Minimum = MetricCriteria_TimeAggregation_ARM("Minimum")
	MetricCriteria_TimeAggregation_ARM_Total   = MetricCriteria_TimeAggregation_ARM("Total")
)

// Mapping from string to MetricCriteria_TimeAggregation_ARM
var metricCriteria_TimeAggregation_ARM_Values = map[string]MetricCriteria_TimeAggregation_ARM{
	"average": MetricCriteria_TimeAggregation_ARM_Average,
	"count":   MetricCriteria_TimeAggregation_ARM_Count,
	"maximum": MetricCriteria_TimeAggregation_ARM_Maximum,
	"minimum": MetricCriteria_TimeAggregation_ARM_Minimum,
	"total":   MetricCriteria_TimeAggregation_ARM_Total,
}

// Specifies a metric dimension.
type MetricDimension_ARM struct {
	// Name: Name of the dimension.
	Name *string `json:"name,omitempty"`

	// Operator: the dimension operator. Only 'Include' and 'Exclude' are supported
	Operator *string `json:"operator,omitempty"`

	// Values: list of dimension values.
	Values []string `json:"values,omitempty"`
}

// +kubebuilder:validation:Enum={"High","Low","Medium"}
type DynamicMetricCriteria_AlertSensitivity_ARM string

const (
	DynamicMetricCriteria_AlertSensitivity_ARM_High   = DynamicMetricCriteria_AlertSensitivity_ARM("High")
	DynamicMetricCriteria_AlertSensitivity_ARM_Low    = DynamicMetricCriteria_AlertSensitivity_ARM("Low")
	DynamicMetricCriteria_AlertSensitivity_ARM_Medium = DynamicMetricCriteria_AlertSensitivity_ARM("Medium")
)

// Mapping from string to DynamicMetricCriteria_AlertSensitivity_ARM
var dynamicMetricCriteria_AlertSensitivity_ARM_Values = map[string]DynamicMetricCriteria_AlertSensitivity_ARM{
	"high":   DynamicMetricCriteria_AlertSensitivity_ARM_High,
	"low":    DynamicMetricCriteria_AlertSensitivity_ARM_Low,
	"medium": DynamicMetricCriteria_AlertSensitivity_ARM_Medium,
}

// +kubebuilder:validation:Enum={"DynamicThresholdCriterion"}
type DynamicMetricCriteria_CriterionType_ARM string

const DynamicMetricCriteria_CriterionType_ARM_DynamicThresholdCriterion = DynamicMetricCriteria_CriterionType_ARM("DynamicThresholdCriterion")

// Mapping from string to DynamicMetricCriteria_CriterionType_ARM
var dynamicMetricCriteria_CriterionType_ARM_Values = map[string]DynamicMetricCriteria_CriterionType_ARM{
	"dynamicthresholdcriterion": DynamicMetricCriteria_CriterionType_ARM_DynamicThresholdCriterion,
}

// +kubebuilder:validation:Enum={"GreaterOrLessThan","GreaterThan","LessThan"}
type DynamicMetricCriteria_Operator_ARM string

const (
	DynamicMetricCriteria_Operator_ARM_GreaterOrLessThan = DynamicMetricCriteria_Operator_ARM("GreaterOrLessThan")
	DynamicMetricCriteria_Operator_ARM_GreaterThan       = DynamicMetricCriteria_Operator_ARM("GreaterThan")
	DynamicMetricCriteria_Operator_ARM_LessThan          = DynamicMetricCriteria_Operator_ARM("LessThan")
)

// Mapping from string to DynamicMetricCriteria_Operator_ARM
var dynamicMetricCriteria_Operator_ARM_Values = map[string]DynamicMetricCriteria_Operator_ARM{
	"greaterorlessthan": DynamicMetricCriteria_Operator_ARM_GreaterOrLessThan,
	"greaterthan":       DynamicMetricCriteria_Operator_ARM_GreaterThan,
	"lessthan":          DynamicMetricCriteria_Operator_ARM_LessThan,
}

// +kubebuilder:validation:Enum={"Average","Count","Maximum","Minimum","Total"}
type DynamicMetricCriteria_TimeAggregation_ARM string

const (
	DynamicMetricCriteria_TimeAggregation_ARM_Average = DynamicMetricCriteria_TimeAggregation_ARM("Average")
	DynamicMetricCriteria_TimeAggregation_ARM_Count   = DynamicMetricCriteria_TimeAggregation_ARM("Count")
	DynamicMetricCriteria_TimeAggregation_ARM_Maximum = DynamicMetricCriteria_TimeAggregation_ARM("Maximum")
	DynamicMetricCriteria_TimeAggregation_ARM_Minimum = DynamicMetricCriteria_TimeAggregation_ARM("Minimum")
	DynamicMetricCriteria_TimeAggregation_ARM_Total   = DynamicMetricCriteria_TimeAggregation_ARM("Total")
)

// Mapping from string to DynamicMetricCriteria_TimeAggregation_ARM
var dynamicMetricCriteria_TimeAggregation_ARM_Values = map[string]DynamicMetricCriteria_TimeAggregation_ARM{
	"average": DynamicMetricCriteria_TimeAggregation_ARM_Average,
	"count":   DynamicMetricCriteria_TimeAggregation_ARM_Count,
	"maximum": DynamicMetricCriteria_TimeAggregation_ARM_Maximum,
	"minimum": DynamicMetricCriteria_TimeAggregation_ARM_Minimum,
	"total":   DynamicMetricCriteria_TimeAggregation_ARM_Total,
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
type DynamicThresholdFailingPeriods_ARM struct {
	// MinFailingPeriodsToAlert: The number of violations to trigger an alert. Should be smaller or equal to
	// numberOfEvaluationPeriods.
	MinFailingPeriodsToAlert *float64 `json:"minFailingPeriodsToAlert,omitempty"`

	// NumberOfEvaluationPeriods: The number of aggregated lookback points. The lookback time window is calculated based on the
	// aggregation granularity (windowSize) and the selected number of aggregated points.
	NumberOfEvaluationPeriods *float64 `json:"numberOfEvaluationPeriods,omitempty"`
}
