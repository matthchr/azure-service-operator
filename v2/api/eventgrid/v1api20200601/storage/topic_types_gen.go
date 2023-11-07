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

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=topics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={topics/status,topics/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20200601.Topic
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Topic_Spec   `json:"spec,omitempty"`
	Status            Topic_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Topic{}

// GetConditions returns the conditions of the resource
func (topic *Topic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *Topic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Topic{}

// AzureName returns the Azure name of the resource
func (topic *Topic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic Topic) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (topic *Topic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *Topic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *Topic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (topic *Topic) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/topics"
func (topic *Topic) GetType() string {
	return "Microsoft.EventGrid/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *Topic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Topic_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (topic *Topic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return topic.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (topic *Topic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Topic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st Topic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// Hub marks that this Topic is the hub type for conversion
func (topic *Topic) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *Topic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "Topic",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20200601.Topic
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Storage version of v1api20200601.Topic_Spec
type Topic_Spec struct {
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

var _ genruntime.ConvertibleSpec = &Topic_Spec{}

// ConvertSpecFrom populates our Topic_Spec from the provided source
func (topic *Topic_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(topic)
}

// ConvertSpecTo populates the provided destination from our Topic_Spec
func (topic *Topic_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(topic)
}

// Storage version of v1api20200601.Topic_STATUS
// EventGrid Topic
type Topic_STATUS struct {
	Conditions                 []conditions.Condition                                       `json:"conditions,omitempty"`
	Endpoint                   *string                                                      `json:"endpoint,omitempty"`
	Id                         *string                                                      `json:"id,omitempty"`
	InboundIpRules             []InboundIpRule_STATUS                                       `json:"inboundIpRules,omitempty"`
	InputSchema                *string                                                      `json:"inputSchema,omitempty"`
	InputSchemaMapping         *InputSchemaMapping_STATUS                                   `json:"inputSchemaMapping,omitempty"`
	Location                   *string                                                      `json:"location,omitempty"`
	MetricResourceId           *string                                                      `json:"metricResourceId,omitempty"`
	Name                       *string                                                      `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                       `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                      `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                      `json:"publicNetworkAccess,omitempty"`
	SystemData                 *SystemData_STATUS                                           `json:"systemData,omitempty"`
	Tags                       map[string]string                                            `json:"tags,omitempty"`
	Type                       *string                                                      `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Topic_STATUS{}

// ConvertStatusFrom populates our Topic_STATUS from the provided source
func (topic *Topic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(topic)
}

// ConvertStatusTo populates the provided destination from our Topic_STATUS
func (topic *Topic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(topic)
}

// Storage version of v1api20200601.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
type PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
