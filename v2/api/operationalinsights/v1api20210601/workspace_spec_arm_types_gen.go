// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Workspace_Spec_ARM struct {
	// Etag: The etag of the workspace.
	Etag *string `json:"etag,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Workspace properties.
	Properties *WorkspaceProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspace_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (workspace Workspace_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (workspace *Workspace_Spec_ARM) GetName() string {
	return workspace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.OperationalInsights/workspaces"
func (workspace *Workspace_Spec_ARM) GetType() string {
	return "Microsoft.OperationalInsights/workspaces"
}

// Workspace properties.
type WorkspaceProperties_ARM struct {
	// Features: Workspace features.
	Features *WorkspaceFeatures_ARM `json:"features,omitempty"`

	// ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.
	ForceCmkForQuery *bool `json:"forceCmkForQuery,omitempty"`

	// ProvisioningState: The provisioning state of the workspace.
	ProvisioningState *WorkspaceProperties_ProvisioningState_ARM `json:"provisioningState,omitempty"`

	// PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion *PublicNetworkAccessType_ARM `json:"publicNetworkAccessForIngestion,omitempty"`

	// PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery *PublicNetworkAccessType_ARM `json:"publicNetworkAccessForQuery,omitempty"`

	// RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
	// documentation for details.
	RetentionInDays *int `json:"retentionInDays,omitempty"`

	// Sku: The SKU of the workspace.
	Sku *WorkspaceSku_ARM `json:"sku,omitempty"`

	// WorkspaceCapping: The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCapping_ARM `json:"workspaceCapping,omitempty"`
}

// The network access type for operating on the Log Analytics Workspace. By default it is Enabled
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type PublicNetworkAccessType_ARM string

const (
	PublicNetworkAccessType_ARM_Disabled = PublicNetworkAccessType_ARM("Disabled")
	PublicNetworkAccessType_ARM_Enabled  = PublicNetworkAccessType_ARM("Enabled")
)

// Mapping from string to PublicNetworkAccessType_ARM
var publicNetworkAccessType_ARM_Values = map[string]PublicNetworkAccessType_ARM{
	"disabled": PublicNetworkAccessType_ARM_Disabled,
	"enabled":  PublicNetworkAccessType_ARM_Enabled,
}

// The daily volume cap for ingestion.
type WorkspaceCapping_ARM struct {
	// DailyQuotaGb: The workspace daily quota for ingestion.
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty"`
}

// Workspace features.
type WorkspaceFeatures_ARM struct {
	ClusterResourceId *string `json:"clusterResourceId,omitempty"`

	// DisableLocalAuth: Disable Non-AAD based Auth.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// EnableDataExport: Flag that indicate if data should be exported.
	EnableDataExport *bool `json:"enableDataExport,omitempty"`

	// EnableLogAccessUsingOnlyResourcePermissions: Flag that indicate which permission to use - resource or workspace or both.
	EnableLogAccessUsingOnlyResourcePermissions *bool `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`

	// ImmediatePurgeDataOn30Days: Flag that describes if we want to remove the data after 30 days.
	ImmediatePurgeDataOn30Days *bool `json:"immediatePurgeDataOn30Days,omitempty"`
}

// +kubebuilder:validation:Enum={"Canceled","Creating","Deleting","Failed","ProvisioningAccount","Succeeded","Updating"}
type WorkspaceProperties_ProvisioningState_ARM string

const (
	WorkspaceProperties_ProvisioningState_ARM_Canceled            = WorkspaceProperties_ProvisioningState_ARM("Canceled")
	WorkspaceProperties_ProvisioningState_ARM_Creating            = WorkspaceProperties_ProvisioningState_ARM("Creating")
	WorkspaceProperties_ProvisioningState_ARM_Deleting            = WorkspaceProperties_ProvisioningState_ARM("Deleting")
	WorkspaceProperties_ProvisioningState_ARM_Failed              = WorkspaceProperties_ProvisioningState_ARM("Failed")
	WorkspaceProperties_ProvisioningState_ARM_ProvisioningAccount = WorkspaceProperties_ProvisioningState_ARM("ProvisioningAccount")
	WorkspaceProperties_ProvisioningState_ARM_Succeeded           = WorkspaceProperties_ProvisioningState_ARM("Succeeded")
	WorkspaceProperties_ProvisioningState_ARM_Updating            = WorkspaceProperties_ProvisioningState_ARM("Updating")
)

// Mapping from string to WorkspaceProperties_ProvisioningState_ARM
var workspaceProperties_ProvisioningState_ARM_Values = map[string]WorkspaceProperties_ProvisioningState_ARM{
	"canceled":            WorkspaceProperties_ProvisioningState_ARM_Canceled,
	"creating":            WorkspaceProperties_ProvisioningState_ARM_Creating,
	"deleting":            WorkspaceProperties_ProvisioningState_ARM_Deleting,
	"failed":              WorkspaceProperties_ProvisioningState_ARM_Failed,
	"provisioningaccount": WorkspaceProperties_ProvisioningState_ARM_ProvisioningAccount,
	"succeeded":           WorkspaceProperties_ProvisioningState_ARM_Succeeded,
	"updating":            WorkspaceProperties_ProvisioningState_ARM_Updating,
}

// The SKU (tier) of a workspace.
type WorkspaceSku_ARM struct {
	// CapacityReservationLevel: The capacity reservation level in GB for this workspace, when CapacityReservation sku is
	// selected.
	CapacityReservationLevel *WorkspaceSku_CapacityReservationLevel_ARM `json:"capacityReservationLevel,omitempty"`

	// Name: The name of the SKU.
	Name *WorkspaceSku_Name_ARM `json:"name,omitempty"`
}

// +kubebuilder:validation:Enum={100,200,300,400,500,1000,2000,5000}
type WorkspaceSku_CapacityReservationLevel_ARM int

const (
	WorkspaceSku_CapacityReservationLevel_ARM_100  = WorkspaceSku_CapacityReservationLevel_ARM(100)
	WorkspaceSku_CapacityReservationLevel_ARM_200  = WorkspaceSku_CapacityReservationLevel_ARM(200)
	WorkspaceSku_CapacityReservationLevel_ARM_300  = WorkspaceSku_CapacityReservationLevel_ARM(300)
	WorkspaceSku_CapacityReservationLevel_ARM_400  = WorkspaceSku_CapacityReservationLevel_ARM(400)
	WorkspaceSku_CapacityReservationLevel_ARM_500  = WorkspaceSku_CapacityReservationLevel_ARM(500)
	WorkspaceSku_CapacityReservationLevel_ARM_1000 = WorkspaceSku_CapacityReservationLevel_ARM(1000)
	WorkspaceSku_CapacityReservationLevel_ARM_2000 = WorkspaceSku_CapacityReservationLevel_ARM(2000)
	WorkspaceSku_CapacityReservationLevel_ARM_5000 = WorkspaceSku_CapacityReservationLevel_ARM(5000)
)

// +kubebuilder:validation:Enum={"CapacityReservation","Free","LACluster","PerGB2018","PerNode","Premium","Standalone","Standard"}
type WorkspaceSku_Name_ARM string

const (
	WorkspaceSku_Name_ARM_CapacityReservation = WorkspaceSku_Name_ARM("CapacityReservation")
	WorkspaceSku_Name_ARM_Free                = WorkspaceSku_Name_ARM("Free")
	WorkspaceSku_Name_ARM_LACluster           = WorkspaceSku_Name_ARM("LACluster")
	WorkspaceSku_Name_ARM_PerGB2018           = WorkspaceSku_Name_ARM("PerGB2018")
	WorkspaceSku_Name_ARM_PerNode             = WorkspaceSku_Name_ARM("PerNode")
	WorkspaceSku_Name_ARM_Premium             = WorkspaceSku_Name_ARM("Premium")
	WorkspaceSku_Name_ARM_Standalone          = WorkspaceSku_Name_ARM("Standalone")
	WorkspaceSku_Name_ARM_Standard            = WorkspaceSku_Name_ARM("Standard")
)

// Mapping from string to WorkspaceSku_Name_ARM
var workspaceSku_Name_ARM_Values = map[string]WorkspaceSku_Name_ARM{
	"capacityreservation": WorkspaceSku_Name_ARM_CapacityReservation,
	"free":                WorkspaceSku_Name_ARM_Free,
	"lacluster":           WorkspaceSku_Name_ARM_LACluster,
	"pergb2018":           WorkspaceSku_Name_ARM_PerGB2018,
	"pernode":             WorkspaceSku_Name_ARM_PerNode,
	"premium":             WorkspaceSku_Name_ARM_Premium,
	"standalone":          WorkspaceSku_Name_ARM_Standalone,
	"standard":            WorkspaceSku_Name_ARM_Standard,
}
