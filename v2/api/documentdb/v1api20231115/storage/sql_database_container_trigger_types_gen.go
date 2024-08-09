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

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontainertriggers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontainertriggers/status,sqldatabasecontainertriggers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20231115.SqlDatabaseContainerTrigger
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}
type SqlDatabaseContainerTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseContainerTrigger_Spec   `json:"spec,omitempty"`
	Status            SqlDatabaseContainerTrigger_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerTrigger{}

// GetConditions returns the conditions of the resource
func (trigger *SqlDatabaseContainerTrigger) GetConditions() conditions.Conditions {
	return trigger.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (trigger *SqlDatabaseContainerTrigger) SetConditions(conditions conditions.Conditions) {
	trigger.Status.Conditions = conditions
}

var _ configmaps.Exporter = &SqlDatabaseContainerTrigger{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (trigger *SqlDatabaseContainerTrigger) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if trigger.Spec.OperatorSpec == nil {
		return nil
	}
	return trigger.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SqlDatabaseContainerTrigger{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (trigger *SqlDatabaseContainerTrigger) SecretDestinationExpressions() []*core.DestinationExpression {
	if trigger.Spec.OperatorSpec == nil {
		return nil
	}
	return trigger.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerTrigger{}

// AzureName returns the Azure name of the resource
func (trigger *SqlDatabaseContainerTrigger) AzureName() string {
	return trigger.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-15"
func (trigger SqlDatabaseContainerTrigger) GetAPIVersion() string {
	return "2023-11-15"
}

// GetResourceScope returns the scope of the resource
func (trigger *SqlDatabaseContainerTrigger) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (trigger *SqlDatabaseContainerTrigger) GetSpec() genruntime.ConvertibleSpec {
	return &trigger.Spec
}

// GetStatus returns the status of this resource
func (trigger *SqlDatabaseContainerTrigger) GetStatus() genruntime.ConvertibleStatus {
	return &trigger.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (trigger *SqlDatabaseContainerTrigger) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
func (trigger *SqlDatabaseContainerTrigger) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
}

// NewEmptyStatus returns a new empty (blank) status
func (trigger *SqlDatabaseContainerTrigger) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseContainerTrigger_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (trigger *SqlDatabaseContainerTrigger) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(trigger.Spec)
	return trigger.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (trigger *SqlDatabaseContainerTrigger) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseContainerTrigger_STATUS); ok {
		trigger.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseContainerTrigger_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	trigger.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainerTrigger is the hub type for conversion
func (trigger *SqlDatabaseContainerTrigger) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (trigger *SqlDatabaseContainerTrigger) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: trigger.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerTrigger",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20231115.SqlDatabaseContainerTrigger
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}
type SqlDatabaseContainerTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerTrigger `json:"items"`
}

// Storage version of v1api20231115.SqlDatabaseContainerTrigger_Spec
type SqlDatabaseContainerTrigger_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                   `json:"azureName,omitempty"`
	Location        *string                                  `json:"location,omitempty"`
	OperatorSpec    *SqlDatabaseContainerTriggerOperatorSpec `json:"operatorSpec,omitempty"`
	Options         *CreateUpdateOptions                     `json:"options,omitempty"`
	OriginalVersion string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlTriggerResource                `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SqlDatabaseContainerTrigger_Spec{}

// ConvertSpecFrom populates our SqlDatabaseContainerTrigger_Spec from the provided source
func (trigger *SqlDatabaseContainerTrigger_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == trigger {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(trigger)
}

// ConvertSpecTo populates the provided destination from our SqlDatabaseContainerTrigger_Spec
func (trigger *SqlDatabaseContainerTrigger_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == trigger {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(trigger)
}

// Storage version of v1api20231115.SqlDatabaseContainerTrigger_STATUS
type SqlDatabaseContainerTrigger_STATUS struct {
	Conditions  []conditions.Condition                   `json:"conditions,omitempty"`
	Id          *string                                  `json:"id,omitempty"`
	Location    *string                                  `json:"location,omitempty"`
	Name        *string                                  `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	Resource    *SqlTriggerGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                        `json:"tags,omitempty"`
	Type        *string                                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseContainerTrigger_STATUS{}

// ConvertStatusFrom populates our SqlDatabaseContainerTrigger_STATUS from the provided source
func (trigger *SqlDatabaseContainerTrigger_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == trigger {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(trigger)
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseContainerTrigger_STATUS
func (trigger *SqlDatabaseContainerTrigger_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == trigger {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(trigger)
}

// Storage version of v1api20231115.SqlDatabaseContainerTriggerOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SqlDatabaseContainerTriggerOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20231115.SqlTriggerGetProperties_Resource_STATUS
type SqlTriggerGetProperties_Resource_STATUS struct {
	Body             *string                `json:"body,omitempty"`
	Etag             *string                `json:"_etag,omitempty"`
	Id               *string                `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid              *string                `json:"_rid,omitempty"`
	TriggerOperation *string                `json:"triggerOperation,omitempty"`
	TriggerType      *string                `json:"triggerType,omitempty"`
	Ts               *float64               `json:"_ts,omitempty"`
}

// Storage version of v1api20231115.SqlTriggerResource
// Cosmos DB SQL trigger resource object
type SqlTriggerResource struct {
	Body             *string                `json:"body,omitempty"`
	Id               *string                `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TriggerOperation *string                `json:"triggerOperation,omitempty"`
	TriggerType      *string                `json:"triggerType,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerTrigger{}, &SqlDatabaseContainerTriggerList{})
}
