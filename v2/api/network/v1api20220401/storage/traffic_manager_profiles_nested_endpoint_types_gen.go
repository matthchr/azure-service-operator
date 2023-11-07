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

// +kubebuilder:rbac:groups=network.azure.com,resources=trafficmanagerprofilesnestedendpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={trafficmanagerprofilesnestedendpoints/status,trafficmanagerprofilesnestedendpoints/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220401.TrafficManagerProfilesNestedEndpoint
// Generator information:
// - Generated from: /trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/trafficmanager.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/NestedEndpoints/{endpointName}
type TrafficManagerProfilesNestedEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Trafficmanagerprofiles_NestedEndpoint_Spec   `json:"spec,omitempty"`
	Status            Trafficmanagerprofiles_NestedEndpoint_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &TrafficManagerProfilesNestedEndpoint{}

// GetConditions returns the conditions of the resource
func (endpoint *TrafficManagerProfilesNestedEndpoint) GetConditions() conditions.Conditions {
	return endpoint.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (endpoint *TrafficManagerProfilesNestedEndpoint) SetConditions(conditions conditions.Conditions) {
	endpoint.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &TrafficManagerProfilesNestedEndpoint{}

// AzureName returns the Azure name of the resource
func (endpoint *TrafficManagerProfilesNestedEndpoint) AzureName() string {
	return endpoint.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-04-01"
func (endpoint TrafficManagerProfilesNestedEndpoint) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (endpoint *TrafficManagerProfilesNestedEndpoint) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (endpoint *TrafficManagerProfilesNestedEndpoint) GetSpec() genruntime.ConvertibleSpec {
	return &endpoint.Spec
}

// GetStatus returns the status of this resource
func (endpoint *TrafficManagerProfilesNestedEndpoint) GetStatus() genruntime.ConvertibleStatus {
	return &endpoint.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (endpoint *TrafficManagerProfilesNestedEndpoint) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/trafficmanagerprofiles/NestedEndpoints"
func (endpoint *TrafficManagerProfilesNestedEndpoint) GetType() string {
	return "Microsoft.Network/trafficmanagerprofiles/NestedEndpoints"
}

// NewEmptyStatus returns a new empty (blank) status
func (endpoint *TrafficManagerProfilesNestedEndpoint) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Trafficmanagerprofiles_NestedEndpoint_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (endpoint *TrafficManagerProfilesNestedEndpoint) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(endpoint.Spec)
	return endpoint.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (endpoint *TrafficManagerProfilesNestedEndpoint) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Trafficmanagerprofiles_NestedEndpoint_STATUS); ok {
		endpoint.Status = *st
		return nil
	}

	// Convert status to required version
	var st Trafficmanagerprofiles_NestedEndpoint_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	endpoint.Status = st
	return nil
}

// Hub marks that this TrafficManagerProfilesNestedEndpoint is the hub type for conversion
func (endpoint *TrafficManagerProfilesNestedEndpoint) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (endpoint *TrafficManagerProfilesNestedEndpoint) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: endpoint.Spec.OriginalVersion,
		Kind:    "TrafficManagerProfilesNestedEndpoint",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220401.TrafficManagerProfilesNestedEndpoint
// Generator information:
// - Generated from: /trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/trafficmanager.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/NestedEndpoints/{endpointName}
type TrafficManagerProfilesNestedEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficManagerProfilesNestedEndpoint `json:"items"`
}

// Storage version of v1api20220401.Trafficmanagerprofiles_NestedEndpoint_Spec
type Trafficmanagerprofiles_NestedEndpoint_Spec struct {
	AlwaysServe *string `json:"alwaysServe,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName             string                             `json:"azureName,omitempty"`
	CustomHeaders         []EndpointProperties_CustomHeaders `json:"customHeaders,omitempty"`
	EndpointLocation      *string                            `json:"endpointLocation,omitempty"`
	EndpointMonitorStatus *string                            `json:"endpointMonitorStatus,omitempty"`
	EndpointStatus        *string                            `json:"endpointStatus,omitempty"`
	GeoMapping            []string                           `json:"geoMapping,omitempty"`
	MinChildEndpoints     *int                               `json:"minChildEndpoints,omitempty"`
	MinChildEndpointsIPv4 *int                               `json:"minChildEndpointsIPv4,omitempty"`
	MinChildEndpointsIPv6 *int                               `json:"minChildEndpointsIPv6,omitempty"`
	OriginalVersion       string                             `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/TrafficManagerProfile resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"TrafficManagerProfile"`
	Priority    *int                               `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Subnets     []EndpointProperties_Subnets       `json:"subnets,omitempty"`
	Target      *string                            `json:"target,omitempty"`

	// TargetResourceReference: The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type
	// 'ExternalEndpoints'.
	TargetResourceReference *genruntime.ResourceReference `armReference:"TargetResourceId" json:"targetResourceReference,omitempty"`
	Type                    *string                       `json:"type,omitempty"`
	Weight                  *int                          `json:"weight,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Trafficmanagerprofiles_NestedEndpoint_Spec{}

// ConvertSpecFrom populates our Trafficmanagerprofiles_NestedEndpoint_Spec from the provided source
func (endpoint *Trafficmanagerprofiles_NestedEndpoint_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(endpoint)
}

// ConvertSpecTo populates the provided destination from our Trafficmanagerprofiles_NestedEndpoint_Spec
func (endpoint *Trafficmanagerprofiles_NestedEndpoint_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(endpoint)
}

// Storage version of v1api20220401.Trafficmanagerprofiles_NestedEndpoint_STATUS
type Trafficmanagerprofiles_NestedEndpoint_STATUS struct {
	AlwaysServe           *string                                   `json:"alwaysServe,omitempty"`
	Conditions            []conditions.Condition                    `json:"conditions,omitempty"`
	CustomHeaders         []EndpointProperties_CustomHeaders_STATUS `json:"customHeaders,omitempty"`
	EndpointLocation      *string                                   `json:"endpointLocation,omitempty"`
	EndpointMonitorStatus *string                                   `json:"endpointMonitorStatus,omitempty"`
	EndpointStatus        *string                                   `json:"endpointStatus,omitempty"`
	GeoMapping            []string                                  `json:"geoMapping,omitempty"`
	Id                    *string                                   `json:"id,omitempty"`
	MinChildEndpoints     *int                                      `json:"minChildEndpoints,omitempty"`
	MinChildEndpointsIPv4 *int                                      `json:"minChildEndpointsIPv4,omitempty"`
	MinChildEndpointsIPv6 *int                                      `json:"minChildEndpointsIPv6,omitempty"`
	Name                  *string                                   `json:"name,omitempty"`
	Priority              *int                                      `json:"priority,omitempty"`
	PropertyBag           genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	Subnets               []EndpointProperties_Subnets_STATUS       `json:"subnets,omitempty"`
	Target                *string                                   `json:"target,omitempty"`
	TargetResourceId      *string                                   `json:"targetResourceId,omitempty"`
	Type                  *string                                   `json:"type,omitempty"`
	Weight                *int                                      `json:"weight,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Trafficmanagerprofiles_NestedEndpoint_STATUS{}

// ConvertStatusFrom populates our Trafficmanagerprofiles_NestedEndpoint_STATUS from the provided source
func (endpoint *Trafficmanagerprofiles_NestedEndpoint_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(endpoint)
}

// ConvertStatusTo populates the provided destination from our Trafficmanagerprofiles_NestedEndpoint_STATUS
func (endpoint *Trafficmanagerprofiles_NestedEndpoint_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(endpoint)
}

func init() {
	SchemeBuilder.Register(&TrafficManagerProfilesNestedEndpoint{}, &TrafficManagerProfilesNestedEndpointList{})
}
