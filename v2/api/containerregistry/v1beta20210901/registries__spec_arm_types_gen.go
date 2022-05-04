// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Registries_SpecARM struct {
	// Identity: Managed identity for the resource.
	Identity *IdentityPropertiesARM `json:"identity,omitempty"`

	// Location: The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// Name: The name of the container registry.
	Name string `json:"name,omitempty"`

	// Properties: The properties of a container registry.
	Properties *RegistryPropertiesARM `json:"properties,omitempty"`

	// Sku: The SKU of a container registry.
	Sku *SkuARM `json:"sku,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Registries_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-09-01"
func (registries Registries_SpecARM) GetAPIVersion() string {
	return "2021-09-01"
}

// GetName returns the Name of the resource
func (registries Registries_SpecARM) GetName() string {
	return registries.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerRegistry/registries"
func (registries Registries_SpecARM) GetType() string {
	return "Microsoft.ContainerRegistry/registries"
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IdentityProperties
type IdentityPropertiesARM struct {
	// PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type.
	Type *IdentityPropertiesType `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user identities associated with the resource. The user identity
	// dictionary key references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/
	// providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]UserIdentityPropertiesARM `json:"userAssignedIdentities,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/RegistryProperties
type RegistryPropertiesARM struct {
	// AdminUserEnabled: The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`

	// DataEndpointEnabled: Enable a single data endpoint per region for serving data.
	DataEndpointEnabled *bool                  `json:"dataEndpointEnabled,omitempty"`
	Encryption          *EncryptionPropertyARM `json:"encryption,omitempty"`

	// NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a network restricted registry.
	NetworkRuleBypassOptions *RegistryPropertiesNetworkRuleBypassOptions `json:"networkRuleBypassOptions,omitempty"`

	// NetworkRuleSet: The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSetARM `json:"networkRuleSet,omitempty"`

	// Policies: The policies for a container registry.
	Policies *PoliciesARM `json:"policies,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess *RegistryPropertiesPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// ZoneRedundancy: Whether or not zone redundancy is enabled for this container registry.
	ZoneRedundancy *RegistryPropertiesZoneRedundancy `json:"zoneRedundancy,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Sku
type SkuARM struct {
	// Name: The SKU name of the container registry. Required for registry creation.
	Name *SkuName `json:"name,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/EncryptionProperty
type EncryptionPropertyARM struct {
	KeyVaultProperties *KeyVaultPropertiesARM `json:"keyVaultProperties,omitempty"`

	// Status: Indicates whether or not the encryption is enabled for container registry.
	Status *EncryptionPropertyStatus `json:"status,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type IdentityPropertiesType string

const (
	IdentityPropertiesTypeNone                       = IdentityPropertiesType("None")
	IdentityPropertiesTypeSystemAssigned             = IdentityPropertiesType("SystemAssigned")
	IdentityPropertiesTypeSystemAssignedUserAssigned = IdentityPropertiesType("SystemAssigned, UserAssigned")
	IdentityPropertiesTypeUserAssigned               = IdentityPropertiesType("UserAssigned")
)

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/NetworkRuleSet
type NetworkRuleSetARM struct {
	// DefaultAction: The default action of allow or deny when no other rules match.
	DefaultAction *NetworkRuleSetDefaultAction `json:"defaultAction,omitempty"`

	// IpRules: The IP ACL rules.
	IpRules []IPRuleARM `json:"ipRules,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Policies
type PoliciesARM struct {
	// ExportPolicy: The export policy for a container registry.
	ExportPolicy *ExportPolicyARM `json:"exportPolicy,omitempty"`

	// QuarantinePolicy: The quarantine policy for a container registry.
	QuarantinePolicy *QuarantinePolicyARM `json:"quarantinePolicy,omitempty"`

	// RetentionPolicy: The retention policy for a container registry.
	RetentionPolicy *RetentionPolicyARM `json:"retentionPolicy,omitempty"`

	// TrustPolicy: The content trust policy for a container registry.
	TrustPolicy *TrustPolicyARM `json:"trustPolicy,omitempty"`
}

// +kubebuilder:validation:Enum={"AzureServices","None"}
type RegistryPropertiesNetworkRuleBypassOptions string

const (
	RegistryPropertiesNetworkRuleBypassOptionsAzureServices = RegistryPropertiesNetworkRuleBypassOptions("AzureServices")
	RegistryPropertiesNetworkRuleBypassOptionsNone          = RegistryPropertiesNetworkRuleBypassOptions("None")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RegistryPropertiesPublicNetworkAccess string

const (
	RegistryPropertiesPublicNetworkAccessDisabled = RegistryPropertiesPublicNetworkAccess("Disabled")
	RegistryPropertiesPublicNetworkAccessEnabled  = RegistryPropertiesPublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RegistryPropertiesZoneRedundancy string

const (
	RegistryPropertiesZoneRedundancyDisabled = RegistryPropertiesZoneRedundancy("Disabled")
	RegistryPropertiesZoneRedundancyEnabled  = RegistryPropertiesZoneRedundancy("Enabled")
)

// +kubebuilder:validation:Enum={"Basic","Classic","Premium","Standard"}
type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNameClassic  = SkuName("Classic")
	SkuNamePremium  = SkuName("Premium")
	SkuNameStandard = SkuName("Standard")
)

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/UserIdentityProperties
type UserIdentityPropertiesARM struct {
	// ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

// +kubebuilder:validation:Enum={"disabled","enabled"}
type EncryptionPropertyStatus string

const (
	EncryptionPropertyStatusDisabled = EncryptionPropertyStatus("disabled")
	EncryptionPropertyStatusEnabled  = EncryptionPropertyStatus("enabled")
)

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/ExportPolicy
type ExportPolicyARM struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *ExportPolicyStatus `json:"status,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IPRule
type IPRuleARM struct {
	// Action: The action of IP ACL rule.
	Action *IPRuleAction `json:"action,omitempty"`

	// Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value *string `json:"value,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/KeyVaultProperties
type KeyVaultPropertiesARM struct {
	// Identity: The client id of the identity which will be used to access key vault.
	Identity *string `json:"identity,omitempty"`

	// KeyIdentifier: Key vault uri to access the encryption key.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}

// +kubebuilder:validation:Enum={"Allow","Deny"}
type NetworkRuleSetDefaultAction string

const (
	NetworkRuleSetDefaultActionAllow = NetworkRuleSetDefaultAction("Allow")
	NetworkRuleSetDefaultActionDeny  = NetworkRuleSetDefaultAction("Deny")
)

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/QuarantinePolicy
type QuarantinePolicyARM struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *QuarantinePolicyStatus `json:"status,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/RetentionPolicy
type RetentionPolicyARM struct {
	// Days: The number of days to retain an untagged manifest after which it gets purged.
	Days *int `json:"days,omitempty"`

	// Status: The value that indicates whether the policy is enabled or not.
	Status *RetentionPolicyStatus `json:"status,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/TrustPolicy
type TrustPolicyARM struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *TrustPolicyStatus `json:"status,omitempty"`

	// Type: The type of trust policy.
	Type *TrustPolicyType `json:"type,omitempty"`
}

// +kubebuilder:validation:Enum={"disabled","enabled"}
type ExportPolicyStatus string

const (
	ExportPolicyStatusDisabled = ExportPolicyStatus("disabled")
	ExportPolicyStatusEnabled  = ExportPolicyStatus("enabled")
)

// +kubebuilder:validation:Enum={"Allow"}
type IPRuleAction string

const IPRuleActionAllow = IPRuleAction("Allow")

// +kubebuilder:validation:Enum={"disabled","enabled"}
type QuarantinePolicyStatus string

const (
	QuarantinePolicyStatusDisabled = QuarantinePolicyStatus("disabled")
	QuarantinePolicyStatusEnabled  = QuarantinePolicyStatus("enabled")
)

// +kubebuilder:validation:Enum={"disabled","enabled"}
type RetentionPolicyStatus string

const (
	RetentionPolicyStatusDisabled = RetentionPolicyStatus("disabled")
	RetentionPolicyStatusEnabled  = RetentionPolicyStatus("enabled")
)

// +kubebuilder:validation:Enum={"disabled","enabled"}
type TrustPolicyStatus string

const (
	TrustPolicyStatusDisabled = TrustPolicyStatus("disabled")
	TrustPolicyStatusEnabled  = TrustPolicyStatus("enabled")
)

// +kubebuilder:validation:Enum={"Notary"}
type TrustPolicyType string

const TrustPolicyTypeNotary = TrustPolicyType("Notary")
