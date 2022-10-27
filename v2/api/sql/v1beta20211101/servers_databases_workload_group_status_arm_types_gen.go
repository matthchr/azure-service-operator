// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

type Servers_Databases_WorkloadGroup_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *WorkloadGroupProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Workload group definition. For more information look at sys.workload_management_workload_groups (DMV).
type WorkloadGroupProperties_STATUS_ARM struct {
	// Importance: The workload group importance level.
	Importance *string `json:"importance,omitempty"`

	// MaxResourcePercent: The workload group cap percentage resource.
	MaxResourcePercent *int `json:"maxResourcePercent,omitempty"`

	// MaxResourcePercentPerRequest: The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest *float64 `json:"maxResourcePercentPerRequest,omitempty"`

	// MinResourcePercent: The workload group minimum percentage resource.
	MinResourcePercent *int `json:"minResourcePercent,omitempty"`

	// MinResourcePercentPerRequest: The workload group request minimum grant percentage.
	MinResourcePercentPerRequest *float64 `json:"minResourcePercentPerRequest,omitempty"`

	// QueryExecutionTimeout: The workload group query execution timeout.
	QueryExecutionTimeout *int `json:"queryExecutionTimeout,omitempty"`
}
