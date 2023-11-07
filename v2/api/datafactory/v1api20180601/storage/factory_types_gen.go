// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=datafactory.azure.com,resources=factories,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=datafactory.azure.com,resources={factories/status,factories/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20180601.Factory
// Generator information:
// - Generated from: /datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}
type Factory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Factory_Spec   `json:"spec,omitempty"`
	Status            Factory_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Factory{}

// GetConditions returns the conditions of the resource
func (factory *Factory) GetConditions() conditions.Conditions {
	return factory.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (factory *Factory) SetConditions(conditions conditions.Conditions) {
	factory.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Factory{}

// AzureName returns the Azure name of the resource
func (factory *Factory) AzureName() string {
	return factory.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (factory Factory) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (factory *Factory) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (factory *Factory) GetSpec() genruntime.ConvertibleSpec {
	return &factory.Spec
}

// GetStatus returns the status of this resource
func (factory *Factory) GetStatus() genruntime.ConvertibleStatus {
	return &factory.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (factory *Factory) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DataFactory/factories"
func (factory *Factory) GetType() string {
	return "Microsoft.DataFactory/factories"
}

// NewEmptyStatus returns a new empty (blank) status
func (factory *Factory) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Factory_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (factory *Factory) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(factory.Spec)
	return factory.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (factory *Factory) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Factory_STATUS); ok {
		factory.Status = *st
		return nil
	}

	// Convert status to required version
	var st Factory_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	factory.Status = st
	return nil
}

// Hub marks that this Factory is the hub type for conversion
func (factory *Factory) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (factory *Factory) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: factory.Spec.OriginalVersion,
		Kind:    "Factory",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20180601.Factory
// Generator information:
// - Generated from: /datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}
type FactoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Factory `json:"items"`
}

// Storage version of v1api20180601.APIVersion
// +kubebuilder:validation:Enum={"2018-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2018-06-01")

// Storage version of v1api20180601.Factory_Spec
type Factory_Spec struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:Pattern="^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string                                  `json:"azureName,omitempty"`
	Encryption       *EncryptionConfiguration                `json:"encryption,omitempty"`
	GlobalParameters map[string]GlobalParameterSpecification `json:"globalParameters,omitempty"`
	Identity         *FactoryIdentity                        `json:"identity,omitempty"`
	Location         *string                                 `json:"location,omitempty"`
	OriginalVersion  string                                  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag          genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess  *string                            `json:"publicNetworkAccess,omitempty"`
	PurviewConfiguration *PurviewConfiguration              `json:"purviewConfiguration,omitempty"`
	RepoConfiguration    *FactoryRepoConfiguration          `json:"repoConfiguration,omitempty"`
	Tags                 map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Factory_Spec{}

// ConvertSpecFrom populates our Factory_Spec from the provided source
func (factory *Factory_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == factory {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(factory)
}

// ConvertSpecTo populates the provided destination from our Factory_Spec
func (factory *Factory_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == factory {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(factory)
}

// Storage version of v1api20180601.Factory_STATUS
// Factory resource type.
type Factory_STATUS struct {
	AdditionalProperties map[string]v1.JSON                             `json:"additionalProperties,omitempty"`
	Conditions           []conditions.Condition                         `json:"conditions,omitempty"`
	CreateTime           *string                                        `json:"createTime,omitempty"`
	ETag                 *string                                        `json:"eTag,omitempty"`
	Encryption           *EncryptionConfiguration_STATUS                `json:"encryption,omitempty"`
	GlobalParameters     map[string]GlobalParameterSpecification_STATUS `json:"globalParameters,omitempty"`
	Id                   *string                                        `json:"id,omitempty"`
	Identity             *FactoryIdentity_STATUS                        `json:"identity,omitempty"`
	Location             *string                                        `json:"location,omitempty"`
	Name                 *string                                        `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag                         `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                                        `json:"provisioningState,omitempty"`
	PublicNetworkAccess  *string                                        `json:"publicNetworkAccess,omitempty"`
	PurviewConfiguration *PurviewConfiguration_STATUS                   `json:"purviewConfiguration,omitempty"`
	RepoConfiguration    *FactoryRepoConfiguration_STATUS               `json:"repoConfiguration,omitempty"`
	Tags                 map[string]string                              `json:"tags,omitempty"`
	Type                 *string                                        `json:"type,omitempty"`
	Version              *string                                        `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Factory_STATUS{}

// ConvertStatusFrom populates our Factory_STATUS from the provided source
func (factory *Factory_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == factory {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(factory)
}

// ConvertStatusTo populates the provided destination from our Factory_STATUS
func (factory *Factory_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == factory {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(factory)
}

// Storage version of v1api20180601.EncryptionConfiguration
// Definition of CMK for the factory.
type EncryptionConfiguration struct {
	Identity     *CMKIdentityDefinition `json:"identity,omitempty"`
	KeyName      *string                `json:"keyName,omitempty"`
	KeyVersion   *string                `json:"keyVersion,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VaultBaseUrl *string                `json:"vaultBaseUrl,omitempty"`
}

// Storage version of v1api20180601.EncryptionConfiguration_STATUS
// Definition of CMK for the factory.
type EncryptionConfiguration_STATUS struct {
	Identity     *CMKIdentityDefinition_STATUS `json:"identity,omitempty"`
	KeyName      *string                       `json:"keyName,omitempty"`
	KeyVersion   *string                       `json:"keyVersion,omitempty"`
	PropertyBag  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	VaultBaseUrl *string                       `json:"vaultBaseUrl,omitempty"`
}

// Storage version of v1api20180601.FactoryIdentity
// Identity properties of the factory resource.
type FactoryIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20180601.FactoryIdentity_STATUS
// Identity properties of the factory resource.
type FactoryIdentity_STATUS struct {
	PrincipalId            *string                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId               *string                `json:"tenantId,omitempty"`
	Type                   *string                `json:"type,omitempty"`
	UserAssignedIdentities map[string]v1.JSON     `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20180601.FactoryRepoConfiguration
type FactoryRepoConfiguration struct {
	FactoryGitHub *FactoryGitHubConfiguration `json:"factoryGitHubConfiguration,omitempty"`
	FactoryVSTS   *FactoryVSTSConfiguration   `json:"factoryVSTSConfiguration,omitempty"`
	PropertyBag   genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20180601.FactoryRepoConfiguration_STATUS
type FactoryRepoConfiguration_STATUS struct {
	FactoryGitHub *FactoryGitHubConfiguration_STATUS `json:"factoryGitHubConfiguration,omitempty"`
	FactoryVSTS   *FactoryVSTSConfiguration_STATUS   `json:"factoryVSTSConfiguration,omitempty"`
	PropertyBag   genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20180601.GlobalParameterSpecification
// Definition of a single parameter for an entity.
type GlobalParameterSpecification struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Value       map[string]v1.JSON     `json:"value,omitempty"`
}

// Storage version of v1api20180601.GlobalParameterSpecification_STATUS
// Definition of a single parameter for an entity.
type GlobalParameterSpecification_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Value       map[string]v1.JSON     `json:"value,omitempty"`
}

// Storage version of v1api20180601.PurviewConfiguration
// Purview configuration.
type PurviewConfiguration struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// PurviewResourceReference: Purview resource id.
	PurviewResourceReference *genruntime.ResourceReference `armReference:"PurviewResourceId" json:"purviewResourceReference,omitempty"`
}

// Storage version of v1api20180601.PurviewConfiguration_STATUS
// Purview configuration.
type PurviewConfiguration_STATUS struct {
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PurviewResourceId *string                `json:"purviewResourceId,omitempty"`
}

// Storage version of v1api20180601.CMKIdentityDefinition
// Managed Identity used for CMK.
type CMKIdentityDefinition struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// UserAssignedIdentityReference: The resource id of the user assigned identity to authenticate to customer's key vault.
	UserAssignedIdentityReference *genruntime.ResourceReference `armReference:"UserAssignedIdentity" json:"userAssignedIdentityReference,omitempty"`
}

// Storage version of v1api20180601.CMKIdentityDefinition_STATUS
// Managed Identity used for CMK.
type CMKIdentityDefinition_STATUS struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

// Storage version of v1api20180601.FactoryGitHubConfiguration
type FactoryGitHubConfiguration struct {
	AccountName         *string                `json:"accountName,omitempty"`
	ClientId            *string                `json:"clientId,omitempty"`
	ClientSecret        *GitHubClientSecret    `json:"clientSecret,omitempty"`
	CollaborationBranch *string                `json:"collaborationBranch,omitempty"`
	DisablePublish      *bool                  `json:"disablePublish,omitempty"`
	HostName            *string                `json:"hostName,omitempty"`
	LastCommitId        *string                `json:"lastCommitId,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RepositoryName      *string                `json:"repositoryName,omitempty"`
	RootFolder          *string                `json:"rootFolder,omitempty"`
	Type                *string                `json:"type,omitempty"`
}

// Storage version of v1api20180601.FactoryGitHubConfiguration_STATUS
type FactoryGitHubConfiguration_STATUS struct {
	AccountName         *string                    `json:"accountName,omitempty"`
	ClientId            *string                    `json:"clientId,omitempty"`
	ClientSecret        *GitHubClientSecret_STATUS `json:"clientSecret,omitempty"`
	CollaborationBranch *string                    `json:"collaborationBranch,omitempty"`
	DisablePublish      *bool                      `json:"disablePublish,omitempty"`
	HostName            *string                    `json:"hostName,omitempty"`
	LastCommitId        *string                    `json:"lastCommitId,omitempty"`
	PropertyBag         genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	RepositoryName      *string                    `json:"repositoryName,omitempty"`
	RootFolder          *string                    `json:"rootFolder,omitempty"`
	Type                *string                    `json:"type,omitempty"`
}

// Storage version of v1api20180601.FactoryVSTSConfiguration
type FactoryVSTSConfiguration struct {
	AccountName         *string                `json:"accountName,omitempty"`
	CollaborationBranch *string                `json:"collaborationBranch,omitempty"`
	DisablePublish      *bool                  `json:"disablePublish,omitempty"`
	LastCommitId        *string                `json:"lastCommitId,omitempty"`
	ProjectName         *string                `json:"projectName,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RepositoryName      *string                `json:"repositoryName,omitempty"`
	RootFolder          *string                `json:"rootFolder,omitempty"`
	TenantId            *string                `json:"tenantId,omitempty"`
	Type                *string                `json:"type,omitempty"`
}

// Storage version of v1api20180601.FactoryVSTSConfiguration_STATUS
type FactoryVSTSConfiguration_STATUS struct {
	AccountName         *string                `json:"accountName,omitempty"`
	CollaborationBranch *string                `json:"collaborationBranch,omitempty"`
	DisablePublish      *bool                  `json:"disablePublish,omitempty"`
	LastCommitId        *string                `json:"lastCommitId,omitempty"`
	ProjectName         *string                `json:"projectName,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RepositoryName      *string                `json:"repositoryName,omitempty"`
	RootFolder          *string                `json:"rootFolder,omitempty"`
	TenantId            *string                `json:"tenantId,omitempty"`
	Type                *string                `json:"type,omitempty"`
}

// Storage version of v1api20180601.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

// Storage version of v1api20180601.GitHubClientSecret
// Client secret information for factory's bring your own app repository configuration.
type GitHubClientSecret struct {
	ByoaSecretAkvUrl *string                `json:"byoaSecretAkvUrl,omitempty"`
	ByoaSecretName   *string                `json:"byoaSecretName,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20180601.GitHubClientSecret_STATUS
// Client secret information for factory's bring your own app repository configuration.
type GitHubClientSecret_STATUS struct {
	ByoaSecretAkvUrl *string                `json:"byoaSecretAkvUrl,omitempty"`
	ByoaSecretName   *string                `json:"byoaSecretName,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Factory{}, &FactoryList{})
}
