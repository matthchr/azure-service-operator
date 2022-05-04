// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

import (
	"fmt"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20210601storage"
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
// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_databases
type FlexibleServersDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersDatabases_Spec `json:"spec,omitempty"`
	Status            Database_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersDatabase{}

// GetConditions returns the conditions of the resource
func (database *FlexibleServersDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *FlexibleServersDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersDatabase{}

// ConvertFrom populates our FlexibleServersDatabase from the provided hub FlexibleServersDatabase
func (database *FlexibleServersDatabase) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210601s.FlexibleServersDatabase)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1beta20210601storage/FlexibleServersDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesFromFlexibleServersDatabase(source)
}

// ConvertTo populates the provided hub FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210601s.FlexibleServersDatabase)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1beta20210601storage/FlexibleServersDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesToFlexibleServersDatabase(destination)
}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1beta20210601-flexibleserversdatabase,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversdatabases,verbs=create;update,versions=v1beta20210601,name=default.v1beta20210601.flexibleserversdatabases.dbforpostgresql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &FlexibleServersDatabase{}

// Default applies defaults to the FlexibleServersDatabase resource
func (database *FlexibleServersDatabase) Default() {
	database.defaultImpl()
	var temp interface{} = database
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (database *FlexibleServersDatabase) defaultAzureName() {
	if database.Spec.AzureName == "" {
		database.Spec.AzureName = database.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersDatabase resource
func (database *FlexibleServersDatabase) defaultImpl() { database.defaultAzureName() }

var _ genruntime.KubernetesResource = &FlexibleServersDatabase{}

// AzureName returns the Azure name of the resource
func (database *FlexibleServersDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (database FlexibleServersDatabase) GetAPIVersion() string {
	return "2021-06-01"
}

// GetResourceKind returns the kind of the resource
func (database *FlexibleServersDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (database *FlexibleServersDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *FlexibleServersDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/databases"
func (database *FlexibleServersDatabase) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *FlexibleServersDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Database_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *FlexibleServersDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *FlexibleServersDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Database_Status); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st Database_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1beta20210601-flexibleserversdatabase,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversdatabases,verbs=create;update,versions=v1beta20210601,name=validate.v1beta20210601.flexibleserversdatabases.dbforpostgresql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &FlexibleServersDatabase{}

// ValidateCreate validates the creation of the resource
func (database *FlexibleServersDatabase) ValidateCreate() error {
	validations := database.createValidations()
	var temp interface{} = database
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
func (database *FlexibleServersDatabase) ValidateDelete() error {
	validations := database.deleteValidations()
	var temp interface{} = database
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
func (database *FlexibleServersDatabase) ValidateUpdate(old runtime.Object) error {
	validations := database.updateValidations()
	var temp interface{} = database
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
func (database *FlexibleServersDatabase) createValidations() []func() error {
	return []func() error{database.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (database *FlexibleServersDatabase) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (database *FlexibleServersDatabase) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return database.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (database *FlexibleServersDatabase) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&database.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromFlexibleServersDatabase populates our FlexibleServersDatabase from the provided source FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignPropertiesFromFlexibleServersDatabase(source *v20210601s.FlexibleServersDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersDatabases_Spec
	err := spec.AssignPropertiesFromFlexibleServersDatabasesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersDatabasesSpec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status Database_Status
	err = status.AssignPropertiesFromDatabaseStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseStatus() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersDatabase populates the provided destination FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignPropertiesToFlexibleServersDatabase(destination *v20210601s.FlexibleServersDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210601s.FlexibleServersDatabases_Spec
	err := database.Spec.AssignPropertiesToFlexibleServersDatabasesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersDatabasesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210601s.Database_Status
	err = database.Status.AssignPropertiesToDatabaseStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *FlexibleServersDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion(),
		Kind:    "FlexibleServersDatabase",
	}
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_databases
type FlexibleServersDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersDatabase `json:"items"`
}

type Database_Status struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_Status `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Database_Status{}

// ConvertStatusFrom populates our Database_Status from the provided source
func (database *Database_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210601s.Database_Status)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.Database_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Database_Status
func (database *Database_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210601s.Database_Status)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.Database_Status{}
	err := database.AssignPropertiesToDatabaseStatus(dst)
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

var _ genruntime.FromARMConverter = &Database_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *Database_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Database_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *Database_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Database_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Database_StatusARM, got %T", armInput)
	}

	// Set property ‘Charset’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Charset != nil {
			charset := *typedInput.Properties.Charset
			database.Charset = &charset
		}
	}

	// Set property ‘Collation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Collation != nil {
			collation := *typedInput.Properties.Collation
			database.Collation = &collation
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		database.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		database.Name = &name
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_Status
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		database.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		database.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromDatabaseStatus populates our Database_Status from the provided source Database_Status
func (database *Database_Status) AssignPropertiesFromDatabaseStatus(source *v20210601s.Database_Status) error {

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		database.SystemData = &systemDatum
	} else {
		database.SystemData = nil
	}

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToDatabaseStatus populates the provided destination Database_Status from our Database_Status
func (database *Database_Status) AssignPropertiesToDatabaseStatus(destination *v20210601s.Database_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// SystemData
	if database.SystemData != nil {
		var systemDatum v20210601s.SystemData_Status
		err := database.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2021-06-01"}
type FlexibleServersDatabasesSpecAPIVersion string

const FlexibleServersDatabasesSpecAPIVersion20210601 = FlexibleServersDatabasesSpecAPIVersion("2021-06-01")

type FlexibleServersDatabases_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &FlexibleServersDatabases_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (databases *FlexibleServersDatabases_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if databases == nil {
		return nil, nil
	}
	var result FlexibleServersDatabases_SpecARM

	// Set property ‘Location’:
	if databases.Location != nil {
		location := *databases.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if databases.Charset != nil || databases.Collation != nil {
		result.Properties = &DatabasePropertiesARM{}
	}
	if databases.Charset != nil {
		charset := *databases.Charset
		result.Properties.Charset = &charset
	}
	if databases.Collation != nil {
		collation := *databases.Collation
		result.Properties.Collation = &collation
	}

	// Set property ‘Tags’:
	if databases.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range databases.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (databases *FlexibleServersDatabases_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServersDatabases_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (databases *FlexibleServersDatabases_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServersDatabases_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServersDatabases_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	databases.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Charset’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Charset != nil {
			charset := *typedInput.Properties.Charset
			databases.Charset = &charset
		}
	}

	// Set property ‘Collation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Collation != nil {
			collation := *typedInput.Properties.Collation
			databases.Collation = &collation
		}
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		databases.Location = &location
	}

	// Set property ‘Owner’:
	databases.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		databases.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			databases.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServersDatabases_Spec{}

// ConvertSpecFrom populates our FlexibleServersDatabases_Spec from the provided source
func (databases *FlexibleServersDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210601s.FlexibleServersDatabases_Spec)
	if ok {
		// Populate our instance from source
		return databases.AssignPropertiesFromFlexibleServersDatabasesSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.FlexibleServersDatabases_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databases.AssignPropertiesFromFlexibleServersDatabasesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersDatabases_Spec
func (databases *FlexibleServersDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210601s.FlexibleServersDatabases_Spec)
	if ok {
		// Populate destination from our instance
		return databases.AssignPropertiesToFlexibleServersDatabasesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.FlexibleServersDatabases_Spec{}
	err := databases.AssignPropertiesToFlexibleServersDatabasesSpec(dst)
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

// AssignPropertiesFromFlexibleServersDatabasesSpec populates our FlexibleServersDatabases_Spec from the provided source FlexibleServersDatabases_Spec
func (databases *FlexibleServersDatabases_Spec) AssignPropertiesFromFlexibleServersDatabasesSpec(source *v20210601s.FlexibleServersDatabases_Spec) error {

	// AzureName
	databases.AzureName = source.AzureName

	// Charset
	databases.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	databases.Collation = genruntime.ClonePointerToString(source.Collation)

	// Location
	databases.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		databases.Owner = &owner
	} else {
		databases.Owner = nil
	}

	// Tags
	databases.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersDatabasesSpec populates the provided destination FlexibleServersDatabases_Spec from our FlexibleServersDatabases_Spec
func (databases *FlexibleServersDatabases_Spec) AssignPropertiesToFlexibleServersDatabasesSpec(destination *v20210601s.FlexibleServersDatabases_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = databases.AzureName

	// Charset
	destination.Charset = genruntime.ClonePointerToString(databases.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(databases.Collation)

	// Location
	destination.Location = genruntime.ClonePointerToString(databases.Location)

	// OriginalVersion
	destination.OriginalVersion = databases.OriginalVersion()

	// Owner
	if databases.Owner != nil {
		owner := databases.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(databases.Tags)

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
func (databases *FlexibleServersDatabases_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (databases *FlexibleServersDatabases_Spec) SetAzureName(azureName string) {
	databases.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&FlexibleServersDatabase{}, &FlexibleServersDatabaseList{})
}
