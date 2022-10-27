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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/JobAgents.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}
type Servers_JobAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_JobAgent_Spec   `json:"spec,omitempty"`
	Status            Servers_JobAgent_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Servers_JobAgent{}

// GetConditions returns the conditions of the resource
func (agent *Servers_JobAgent) GetConditions() conditions.Conditions {
	return agent.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (agent *Servers_JobAgent) SetConditions(conditions conditions.Conditions) {
	agent.Status.Conditions = conditions
}

var _ conversion.Convertible = &Servers_JobAgent{}

// ConvertFrom populates our Servers_JobAgent from the provided hub Servers_JobAgent
func (agent *Servers_JobAgent) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.Servers_JobAgent)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/Servers_JobAgent but received %T instead", hub)
	}

	return agent.AssignProperties_From_Servers_JobAgent(source)
}

// ConvertTo populates the provided hub Servers_JobAgent from our Servers_JobAgent
func (agent *Servers_JobAgent) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.Servers_JobAgent)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/Servers_JobAgent but received %T instead", hub)
	}

	return agent.AssignProperties_To_Servers_JobAgent(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1beta20211101-servers_jobagent,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=servers_jobagents,verbs=create;update,versions=v1beta20211101,name=default.v1beta20211101.servers_jobagents.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &Servers_JobAgent{}

// Default applies defaults to the Servers_JobAgent resource
func (agent *Servers_JobAgent) Default() {
	agent.defaultImpl()
	var temp interface{} = agent
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (agent *Servers_JobAgent) defaultAzureName() {
	if agent.Spec.AzureName == "" {
		agent.Spec.AzureName = agent.Name
	}
}

// defaultImpl applies the code generated defaults to the Servers_JobAgent resource
func (agent *Servers_JobAgent) defaultImpl() { agent.defaultAzureName() }

var _ genruntime.KubernetesResource = &Servers_JobAgent{}

// AzureName returns the Azure name of the resource
func (agent *Servers_JobAgent) AzureName() string {
	return agent.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (agent Servers_JobAgent) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (agent *Servers_JobAgent) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (agent *Servers_JobAgent) GetSpec() genruntime.ConvertibleSpec {
	return &agent.Spec
}

// GetStatus returns the status of this resource
func (agent *Servers_JobAgent) GetStatus() genruntime.ConvertibleStatus {
	return &agent.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/jobAgents"
func (agent *Servers_JobAgent) GetType() string {
	return "Microsoft.Sql/servers/jobAgents"
}

// NewEmptyStatus returns a new empty (blank) status
func (agent *Servers_JobAgent) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_JobAgent_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (agent *Servers_JobAgent) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(agent.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  agent.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (agent *Servers_JobAgent) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_JobAgent_STATUS); ok {
		agent.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_JobAgent_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	agent.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1beta20211101-servers_jobagent,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=servers_jobagents,verbs=create;update,versions=v1beta20211101,name=validate.v1beta20211101.servers_jobagents.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &Servers_JobAgent{}

// ValidateCreate validates the creation of the resource
func (agent *Servers_JobAgent) ValidateCreate() error {
	validations := agent.createValidations()
	var temp interface{} = agent
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
func (agent *Servers_JobAgent) ValidateDelete() error {
	validations := agent.deleteValidations()
	var temp interface{} = agent
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
func (agent *Servers_JobAgent) ValidateUpdate(old runtime.Object) error {
	validations := agent.updateValidations()
	var temp interface{} = agent
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
func (agent *Servers_JobAgent) createValidations() []func() error {
	return []func() error{agent.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (agent *Servers_JobAgent) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (agent *Servers_JobAgent) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return agent.validateResourceReferences()
		},
		agent.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (agent *Servers_JobAgent) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&agent.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (agent *Servers_JobAgent) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*Servers_JobAgent)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, agent)
}

// AssignProperties_From_Servers_JobAgent populates our Servers_JobAgent from the provided source Servers_JobAgent
func (agent *Servers_JobAgent) AssignProperties_From_Servers_JobAgent(source *v20211101s.Servers_JobAgent) error {

	// ObjectMeta
	agent.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_JobAgent_Spec
	err := spec.AssignProperties_From_Servers_JobAgent_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_JobAgent_Spec() to populate field Spec")
	}
	agent.Spec = spec

	// Status
	var status Servers_JobAgent_STATUS
	err = status.AssignProperties_From_Servers_JobAgent_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_JobAgent_STATUS() to populate field Status")
	}
	agent.Status = status

	// No error
	return nil
}

// AssignProperties_To_Servers_JobAgent populates the provided destination Servers_JobAgent from our Servers_JobAgent
func (agent *Servers_JobAgent) AssignProperties_To_Servers_JobAgent(destination *v20211101s.Servers_JobAgent) error {

	// ObjectMeta
	destination.ObjectMeta = *agent.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_JobAgent_Spec
	err := agent.Spec.AssignProperties_To_Servers_JobAgent_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_JobAgent_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_JobAgent_STATUS
	err = agent.Status.AssignProperties_To_Servers_JobAgent_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_JobAgent_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (agent *Servers_JobAgent) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: agent.Spec.OriginalVersion(),
		Kind:    "Servers_JobAgent",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/JobAgents.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}
type Servers_JobAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Servers_JobAgent `json:"items"`
}

type Servers_JobAgent_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// DatabaseReference: Resource ID of the database to store job metadata in.
	DatabaseReference *genruntime.ResourceReference `armReference:"DatabaseId" json:"databaseReference,omitempty"`

	// +kubebuilder:validation:Required
	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`

	// Sku: The name and tier of the SKU.
	Sku *Sku `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_JobAgent_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (agent *Servers_JobAgent_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if agent == nil {
		return nil, nil
	}
	result := &Servers_JobAgent_Spec_ARM{}

	// Set property ‘Location’:
	if agent.Location != nil {
		location := *agent.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if agent.DatabaseReference != nil {
		result.Properties = &JobAgentProperties_ARM{}
	}
	if agent.DatabaseReference != nil {
		databaseIdARMID, err := resolved.ResolvedReferences.Lookup(*agent.DatabaseReference)
		if err != nil {
			return nil, err
		}
		databaseId := databaseIdARMID
		result.Properties.DatabaseId = &databaseId
	}

	// Set property ‘Sku’:
	if agent.Sku != nil {
		sku_ARM, err := (*agent.Sku).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		sku := *sku_ARM.(*Sku_ARM)
		result.Sku = &sku
	}

	// Set property ‘Tags’:
	if agent.Tags != nil {
		result.Tags = make(map[string]string, len(agent.Tags))
		for key, value := range agent.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (agent *Servers_JobAgent_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_JobAgent_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (agent *Servers_JobAgent_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_JobAgent_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_JobAgent_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	agent.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// no assignment for property ‘DatabaseReference’

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		agent.Location = &location
	}

	// Set property ‘Owner’:
	agent.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘Sku’:
	if typedInput.Sku != nil {
		var sku1 Sku
		err := sku1.PopulateFromARM(owner, *typedInput.Sku)
		if err != nil {
			return err
		}
		sku := sku1
		agent.Sku = &sku
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		agent.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			agent.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_JobAgent_Spec{}

// ConvertSpecFrom populates our Servers_JobAgent_Spec from the provided source
func (agent *Servers_JobAgent_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_JobAgent_Spec)
	if ok {
		// Populate our instance from source
		return agent.AssignProperties_From_Servers_JobAgent_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_JobAgent_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = agent.AssignProperties_From_Servers_JobAgent_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_JobAgent_Spec
func (agent *Servers_JobAgent_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_JobAgent_Spec)
	if ok {
		// Populate destination from our instance
		return agent.AssignProperties_To_Servers_JobAgent_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_JobAgent_Spec{}
	err := agent.AssignProperties_To_Servers_JobAgent_Spec(dst)
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

// AssignProperties_From_Servers_JobAgent_Spec populates our Servers_JobAgent_Spec from the provided source Servers_JobAgent_Spec
func (agent *Servers_JobAgent_Spec) AssignProperties_From_Servers_JobAgent_Spec(source *v20211101s.Servers_JobAgent_Spec) error {

	// AzureName
	agent.AzureName = source.AzureName

	// DatabaseReference
	if source.DatabaseReference != nil {
		databaseReference := source.DatabaseReference.Copy()
		agent.DatabaseReference = &databaseReference
	} else {
		agent.DatabaseReference = nil
	}

	// Location
	agent.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		agent.Owner = &owner
	} else {
		agent.Owner = nil
	}

	// Sku
	if source.Sku != nil {
		var sku Sku
		err := sku.AssignProperties_From_Sku(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku() to populate field Sku")
		}
		agent.Sku = &sku
	} else {
		agent.Sku = nil
	}

	// Tags
	agent.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_Servers_JobAgent_Spec populates the provided destination Servers_JobAgent_Spec from our Servers_JobAgent_Spec
func (agent *Servers_JobAgent_Spec) AssignProperties_To_Servers_JobAgent_Spec(destination *v20211101s.Servers_JobAgent_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = agent.AzureName

	// DatabaseReference
	if agent.DatabaseReference != nil {
		databaseReference := agent.DatabaseReference.Copy()
		destination.DatabaseReference = &databaseReference
	} else {
		destination.DatabaseReference = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(agent.Location)

	// OriginalVersion
	destination.OriginalVersion = agent.OriginalVersion()

	// Owner
	if agent.Owner != nil {
		owner := agent.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Sku
	if agent.Sku != nil {
		var sku v20211101s.Sku
		err := agent.Sku.AssignProperties_To_Sku(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku() to populate field Sku")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(agent.Tags)

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
func (agent *Servers_JobAgent_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (agent *Servers_JobAgent_Spec) SetAzureName(azureName string) { agent.AzureName = azureName }

type Servers_JobAgent_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// DatabaseId: Resource ID of the database to store job metadata in.
	DatabaseId *string `json:"databaseId,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Sku: The name and tier of the SKU.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// State: The state of the job agent.
	State *JobAgentProperties_State_STATUS `json:"state,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_JobAgent_STATUS{}

// ConvertStatusFrom populates our Servers_JobAgent_STATUS from the provided source
func (agent *Servers_JobAgent_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_JobAgent_STATUS)
	if ok {
		// Populate our instance from source
		return agent.AssignProperties_From_Servers_JobAgent_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_JobAgent_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = agent.AssignProperties_From_Servers_JobAgent_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_JobAgent_STATUS
func (agent *Servers_JobAgent_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_JobAgent_STATUS)
	if ok {
		// Populate destination from our instance
		return agent.AssignProperties_To_Servers_JobAgent_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_JobAgent_STATUS{}
	err := agent.AssignProperties_To_Servers_JobAgent_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_JobAgent_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (agent *Servers_JobAgent_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_JobAgent_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (agent *Servers_JobAgent_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_JobAgent_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_JobAgent_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘DatabaseId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DatabaseId != nil {
			databaseId := *typedInput.Properties.DatabaseId
			agent.DatabaseId = &databaseId
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		agent.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		agent.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		agent.Name = &name
	}

	// Set property ‘Sku’:
	if typedInput.Sku != nil {
		var sku1 Sku_STATUS
		err := sku1.PopulateFromARM(owner, *typedInput.Sku)
		if err != nil {
			return err
		}
		sku := sku1
		agent.Sku = &sku
	}

	// Set property ‘State’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			agent.State = &state
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		agent.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			agent.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		agent.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_JobAgent_STATUS populates our Servers_JobAgent_STATUS from the provided source Servers_JobAgent_STATUS
func (agent *Servers_JobAgent_STATUS) AssignProperties_From_Servers_JobAgent_STATUS(source *v20211101s.Servers_JobAgent_STATUS) error {

	// Conditions
	agent.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DatabaseId
	agent.DatabaseId = genruntime.ClonePointerToString(source.DatabaseId)

	// Id
	agent.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	agent.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	agent.Name = genruntime.ClonePointerToString(source.Name)

	// Sku
	if source.Sku != nil {
		var sku Sku_STATUS
		err := sku.AssignProperties_From_Sku_STATUS(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku_STATUS() to populate field Sku")
		}
		agent.Sku = &sku
	} else {
		agent.Sku = nil
	}

	// State
	if source.State != nil {
		state := JobAgentProperties_State_STATUS(*source.State)
		agent.State = &state
	} else {
		agent.State = nil
	}

	// Tags
	agent.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	agent.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_JobAgent_STATUS populates the provided destination Servers_JobAgent_STATUS from our Servers_JobAgent_STATUS
func (agent *Servers_JobAgent_STATUS) AssignProperties_To_Servers_JobAgent_STATUS(destination *v20211101s.Servers_JobAgent_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(agent.Conditions)

	// DatabaseId
	destination.DatabaseId = genruntime.ClonePointerToString(agent.DatabaseId)

	// Id
	destination.Id = genruntime.ClonePointerToString(agent.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(agent.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(agent.Name)

	// Sku
	if agent.Sku != nil {
		var sku v20211101s.Sku_STATUS
		err := agent.Sku.AssignProperties_To_Sku_STATUS(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku_STATUS() to populate field Sku")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// State
	if agent.State != nil {
		state := string(*agent.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(agent.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(agent.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type JobAgentProperties_State_STATUS string

const (
	JobAgentProperties_State_STATUS_Creating = JobAgentProperties_State_STATUS("Creating")
	JobAgentProperties_State_STATUS_Deleting = JobAgentProperties_State_STATUS("Deleting")
	JobAgentProperties_State_STATUS_Disabled = JobAgentProperties_State_STATUS("Disabled")
	JobAgentProperties_State_STATUS_Ready    = JobAgentProperties_State_STATUS("Ready")
	JobAgentProperties_State_STATUS_Updating = JobAgentProperties_State_STATUS("Updating")
)

func init() {
	SchemeBuilder.Register(&Servers_JobAgent{}, &Servers_JobAgentList{})
}
