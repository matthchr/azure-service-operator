// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_JobAgent_Spec_ARM struct {
	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *JobAgentProperties_ARM `json:"properties,omitempty"`

	// Sku: The name and tier of the SKU.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_JobAgent_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (agent Servers_JobAgent_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (agent *Servers_JobAgent_Spec_ARM) GetName() string {
	return agent.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/jobAgents"
func (agent *Servers_JobAgent_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/jobAgents"
}

// Properties of a job agent.
type JobAgentProperties_ARM struct {
	DatabaseId *string `json:"databaseId,omitempty"`
}
