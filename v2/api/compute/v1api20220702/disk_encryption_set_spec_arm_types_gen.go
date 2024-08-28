// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220702

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DiskEncryptionSet_Spec_ARM struct {
	// Identity: The managed identity for the disk encryption set. It should be given permission on the key vault before it can
	// be used  to encrypt disks.
	Identity *EncryptionSetIdentity_ARM `json:"identity,omitempty"`

	// Location: Resource location
	Location   *string                      `json:"location,omitempty"`
	Name       string                       `json:"name,omitempty"`
	Properties *EncryptionSetProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DiskEncryptionSet_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-02"
func (encryptionSet DiskEncryptionSet_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (encryptionSet *DiskEncryptionSet_Spec_ARM) GetName() string {
	return encryptionSet.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/diskEncryptionSets"
func (encryptionSet *DiskEncryptionSet_Spec_ARM) GetType() string {
	return "Microsoft.Compute/diskEncryptionSets"
}

// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used
// to encrypt disks.
type EncryptionSetIdentity_ARM struct {
	// Type: The type of Managed Identity used by the DiskEncryptionSet. Only SystemAssigned is supported for new creations.
	// Disk Encryption Sets can be updated with Identity type None during migration of subscription to a new Azure Active
	// Directory tenant; it will cause the encrypted resources to lose access to the keys.
	Type                   *EncryptionSetIdentity_Type_ARM            `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

type EncryptionSetProperties_ARM struct {
	// ActiveKey: The key vault key which is currently used by this disk encryption set.
	ActiveKey *KeyForDiskEncryptionSet_ARM `json:"activeKey,omitempty"`

	// EncryptionType: The type of key used to encrypt the data of the disk.
	EncryptionType *DiskEncryptionSetType_ARM `json:"encryptionType,omitempty"`

	// FederatedClientId: Multi-tenant application client id to access key vault in a different tenant. Setting the value to
	// 'None' will clear the property.
	FederatedClientId *string `json:"federatedClientId,omitempty" optionalConfigMapPair:"FederatedClientId"`

	// RotationToLatestKeyVersionEnabled: Set this flag to true to enable auto-updating of this disk encryption set to the
	// latest key version.
	RotationToLatestKeyVersionEnabled *bool `json:"rotationToLatestKeyVersionEnabled,omitempty"`
}

// The type of key used to encrypt the data of the disk.
// +kubebuilder:validation:Enum={"ConfidentialVmEncryptedWithCustomerKey","EncryptionAtRestWithCustomerKey","EncryptionAtRestWithPlatformAndCustomerKeys"}
type DiskEncryptionSetType_ARM string

const (
	DiskEncryptionSetType_ARM_ConfidentialVmEncryptedWithCustomerKey      = DiskEncryptionSetType_ARM("ConfidentialVmEncryptedWithCustomerKey")
	DiskEncryptionSetType_ARM_EncryptionAtRestWithCustomerKey             = DiskEncryptionSetType_ARM("EncryptionAtRestWithCustomerKey")
	DiskEncryptionSetType_ARM_EncryptionAtRestWithPlatformAndCustomerKeys = DiskEncryptionSetType_ARM("EncryptionAtRestWithPlatformAndCustomerKeys")
)

// Mapping from string to DiskEncryptionSetType_ARM
var diskEncryptionSetType_ARM_Values = map[string]DiskEncryptionSetType_ARM{
	"confidentialvmencryptedwithcustomerkey":      DiskEncryptionSetType_ARM_ConfidentialVmEncryptedWithCustomerKey,
	"encryptionatrestwithcustomerkey":             DiskEncryptionSetType_ARM_EncryptionAtRestWithCustomerKey,
	"encryptionatrestwithplatformandcustomerkeys": DiskEncryptionSetType_ARM_EncryptionAtRestWithPlatformAndCustomerKeys,
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type EncryptionSetIdentity_Type_ARM string

const (
	EncryptionSetIdentity_Type_ARM_None                       = EncryptionSetIdentity_Type_ARM("None")
	EncryptionSetIdentity_Type_ARM_SystemAssigned             = EncryptionSetIdentity_Type_ARM("SystemAssigned")
	EncryptionSetIdentity_Type_ARM_SystemAssignedUserAssigned = EncryptionSetIdentity_Type_ARM("SystemAssigned, UserAssigned")
	EncryptionSetIdentity_Type_ARM_UserAssigned               = EncryptionSetIdentity_Type_ARM("UserAssigned")
)

// Mapping from string to EncryptionSetIdentity_Type_ARM
var encryptionSetIdentity_Type_ARM_Values = map[string]EncryptionSetIdentity_Type_ARM{
	"none":                         EncryptionSetIdentity_Type_ARM_None,
	"systemassigned":               EncryptionSetIdentity_Type_ARM_SystemAssigned,
	"systemassigned, userassigned": EncryptionSetIdentity_Type_ARM_SystemAssignedUserAssigned,
	"userassigned":                 EncryptionSetIdentity_Type_ARM_UserAssigned,
}

// Key Vault Key Url to be used for server side encryption of Managed Disks and Snapshots
type KeyForDiskEncryptionSet_ARM struct {
	// KeyUrl: Fully versioned Key Url pointing to a key in KeyVault. Version segment of the Url is required regardless of
	// rotationToLatestKeyVersionEnabled value.
	KeyUrl *string `json:"keyUrl,omitempty" optionalConfigMapPair:"KeyUrl"`

	// SourceVault: Resource id of the KeyVault containing the key or secret. This property is optional and cannot be used if
	// the KeyVault subscription is not the same as the Disk Encryption Set subscription.
	SourceVault *SourceVault_ARM `json:"sourceVault,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// The vault id is an Azure Resource Manager Resource id in the form
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
type SourceVault_ARM struct {
	Id *string `json:"id,omitempty"`
}
