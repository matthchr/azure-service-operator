// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources=workspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources={workspaces/status,workspaces/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240401.Workspace
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Workspace_Spec   `json:"spec,omitempty"`
	Status            Workspace_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Workspace{}

// GetConditions returns the conditions of the resource
func (workspace *Workspace) GetConditions() conditions.Conditions {
	return workspace.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (workspace *Workspace) SetConditions(conditions conditions.Conditions) {
	workspace.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Workspace{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (workspace *Workspace) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if workspace.Spec.OperatorSpec == nil {
		return nil
	}
	return workspace.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Workspace{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (workspace *Workspace) SecretDestinationExpressions() []*core.DestinationExpression {
	if workspace.Spec.OperatorSpec == nil {
		return nil
	}
	return workspace.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Workspace{}

// AzureName returns the Azure name of the resource
func (workspace *Workspace) AzureName() string {
	return workspace.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-04-01"
func (workspace Workspace) GetAPIVersion() string {
	return "2024-04-01"
}

// GetResourceScope returns the scope of the resource
func (workspace *Workspace) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (workspace *Workspace) GetSpec() genruntime.ConvertibleSpec {
	return &workspace.Spec
}

// GetStatus returns the status of this resource
func (workspace *Workspace) GetStatus() genruntime.ConvertibleStatus {
	return &workspace.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (workspace *Workspace) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces"
func (workspace *Workspace) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces"
}

// NewEmptyStatus returns a new empty (blank) status
func (workspace *Workspace) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Workspace_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (workspace *Workspace) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(workspace.Spec)
	return workspace.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (workspace *Workspace) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Workspace_STATUS); ok {
		workspace.Status = *st
		return nil
	}

	// Convert status to required version
	var st Workspace_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	workspace.Status = st
	return nil
}

// Hub marks that this Workspace is the hub type for conversion
func (workspace *Workspace) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (workspace *Workspace) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: workspace.Spec.OriginalVersion,
		Kind:    "Workspace",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240401.Workspace
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Storage version of v1api20240401.Workspace_Spec
type Workspace_Spec struct {
	AllowPublicAccessWhenBehindVnet *bool `json:"allowPublicAccessWhenBehindVnet,omitempty"`

	// ApplicationInsightsReference: ARM id of the application insights associated with this workspace.
	ApplicationInsightsReference *genruntime.ResourceReference `armReference:"ApplicationInsights" json:"applicationInsightsReference,omitempty"`
	AssociatedWorkspaces         []string                      `json:"associatedWorkspaces,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// ContainerRegistryReference: ARM id of the container registry associated with this workspace.
	ContainerRegistryReference *genruntime.ResourceReference `armReference:"ContainerRegistry" json:"containerRegistryReference,omitempty"`
	Description                *string                       `json:"description,omitempty"`
	DiscoveryUrl               *string                       `json:"discoveryUrl,omitempty"`
	EnableDataIsolation        *bool                         `json:"enableDataIsolation,omitempty"`
	Encryption                 *EncryptionProperty           `json:"encryption,omitempty"`
	FeatureStoreSettings       *FeatureStoreSettings         `json:"featureStoreSettings,omitempty"`
	FriendlyName               *string                       `json:"friendlyName,omitempty"`
	HbiWorkspace               *bool                         `json:"hbiWorkspace,omitempty"`
	HubResourceReference       *genruntime.ResourceReference `armReference:"HubResourceId" json:"hubResourceReference,omitempty"`
	Identity                   *ManagedServiceIdentity       `json:"identity,omitempty"`
	ImageBuildCompute          *string                       `json:"imageBuildCompute,omitempty"`

	// KeyVaultReference: ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has
	// been created
	KeyVaultReference *genruntime.ResourceReference `armReference:"KeyVault" json:"keyVaultReference,omitempty"`
	Kind              *string                       `json:"kind,omitempty"`
	Location          *string                       `json:"location,omitempty"`
	ManagedNetwork    *ManagedNetworkSettings       `json:"managedNetwork,omitempty"`
	OperatorSpec      *WorkspaceOperatorSpec        `json:"operatorSpec,omitempty"`
	OriginalVersion   string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// PrimaryUserAssignedIdentityReference: The user assigned identity resource id that represents the workspace identity.
	PrimaryUserAssignedIdentityReference *genruntime.ResourceReference    `armReference:"PrimaryUserAssignedIdentity" json:"primaryUserAssignedIdentityReference,omitempty"`
	PropertyBag                          genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	PublicNetworkAccess                  *string                          `json:"publicNetworkAccess,omitempty"`
	ServerlessComputeSettings            *ServerlessComputeSettings       `json:"serverlessComputeSettings,omitempty"`
	ServiceManagedResourcesSettings      *ServiceManagedResourcesSettings `json:"serviceManagedResourcesSettings,omitempty"`
	SharedPrivateLinkResources           []SharedPrivateLinkResource      `json:"sharedPrivateLinkResources,omitempty"`
	Sku                                  *Sku                             `json:"sku,omitempty"`

	// StorageAccountReference: ARM id of the storage account associated with this workspace. This cannot be changed once the
	// workspace has been created
	StorageAccountReference *genruntime.ResourceReference `armReference:"StorageAccount" json:"storageAccountReference,omitempty"`
	Tags                    map[string]string             `json:"tags,omitempty"`
	V1LegacyMode            *bool                         `json:"v1LegacyMode,omitempty"`
	WorkspaceHubConfig      *WorkspaceHubConfig           `json:"workspaceHubConfig,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Workspace_Spec{}

// ConvertSpecFrom populates our Workspace_Spec from the provided source
func (workspace *Workspace_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(workspace)
}

// ConvertSpecTo populates the provided destination from our Workspace_Spec
func (workspace *Workspace_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(workspace)
}

// Storage version of v1api20240401.Workspace_STATUS
// An object that represents a machine learning workspace.
type Workspace_STATUS struct {
	AllowPublicAccessWhenBehindVnet *bool                                   `json:"allowPublicAccessWhenBehindVnet,omitempty"`
	ApplicationInsights             *string                                 `json:"applicationInsights,omitempty"`
	AssociatedWorkspaces            []string                                `json:"associatedWorkspaces,omitempty"`
	Conditions                      []conditions.Condition                  `json:"conditions,omitempty"`
	ContainerRegistry               *string                                 `json:"containerRegistry,omitempty"`
	Description                     *string                                 `json:"description,omitempty"`
	DiscoveryUrl                    *string                                 `json:"discoveryUrl,omitempty"`
	EnableDataIsolation             *bool                                   `json:"enableDataIsolation,omitempty"`
	Encryption                      *EncryptionProperty_STATUS              `json:"encryption,omitempty"`
	FeatureStoreSettings            *FeatureStoreSettings_STATUS            `json:"featureStoreSettings,omitempty"`
	FriendlyName                    *string                                 `json:"friendlyName,omitempty"`
	HbiWorkspace                    *bool                                   `json:"hbiWorkspace,omitempty"`
	HubResourceId                   *string                                 `json:"hubResourceId,omitempty"`
	Id                              *string                                 `json:"id,omitempty"`
	Identity                        *ManagedServiceIdentity_STATUS          `json:"identity,omitempty"`
	ImageBuildCompute               *string                                 `json:"imageBuildCompute,omitempty"`
	KeyVault                        *string                                 `json:"keyVault,omitempty"`
	Kind                            *string                                 `json:"kind,omitempty"`
	Location                        *string                                 `json:"location,omitempty"`
	ManagedNetwork                  *ManagedNetworkSettings_STATUS          `json:"managedNetwork,omitempty"`
	MlFlowTrackingUri               *string                                 `json:"mlFlowTrackingUri,omitempty"`
	Name                            *string                                 `json:"name,omitempty"`
	NotebookInfo                    *NotebookResourceInfo_STATUS            `json:"notebookInfo,omitempty"`
	PrimaryUserAssignedIdentity     *string                                 `json:"primaryUserAssignedIdentity,omitempty"`
	PrivateEndpointConnections      []PrivateEndpointConnection_STATUS      `json:"privateEndpointConnections,omitempty"`
	PrivateLinkCount                *int                                    `json:"privateLinkCount,omitempty"`
	PropertyBag                     genruntime.PropertyBag                  `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                                 `json:"provisioningState,omitempty"`
	PublicNetworkAccess             *string                                 `json:"publicNetworkAccess,omitempty"`
	ServerlessComputeSettings       *ServerlessComputeSettings_STATUS       `json:"serverlessComputeSettings,omitempty"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings_STATUS `json:"serviceManagedResourcesSettings,omitempty"`
	ServiceProvisionedResourceGroup *string                                 `json:"serviceProvisionedResourceGroup,omitempty"`
	SharedPrivateLinkResources      []SharedPrivateLinkResource_STATUS      `json:"sharedPrivateLinkResources,omitempty"`
	Sku                             *Sku_STATUS                             `json:"sku,omitempty"`
	StorageAccount                  *string                                 `json:"storageAccount,omitempty"`
	StorageHnsEnabled               *bool                                   `json:"storageHnsEnabled,omitempty"`
	SystemData                      *SystemData_STATUS                      `json:"systemData,omitempty"`
	Tags                            map[string]string                       `json:"tags,omitempty"`
	TenantId                        *string                                 `json:"tenantId,omitempty"`
	Type                            *string                                 `json:"type,omitempty"`
	V1LegacyMode                    *bool                                   `json:"v1LegacyMode,omitempty"`
	WorkspaceHubConfig              *WorkspaceHubConfig_STATUS              `json:"workspaceHubConfig,omitempty"`
	WorkspaceId                     *string                                 `json:"workspaceId,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Workspace_STATUS{}

// ConvertStatusFrom populates our Workspace_STATUS from the provided source
func (workspace *Workspace_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(workspace)
}

// ConvertStatusTo populates the provided destination from our Workspace_STATUS
func (workspace *Workspace_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(workspace)
}

// Storage version of v1api20240401.EncryptionProperty
type EncryptionProperty struct {
	Identity           *IdentityForCmk               `json:"identity,omitempty"`
	KeyVaultProperties *EncryptionKeyVaultProperties `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.EncryptionProperty_STATUS
type EncryptionProperty_STATUS struct {
	Identity           *IdentityForCmk_STATUS               `json:"identity,omitempty"`
	KeyVaultProperties *EncryptionKeyVaultProperties_STATUS `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	Status             *string                              `json:"status,omitempty"`
}

// Storage version of v1api20240401.FeatureStoreSettings
// Settings for feature store type workspace.
type FeatureStoreSettings struct {
	ComputeRuntime             *ComputeRuntimeDto     `json:"computeRuntime,omitempty"`
	OfflineStoreConnectionName *string                `json:"offlineStoreConnectionName,omitempty"`
	OnlineStoreConnectionName  *string                `json:"onlineStoreConnectionName,omitempty"`
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.FeatureStoreSettings_STATUS
// Settings for feature store type workspace.
type FeatureStoreSettings_STATUS struct {
	ComputeRuntime             *ComputeRuntimeDto_STATUS `json:"computeRuntime,omitempty"`
	OfflineStoreConnectionName *string                   `json:"offlineStoreConnectionName,omitempty"`
	OnlineStoreConnectionName  *string                   `json:"onlineStoreConnectionName,omitempty"`
	PropertyBag                genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.ManagedNetworkSettings
// Managed Network settings for a machine learning workspace.
type ManagedNetworkSettings struct {
	IsolationMode *string                        `json:"isolationMode,omitempty"`
	OutboundRules map[string]OutboundRule        `json:"outboundRules,omitempty"`
	PropertyBag   genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	Status        *ManagedNetworkProvisionStatus `json:"status,omitempty"`
}

// Storage version of v1api20240401.ManagedNetworkSettings_STATUS
// Managed Network settings for a machine learning workspace.
type ManagedNetworkSettings_STATUS struct {
	IsolationMode *string                               `json:"isolationMode,omitempty"`
	NetworkId     *string                               `json:"networkId,omitempty"`
	OutboundRules map[string]OutboundRule_STATUS        `json:"outboundRules,omitempty"`
	PropertyBag   genruntime.PropertyBag                `json:"$propertyBag,omitempty"`
	Status        *ManagedNetworkProvisionStatus_STATUS `json:"status,omitempty"`
}

// Storage version of v1api20240401.NotebookResourceInfo_STATUS
type NotebookResourceInfo_STATUS struct {
	Fqdn                     *string                          `json:"fqdn,omitempty"`
	NotebookPreparationError *NotebookPreparationError_STATUS `json:"notebookPreparationError,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ResourceId               *string                          `json:"resourceId,omitempty"`
}

// Storage version of v1api20240401.PrivateEndpointConnection_STATUS
// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.ServerlessComputeSettings
type ServerlessComputeSettings struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ServerlessComputeCustomSubnetReference: The resource ID of an existing virtual network subnet in which serverless
	// compute nodes should be deployed
	ServerlessComputeCustomSubnetReference *genruntime.ResourceReference `armReference:"ServerlessComputeCustomSubnet" json:"serverlessComputeCustomSubnetReference,omitempty"`
	ServerlessComputeNoPublicIP            *bool                         `json:"serverlessComputeNoPublicIP,omitempty"`
}

// Storage version of v1api20240401.ServerlessComputeSettings_STATUS
type ServerlessComputeSettings_STATUS struct {
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServerlessComputeCustomSubnet *string                `json:"serverlessComputeCustomSubnet,omitempty"`
	ServerlessComputeNoPublicIP   *bool                  `json:"serverlessComputeNoPublicIP,omitempty"`
}

// Storage version of v1api20240401.ServiceManagedResourcesSettings
type ServiceManagedResourcesSettings struct {
	CosmosDb    *CosmosDbSettings      `json:"cosmosDb,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.ServiceManagedResourcesSettings_STATUS
type ServiceManagedResourcesSettings_STATUS struct {
	CosmosDb    *CosmosDbSettings_STATUS `json:"cosmosDb,omitempty"`
	PropertyBag genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.SharedPrivateLinkResource
type SharedPrivateLinkResource struct {
	GroupId *string `json:"groupId,omitempty"`
	Name    *string `json:"name,omitempty"`

	// PrivateLinkResourceReference: The resource id that private link links to.
	PrivateLinkResourceReference *genruntime.ResourceReference `armReference:"PrivateLinkResourceId" json:"privateLinkResourceReference,omitempty"`
	PropertyBag                  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RequestMessage               *string                       `json:"requestMessage,omitempty"`
	Status                       *string                       `json:"status,omitempty"`
}

// Storage version of v1api20240401.SharedPrivateLinkResource_STATUS
type SharedPrivateLinkResource_STATUS struct {
	GroupId               *string                `json:"groupId,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	PrivateLinkResourceId *string                `json:"privateLinkResourceId,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestMessage        *string                `json:"requestMessage,omitempty"`
	Status                *string                `json:"status,omitempty"`
}

// Storage version of v1api20240401.WorkspaceHubConfig
// WorkspaceHub's configuration object.
type WorkspaceHubConfig struct {
	AdditionalWorkspaceStorageAccounts []string               `json:"additionalWorkspaceStorageAccounts,omitempty"`
	DefaultWorkspaceResourceGroup      *string                `json:"defaultWorkspaceResourceGroup,omitempty"`
	PropertyBag                        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.WorkspaceHubConfig_STATUS
// WorkspaceHub's configuration object.
type WorkspaceHubConfig_STATUS struct {
	AdditionalWorkspaceStorageAccounts []string               `json:"additionalWorkspaceStorageAccounts,omitempty"`
	DefaultWorkspaceResourceGroup      *string                `json:"defaultWorkspaceResourceGroup,omitempty"`
	PropertyBag                        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.WorkspaceOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type WorkspaceOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
	Secrets              *WorkspaceOperatorSecrets     `json:"secrets,omitempty"`
}

// Storage version of v1api20240401.ComputeRuntimeDto
// Compute runtime config for feature store type workspace.
type ComputeRuntimeDto struct {
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SparkRuntimeVersion *string                `json:"sparkRuntimeVersion,omitempty"`
}

// Storage version of v1api20240401.ComputeRuntimeDto_STATUS
// Compute runtime config for feature store type workspace.
type ComputeRuntimeDto_STATUS struct {
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SparkRuntimeVersion *string                `json:"sparkRuntimeVersion,omitempty"`
}

// Storage version of v1api20240401.CosmosDbSettings
type CosmosDbSettings struct {
	CollectionsThroughput *int                   `json:"collectionsThroughput,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.CosmosDbSettings_STATUS
type CosmosDbSettings_STATUS struct {
	CollectionsThroughput *int                   `json:"collectionsThroughput,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.EncryptionKeyVaultProperties
type EncryptionKeyVaultProperties struct {
	IdentityClientId           *string                        `json:"identityClientId,omitempty" optionalConfigMapPair:"IdentityClientId"`
	IdentityClientIdFromConfig *genruntime.ConfigMapReference `json:"identityClientIdFromConfig,omitempty" optionalConfigMapPair:"IdentityClientId"`
	KeyIdentifier              *string                        `json:"keyIdentifier,omitempty"`

	// +kubebuilder:validation:Required
	// KeyVaultArmReference: The ArmId of the keyVault where the customer owned encryption key is present.
	KeyVaultArmReference *genruntime.ResourceReference `armReference:"KeyVaultArmId" json:"keyVaultArmReference,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.EncryptionKeyVaultProperties_STATUS
type EncryptionKeyVaultProperties_STATUS struct {
	IdentityClientId *string                `json:"identityClientId,omitempty"`
	KeyIdentifier    *string                `json:"keyIdentifier,omitempty"`
	KeyVaultArmId    *string                `json:"keyVaultArmId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.IdentityForCmk
// Identity that will be used to access key vault for encryption at rest
type IdentityForCmk struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// UserAssignedIdentityReference: The ArmId of the user assigned identity that will be used to access the customer managed
	// key vault
	UserAssignedIdentityReference *genruntime.ResourceReference `armReference:"UserAssignedIdentity" json:"userAssignedIdentityReference,omitempty"`
}

// Storage version of v1api20240401.IdentityForCmk_STATUS
// Identity that will be used to access key vault for encryption at rest
type IdentityForCmk_STATUS struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

// Storage version of v1api20240401.ManagedNetworkProvisionStatus
// Status of the Provisioning for the managed network of a machine learning workspace.
type ManagedNetworkProvisionStatus struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SparkReady  *bool                  `json:"sparkReady,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

// Storage version of v1api20240401.ManagedNetworkProvisionStatus_STATUS
// Status of the Provisioning for the managed network of a machine learning workspace.
type ManagedNetworkProvisionStatus_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SparkReady  *bool                  `json:"sparkReady,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

// Storage version of v1api20240401.NotebookPreparationError_STATUS
type NotebookPreparationError_STATUS struct {
	ErrorMessage *string                `json:"errorMessage,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StatusCode   *int                   `json:"statusCode,omitempty"`
}

// Storage version of v1api20240401.OutboundRule
type OutboundRule struct {
	FQDN            *FqdnOutboundRule            `json:"fqdn,omitempty"`
	PrivateEndpoint *PrivateEndpointOutboundRule `json:"privateEndpoint,omitempty"`
	PropertyBag     genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	ServiceTag      *ServiceTagOutboundRule      `json:"serviceTag,omitempty"`
}

// Storage version of v1api20240401.OutboundRule_STATUS
type OutboundRule_STATUS struct {
	FQDN            *FqdnOutboundRule_STATUS            `json:"fqdn,omitempty"`
	PrivateEndpoint *PrivateEndpointOutboundRule_STATUS `json:"privateEndpoint,omitempty"`
	PropertyBag     genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	ServiceTag      *ServiceTagOutboundRule_STATUS      `json:"serviceTag,omitempty"`
}

// Storage version of v1api20240401.WorkspaceOperatorSecrets
type WorkspaceOperatorSecrets struct {
	AppInsightsInstrumentationKey *genruntime.SecretDestination `json:"appInsightsInstrumentationKey,omitempty"`
	ContainerRegistryPassword     *genruntime.SecretDestination `json:"containerRegistryPassword,omitempty"`
	ContainerRegistryPassword2    *genruntime.SecretDestination `json:"containerRegistryPassword2,omitempty"`
	ContainerRegistryUserName     *genruntime.SecretDestination `json:"containerRegistryUserName,omitempty"`
	PrimaryNotebookAccessKey      *genruntime.SecretDestination `json:"primaryNotebookAccessKey,omitempty"`
	PropertyBag                   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecondaryNotebookAccessKey    *genruntime.SecretDestination `json:"secondaryNotebookAccessKey,omitempty"`
	UserStorageKey                *genruntime.SecretDestination `json:"userStorageKey,omitempty"`
}

// Storage version of v1api20240401.FqdnOutboundRule
type FqdnOutboundRule struct {
	Category    *string                `json:"category,omitempty"`
	Destination *string                `json:"destination,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20240401.FqdnOutboundRule_STATUS
type FqdnOutboundRule_STATUS struct {
	Category    *string                `json:"category,omitempty"`
	Destination *string                `json:"destination,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20240401.PrivateEndpointOutboundRule
type PrivateEndpointOutboundRule struct {
	Category    *string                     `json:"category,omitempty"`
	Destination *PrivateEndpointDestination `json:"destination,omitempty"`
	PropertyBag genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	Status      *string                     `json:"status,omitempty"`
	Type        *string                     `json:"type,omitempty"`
}

// Storage version of v1api20240401.PrivateEndpointOutboundRule_STATUS
type PrivateEndpointOutboundRule_STATUS struct {
	Category    *string                            `json:"category,omitempty"`
	Destination *PrivateEndpointDestination_STATUS `json:"destination,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Status      *string                            `json:"status,omitempty"`
	Type        *string                            `json:"type,omitempty"`
}

// Storage version of v1api20240401.ServiceTagOutboundRule
type ServiceTagOutboundRule struct {
	Category    *string                `json:"category,omitempty"`
	Destination *ServiceTagDestination `json:"destination,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20240401.ServiceTagOutboundRule_STATUS
type ServiceTagOutboundRule_STATUS struct {
	Category    *string                       `json:"category,omitempty"`
	Destination *ServiceTagDestination_STATUS `json:"destination,omitempty"`
	PropertyBag genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Status      *string                       `json:"status,omitempty"`
	Type        *string                       `json:"type,omitempty"`
}

// Storage version of v1api20240401.PrivateEndpointDestination
// Private Endpoint destination for a Private Endpoint Outbound Rule for the managed network of a machine learning
// workspace.
type PrivateEndpointDestination struct {
	PropertyBag                genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	ServiceResourceReference   *genruntime.ResourceReference `armReference:"ServiceResourceId" json:"serviceResourceReference,omitempty"`
	SparkEnabled               *bool                         `json:"sparkEnabled,omitempty"`
	SparkStatus                *string                       `json:"sparkStatus,omitempty"`
	SubresourceTargetReference *genruntime.ResourceReference `armReference:"SubresourceTarget" json:"subresourceTargetReference,omitempty"`
}

// Storage version of v1api20240401.PrivateEndpointDestination_STATUS
// Private Endpoint destination for a Private Endpoint Outbound Rule for the managed network of a machine learning
// workspace.
type PrivateEndpointDestination_STATUS struct {
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServiceResourceId *string                `json:"serviceResourceId,omitempty"`
	SparkEnabled      *bool                  `json:"sparkEnabled,omitempty"`
	SparkStatus       *string                `json:"sparkStatus,omitempty"`
	SubresourceTarget *string                `json:"subresourceTarget,omitempty"`
}

// Storage version of v1api20240401.ServiceTagDestination
// Service Tag destination for a Service Tag Outbound Rule for the managed network of a machine learning workspace.
type ServiceTagDestination struct {
	Action      *string                `json:"action,omitempty"`
	PortRanges  *string                `json:"portRanges,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol    *string                `json:"protocol,omitempty"`
	ServiceTag  *string                `json:"serviceTag,omitempty"`
}

// Storage version of v1api20240401.ServiceTagDestination_STATUS
// Service Tag destination for a Service Tag Outbound Rule for the managed network of a machine learning workspace.
type ServiceTagDestination_STATUS struct {
	Action          *string                `json:"action,omitempty"`
	AddressPrefixes []string               `json:"addressPrefixes,omitempty"`
	PortRanges      *string                `json:"portRanges,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol        *string                `json:"protocol,omitempty"`
	ServiceTag      *string                `json:"serviceTag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
