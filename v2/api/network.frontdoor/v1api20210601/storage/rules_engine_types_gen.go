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

// +kubebuilder:rbac:groups=network.frontdoor.azure.com,resources=rulesengines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.frontdoor.azure.com,resources={rulesengines/status,rulesengines/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210601.RulesEngine
// Generator information:
// - Generated from: /frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/frontdoor.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}
type RulesEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontDoors_RulesEngine_Spec   `json:"spec,omitempty"`
	Status            FrontDoors_RulesEngine_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RulesEngine{}

// GetConditions returns the conditions of the resource
func (engine *RulesEngine) GetConditions() conditions.Conditions {
	return engine.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (engine *RulesEngine) SetConditions(conditions conditions.Conditions) {
	engine.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RulesEngine{}

// AzureName returns the Azure name of the resource
func (engine *RulesEngine) AzureName() string {
	return engine.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (engine RulesEngine) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (engine *RulesEngine) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (engine *RulesEngine) GetSpec() genruntime.ConvertibleSpec {
	return &engine.Spec
}

// GetStatus returns the status of this resource
func (engine *RulesEngine) GetStatus() genruntime.ConvertibleStatus {
	return &engine.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (engine *RulesEngine) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/frontDoors/rulesEngines"
func (engine *RulesEngine) GetType() string {
	return "Microsoft.Network/frontDoors/rulesEngines"
}

// NewEmptyStatus returns a new empty (blank) status
func (engine *RulesEngine) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FrontDoors_RulesEngine_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (engine *RulesEngine) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(engine.Spec)
	return engine.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (engine *RulesEngine) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FrontDoors_RulesEngine_STATUS); ok {
		engine.Status = *st
		return nil
	}

	// Convert status to required version
	var st FrontDoors_RulesEngine_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	engine.Status = st
	return nil
}

// Hub marks that this RulesEngine is the hub type for conversion
func (engine *RulesEngine) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (engine *RulesEngine) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: engine.Spec.OriginalVersion,
		Kind:    "RulesEngine",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210601.RulesEngine
// Generator information:
// - Generated from: /frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/frontdoor.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}
type RulesEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RulesEngine `json:"items"`
}

// Storage version of v1api20210601.FrontDoors_RulesEngine_Spec
type FrontDoors_RulesEngine_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.frontdoor.azure.com/FrontDoor resource
	Owner       *genruntime.KnownResourceReference `group:"network.frontdoor.azure.com" json:"owner,omitempty" kind:"FrontDoor"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Rules       []RulesEngineRule                  `json:"rules,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FrontDoors_RulesEngine_Spec{}

// ConvertSpecFrom populates our FrontDoors_RulesEngine_Spec from the provided source
func (engine *FrontDoors_RulesEngine_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == engine {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(engine)
}

// ConvertSpecTo populates the provided destination from our FrontDoors_RulesEngine_Spec
func (engine *FrontDoors_RulesEngine_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == engine {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(engine)
}

// Storage version of v1api20210601.FrontDoors_RulesEngine_STATUS
type FrontDoors_RulesEngine_STATUS struct {
	Conditions    []conditions.Condition   `json:"conditions,omitempty"`
	Id            *string                  `json:"id,omitempty"`
	Name          *string                  `json:"name,omitempty"`
	PropertyBag   genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	ResourceState *string                  `json:"resourceState,omitempty"`
	Rules         []RulesEngineRule_STATUS `json:"rules,omitempty"`
	Type          *string                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FrontDoors_RulesEngine_STATUS{}

// ConvertStatusFrom populates our FrontDoors_RulesEngine_STATUS from the provided source
func (engine *FrontDoors_RulesEngine_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == engine {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(engine)
}

// ConvertStatusTo populates the provided destination from our FrontDoors_RulesEngine_STATUS
func (engine *FrontDoors_RulesEngine_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == engine {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(engine)
}

// Storage version of v1api20210601.RulesEngineRule
// Contains a list of match conditions, and an action on how to modify the request/response. If multiple rules match, the
// actions from one rule that conflict with a previous rule overwrite for a singular action, or append in the case of
// headers manipulation.
type RulesEngineRule struct {
	Action                  *RulesEngineAction          `json:"action,omitempty"`
	MatchConditions         []RulesEngineMatchCondition `json:"matchConditions,omitempty"`
	MatchProcessingBehavior *string                     `json:"matchProcessingBehavior,omitempty"`
	Name                    *string                     `json:"name,omitempty"`
	Priority                *int                        `json:"priority,omitempty"`
	PropertyBag             genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.RulesEngineRule_STATUS
// Contains a list of match conditions, and an action on how to modify the request/response. If multiple rules match, the
// actions from one rule that conflict with a previous rule overwrite for a singular action, or append in the case of
// headers manipulation.
type RulesEngineRule_STATUS struct {
	Action                  *RulesEngineAction_STATUS          `json:"action,omitempty"`
	MatchConditions         []RulesEngineMatchCondition_STATUS `json:"matchConditions,omitempty"`
	MatchProcessingBehavior *string                            `json:"matchProcessingBehavior,omitempty"`
	Name                    *string                            `json:"name,omitempty"`
	Priority                *int                               `json:"priority,omitempty"`
	PropertyBag             genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.RulesEngineAction
// One or more actions that will execute, modifying the request and/or response.
type RulesEngineAction struct {
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestHeaderActions       []HeaderAction         `json:"requestHeaderActions,omitempty"`
	ResponseHeaderActions      []HeaderAction         `json:"responseHeaderActions,omitempty"`
	RouteConfigurationOverride *RouteConfiguration    `json:"routeConfigurationOverride,omitempty"`
}

// Storage version of v1api20210601.RulesEngineAction_STATUS
// One or more actions that will execute, modifying the request and/or response.
type RulesEngineAction_STATUS struct {
	PropertyBag                genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	RequestHeaderActions       []HeaderAction_STATUS      `json:"requestHeaderActions,omitempty"`
	ResponseHeaderActions      []HeaderAction_STATUS      `json:"responseHeaderActions,omitempty"`
	RouteConfigurationOverride *RouteConfiguration_STATUS `json:"routeConfigurationOverride,omitempty"`
}

// Storage version of v1api20210601.RulesEngineMatchCondition
// Define a match condition
type RulesEngineMatchCondition struct {
	NegateCondition          *bool                  `json:"negateCondition,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RulesEngineMatchValue    []string               `json:"rulesEngineMatchValue,omitempty"`
	RulesEngineMatchVariable *string                `json:"rulesEngineMatchVariable,omitempty"`
	RulesEngineOperator      *string                `json:"rulesEngineOperator,omitempty"`
	Selector                 *string                `json:"selector,omitempty"`
	Transforms               []string               `json:"transforms,omitempty"`
}

// Storage version of v1api20210601.RulesEngineMatchCondition_STATUS
// Define a match condition
type RulesEngineMatchCondition_STATUS struct {
	NegateCondition          *bool                  `json:"negateCondition,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RulesEngineMatchValue    []string               `json:"rulesEngineMatchValue,omitempty"`
	RulesEngineMatchVariable *string                `json:"rulesEngineMatchVariable,omitempty"`
	RulesEngineOperator      *string                `json:"rulesEngineOperator,omitempty"`
	Selector                 *string                `json:"selector,omitempty"`
	Transforms               []string               `json:"transforms,omitempty"`
}

// Storage version of v1api20210601.HeaderAction
// An action that can manipulate an http header.
type HeaderAction struct {
	HeaderActionType *string                `json:"headerActionType,omitempty"`
	HeaderName       *string                `json:"headerName,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value            *string                `json:"value,omitempty"`
}

// Storage version of v1api20210601.HeaderAction_STATUS
// An action that can manipulate an http header.
type HeaderAction_STATUS struct {
	HeaderActionType *string                `json:"headerActionType,omitempty"`
	HeaderName       *string                `json:"headerName,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value            *string                `json:"value,omitempty"`
}

// Storage version of v1api20210601.RouteConfiguration_STATUS
type RouteConfiguration_STATUS struct {
	MicrosoftAzureFrontDoorModelsFrontdoorForwarding *ForwardingConfiguration_STATUS `json:"microsoftAzureFrontDoorModelsFrontdoorForwardingConfiguration,omitempty"`
	MicrosoftAzureFrontDoorModelsFrontdoorRedirect   *RedirectConfiguration_STATUS   `json:"microsoftAzureFrontDoorModelsFrontdoorRedirectConfiguration,omitempty"`
	PropertyBag                                      genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.ForwardingConfiguration_STATUS
type ForwardingConfiguration_STATUS struct {
	BackendPool          *SubResource_STATUS        `json:"backendPool,omitempty"`
	CacheConfiguration   *CacheConfiguration_STATUS `json:"cacheConfiguration,omitempty"`
	CustomForwardingPath *string                    `json:"customForwardingPath,omitempty"`
	ForwardingProtocol   *string                    `json:"forwardingProtocol,omitempty"`
	OdataType            *string                    `json:"@odata.type,omitempty"`
	PropertyBag          genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.RedirectConfiguration_STATUS
type RedirectConfiguration_STATUS struct {
	CustomFragment    *string                `json:"customFragment,omitempty"`
	CustomHost        *string                `json:"customHost,omitempty"`
	CustomPath        *string                `json:"customPath,omitempty"`
	CustomQueryString *string                `json:"customQueryString,omitempty"`
	OdataType         *string                `json:"@odata.type,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RedirectProtocol  *string                `json:"redirectProtocol,omitempty"`
	RedirectType      *string                `json:"redirectType,omitempty"`
}

// Storage version of v1api20210601.CacheConfiguration_STATUS
// Caching settings for a caching-type route. To disable caching, do not provide a cacheConfiguration object.
type CacheConfiguration_STATUS struct {
	CacheDuration                *string                `json:"cacheDuration,omitempty"`
	DynamicCompression           *string                `json:"dynamicCompression,omitempty"`
	PropertyBag                  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QueryParameterStripDirective *string                `json:"queryParameterStripDirective,omitempty"`
	QueryParameters              *string                `json:"queryParameters,omitempty"`
}

// Storage version of v1api20210601.SubResource_STATUS
// Reference to another subresource.
type SubResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RulesEngine{}, &RulesEngineList{})
}
