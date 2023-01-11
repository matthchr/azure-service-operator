// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerConnectionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/default
type ServersConnectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_ConnectionPolicy_Spec   `json:"spec,omitempty"`
	Status            Servers_ConnectionPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersConnectionPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ServersConnectionPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ServersConnectionPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersConnectionPolicy{}

// ConvertFrom populates our ServersConnectionPolicy from the provided hub ServersConnectionPolicy
func (policy *ServersConnectionPolicy) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.ServersConnectionPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/ServersConnectionPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_From_ServersConnectionPolicy(source)
}

// ConvertTo populates the provided hub ServersConnectionPolicy from our ServersConnectionPolicy
func (policy *ServersConnectionPolicy) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.ServersConnectionPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/ServersConnectionPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_To_ServersConnectionPolicy(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1beta20211101-serversconnectionpolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversconnectionpolicies,verbs=create;update,versions=v1beta20211101,name=default.v1beta20211101.serversconnectionpolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersConnectionPolicy{}

// Default applies defaults to the ServersConnectionPolicy resource
func (policy *ServersConnectionPolicy) Default() {
	policy.defaultImpl()
	var temp interface{} = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersConnectionPolicy resource
func (policy *ServersConnectionPolicy) defaultImpl() {}

var _ genruntime.KubernetesResource = &ServersConnectionPolicy{}

// AzureName returns the Azure name of the resource (always "default")
func (policy *ServersConnectionPolicy) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersConnectionPolicy) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (policy *ServersConnectionPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ServersConnectionPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ServersConnectionPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/connectionPolicies"
func (policy *ServersConnectionPolicy) GetType() string {
	return "Microsoft.Sql/servers/connectionPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ServersConnectionPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_ConnectionPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (policy *ServersConnectionPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  policy.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (policy *ServersConnectionPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_ConnectionPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_ConnectionPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1beta20211101-serversconnectionpolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversconnectionpolicies,verbs=create;update,versions=v1beta20211101,name=validate.v1beta20211101.serversconnectionpolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersConnectionPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *ServersConnectionPolicy) ValidateCreate() error {
	validations := policy.createValidations()
	var temp interface{} = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (policy *ServersConnectionPolicy) ValidateDelete() error {
	validations := policy.deleteValidations()
	var temp interface{} = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (policy *ServersConnectionPolicy) ValidateUpdate(old runtime.Object) error {
	validations := policy.updateValidations()
	var temp interface{} = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (policy *ServersConnectionPolicy) createValidations() []func() error {
	return []func() error{policy.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (policy *ServersConnectionPolicy) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (policy *ServersConnectionPolicy) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return policy.validateResourceReferences()
		},
		policy.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (policy *ServersConnectionPolicy) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&policy.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *ServersConnectionPolicy) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*ServersConnectionPolicy)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, policy)
}

// AssignProperties_From_ServersConnectionPolicy populates our ServersConnectionPolicy from the provided source ServersConnectionPolicy
func (policy *ServersConnectionPolicy) AssignProperties_From_ServersConnectionPolicy(source *v20211101s.ServersConnectionPolicy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_ConnectionPolicy_Spec
	err := spec.AssignProperties_From_Servers_ConnectionPolicy_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_ConnectionPolicy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Servers_ConnectionPolicy_STATUS
	err = status.AssignProperties_From_Servers_ConnectionPolicy_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_ConnectionPolicy_STATUS() to populate field Status")
	}
	policy.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersConnectionPolicy populates the provided destination ServersConnectionPolicy from our ServersConnectionPolicy
func (policy *ServersConnectionPolicy) AssignProperties_To_ServersConnectionPolicy(destination *v20211101s.ServersConnectionPolicy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_ConnectionPolicy_Spec
	err := policy.Spec.AssignProperties_To_Servers_ConnectionPolicy_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_ConnectionPolicy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_ConnectionPolicy_STATUS
	err = policy.Status.AssignProperties_To_Servers_ConnectionPolicy_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_ConnectionPolicy_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ServersConnectionPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion(),
		Kind:    "ServersConnectionPolicy",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerConnectionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/default
type ServersConnectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersConnectionPolicy `json:"items"`
}

type Servers_ConnectionPolicy_Spec struct {
	// +kubebuilder:validation:Required
	// ConnectionType: The server connection type.
	ConnectionType *ServerConnectionPolicyProperties_ConnectionType `json:"connectionType,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`
}

var _ genruntime.ARMTransformer = &Servers_ConnectionPolicy_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (policy *Servers_ConnectionPolicy_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if policy == nil {
		return nil, nil
	}
	result := &Servers_ConnectionPolicy_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if policy.ConnectionType != nil {
		result.Properties = &ServerConnectionPolicyProperties_ARM{}
	}
	if policy.ConnectionType != nil {
		connectionType := *policy.ConnectionType
		result.Properties.ConnectionType = &connectionType
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_ConnectionPolicy_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_ConnectionPolicy_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_ConnectionPolicy_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_ConnectionPolicy_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_ConnectionPolicy_Spec_ARM, got %T", armInput)
	}

	// Set property ‘ConnectionType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConnectionType != nil {
			connectionType := *typedInput.Properties.ConnectionType
			policy.ConnectionType = &connectionType
		}
	}

	// Set property ‘Owner’:
	policy.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_ConnectionPolicy_Spec{}

// ConvertSpecFrom populates our Servers_ConnectionPolicy_Spec from the provided source
func (policy *Servers_ConnectionPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_ConnectionPolicy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_ConnectionPolicy_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_ConnectionPolicy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_ConnectionPolicy_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_ConnectionPolicy_Spec
func (policy *Servers_ConnectionPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_ConnectionPolicy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_ConnectionPolicy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_ConnectionPolicy_Spec{}
	err := policy.AssignProperties_To_Servers_ConnectionPolicy_Spec(dst)
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

// AssignProperties_From_Servers_ConnectionPolicy_Spec populates our Servers_ConnectionPolicy_Spec from the provided source Servers_ConnectionPolicy_Spec
func (policy *Servers_ConnectionPolicy_Spec) AssignProperties_From_Servers_ConnectionPolicy_Spec(source *v20211101s.Servers_ConnectionPolicy_Spec) error {

	// ConnectionType
	if source.ConnectionType != nil {
		connectionType := ServerConnectionPolicyProperties_ConnectionType(*source.ConnectionType)
		policy.ConnectionType = &connectionType
	} else {
		policy.ConnectionType = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Servers_ConnectionPolicy_Spec populates the provided destination Servers_ConnectionPolicy_Spec from our Servers_ConnectionPolicy_Spec
func (policy *Servers_ConnectionPolicy_Spec) AssignProperties_To_Servers_ConnectionPolicy_Spec(destination *v20211101s.Servers_ConnectionPolicy_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConnectionType
	if policy.ConnectionType != nil {
		connectionType := string(*policy.ConnectionType)
		destination.ConnectionType = &connectionType
	} else {
		destination.ConnectionType = nil
	}

	// OriginalVersion
	destination.OriginalVersion = policy.OriginalVersion()

	// Owner
	if policy.Owner != nil {
		owner := policy.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (policy *Servers_ConnectionPolicy_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_ConnectionPolicy_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// ConnectionType: The server connection type.
	ConnectionType *ServerConnectionPolicyProperties_ConnectionType_STATUS `json:"connectionType,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Kind: Metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_ConnectionPolicy_STATUS{}

// ConvertStatusFrom populates our Servers_ConnectionPolicy_STATUS from the provided source
func (policy *Servers_ConnectionPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_ConnectionPolicy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_ConnectionPolicy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_ConnectionPolicy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_ConnectionPolicy_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_ConnectionPolicy_STATUS
func (policy *Servers_ConnectionPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_ConnectionPolicy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_ConnectionPolicy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_ConnectionPolicy_STATUS{}
	err := policy.AssignProperties_To_Servers_ConnectionPolicy_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_ConnectionPolicy_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_ConnectionPolicy_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_ConnectionPolicy_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_ConnectionPolicy_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_ConnectionPolicy_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_ConnectionPolicy_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘ConnectionType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConnectionType != nil {
			connectionType := *typedInput.Properties.ConnectionType
			policy.ConnectionType = &connectionType
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		policy.Id = &id
	}

	// Set property ‘Kind’:
	if typedInput.Kind != nil {
		kind := *typedInput.Kind
		policy.Kind = &kind
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		policy.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		policy.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		policy.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_ConnectionPolicy_STATUS populates our Servers_ConnectionPolicy_STATUS from the provided source Servers_ConnectionPolicy_STATUS
func (policy *Servers_ConnectionPolicy_STATUS) AssignProperties_From_Servers_ConnectionPolicy_STATUS(source *v20211101s.Servers_ConnectionPolicy_STATUS) error {

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// ConnectionType
	if source.ConnectionType != nil {
		connectionType := ServerConnectionPolicyProperties_ConnectionType_STATUS(*source.ConnectionType)
		policy.ConnectionType = &connectionType
	} else {
		policy.ConnectionType = nil
	}

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// Kind
	policy.Kind = genruntime.ClonePointerToString(source.Kind)

	// Location
	policy.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_ConnectionPolicy_STATUS populates the provided destination Servers_ConnectionPolicy_STATUS from our Servers_ConnectionPolicy_STATUS
func (policy *Servers_ConnectionPolicy_STATUS) AssignProperties_To_Servers_ConnectionPolicy_STATUS(destination *v20211101s.Servers_ConnectionPolicy_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// ConnectionType
	if policy.ConnectionType != nil {
		connectionType := string(*policy.ConnectionType)
		destination.ConnectionType = &connectionType
	} else {
		destination.ConnectionType = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(policy.Id)

	// Kind
	destination.Kind = genruntime.ClonePointerToString(policy.Kind)

	// Location
	destination.Location = genruntime.ClonePointerToString(policy.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(policy.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(policy.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Default","Proxy","Redirect"}
type ServerConnectionPolicyProperties_ConnectionType string

const (
	ServerConnectionPolicyProperties_ConnectionType_Default  = ServerConnectionPolicyProperties_ConnectionType("Default")
	ServerConnectionPolicyProperties_ConnectionType_Proxy    = ServerConnectionPolicyProperties_ConnectionType("Proxy")
	ServerConnectionPolicyProperties_ConnectionType_Redirect = ServerConnectionPolicyProperties_ConnectionType("Redirect")
)

type ServerConnectionPolicyProperties_ConnectionType_STATUS string

const (
	ServerConnectionPolicyProperties_ConnectionType_STATUS_Default  = ServerConnectionPolicyProperties_ConnectionType_STATUS("Default")
	ServerConnectionPolicyProperties_ConnectionType_STATUS_Proxy    = ServerConnectionPolicyProperties_ConnectionType_STATUS("Proxy")
	ServerConnectionPolicyProperties_ConnectionType_STATUS_Redirect = ServerConnectionPolicyProperties_ConnectionType_STATUS("Redirect")
)

func init() {
	SchemeBuilder.Register(&ServersConnectionPolicy{}, &ServersConnectionPolicyList{})
}
