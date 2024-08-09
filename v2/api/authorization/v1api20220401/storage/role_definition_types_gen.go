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

// +kubebuilder:rbac:groups=authorization.azure.com,resources=roledefinitions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=authorization.azure.com,resources={roledefinitions/status,roledefinitions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220401.RoleDefinition
// Generator information:
// - Generated from: /authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/authorization-RoleDefinitionsCalls.json
// - ARM URI: /{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}
type RoleDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleDefinition_Spec   `json:"spec,omitempty"`
	Status            RoleDefinition_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RoleDefinition{}

// GetConditions returns the conditions of the resource
func (definition *RoleDefinition) GetConditions() conditions.Conditions {
	return definition.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (definition *RoleDefinition) SetConditions(conditions conditions.Conditions) {
	definition.Status.Conditions = conditions
}

var _ configmaps.Exporter = &RoleDefinition{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (definition *RoleDefinition) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if definition.Spec.OperatorSpec == nil {
		return nil
	}
	return definition.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &RoleDefinition{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (definition *RoleDefinition) SecretDestinationExpressions() []*core.DestinationExpression {
	if definition.Spec.OperatorSpec == nil {
		return nil
	}
	return definition.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &RoleDefinition{}

// AzureName returns the Azure name of the resource
func (definition *RoleDefinition) AzureName() string {
	return definition.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-04-01"
func (definition RoleDefinition) GetAPIVersion() string {
	return "2022-04-01"
}

// GetResourceScope returns the scope of the resource
func (definition *RoleDefinition) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeExtension
}

// GetSpec returns the specification of this resource
func (definition *RoleDefinition) GetSpec() genruntime.ConvertibleSpec {
	return &definition.Spec
}

// GetStatus returns the status of this resource
func (definition *RoleDefinition) GetStatus() genruntime.ConvertibleStatus {
	return &definition.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (definition *RoleDefinition) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Authorization/roleDefinitions"
func (definition *RoleDefinition) GetType() string {
	return "Microsoft.Authorization/roleDefinitions"
}

// NewEmptyStatus returns a new empty (blank) status
func (definition *RoleDefinition) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RoleDefinition_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (definition *RoleDefinition) Owner() *genruntime.ResourceReference {
	return definition.Spec.Owner.AsResourceReference()
}

// SetStatus sets the status of this resource
func (definition *RoleDefinition) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RoleDefinition_STATUS); ok {
		definition.Status = *st
		return nil
	}

	// Convert status to required version
	var st RoleDefinition_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	definition.Status = st
	return nil
}

// Hub marks that this RoleDefinition is the hub type for conversion
func (definition *RoleDefinition) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (definition *RoleDefinition) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: definition.Spec.OriginalVersion,
		Kind:    "RoleDefinition",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220401.RoleDefinition
// Generator information:
// - Generated from: /authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/authorization-RoleDefinitionsCalls.json
// - ARM URI: /{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}
type RoleDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleDefinition `json:"items"`
}

// Storage version of v1api20220401.RoleDefinition_Spec
type RoleDefinition_Spec struct {
	AssignableScopesReferences []genruntime.ResourceReference `armReference:"AssignableScopes" json:"assignableScopesReferences,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                      `json:"azureName,omitempty"`
	Description     *string                     `json:"description,omitempty"`
	OperatorSpec    *RoleDefinitionOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                      `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
	// extension resource, which means that any other Azure resource can be its owner.
	Owner       *genruntime.ArbitraryOwnerReference `json:"owner,omitempty"`
	Permissions []Permission                        `json:"permissions,omitempty"`
	PropertyBag genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	RoleName    *string                             `json:"roleName,omitempty"`
	Type        *string                             `json:"type,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RoleDefinition_Spec{}

// ConvertSpecFrom populates our RoleDefinition_Spec from the provided source
func (definition *RoleDefinition_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == definition {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(definition)
}

// ConvertSpecTo populates the provided destination from our RoleDefinition_Spec
func (definition *RoleDefinition_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == definition {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(definition)
}

// Storage version of v1api20220401.RoleDefinition_STATUS
// Role definition.
type RoleDefinition_STATUS struct {
	AssignableScopes []string               `json:"assignableScopes,omitempty"`
	Conditions       []conditions.Condition `json:"conditions,omitempty"`
	CreatedBy        *string                `json:"createdBy,omitempty"`
	CreatedOn        *string                `json:"createdOn,omitempty"`
	Description      *string                `json:"description,omitempty"`
	Id               *string                `json:"id,omitempty"`
	Name             *string                `json:"name,omitempty"`
	Permissions      []Permission_STATUS    `json:"permissions,omitempty"`
	PropertiesType   *string                `json:"properties_type,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RoleName         *string                `json:"roleName,omitempty"`
	Type             *string                `json:"type,omitempty"`
	UpdatedBy        *string                `json:"updatedBy,omitempty"`
	UpdatedOn        *string                `json:"updatedOn,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RoleDefinition_STATUS{}

// ConvertStatusFrom populates our RoleDefinition_STATUS from the provided source
func (definition *RoleDefinition_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == definition {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(definition)
}

// ConvertStatusTo populates the provided destination from our RoleDefinition_STATUS
func (definition *RoleDefinition_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == definition {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(definition)
}

// Storage version of v1api20220401.Permission
// Role definition permissions.
type Permission struct {
	Actions        []string               `json:"actions,omitempty"`
	DataActions    []string               `json:"dataActions,omitempty"`
	NotActions     []string               `json:"notActions,omitempty"`
	NotDataActions []string               `json:"notDataActions,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220401.Permission_STATUS
// Role definition permissions.
type Permission_STATUS struct {
	Actions        []string               `json:"actions,omitempty"`
	DataActions    []string               `json:"dataActions,omitempty"`
	NotActions     []string               `json:"notActions,omitempty"`
	NotDataActions []string               `json:"notDataActions,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220401.RoleDefinitionOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RoleDefinitionOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	NamingConvention     *string                       `json:"namingConvention,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RoleDefinition{}, &RoleDefinitionList{})
}
