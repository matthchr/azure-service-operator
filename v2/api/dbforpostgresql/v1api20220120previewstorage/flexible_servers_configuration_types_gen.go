// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220120previewstorage

import (
	"fmt"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220120preview.FlexibleServersConfiguration
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/configurations/{configurationName}
type FlexibleServersConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServers_Configuration_Spec   `json:"spec,omitempty"`
	Status            FlexibleServers_Configuration_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersConfiguration{}

// GetConditions returns the conditions of the resource
func (configuration *FlexibleServersConfiguration) GetConditions() conditions.Conditions {
	return configuration.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (configuration *FlexibleServersConfiguration) SetConditions(conditions conditions.Conditions) {
	configuration.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersConfiguration{}

// ConvertFrom populates our FlexibleServersConfiguration from the provided hub FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210601s.FlexibleServersConfiguration)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1api20210601storage/FlexibleServersConfiguration but received %T instead", hub)
	}

	return configuration.AssignProperties_From_FlexibleServersConfiguration(source)
}

// ConvertTo populates the provided hub FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210601s.FlexibleServersConfiguration)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1api20210601storage/FlexibleServersConfiguration but received %T instead", hub)
	}

	return configuration.AssignProperties_To_FlexibleServersConfiguration(destination)
}

var _ genruntime.KubernetesResource = &FlexibleServersConfiguration{}

// AzureName returns the Azure name of the resource
func (configuration *FlexibleServersConfiguration) AzureName() string {
	return configuration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-20-preview"
func (configuration FlexibleServersConfiguration) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (configuration *FlexibleServersConfiguration) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (configuration *FlexibleServersConfiguration) GetSpec() genruntime.ConvertibleSpec {
	return &configuration.Spec
}

// GetStatus returns the status of this resource
func (configuration *FlexibleServersConfiguration) GetStatus() genruntime.ConvertibleStatus {
	return &configuration.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (configuration *FlexibleServersConfiguration) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// NewEmptyStatus returns a new empty (blank) status
func (configuration *FlexibleServersConfiguration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServers_Configuration_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (configuration *FlexibleServersConfiguration) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(configuration.Spec)
	return configuration.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (configuration *FlexibleServersConfiguration) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServers_Configuration_STATUS); ok {
		configuration.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServers_Configuration_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	configuration.Status = st
	return nil
}

// AssignProperties_From_FlexibleServersConfiguration populates our FlexibleServersConfiguration from the provided source FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignProperties_From_FlexibleServersConfiguration(source *v20210601s.FlexibleServersConfiguration) error {

	// ObjectMeta
	configuration.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServers_Configuration_Spec
	err := spec.AssignProperties_From_FlexibleServers_Configuration_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServers_Configuration_Spec() to populate field Spec")
	}
	configuration.Spec = spec

	// Status
	var status FlexibleServers_Configuration_STATUS
	err = status.AssignProperties_From_FlexibleServers_Configuration_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServers_Configuration_STATUS() to populate field Status")
	}
	configuration.Status = status

	// Invoke the augmentConversionForFlexibleServersConfiguration interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServersConfiguration); ok {
		err := augmentedConfiguration.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersConfiguration populates the provided destination FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignProperties_To_FlexibleServersConfiguration(destination *v20210601s.FlexibleServersConfiguration) error {

	// ObjectMeta
	destination.ObjectMeta = *configuration.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210601s.FlexibleServers_Configuration_Spec
	err := configuration.Spec.AssignProperties_To_FlexibleServers_Configuration_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServers_Configuration_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210601s.FlexibleServers_Configuration_STATUS
	err = configuration.Status.AssignProperties_To_FlexibleServers_Configuration_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServers_Configuration_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForFlexibleServersConfiguration interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServersConfiguration); ok {
		err := augmentedConfiguration.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (configuration *FlexibleServersConfiguration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: configuration.Spec.OriginalVersion,
		Kind:    "FlexibleServersConfiguration",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220120preview.FlexibleServersConfiguration
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/configurations/{configurationName}
type FlexibleServersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersConfiguration `json:"items"`
}

type augmentConversionForFlexibleServersConfiguration interface {
	AssignPropertiesFrom(src *v20210601s.FlexibleServersConfiguration) error
	AssignPropertiesTo(dst *v20210601s.FlexibleServersConfiguration) error
}

// Storage version of v1api20220120preview.FlexibleServers_Configuration_Spec
type FlexibleServers_Configuration_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner       *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Source      *string                            `json:"source,omitempty"`
	Value       *string                            `json:"value,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServers_Configuration_Spec{}

// ConvertSpecFrom populates our FlexibleServers_Configuration_Spec from the provided source
func (configuration *FlexibleServers_Configuration_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210601s.FlexibleServers_Configuration_Spec)
	if ok {
		// Populate our instance from source
		return configuration.AssignProperties_From_FlexibleServers_Configuration_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.FlexibleServers_Configuration_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = configuration.AssignProperties_From_FlexibleServers_Configuration_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServers_Configuration_Spec
func (configuration *FlexibleServers_Configuration_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210601s.FlexibleServers_Configuration_Spec)
	if ok {
		// Populate destination from our instance
		return configuration.AssignProperties_To_FlexibleServers_Configuration_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.FlexibleServers_Configuration_Spec{}
	err := configuration.AssignProperties_To_FlexibleServers_Configuration_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_FlexibleServers_Configuration_Spec populates our FlexibleServers_Configuration_Spec from the provided source FlexibleServers_Configuration_Spec
func (configuration *FlexibleServers_Configuration_Spec) AssignProperties_From_FlexibleServers_Configuration_Spec(source *v20210601s.FlexibleServers_Configuration_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	configuration.AzureName = source.AzureName

	// OriginalVersion
	configuration.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		configuration.Owner = &owner
	} else {
		configuration.Owner = nil
	}

	// Source
	configuration.Source = genruntime.ClonePointerToString(source.Source)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		configuration.PropertyBag = propertyBag
	} else {
		configuration.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServers_Configuration_Spec interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServers_Configuration_Spec); ok {
		err := augmentedConfiguration.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServers_Configuration_Spec populates the provided destination FlexibleServers_Configuration_Spec from our FlexibleServers_Configuration_Spec
func (configuration *FlexibleServers_Configuration_Spec) AssignProperties_To_FlexibleServers_Configuration_Spec(destination *v20210601s.FlexibleServers_Configuration_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(configuration.PropertyBag)

	// AzureName
	destination.AzureName = configuration.AzureName

	// OriginalVersion
	destination.OriginalVersion = configuration.OriginalVersion

	// Owner
	if configuration.Owner != nil {
		owner := configuration.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Source
	destination.Source = genruntime.ClonePointerToString(configuration.Source)

	// Value
	destination.Value = genruntime.ClonePointerToString(configuration.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServers_Configuration_Spec interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServers_Configuration_Spec); ok {
		err := augmentedConfiguration.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220120preview.FlexibleServers_Configuration_STATUS
type FlexibleServers_Configuration_STATUS struct {
	AllowedValues          *string                `json:"allowedValues,omitempty"`
	Conditions             []conditions.Condition `json:"conditions,omitempty"`
	DataType               *string                `json:"dataType,omitempty"`
	DefaultValue           *string                `json:"defaultValue,omitempty"`
	Description            *string                `json:"description,omitempty"`
	DocumentationLink      *string                `json:"documentationLink,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	IsConfigPendingRestart *bool                  `json:"isConfigPendingRestart,omitempty"`
	IsDynamicConfig        *bool                  `json:"isDynamicConfig,omitempty"`
	IsReadOnly             *bool                  `json:"isReadOnly,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Source                 *string                `json:"source,omitempty"`
	SystemData             *SystemData_STATUS     `json:"systemData,omitempty"`
	Type                   *string                `json:"type,omitempty"`
	Unit                   *string                `json:"unit,omitempty"`
	Value                  *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServers_Configuration_STATUS{}

// ConvertStatusFrom populates our FlexibleServers_Configuration_STATUS from the provided source
func (configuration *FlexibleServers_Configuration_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210601s.FlexibleServers_Configuration_STATUS)
	if ok {
		// Populate our instance from source
		return configuration.AssignProperties_From_FlexibleServers_Configuration_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.FlexibleServers_Configuration_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = configuration.AssignProperties_From_FlexibleServers_Configuration_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServers_Configuration_STATUS
func (configuration *FlexibleServers_Configuration_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210601s.FlexibleServers_Configuration_STATUS)
	if ok {
		// Populate destination from our instance
		return configuration.AssignProperties_To_FlexibleServers_Configuration_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.FlexibleServers_Configuration_STATUS{}
	err := configuration.AssignProperties_To_FlexibleServers_Configuration_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_FlexibleServers_Configuration_STATUS populates our FlexibleServers_Configuration_STATUS from the provided source FlexibleServers_Configuration_STATUS
func (configuration *FlexibleServers_Configuration_STATUS) AssignProperties_From_FlexibleServers_Configuration_STATUS(source *v20210601s.FlexibleServers_Configuration_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AllowedValues
	configuration.AllowedValues = genruntime.ClonePointerToString(source.AllowedValues)

	// Conditions
	configuration.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DataType
	configuration.DataType = genruntime.ClonePointerToString(source.DataType)

	// DefaultValue
	configuration.DefaultValue = genruntime.ClonePointerToString(source.DefaultValue)

	// Description
	configuration.Description = genruntime.ClonePointerToString(source.Description)

	// DocumentationLink
	configuration.DocumentationLink = genruntime.ClonePointerToString(source.DocumentationLink)

	// Id
	configuration.Id = genruntime.ClonePointerToString(source.Id)

	// IsConfigPendingRestart
	if source.IsConfigPendingRestart != nil {
		isConfigPendingRestart := *source.IsConfigPendingRestart
		configuration.IsConfigPendingRestart = &isConfigPendingRestart
	} else {
		configuration.IsConfigPendingRestart = nil
	}

	// IsDynamicConfig
	if source.IsDynamicConfig != nil {
		isDynamicConfig := *source.IsDynamicConfig
		configuration.IsDynamicConfig = &isDynamicConfig
	} else {
		configuration.IsDynamicConfig = nil
	}

	// IsReadOnly
	if source.IsReadOnly != nil {
		isReadOnly := *source.IsReadOnly
		configuration.IsReadOnly = &isReadOnly
	} else {
		configuration.IsReadOnly = nil
	}

	// Name
	configuration.Name = genruntime.ClonePointerToString(source.Name)

	// Source
	configuration.Source = genruntime.ClonePointerToString(source.Source)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		configuration.SystemData = &systemDatum
	} else {
		configuration.SystemData = nil
	}

	// Type
	configuration.Type = genruntime.ClonePointerToString(source.Type)

	// Unit
	configuration.Unit = genruntime.ClonePointerToString(source.Unit)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		configuration.PropertyBag = propertyBag
	} else {
		configuration.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServers_Configuration_STATUS interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServers_Configuration_STATUS); ok {
		err := augmentedConfiguration.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServers_Configuration_STATUS populates the provided destination FlexibleServers_Configuration_STATUS from our FlexibleServers_Configuration_STATUS
func (configuration *FlexibleServers_Configuration_STATUS) AssignProperties_To_FlexibleServers_Configuration_STATUS(destination *v20210601s.FlexibleServers_Configuration_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(configuration.PropertyBag)

	// AllowedValues
	destination.AllowedValues = genruntime.ClonePointerToString(configuration.AllowedValues)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(configuration.Conditions)

	// DataType
	destination.DataType = genruntime.ClonePointerToString(configuration.DataType)

	// DefaultValue
	destination.DefaultValue = genruntime.ClonePointerToString(configuration.DefaultValue)

	// Description
	destination.Description = genruntime.ClonePointerToString(configuration.Description)

	// DocumentationLink
	destination.DocumentationLink = genruntime.ClonePointerToString(configuration.DocumentationLink)

	// Id
	destination.Id = genruntime.ClonePointerToString(configuration.Id)

	// IsConfigPendingRestart
	if configuration.IsConfigPendingRestart != nil {
		isConfigPendingRestart := *configuration.IsConfigPendingRestart
		destination.IsConfigPendingRestart = &isConfigPendingRestart
	} else {
		destination.IsConfigPendingRestart = nil
	}

	// IsDynamicConfig
	if configuration.IsDynamicConfig != nil {
		isDynamicConfig := *configuration.IsDynamicConfig
		destination.IsDynamicConfig = &isDynamicConfig
	} else {
		destination.IsDynamicConfig = nil
	}

	// IsReadOnly
	if configuration.IsReadOnly != nil {
		isReadOnly := *configuration.IsReadOnly
		destination.IsReadOnly = &isReadOnly
	} else {
		destination.IsReadOnly = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(configuration.Name)

	// Source
	destination.Source = genruntime.ClonePointerToString(configuration.Source)

	// SystemData
	if configuration.SystemData != nil {
		var systemDatum v20210601s.SystemData_STATUS
		err := configuration.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(configuration.Type)

	// Unit
	destination.Unit = genruntime.ClonePointerToString(configuration.Unit)

	// Value
	destination.Value = genruntime.ClonePointerToString(configuration.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServers_Configuration_STATUS interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServers_Configuration_STATUS); ok {
		err := augmentedConfiguration.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServers_Configuration_Spec interface {
	AssignPropertiesFrom(src *v20210601s.FlexibleServers_Configuration_Spec) error
	AssignPropertiesTo(dst *v20210601s.FlexibleServers_Configuration_Spec) error
}

type augmentConversionForFlexibleServers_Configuration_STATUS interface {
	AssignPropertiesFrom(src *v20210601s.FlexibleServers_Configuration_STATUS) error
	AssignPropertiesTo(dst *v20210601s.FlexibleServers_Configuration_STATUS) error
}

func init() {
	SchemeBuilder.Register(&FlexibleServersConfiguration{}, &FlexibleServersConfigurationList{})
}
