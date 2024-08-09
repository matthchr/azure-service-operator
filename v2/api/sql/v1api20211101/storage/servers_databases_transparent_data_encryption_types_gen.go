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

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversdatabasestransparentdataencryptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversdatabasestransparentdataencryptions/status,serversdatabasestransparentdataencryptions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersDatabasesTransparentDataEncryption
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/TransparentDataEncryptions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{tdeName}
type ServersDatabasesTransparentDataEncryption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServersDatabasesTransparentDataEncryption_Spec   `json:"spec,omitempty"`
	Status            ServersDatabasesTransparentDataEncryption_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesTransparentDataEncryption{}

// GetConditions returns the conditions of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetConditions() conditions.Conditions {
	return encryption.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (encryption *ServersDatabasesTransparentDataEncryption) SetConditions(conditions conditions.Conditions) {
	encryption.Status.Conditions = conditions
}

var _ configmaps.Exporter = &ServersDatabasesTransparentDataEncryption{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (encryption *ServersDatabasesTransparentDataEncryption) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if encryption.Spec.OperatorSpec == nil {
		return nil
	}
	return encryption.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ServersDatabasesTransparentDataEncryption{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (encryption *ServersDatabasesTransparentDataEncryption) SecretDestinationExpressions() []*core.DestinationExpression {
	if encryption.Spec.OperatorSpec == nil {
		return nil
	}
	return encryption.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ServersDatabasesTransparentDataEncryption{}

// AzureName returns the Azure name of the resource (always "current")
func (encryption *ServersDatabasesTransparentDataEncryption) AzureName() string {
	return "current"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (encryption ServersDatabasesTransparentDataEncryption) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetSpec() genruntime.ConvertibleSpec {
	return &encryption.Spec
}

// GetStatus returns the status of this resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetStatus() genruntime.ConvertibleStatus {
	return &encryption.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/transparentDataEncryption"
func (encryption *ServersDatabasesTransparentDataEncryption) GetType() string {
	return "Microsoft.Sql/servers/databases/transparentDataEncryption"
}

// NewEmptyStatus returns a new empty (blank) status
func (encryption *ServersDatabasesTransparentDataEncryption) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ServersDatabasesTransparentDataEncryption_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (encryption *ServersDatabasesTransparentDataEncryption) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(encryption.Spec)
	return encryption.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (encryption *ServersDatabasesTransparentDataEncryption) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ServersDatabasesTransparentDataEncryption_STATUS); ok {
		encryption.Status = *st
		return nil
	}

	// Convert status to required version
	var st ServersDatabasesTransparentDataEncryption_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	encryption.Status = st
	return nil
}

// Hub marks that this ServersDatabasesTransparentDataEncryption is the hub type for conversion
func (encryption *ServersDatabasesTransparentDataEncryption) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (encryption *ServersDatabasesTransparentDataEncryption) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: encryption.Spec.OriginalVersion,
		Kind:    "ServersDatabasesTransparentDataEncryption",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersDatabasesTransparentDataEncryption
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/TransparentDataEncryptions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{tdeName}
type ServersDatabasesTransparentDataEncryptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesTransparentDataEncryption `json:"items"`
}

// Storage version of v1api20211101.ServersDatabasesTransparentDataEncryption_Spec
type ServersDatabasesTransparentDataEncryption_Spec struct {
	OperatorSpec    *ServersDatabasesTransparentDataEncryptionOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                                 `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	State       *string                            `json:"state,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ServersDatabasesTransparentDataEncryption_Spec{}

// ConvertSpecFrom populates our ServersDatabasesTransparentDataEncryption_Spec from the provided source
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == encryption {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(encryption)
}

// ConvertSpecTo populates the provided destination from our ServersDatabasesTransparentDataEncryption_Spec
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == encryption {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(encryption)
}

// Storage version of v1api20211101.ServersDatabasesTransparentDataEncryption_STATUS
type ServersDatabasesTransparentDataEncryption_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ServersDatabasesTransparentDataEncryption_STATUS{}

// ConvertStatusFrom populates our ServersDatabasesTransparentDataEncryption_STATUS from the provided source
func (encryption *ServersDatabasesTransparentDataEncryption_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == encryption {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(encryption)
}

// ConvertStatusTo populates the provided destination from our ServersDatabasesTransparentDataEncryption_STATUS
func (encryption *ServersDatabasesTransparentDataEncryption_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == encryption {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(encryption)
}

// Storage version of v1api20211101.ServersDatabasesTransparentDataEncryptionOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServersDatabasesTransparentDataEncryptionOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServersDatabasesTransparentDataEncryption{}, &ServersDatabasesTransparentDataEncryptionList{})
}
