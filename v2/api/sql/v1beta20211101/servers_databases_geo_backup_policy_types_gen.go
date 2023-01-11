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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/GeoBackupPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/Default
type ServersDatabasesGeoBackupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Databases_GeoBackupPolicy_Spec   `json:"spec,omitempty"`
	Status            Servers_Databases_GeoBackupPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesGeoBackupPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ServersDatabasesGeoBackupPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ServersDatabasesGeoBackupPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersDatabasesGeoBackupPolicy{}

// ConvertFrom populates our ServersDatabasesGeoBackupPolicy from the provided hub ServersDatabasesGeoBackupPolicy
func (policy *ServersDatabasesGeoBackupPolicy) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.ServersDatabasesGeoBackupPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/ServersDatabasesGeoBackupPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_From_ServersDatabasesGeoBackupPolicy(source)
}

// ConvertTo populates the provided hub ServersDatabasesGeoBackupPolicy from our ServersDatabasesGeoBackupPolicy
func (policy *ServersDatabasesGeoBackupPolicy) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.ServersDatabasesGeoBackupPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/ServersDatabasesGeoBackupPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_To_ServersDatabasesGeoBackupPolicy(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1beta20211101-serversdatabasesgeobackuppolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesgeobackuppolicies,verbs=create;update,versions=v1beta20211101,name=default.v1beta20211101.serversdatabasesgeobackuppolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersDatabasesGeoBackupPolicy{}

// Default applies defaults to the ServersDatabasesGeoBackupPolicy resource
func (policy *ServersDatabasesGeoBackupPolicy) Default() {
	policy.defaultImpl()
	var temp interface{} = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersDatabasesGeoBackupPolicy resource
func (policy *ServersDatabasesGeoBackupPolicy) defaultImpl() {}

var _ genruntime.KubernetesResource = &ServersDatabasesGeoBackupPolicy{}

// AzureName returns the Azure name of the resource (always "Default")
func (policy *ServersDatabasesGeoBackupPolicy) AzureName() string {
	return "Default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersDatabasesGeoBackupPolicy) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (policy *ServersDatabasesGeoBackupPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ServersDatabasesGeoBackupPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ServersDatabasesGeoBackupPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/geoBackupPolicies"
func (policy *ServersDatabasesGeoBackupPolicy) GetType() string {
	return "Microsoft.Sql/servers/databases/geoBackupPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ServersDatabasesGeoBackupPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_Databases_GeoBackupPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (policy *ServersDatabasesGeoBackupPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  policy.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (policy *ServersDatabasesGeoBackupPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_Databases_GeoBackupPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Databases_GeoBackupPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1beta20211101-serversdatabasesgeobackuppolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesgeobackuppolicies,verbs=create;update,versions=v1beta20211101,name=validate.v1beta20211101.serversdatabasesgeobackuppolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersDatabasesGeoBackupPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *ServersDatabasesGeoBackupPolicy) ValidateCreate() error {
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
func (policy *ServersDatabasesGeoBackupPolicy) ValidateDelete() error {
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
func (policy *ServersDatabasesGeoBackupPolicy) ValidateUpdate(old runtime.Object) error {
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
func (policy *ServersDatabasesGeoBackupPolicy) createValidations() []func() error {
	return []func() error{policy.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (policy *ServersDatabasesGeoBackupPolicy) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (policy *ServersDatabasesGeoBackupPolicy) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return policy.validateResourceReferences()
		},
		policy.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (policy *ServersDatabasesGeoBackupPolicy) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&policy.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *ServersDatabasesGeoBackupPolicy) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*ServersDatabasesGeoBackupPolicy)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, policy)
}

// AssignProperties_From_ServersDatabasesGeoBackupPolicy populates our ServersDatabasesGeoBackupPolicy from the provided source ServersDatabasesGeoBackupPolicy
func (policy *ServersDatabasesGeoBackupPolicy) AssignProperties_From_ServersDatabasesGeoBackupPolicy(source *v20211101s.ServersDatabasesGeoBackupPolicy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Databases_GeoBackupPolicy_Spec
	err := spec.AssignProperties_From_Servers_Databases_GeoBackupPolicy_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_GeoBackupPolicy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Servers_Databases_GeoBackupPolicy_STATUS
	err = status.AssignProperties_From_Servers_Databases_GeoBackupPolicy_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_GeoBackupPolicy_STATUS() to populate field Status")
	}
	policy.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersDatabasesGeoBackupPolicy populates the provided destination ServersDatabasesGeoBackupPolicy from our ServersDatabasesGeoBackupPolicy
func (policy *ServersDatabasesGeoBackupPolicy) AssignProperties_To_ServersDatabasesGeoBackupPolicy(destination *v20211101s.ServersDatabasesGeoBackupPolicy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_Databases_GeoBackupPolicy_Spec
	err := policy.Spec.AssignProperties_To_Servers_Databases_GeoBackupPolicy_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_GeoBackupPolicy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_Databases_GeoBackupPolicy_STATUS
	err = policy.Status.AssignProperties_To_Servers_Databases_GeoBackupPolicy_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_GeoBackupPolicy_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ServersDatabasesGeoBackupPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion(),
		Kind:    "ServersDatabasesGeoBackupPolicy",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/GeoBackupPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/Default
type ServersDatabasesGeoBackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesGeoBackupPolicy `json:"items"`
}

type Servers_Databases_GeoBackupPolicy_Spec struct {
	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`

	// +kubebuilder:validation:Required
	// State: The state of the geo backup policy.
	State *GeoBackupPolicyProperties_State `json:"state,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_Databases_GeoBackupPolicy_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (policy *Servers_Databases_GeoBackupPolicy_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if policy == nil {
		return nil, nil
	}
	result := &Servers_Databases_GeoBackupPolicy_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if policy.State != nil {
		result.Properties = &GeoBackupPolicyProperties_ARM{}
	}
	if policy.State != nil {
		state := *policy.State
		result.Properties.State = &state
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_Databases_GeoBackupPolicy_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_GeoBackupPolicy_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_Databases_GeoBackupPolicy_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_GeoBackupPolicy_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_GeoBackupPolicy_Spec_ARM, got %T", armInput)
	}

	// Set property ‘Owner’:
	policy.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘State’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			policy.State = &state
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_Databases_GeoBackupPolicy_Spec{}

// ConvertSpecFrom populates our Servers_Databases_GeoBackupPolicy_Spec from the provided source
func (policy *Servers_Databases_GeoBackupPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_Databases_GeoBackupPolicy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_Databases_GeoBackupPolicy_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_Databases_GeoBackupPolicy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_Databases_GeoBackupPolicy_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Databases_GeoBackupPolicy_Spec
func (policy *Servers_Databases_GeoBackupPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_Databases_GeoBackupPolicy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_Databases_GeoBackupPolicy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_Databases_GeoBackupPolicy_Spec{}
	err := policy.AssignProperties_To_Servers_Databases_GeoBackupPolicy_Spec(dst)
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

// AssignProperties_From_Servers_Databases_GeoBackupPolicy_Spec populates our Servers_Databases_GeoBackupPolicy_Spec from the provided source Servers_Databases_GeoBackupPolicy_Spec
func (policy *Servers_Databases_GeoBackupPolicy_Spec) AssignProperties_From_Servers_Databases_GeoBackupPolicy_Spec(source *v20211101s.Servers_Databases_GeoBackupPolicy_Spec) error {

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// State
	if source.State != nil {
		state := GeoBackupPolicyProperties_State(*source.State)
		policy.State = &state
	} else {
		policy.State = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_GeoBackupPolicy_Spec populates the provided destination Servers_Databases_GeoBackupPolicy_Spec from our Servers_Databases_GeoBackupPolicy_Spec
func (policy *Servers_Databases_GeoBackupPolicy_Spec) AssignProperties_To_Servers_Databases_GeoBackupPolicy_Spec(destination *v20211101s.Servers_Databases_GeoBackupPolicy_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// OriginalVersion
	destination.OriginalVersion = policy.OriginalVersion()

	// Owner
	if policy.Owner != nil {
		owner := policy.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// State
	if policy.State != nil {
		state := string(*policy.State)
		destination.State = &state
	} else {
		destination.State = nil
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
func (policy *Servers_Databases_GeoBackupPolicy_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_Databases_GeoBackupPolicy_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Kind: Kind of geo backup policy.  This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	// Location: Backup policy location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// State: The state of the geo backup policy.
	State *GeoBackupPolicyProperties_State_STATUS `json:"state,omitempty"`

	// StorageType: The storage type of the geo backup policy.
	StorageType *string `json:"storageType,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Databases_GeoBackupPolicy_STATUS{}

// ConvertStatusFrom populates our Servers_Databases_GeoBackupPolicy_STATUS from the provided source
func (policy *Servers_Databases_GeoBackupPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_Databases_GeoBackupPolicy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_Databases_GeoBackupPolicy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_Databases_GeoBackupPolicy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_Databases_GeoBackupPolicy_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_Databases_GeoBackupPolicy_STATUS
func (policy *Servers_Databases_GeoBackupPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_Databases_GeoBackupPolicy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_Databases_GeoBackupPolicy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_Databases_GeoBackupPolicy_STATUS{}
	err := policy.AssignProperties_To_Servers_Databases_GeoBackupPolicy_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_Databases_GeoBackupPolicy_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_Databases_GeoBackupPolicy_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_GeoBackupPolicy_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_Databases_GeoBackupPolicy_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_GeoBackupPolicy_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_GeoBackupPolicy_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

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

	// Set property ‘State’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			policy.State = &state
		}
	}

	// Set property ‘StorageType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StorageType != nil {
			storageType := *typedInput.Properties.StorageType
			policy.StorageType = &storageType
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		policy.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_Databases_GeoBackupPolicy_STATUS populates our Servers_Databases_GeoBackupPolicy_STATUS from the provided source Servers_Databases_GeoBackupPolicy_STATUS
func (policy *Servers_Databases_GeoBackupPolicy_STATUS) AssignProperties_From_Servers_Databases_GeoBackupPolicy_STATUS(source *v20211101s.Servers_Databases_GeoBackupPolicy_STATUS) error {

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// Kind
	policy.Kind = genruntime.ClonePointerToString(source.Kind)

	// Location
	policy.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// State
	if source.State != nil {
		state := GeoBackupPolicyProperties_State_STATUS(*source.State)
		policy.State = &state
	} else {
		policy.State = nil
	}

	// StorageType
	policy.StorageType = genruntime.ClonePointerToString(source.StorageType)

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_GeoBackupPolicy_STATUS populates the provided destination Servers_Databases_GeoBackupPolicy_STATUS from our Servers_Databases_GeoBackupPolicy_STATUS
func (policy *Servers_Databases_GeoBackupPolicy_STATUS) AssignProperties_To_Servers_Databases_GeoBackupPolicy_STATUS(destination *v20211101s.Servers_Databases_GeoBackupPolicy_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(policy.Id)

	// Kind
	destination.Kind = genruntime.ClonePointerToString(policy.Kind)

	// Location
	destination.Location = genruntime.ClonePointerToString(policy.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(policy.Name)

	// State
	if policy.State != nil {
		state := string(*policy.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

	// StorageType
	destination.StorageType = genruntime.ClonePointerToString(policy.StorageType)

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

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type GeoBackupPolicyProperties_State string

const (
	GeoBackupPolicyProperties_State_Disabled = GeoBackupPolicyProperties_State("Disabled")
	GeoBackupPolicyProperties_State_Enabled  = GeoBackupPolicyProperties_State("Enabled")
)

type GeoBackupPolicyProperties_State_STATUS string

const (
	GeoBackupPolicyProperties_State_STATUS_Disabled = GeoBackupPolicyProperties_State_STATUS("Disabled")
	GeoBackupPolicyProperties_State_STATUS_Enabled  = GeoBackupPolicyProperties_State_STATUS("Enabled")
)

func init() {
	SchemeBuilder.Register(&ServersDatabasesGeoBackupPolicy{}, &ServersDatabasesGeoBackupPolicyList{})
}
