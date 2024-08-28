// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *BackupShortTermRetentionPolicyProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"
}

// Properties of a short term retention policy
type BackupShortTermRetentionPolicyProperties_ARM struct {
	// DiffBackupIntervalInHours: The differential backup interval in hours. This is how many interval hours between each
	// differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_ARM `json:"diffBackupIntervalInHours,omitempty"`

	// RetentionDays: The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `json:"retentionDays,omitempty"`
}

// +kubebuilder:validation:Enum={12,24}
type BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_ARM int

const (
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_ARM_12 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_ARM(12)
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_ARM_24 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_ARM(24)
)
