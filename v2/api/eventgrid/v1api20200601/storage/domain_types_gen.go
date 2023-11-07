// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=domains,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={domains/status,domains/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20200601.Domain
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Domain_Spec   `json:"spec,omitempty"`
	Status            Domain_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Domain{}

// GetConditions returns the conditions of the resource
func (domain *Domain) GetConditions() conditions.Conditions {
	return domain.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (domain *Domain) SetConditions(conditions conditions.Conditions) {
	domain.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Domain{}

// AzureName returns the Azure name of the resource
func (domain *Domain) AzureName() string {
	return domain.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (domain Domain) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (domain *Domain) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (domain *Domain) GetSpec() genruntime.ConvertibleSpec {
	return &domain.Spec
}

// GetStatus returns the status of this resource
func (domain *Domain) GetStatus() genruntime.ConvertibleStatus {
	return &domain.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (domain *Domain) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains"
func (domain *Domain) GetType() string {
	return "Microsoft.EventGrid/domains"
}

// NewEmptyStatus returns a new empty (blank) status
func (domain *Domain) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Domain_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (domain *Domain) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(domain.Spec)
	return domain.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (domain *Domain) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Domain_STATUS); ok {
		domain.Status = *st
		return nil
	}

	// Convert status to required version
	var st Domain_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	domain.Status = st
	return nil
}

// Hub marks that this Domain is the hub type for conversion
func (domain *Domain) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (domain *Domain) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: domain.Spec.OriginalVersion,
		Kind:    "Domain",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20200601.Domain
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Storage version of v1api20200601.APIVersion
// +kubebuilder:validation:Enum={"2020-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-06-01")

// Storage version of v1api20200601.Domain_Spec
type Domain_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName          string              `json:"azureName,omitempty"`
	InboundIpRules     []InboundIpRule     `json:"inboundIpRules,omitempty"`
	InputSchema        *string             `json:"inputSchema,omitempty"`
	InputSchemaMapping *InputSchemaMapping `json:"inputSchemaMapping,omitempty"`
	Location           *string             `json:"location,omitempty"`
	OriginalVersion    string              `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                            `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Domain_Spec{}

// ConvertSpecFrom populates our Domain_Spec from the provided source
func (domain *Domain_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(domain)
}

// ConvertSpecTo populates the provided destination from our Domain_Spec
func (domain *Domain_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(domain)
}

// Storage version of v1api20200601.Domain_STATUS
// EventGrid Domain.
type Domain_STATUS struct {
	Conditions                 []conditions.Condition                                        `json:"conditions,omitempty"`
	Endpoint                   *string                                                       `json:"endpoint,omitempty"`
	Id                         *string                                                       `json:"id,omitempty"`
	InboundIpRules             []InboundIpRule_STATUS                                        `json:"inboundIpRules,omitempty"`
	InputSchema                *string                                                       `json:"inputSchema,omitempty"`
	InputSchemaMapping         *InputSchemaMapping_STATUS                                    `json:"inputSchemaMapping,omitempty"`
	Location                   *string                                                       `json:"location,omitempty"`
	MetricResourceId           *string                                                       `json:"metricResourceId,omitempty"`
	Name                       *string                                                       `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_Domain_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                        `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                       `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                       `json:"publicNetworkAccess,omitempty"`
	SystemData                 *SystemData_STATUS                                            `json:"systemData,omitempty"`
	Tags                       map[string]string                                             `json:"tags,omitempty"`
	Type                       *string                                                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Domain_STATUS{}

// ConvertStatusFrom populates our Domain_STATUS from the provided source
func (domain *Domain_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(domain)
}

// ConvertStatusTo populates the provided destination from our Domain_STATUS
func (domain *Domain_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(domain)
}

// Storage version of v1api20200601.InboundIpRule
type InboundIpRule struct {
	Action      *string                `json:"action,omitempty"`
	IpMask      *string                `json:"ipMask,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.InboundIpRule_STATUS
type InboundIpRule_STATUS struct {
	Action      *string                `json:"action,omitempty"`
	IpMask      *string                `json:"ipMask,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.InputSchemaMapping
type InputSchemaMapping struct {
	Json        *JsonInputSchemaMapping `json:"json,omitempty"`
	PropertyBag genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.InputSchemaMapping_STATUS
type InputSchemaMapping_STATUS struct {
	Json        *JsonInputSchemaMapping_STATUS `json:"json,omitempty"`
	PropertyBag genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.PrivateEndpointConnection_STATUS_Domain_SubResourceEmbedded
type PrivateEndpointConnection_STATUS_Domain_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.SystemData_STATUS
// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.JsonInputSchemaMapping
type JsonInputSchemaMapping struct {
	DataVersion            *JsonFieldWithDefault  `json:"dataVersion,omitempty"`
	EventTime              *JsonField             `json:"eventTime,omitempty"`
	EventType              *JsonFieldWithDefault  `json:"eventType,omitempty"`
	Id                     *JsonField             `json:"id,omitempty"`
	InputSchemaMappingType *string                `json:"inputSchemaMappingType,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subject                *JsonFieldWithDefault  `json:"subject,omitempty"`
	Topic                  *JsonField             `json:"topic,omitempty"`
}

// Storage version of v1api20200601.JsonInputSchemaMapping_STATUS
type JsonInputSchemaMapping_STATUS struct {
	DataVersion            *JsonFieldWithDefault_STATUS `json:"dataVersion,omitempty"`
	EventTime              *JsonField_STATUS            `json:"eventTime,omitempty"`
	EventType              *JsonFieldWithDefault_STATUS `json:"eventType,omitempty"`
	Id                     *JsonField_STATUS            `json:"id,omitempty"`
	InputSchemaMappingType *string                      `json:"inputSchemaMappingType,omitempty"`
	PropertyBag            genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Subject                *JsonFieldWithDefault_STATUS `json:"subject,omitempty"`
	Topic                  *JsonField_STATUS            `json:"topic,omitempty"`
}

// Storage version of v1api20200601.JsonField
// This is used to express the source of an input schema mapping for a single target field in the Event Grid Event schema.
// This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a field in the
// input event schema.
type JsonField struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceField *string                `json:"sourceField,omitempty"`
}

// Storage version of v1api20200601.JsonField_STATUS
// This is used to express the source of an input schema mapping for a single target field in the Event Grid Event schema.
// This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a field in the
// input event schema.
type JsonField_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceField *string                `json:"sourceField,omitempty"`
}

// Storage version of v1api20200601.JsonFieldWithDefault
// This is used to express the source of an input schema mapping for a single target field
// in the Event Grid Event schema.
// This is currently used in the mappings for the 'subject',
// 'eventtype' and 'dataversion' properties. This represents a
// field in the input event schema
// along with a default value to be used, and at least one of these two properties should
// be provided.
type JsonFieldWithDefault struct {
	DefaultValue *string                `json:"defaultValue,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceField  *string                `json:"sourceField,omitempty"`
}

// Storage version of v1api20200601.JsonFieldWithDefault_STATUS
// This is used to express the source of an input schema mapping for a single target field
// in the Event Grid Event schema.
// This is currently used in the mappings for the 'subject',
// 'eventtype' and 'dataversion' properties. This represents a
// field in the input event schema
// along with a default value to be used, and at least one of these two properties should
// be provided.
type JsonFieldWithDefault_STATUS struct {
	DefaultValue *string                `json:"defaultValue,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceField  *string                `json:"sourceField,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
