// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231115

import "encoding/json"

type DatabaseAccount_STATUS_ARM struct {
	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Identity: Identity for the resource.
	Identity *ManagedServiceIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Kind: Indicates the type of database account. This can only be set at database account creation.
	Kind *DatabaseAccount_Kind_STATUS_ARM `json:"kind,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties for the database account.
	Properties *DatabaseAccountGetProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`
	Tags       map[string]string      `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

type DatabaseAccount_Kind_STATUS_ARM string

const (
	DatabaseAccount_Kind_STATUS_ARM_GlobalDocumentDB = DatabaseAccount_Kind_STATUS_ARM("GlobalDocumentDB")
	DatabaseAccount_Kind_STATUS_ARM_MongoDB          = DatabaseAccount_Kind_STATUS_ARM("MongoDB")
	DatabaseAccount_Kind_STATUS_ARM_Parse            = DatabaseAccount_Kind_STATUS_ARM("Parse")
)

// Mapping from string to DatabaseAccount_Kind_STATUS_ARM
var databaseAccount_Kind_STATUS_ARM_Values = map[string]DatabaseAccount_Kind_STATUS_ARM{
	"globaldocumentdb": DatabaseAccount_Kind_STATUS_ARM_GlobalDocumentDB,
	"mongodb":          DatabaseAccount_Kind_STATUS_ARM_MongoDB,
	"parse":            DatabaseAccount_Kind_STATUS_ARM_Parse,
}

// Properties for the database account.
type DatabaseAccountGetProperties_STATUS_ARM struct {
	// AnalyticalStorageConfiguration: Analytical storage specific properties.
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration_STATUS_ARM `json:"analyticalStorageConfiguration,omitempty"`

	// ApiProperties: API specific properties.
	ApiProperties *ApiProperties_STATUS_ARM `json:"apiProperties,omitempty"`

	// BackupPolicy: The object representing the policy for taking backups on an account.
	BackupPolicy *BackupPolicy_STATUS_ARM `json:"backupPolicy,omitempty"`

	// Capabilities: List of Cosmos DB capabilities for the account
	Capabilities []Capability_STATUS_ARM `json:"capabilities,omitempty"`

	// Capacity: The object that represents all properties related to capacity enforcement on an account.
	Capacity *Capacity_STATUS_ARM `json:"capacity,omitempty"`

	// ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer *ConnectorOffer_STATUS_ARM `json:"connectorOffer,omitempty"`

	// ConsistencyPolicy: The consistency policy for the Cosmos DB database account.
	ConsistencyPolicy *ConsistencyPolicy_STATUS_ARM `json:"consistencyPolicy,omitempty"`

	// Cors: The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicy_STATUS_ARM `json:"cors,omitempty"`

	// CreateMode: Enum to indicate the mode of account creation.
	CreateMode *CreateMode_STATUS_ARM `json:"createMode,omitempty"`

	// CustomerManagedKeyStatus: Indicates the status of the Customer Managed Key feature on the account. In case there are
	// errors, the property provides troubleshooting guidance.
	CustomerManagedKeyStatus *string `json:"customerManagedKeyStatus,omitempty"`

	// DatabaseAccountOfferType: The offer type for the Cosmos DB database account. Default value: Standard.
	DatabaseAccountOfferType *DatabaseAccountOfferType_STATUS_ARM `json:"databaseAccountOfferType,omitempty"`

	// DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
	// identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity *string `json:"defaultIdentity,omitempty"`

	// DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
	// via account keys
	DisableKeyBasedMetadataWriteAccess *bool `json:"disableKeyBasedMetadataWriteAccess,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure only MSI and AAD can be used exclusively for authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// DocumentEndpoint: The connection endpoint for the Cosmos DB database account.
	DocumentEndpoint *string `json:"documentEndpoint,omitempty"`

	// EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `json:"enableAnalyticalStorage,omitempty"`

	// EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
	// due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
	// failover priorities configured for the account.
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty"`

	// EnableBurstCapacity: Flag to indicate enabling/disabling of Burst Capacity feature on the account
	EnableBurstCapacity *bool `json:"enableBurstCapacity,omitempty"`

	// EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector *bool `json:"enableCassandraConnector,omitempty"`

	// EnableFreeTier: Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `json:"enableFreeTier,omitempty"`

	// EnableMultipleWriteLocations: Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty"`

	// EnablePartitionMerge: Flag to indicate enabling/disabling of Partition Merge feature on the account
	EnablePartitionMerge *bool `json:"enablePartitionMerge,omitempty"`

	// FailoverPolicies: An array that contains the regions ordered by their failover priorities.
	FailoverPolicies []FailoverPolicy_STATUS_ARM `json:"failoverPolicies,omitempty"`

	// InstanceId: A unique identifier assigned to the database account
	InstanceId *string `json:"instanceId,omitempty"`

	// IpRules: List of IpRules.
	IpRules []IpAddressOrRange_STATUS_ARM `json:"ipRules,omitempty"`

	// IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty"`

	// KeyVaultKeyUri: The URI of the key vault
	KeyVaultKeyUri *string `json:"keyVaultKeyUri,omitempty"`

	// KeysMetadata: The object that represents the metadata for the Account Keys of the Cosmos DB account.
	KeysMetadata *DatabaseAccountKeysMetadata_STATUS_ARM `json:"keysMetadata,omitempty"`

	// Locations: An array that contains all of the locations enabled for the Cosmos DB account.
	Locations []Location_STATUS_ARM `json:"locations,omitempty"`

	// MinimalTlsVersion: Indicates the minimum allowed Tls version. The default value is Tls 1.2. Cassandra and Mongo APIs
	// only work with Tls 1.2.
	MinimalTlsVersion *MinimalTlsVersion_STATUS_ARM `json:"minimalTlsVersion,omitempty"`

	// NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass *NetworkAclBypass_STATUS_ARM `json:"networkAclBypass,omitempty"`

	// NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds []string `json:"networkAclBypassResourceIds,omitempty"`

	// PrivateEndpointConnections: List of Private Endpoint Connections configured for the Cosmos DB account.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *string                                `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether requests from Public Network are allowed
	PublicNetworkAccess *PublicNetworkAccess_STATUS_ARM `json:"publicNetworkAccess,omitempty"`

	// ReadLocations: An array that contains of the read locations enabled for the Cosmos DB account.
	ReadLocations []Location_STATUS_ARM `json:"readLocations,omitempty"`

	// RestoreParameters: Parameters to indicate the information about the restore.
	RestoreParameters *RestoreParameters_STATUS_ARM `json:"restoreParameters,omitempty"`

	// VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules []VirtualNetworkRule_STATUS_ARM `json:"virtualNetworkRules,omitempty"`

	// WriteLocations: An array that contains the write location for the Cosmos DB account.
	WriteLocations []Location_STATUS_ARM `json:"writeLocations,omitempty"`
}

// Identity for the resource.
type ManagedServiceIdentity_STATUS_ARM struct {
	// PrincipalId: The principal id of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant id of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly
	// created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
	Type *ManagedServiceIdentity_Type_STATUS_ARM `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user identities associated with resource. The user identity dictionary key
	// references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]ManagedServiceIdentity_UserAssignedIdentities_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS_ARM `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS_ARM `json:"lastModifiedByType,omitempty"`
}

// Analytical storage specific properties.
type AnalyticalStorageConfiguration_STATUS_ARM struct {
	// SchemaType: Describes the types of schema for analytical storage.
	SchemaType *AnalyticalStorageSchemaType_STATUS_ARM `json:"schemaType,omitempty"`
}

type ApiProperties_STATUS_ARM struct {
	// ServerVersion: Describes the ServerVersion of an a MongoDB account.
	ServerVersion *ApiProperties_ServerVersion_STATUS_ARM `json:"serverVersion,omitempty"`
}

type BackupPolicy_STATUS_ARM struct {
	// Continuous: Mutually exclusive with all other properties
	Continuous *ContinuousModeBackupPolicy_STATUS_ARM `json:"continuous,omitempty"`

	// Periodic: Mutually exclusive with all other properties
	Periodic *PeriodicModeBackupPolicy_STATUS_ARM `json:"periodic,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupPolicy_STATUS_ARM represents a discriminated union (JSON OneOf)
func (policy BackupPolicy_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if policy.Continuous != nil {
		return json.Marshal(policy.Continuous)
	}
	if policy.Periodic != nil {
		return json.Marshal(policy.Periodic)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupPolicy_STATUS_ARM
func (policy *BackupPolicy_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "Continuous" {
		policy.Continuous = &ContinuousModeBackupPolicy_STATUS_ARM{}
		return json.Unmarshal(data, policy.Continuous)
	}
	if discriminator == "Periodic" {
		policy.Periodic = &PeriodicModeBackupPolicy_STATUS_ARM{}
		return json.Unmarshal(data, policy.Periodic)
	}

	// No error
	return nil
}

// Cosmos DB capability object
type Capability_STATUS_ARM struct {
	// Name: Name of the Cosmos DB capability. For example, "name": "EnableCassandra". Current values also include
	// "EnableTable" and "EnableGremlin".
	Name *string `json:"name,omitempty"`
}

// The object that represents all properties related to capacity enforcement on an account.
type Capacity_STATUS_ARM struct {
	// TotalThroughputLimit: The total throughput limit imposed on the account. A totalThroughputLimit of 2000 imposes a strict
	// limit of max throughput that can be provisioned on that account to be 2000. A totalThroughputLimit of -1 indicates no
	// limits on provisioning of throughput.
	TotalThroughputLimit *int `json:"totalThroughputLimit,omitempty"`
}

// The cassandra connector offer type for the Cosmos DB C* database account.
type ConnectorOffer_STATUS_ARM string

const ConnectorOffer_STATUS_ARM_Small = ConnectorOffer_STATUS_ARM("Small")

// Mapping from string to ConnectorOffer_STATUS_ARM
var connectorOffer_STATUS_ARM_Values = map[string]ConnectorOffer_STATUS_ARM{
	"small": ConnectorOffer_STATUS_ARM_Small,
}

// The consistency policy for the Cosmos DB database account.
type ConsistencyPolicy_STATUS_ARM struct {
	// DefaultConsistencyLevel: The default consistency level and configuration settings of the Cosmos DB account.
	DefaultConsistencyLevel *ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM `json:"defaultConsistencyLevel,omitempty"`

	// MaxIntervalInSeconds: When used with the Bounded Staleness consistency level, this value represents the time amount of
	// staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is
	// set to 'BoundedStaleness'.
	MaxIntervalInSeconds *int `json:"maxIntervalInSeconds,omitempty"`

	// MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this value represents the number of stale
	// requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set
	// to 'BoundedStaleness'.
	MaxStalenessPrefix *int `json:"maxStalenessPrefix,omitempty"`
}

// The CORS policy for the Cosmos DB database account.
type CorsPolicy_STATUS_ARM struct {
	// AllowedHeaders: The request headers that the origin domain may specify on the CORS request.
	AllowedHeaders *string `json:"allowedHeaders,omitempty"`

	// AllowedMethods: The methods (HTTP request verbs) that the origin domain may use for a CORS request.
	AllowedMethods *string `json:"allowedMethods,omitempty"`

	// AllowedOrigins: The origin domains that are permitted to make a request against the service via CORS.
	AllowedOrigins *string `json:"allowedOrigins,omitempty"`

	// ExposedHeaders: The response headers that may be sent in the response to the CORS request and exposed by the browser to
	// the request issuer.
	ExposedHeaders *string `json:"exposedHeaders,omitempty"`

	// MaxAgeInSeconds: The maximum amount time that a browser should cache the preflight OPTIONS request.
	MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
}

// Enum to indicate the mode of account creation.
type CreateMode_STATUS_ARM string

const (
	CreateMode_STATUS_ARM_Default = CreateMode_STATUS_ARM("Default")
	CreateMode_STATUS_ARM_Restore = CreateMode_STATUS_ARM("Restore")
)

// Mapping from string to CreateMode_STATUS_ARM
var createMode_STATUS_ARM_Values = map[string]CreateMode_STATUS_ARM{
	"default": CreateMode_STATUS_ARM_Default,
	"restore": CreateMode_STATUS_ARM_Restore,
}

// The metadata related to each access key for the given Cosmos DB database account.
type DatabaseAccountKeysMetadata_STATUS_ARM struct {
	// PrimaryMasterKey: The metadata related to the Primary Read-Write Key for the given Cosmos DB database account.
	PrimaryMasterKey *AccountKeyMetadata_STATUS_ARM `json:"primaryMasterKey,omitempty"`

	// PrimaryReadonlyMasterKey: The metadata related to the Primary Read-Only Key for the given Cosmos DB database account.
	PrimaryReadonlyMasterKey *AccountKeyMetadata_STATUS_ARM `json:"primaryReadonlyMasterKey,omitempty"`

	// SecondaryMasterKey: The metadata related to the Secondary Read-Write Key for the given Cosmos DB database account.
	SecondaryMasterKey *AccountKeyMetadata_STATUS_ARM `json:"secondaryMasterKey,omitempty"`

	// SecondaryReadonlyMasterKey: The metadata related to the Secondary Read-Only Key for the given Cosmos DB database account.
	SecondaryReadonlyMasterKey *AccountKeyMetadata_STATUS_ARM `json:"secondaryReadonlyMasterKey,omitempty"`
}

// The offer type for the Cosmos DB database account.
type DatabaseAccountOfferType_STATUS_ARM string

const DatabaseAccountOfferType_STATUS_ARM_Standard = DatabaseAccountOfferType_STATUS_ARM("Standard")

// Mapping from string to DatabaseAccountOfferType_STATUS_ARM
var databaseAccountOfferType_STATUS_ARM_Values = map[string]DatabaseAccountOfferType_STATUS_ARM{
	"standard": DatabaseAccountOfferType_STATUS_ARM_Standard,
}

// The failover policy for a given region of a database account.
type FailoverPolicy_STATUS_ARM struct {
	// FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
	// value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
	// regions in which the database account exists.
	FailoverPriority *int `json:"failoverPriority,omitempty"`

	// Id: The unique identifier of the region in which the database account replicates to. Example:
	// &lt;accountName&gt;-&lt;locationName&gt;.
	Id *string `json:"id,omitempty"`

	// LocationName: The name of the region in which the database account exists.
	LocationName *string `json:"locationName,omitempty"`
}

// IpAddressOrRange object
type IpAddressOrRange_STATUS_ARM struct {
	// IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR format. Provided IPs must be
	// well-formatted and cannot be contained in one of the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12,
	// 192.168.0.0/16, since these are not enforceable by the IP address filter. Example of valid inputs: “23.40.210.245”
	// or “23.40.210.0/8”.
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

// A region in which the Azure Cosmos DB database account is deployed.
type Location_STATUS_ARM struct {
	// DocumentEndpoint: The connection endpoint for the specific region. Example:
	// https://&lt;accountName&gt;-&lt;locationName&gt;.documents.azure.com:443/
	DocumentEndpoint *string `json:"documentEndpoint,omitempty"`

	// FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
	// value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
	// regions in which the database account exists.
	FailoverPriority *int `json:"failoverPriority,omitempty"`

	// Id: The unique identifier of the region within the database account. Example: &lt;accountName&gt;-&lt;locationName&gt;.
	Id *string `json:"id,omitempty"`

	// IsZoneRedundant: Flag to indicate whether or not this region is an AvailabilityZone region
	IsZoneRedundant *bool `json:"isZoneRedundant,omitempty"`

	// LocationName: The name of the region.
	LocationName      *string `json:"locationName,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

type ManagedServiceIdentity_Type_STATUS_ARM string

const (
	ManagedServiceIdentity_Type_STATUS_ARM_None                       = ManagedServiceIdentity_Type_STATUS_ARM("None")
	ManagedServiceIdentity_Type_STATUS_ARM_SystemAssigned             = ManagedServiceIdentity_Type_STATUS_ARM("SystemAssigned")
	ManagedServiceIdentity_Type_STATUS_ARM_SystemAssignedUserAssigned = ManagedServiceIdentity_Type_STATUS_ARM("SystemAssigned,UserAssigned")
	ManagedServiceIdentity_Type_STATUS_ARM_UserAssigned               = ManagedServiceIdentity_Type_STATUS_ARM("UserAssigned")
)

// Mapping from string to ManagedServiceIdentity_Type_STATUS_ARM
var managedServiceIdentity_Type_STATUS_ARM_Values = map[string]ManagedServiceIdentity_Type_STATUS_ARM{
	"none":                        ManagedServiceIdentity_Type_STATUS_ARM_None,
	"systemassigned":              ManagedServiceIdentity_Type_STATUS_ARM_SystemAssigned,
	"systemassigned,userassigned": ManagedServiceIdentity_Type_STATUS_ARM_SystemAssignedUserAssigned,
	"userassigned":                ManagedServiceIdentity_Type_STATUS_ARM_UserAssigned,
}

type ManagedServiceIdentity_UserAssignedIdentities_STATUS_ARM struct {
	// ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

// Indicates the minimum allowed Tls version. The default value is Tls 1.2. Cassandra and Mongo APIs only work with Tls 1.2.
type MinimalTlsVersion_STATUS_ARM string

const (
	MinimalTlsVersion_STATUS_ARM_Tls   = MinimalTlsVersion_STATUS_ARM("Tls")
	MinimalTlsVersion_STATUS_ARM_Tls11 = MinimalTlsVersion_STATUS_ARM("Tls11")
	MinimalTlsVersion_STATUS_ARM_Tls12 = MinimalTlsVersion_STATUS_ARM("Tls12")
)

// Mapping from string to MinimalTlsVersion_STATUS_ARM
var minimalTlsVersion_STATUS_ARM_Values = map[string]MinimalTlsVersion_STATUS_ARM{
	"tls":   MinimalTlsVersion_STATUS_ARM_Tls,
	"tls11": MinimalTlsVersion_STATUS_ARM_Tls11,
	"tls12": MinimalTlsVersion_STATUS_ARM_Tls12,
}

// Indicates what services are allowed to bypass firewall checks.
type NetworkAclBypass_STATUS_ARM string

const (
	NetworkAclBypass_STATUS_ARM_AzureServices = NetworkAclBypass_STATUS_ARM("AzureServices")
	NetworkAclBypass_STATUS_ARM_None          = NetworkAclBypass_STATUS_ARM("None")
)

// Mapping from string to NetworkAclBypass_STATUS_ARM
var networkAclBypass_STATUS_ARM_Values = map[string]NetworkAclBypass_STATUS_ARM{
	"azureservices": NetworkAclBypass_STATUS_ARM_AzureServices,
	"none":          NetworkAclBypass_STATUS_ARM_None,
}

// A private endpoint connection
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

// Whether requests from Public Network are allowed
type PublicNetworkAccess_STATUS_ARM string

const (
	PublicNetworkAccess_STATUS_ARM_Disabled           = PublicNetworkAccess_STATUS_ARM("Disabled")
	PublicNetworkAccess_STATUS_ARM_Enabled            = PublicNetworkAccess_STATUS_ARM("Enabled")
	PublicNetworkAccess_STATUS_ARM_SecuredByPerimeter = PublicNetworkAccess_STATUS_ARM("SecuredByPerimeter")
)

// Mapping from string to PublicNetworkAccess_STATUS_ARM
var publicNetworkAccess_STATUS_ARM_Values = map[string]PublicNetworkAccess_STATUS_ARM{
	"disabled":           PublicNetworkAccess_STATUS_ARM_Disabled,
	"enabled":            PublicNetworkAccess_STATUS_ARM_Enabled,
	"securedbyperimeter": PublicNetworkAccess_STATUS_ARM_SecuredByPerimeter,
}

// Parameters to indicate the information about the restore.
type RestoreParameters_STATUS_ARM struct {
	// DatabasesToRestore: List of specific databases available for restore.
	DatabasesToRestore []DatabaseRestoreResource_STATUS_ARM `json:"databasesToRestore,omitempty"`

	// GremlinDatabasesToRestore: List of specific gremlin databases available for restore.
	GremlinDatabasesToRestore []GremlinDatabaseRestoreResource_STATUS_ARM `json:"gremlinDatabasesToRestore,omitempty"`

	// RestoreMode: Describes the mode of the restore.
	RestoreMode *RestoreParameters_RestoreMode_STATUS_ARM `json:"restoreMode,omitempty"`

	// RestoreSource: The id of the restorable database account from which the restore has to be initiated. For example:
	// /subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{restorableDatabaseAccountName}
	RestoreSource *string `json:"restoreSource,omitempty"`

	// RestoreTimestampInUtc: Time to which the account has to be restored (ISO-8601 format).
	RestoreTimestampInUtc *string `json:"restoreTimestampInUtc,omitempty"`

	// TablesToRestore: List of specific tables available for restore.
	TablesToRestore []string `json:"tablesToRestore,omitempty"`
}

type SystemData_CreatedByType_STATUS_ARM string

const (
	SystemData_CreatedByType_STATUS_ARM_Application     = SystemData_CreatedByType_STATUS_ARM("Application")
	SystemData_CreatedByType_STATUS_ARM_Key             = SystemData_CreatedByType_STATUS_ARM("Key")
	SystemData_CreatedByType_STATUS_ARM_ManagedIdentity = SystemData_CreatedByType_STATUS_ARM("ManagedIdentity")
	SystemData_CreatedByType_STATUS_ARM_User            = SystemData_CreatedByType_STATUS_ARM("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS_ARM
var systemData_CreatedByType_STATUS_ARM_Values = map[string]SystemData_CreatedByType_STATUS_ARM{
	"application":     SystemData_CreatedByType_STATUS_ARM_Application,
	"key":             SystemData_CreatedByType_STATUS_ARM_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ARM_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_ARM_User,
}

type SystemData_LastModifiedByType_STATUS_ARM string

const (
	SystemData_LastModifiedByType_STATUS_ARM_Application     = SystemData_LastModifiedByType_STATUS_ARM("Application")
	SystemData_LastModifiedByType_STATUS_ARM_Key             = SystemData_LastModifiedByType_STATUS_ARM("Key")
	SystemData_LastModifiedByType_STATUS_ARM_ManagedIdentity = SystemData_LastModifiedByType_STATUS_ARM("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_ARM_User            = SystemData_LastModifiedByType_STATUS_ARM("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS_ARM
var systemData_LastModifiedByType_STATUS_ARM_Values = map[string]SystemData_LastModifiedByType_STATUS_ARM{
	"application":     SystemData_LastModifiedByType_STATUS_ARM_Application,
	"key":             SystemData_LastModifiedByType_STATUS_ARM_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ARM_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_ARM_User,
}

// Virtual Network ACL Rule object
type VirtualNetworkRule_STATUS_ARM struct {
	// Id: Resource ID of a subnet, for example:
	// /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.
	Id *string `json:"id,omitempty"`

	// IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVNetServiceEndpoint *bool `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}

// The metadata related to an access key for a given database account.
type AccountKeyMetadata_STATUS_ARM struct {
	// GenerationTime: Generation time in UTC of the key in ISO-8601 format. If the value is missing from the object, it means
	// that the last key regeneration was triggered before 2022-06-18.
	GenerationTime *string `json:"generationTime,omitempty"`
}

// Describes the types of schema for analytical storage.
type AnalyticalStorageSchemaType_STATUS_ARM string

const (
	AnalyticalStorageSchemaType_STATUS_ARM_FullFidelity = AnalyticalStorageSchemaType_STATUS_ARM("FullFidelity")
	AnalyticalStorageSchemaType_STATUS_ARM_WellDefined  = AnalyticalStorageSchemaType_STATUS_ARM("WellDefined")
)

// Mapping from string to AnalyticalStorageSchemaType_STATUS_ARM
var analyticalStorageSchemaType_STATUS_ARM_Values = map[string]AnalyticalStorageSchemaType_STATUS_ARM{
	"fullfidelity": AnalyticalStorageSchemaType_STATUS_ARM_FullFidelity,
	"welldefined":  AnalyticalStorageSchemaType_STATUS_ARM_WellDefined,
}

type ApiProperties_ServerVersion_STATUS_ARM string

const (
	ApiProperties_ServerVersion_STATUS_ARM_32 = ApiProperties_ServerVersion_STATUS_ARM("3.2")
	ApiProperties_ServerVersion_STATUS_ARM_36 = ApiProperties_ServerVersion_STATUS_ARM("3.6")
	ApiProperties_ServerVersion_STATUS_ARM_40 = ApiProperties_ServerVersion_STATUS_ARM("4.0")
	ApiProperties_ServerVersion_STATUS_ARM_42 = ApiProperties_ServerVersion_STATUS_ARM("4.2")
)

// Mapping from string to ApiProperties_ServerVersion_STATUS_ARM
var apiProperties_ServerVersion_STATUS_ARM_Values = map[string]ApiProperties_ServerVersion_STATUS_ARM{
	"3.2": ApiProperties_ServerVersion_STATUS_ARM_32,
	"3.6": ApiProperties_ServerVersion_STATUS_ARM_36,
	"4.0": ApiProperties_ServerVersion_STATUS_ARM_40,
	"4.2": ApiProperties_ServerVersion_STATUS_ARM_42,
}

type ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM string

const (
	ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_BoundedStaleness = ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM("BoundedStaleness")
	ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_ConsistentPrefix = ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM("ConsistentPrefix")
	ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_Eventual         = ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM("Eventual")
	ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_Session          = ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM("Session")
	ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_Strong           = ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM("Strong")
)

// Mapping from string to ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM
var consistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_Values = map[string]ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM{
	"boundedstaleness": ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_BoundedStaleness,
	"consistentprefix": ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_ConsistentPrefix,
	"eventual":         ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_Eventual,
	"session":          ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_Session,
	"strong":           ConsistencyPolicy_DefaultConsistencyLevel_STATUS_ARM_Strong,
}

type ContinuousModeBackupPolicy_STATUS_ARM struct {
	// ContinuousModeProperties: Configuration values for continuous mode backup
	ContinuousModeProperties *ContinuousModeProperties_STATUS_ARM `json:"continuousModeProperties,omitempty"`

	// MigrationState: The object representing the state of the migration between the backup policies.
	MigrationState *BackupPolicyMigrationState_STATUS_ARM `json:"migrationState,omitempty"`

	// Type: Describes the mode of backups.
	Type ContinuousModeBackupPolicy_Type_STATUS_ARM `json:"type,omitempty"`
}

// Specific Databases to restore.
type DatabaseRestoreResource_STATUS_ARM struct {
	// CollectionNames: The names of the collections available for restore.
	CollectionNames []string `json:"collectionNames,omitempty"`

	// DatabaseName: The name of the database available for restore.
	DatabaseName *string `json:"databaseName,omitempty"`
}

// Specific Gremlin Databases to restore.
type GremlinDatabaseRestoreResource_STATUS_ARM struct {
	// DatabaseName: The name of the gremlin database available for restore.
	DatabaseName *string `json:"databaseName,omitempty"`

	// GraphNames: The names of the graphs available for restore.
	GraphNames []string `json:"graphNames,omitempty"`
}

type PeriodicModeBackupPolicy_STATUS_ARM struct {
	// MigrationState: The object representing the state of the migration between the backup policies.
	MigrationState *BackupPolicyMigrationState_STATUS_ARM `json:"migrationState,omitempty"`

	// PeriodicModeProperties: Configuration values for periodic mode backup
	PeriodicModeProperties *PeriodicModeProperties_STATUS_ARM `json:"periodicModeProperties,omitempty"`

	// Type: Describes the mode of backups.
	Type PeriodicModeBackupPolicy_Type_STATUS_ARM `json:"type,omitempty"`
}

type RestoreParameters_RestoreMode_STATUS_ARM string

const RestoreParameters_RestoreMode_STATUS_ARM_PointInTime = RestoreParameters_RestoreMode_STATUS_ARM("PointInTime")

// Mapping from string to RestoreParameters_RestoreMode_STATUS_ARM
var restoreParameters_RestoreMode_STATUS_ARM_Values = map[string]RestoreParameters_RestoreMode_STATUS_ARM{
	"pointintime": RestoreParameters_RestoreMode_STATUS_ARM_PointInTime,
}

// The object representing the state of the migration between the backup policies.
type BackupPolicyMigrationState_STATUS_ARM struct {
	// StartTime: Time at which the backup policy migration started (ISO-8601 format).
	StartTime *string `json:"startTime,omitempty"`

	// Status: Describes the status of migration between backup policy types.
	Status *BackupPolicyMigrationStatus_STATUS_ARM `json:"status,omitempty"`

	// TargetType: Describes the target backup policy type of the backup policy migration.
	TargetType *BackupPolicyType_STATUS_ARM `json:"targetType,omitempty"`
}

type ContinuousModeBackupPolicy_Type_STATUS_ARM string

const ContinuousModeBackupPolicy_Type_STATUS_ARM_Continuous = ContinuousModeBackupPolicy_Type_STATUS_ARM("Continuous")

// Mapping from string to ContinuousModeBackupPolicy_Type_STATUS_ARM
var continuousModeBackupPolicy_Type_STATUS_ARM_Values = map[string]ContinuousModeBackupPolicy_Type_STATUS_ARM{
	"continuous": ContinuousModeBackupPolicy_Type_STATUS_ARM_Continuous,
}

// Configuration values for periodic mode backup
type ContinuousModeProperties_STATUS_ARM struct {
	// Tier: Enum to indicate type of Continuous backup mode
	Tier *ContinuousTier_STATUS_ARM `json:"tier,omitempty"`
}

type PeriodicModeBackupPolicy_Type_STATUS_ARM string

const PeriodicModeBackupPolicy_Type_STATUS_ARM_Periodic = PeriodicModeBackupPolicy_Type_STATUS_ARM("Periodic")

// Mapping from string to PeriodicModeBackupPolicy_Type_STATUS_ARM
var periodicModeBackupPolicy_Type_STATUS_ARM_Values = map[string]PeriodicModeBackupPolicy_Type_STATUS_ARM{
	"periodic": PeriodicModeBackupPolicy_Type_STATUS_ARM_Periodic,
}

// Configuration values for periodic mode backup
type PeriodicModeProperties_STATUS_ARM struct {
	// BackupIntervalInMinutes: An integer representing the interval in minutes between two backups
	BackupIntervalInMinutes *int `json:"backupIntervalInMinutes,omitempty"`

	// BackupRetentionIntervalInHours: An integer representing the time (in hours) that each backup is retained
	BackupRetentionIntervalInHours *int `json:"backupRetentionIntervalInHours,omitempty"`

	// BackupStorageRedundancy: Enum to indicate type of backup residency
	BackupStorageRedundancy *BackupStorageRedundancy_STATUS_ARM `json:"backupStorageRedundancy,omitempty"`
}

// Describes the status of migration between backup policy types.
type BackupPolicyMigrationStatus_STATUS_ARM string

const (
	BackupPolicyMigrationStatus_STATUS_ARM_Completed  = BackupPolicyMigrationStatus_STATUS_ARM("Completed")
	BackupPolicyMigrationStatus_STATUS_ARM_Failed     = BackupPolicyMigrationStatus_STATUS_ARM("Failed")
	BackupPolicyMigrationStatus_STATUS_ARM_InProgress = BackupPolicyMigrationStatus_STATUS_ARM("InProgress")
	BackupPolicyMigrationStatus_STATUS_ARM_Invalid    = BackupPolicyMigrationStatus_STATUS_ARM("Invalid")
)

// Mapping from string to BackupPolicyMigrationStatus_STATUS_ARM
var backupPolicyMigrationStatus_STATUS_ARM_Values = map[string]BackupPolicyMigrationStatus_STATUS_ARM{
	"completed":  BackupPolicyMigrationStatus_STATUS_ARM_Completed,
	"failed":     BackupPolicyMigrationStatus_STATUS_ARM_Failed,
	"inprogress": BackupPolicyMigrationStatus_STATUS_ARM_InProgress,
	"invalid":    BackupPolicyMigrationStatus_STATUS_ARM_Invalid,
}

// Describes the mode of backups.
type BackupPolicyType_STATUS_ARM string

const (
	BackupPolicyType_STATUS_ARM_Continuous = BackupPolicyType_STATUS_ARM("Continuous")
	BackupPolicyType_STATUS_ARM_Periodic   = BackupPolicyType_STATUS_ARM("Periodic")
)

// Mapping from string to BackupPolicyType_STATUS_ARM
var backupPolicyType_STATUS_ARM_Values = map[string]BackupPolicyType_STATUS_ARM{
	"continuous": BackupPolicyType_STATUS_ARM_Continuous,
	"periodic":   BackupPolicyType_STATUS_ARM_Periodic,
}

// Enum to indicate type of backup storage redundancy.
type BackupStorageRedundancy_STATUS_ARM string

const (
	BackupStorageRedundancy_STATUS_ARM_Geo   = BackupStorageRedundancy_STATUS_ARM("Geo")
	BackupStorageRedundancy_STATUS_ARM_Local = BackupStorageRedundancy_STATUS_ARM("Local")
	BackupStorageRedundancy_STATUS_ARM_Zone  = BackupStorageRedundancy_STATUS_ARM("Zone")
)

// Mapping from string to BackupStorageRedundancy_STATUS_ARM
var backupStorageRedundancy_STATUS_ARM_Values = map[string]BackupStorageRedundancy_STATUS_ARM{
	"geo":   BackupStorageRedundancy_STATUS_ARM_Geo,
	"local": BackupStorageRedundancy_STATUS_ARM_Local,
	"zone":  BackupStorageRedundancy_STATUS_ARM_Zone,
}

// Enum to indicate type of Continuous backup tier.
type ContinuousTier_STATUS_ARM string

const (
	ContinuousTier_STATUS_ARM_Continuous30Days = ContinuousTier_STATUS_ARM("Continuous30Days")
	ContinuousTier_STATUS_ARM_Continuous7Days  = ContinuousTier_STATUS_ARM("Continuous7Days")
)

// Mapping from string to ContinuousTier_STATUS_ARM
var continuousTier_STATUS_ARM_Values = map[string]ContinuousTier_STATUS_ARM{
	"continuous30days": ContinuousTier_STATUS_ARM_Continuous30Days,
	"continuous7days":  ContinuousTier_STATUS_ARM_Continuous7Days,
}
