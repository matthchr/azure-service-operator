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

// +kubebuilder:rbac:groups=network.azure.com,resources=dnsresolvers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnsresolvers/status,dnsresolvers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.DnsResolver
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}
type DnsResolver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsResolver_Spec   `json:"spec,omitempty"`
	Status            DnsResolver_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsResolver{}

// GetConditions returns the conditions of the resource
func (resolver *DnsResolver) GetConditions() conditions.Conditions {
	return resolver.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (resolver *DnsResolver) SetConditions(conditions conditions.Conditions) {
	resolver.Status.Conditions = conditions
}

var _ configmaps.Exporter = &DnsResolver{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (resolver *DnsResolver) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if resolver.Spec.OperatorSpec == nil {
		return nil
	}
	return resolver.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DnsResolver{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (resolver *DnsResolver) SecretDestinationExpressions() []*core.DestinationExpression {
	if resolver.Spec.OperatorSpec == nil {
		return nil
	}
	return resolver.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &DnsResolver{}

// AzureName returns the Azure name of the resource
func (resolver *DnsResolver) AzureName() string {
	return resolver.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (resolver DnsResolver) GetAPIVersion() string {
	return "2022-07-01"
}

// GetResourceScope returns the scope of the resource
func (resolver *DnsResolver) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (resolver *DnsResolver) GetSpec() genruntime.ConvertibleSpec {
	return &resolver.Spec
}

// GetStatus returns the status of this resource
func (resolver *DnsResolver) GetStatus() genruntime.ConvertibleStatus {
	return &resolver.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (resolver *DnsResolver) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsResolvers"
func (resolver *DnsResolver) GetType() string {
	return "Microsoft.Network/dnsResolvers"
}

// NewEmptyStatus returns a new empty (blank) status
func (resolver *DnsResolver) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsResolver_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (resolver *DnsResolver) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(resolver.Spec)
	return resolver.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (resolver *DnsResolver) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsResolver_STATUS); ok {
		resolver.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsResolver_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	resolver.Status = st
	return nil
}

// Hub marks that this DnsResolver is the hub type for conversion
func (resolver *DnsResolver) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (resolver *DnsResolver) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: resolver.Spec.OriginalVersion,
		Kind:    "DnsResolver",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.DnsResolver
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}
type DnsResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsResolver `json:"items"`
}

// Storage version of v1api20220701.DnsResolver_Spec
type DnsResolver_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                   `json:"azureName,omitempty"`
	Location        *string                  `json:"location,omitempty"`
	OperatorSpec    *DnsResolverOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner          *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags           map[string]string                  `json:"tags,omitempty"`
	VirtualNetwork *DnsresolverSubResource            `json:"virtualNetwork,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DnsResolver_Spec{}

// ConvertSpecFrom populates our DnsResolver_Spec from the provided source
func (resolver *DnsResolver_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == resolver {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(resolver)
}

// ConvertSpecTo populates the provided destination from our DnsResolver_Spec
func (resolver *DnsResolver_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == resolver {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(resolver)
}

// Storage version of v1api20220701.DnsResolver_STATUS
// Describes a DNS resolver.
type DnsResolver_STATUS struct {
	Conditions        []conditions.Condition         `json:"conditions,omitempty"`
	DnsResolverState  *string                        `json:"dnsResolverState,omitempty"`
	Etag              *string                        `json:"etag,omitempty"`
	Id                *string                        `json:"id,omitempty"`
	Location          *string                        `json:"location,omitempty"`
	Name              *string                        `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	ProvisioningState *string                        `json:"provisioningState,omitempty"`
	ResourceGuid      *string                        `json:"resourceGuid,omitempty"`
	SystemData        *SystemData_STATUS             `json:"systemData,omitempty"`
	Tags              map[string]string              `json:"tags,omitempty"`
	Type              *string                        `json:"type,omitempty"`
	VirtualNetwork    *DnsresolverSubResource_STATUS `json:"virtualNetwork,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsResolver_STATUS{}

// ConvertStatusFrom populates our DnsResolver_STATUS from the provided source
func (resolver *DnsResolver_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == resolver {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(resolver)
}

// ConvertStatusTo populates the provided destination from our DnsResolver_STATUS
func (resolver *DnsResolver_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == resolver {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(resolver)
}

// Storage version of v1api20220701.DnsResolverOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DnsResolverOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DnsResolver{}, &DnsResolverList{})
}
