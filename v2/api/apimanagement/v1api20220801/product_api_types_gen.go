// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis/{apiId}
type ProductApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductApi_Spec   `json:"spec,omitempty"`
	Status            ProductApi_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ProductApi{}

// GetConditions returns the conditions of the resource
func (productApi *ProductApi) GetConditions() conditions.Conditions {
	return productApi.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (productApi *ProductApi) SetConditions(conditions conditions.Conditions) {
	productApi.Status.Conditions = conditions
}

var _ conversion.Convertible = &ProductApi{}

// ConvertFrom populates our ProductApi from the provided hub ProductApi
func (productApi *ProductApi) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.ProductApi)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/ProductApi but received %T instead", hub)
	}

	return productApi.AssignProperties_From_ProductApi(source)
}

// ConvertTo populates the provided hub ProductApi from our ProductApi
func (productApi *ProductApi) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.ProductApi)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/ProductApi but received %T instead", hub)
	}

	return productApi.AssignProperties_To_ProductApi(destination)
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20220801-productapi,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=productapis,verbs=create;update,versions=v1api20220801,name=default.v1api20220801.productapis.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ProductApi{}

// Default applies defaults to the ProductApi resource
func (productApi *ProductApi) Default() {
	productApi.defaultImpl()
	var temp any = productApi
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (productApi *ProductApi) defaultAzureName() {
	if productApi.Spec.AzureName == "" {
		productApi.Spec.AzureName = productApi.Name
	}
}

// defaultImpl applies the code generated defaults to the ProductApi resource
func (productApi *ProductApi) defaultImpl() { productApi.defaultAzureName() }

var _ configmaps.Exporter = &ProductApi{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (productApi *ProductApi) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if productApi.Spec.OperatorSpec == nil {
		return nil
	}
	return productApi.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ProductApi{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (productApi *ProductApi) SecretDestinationExpressions() []*core.DestinationExpression {
	if productApi.Spec.OperatorSpec == nil {
		return nil
	}
	return productApi.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &ProductApi{}

// InitializeSpec initializes the spec for this resource from the given status
func (productApi *ProductApi) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*ProductApi_STATUS); ok {
		return productApi.Spec.Initialize_From_ProductApi_STATUS(s)
	}

	return fmt.Errorf("expected Status of type ProductApi_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ProductApi{}

// AzureName returns the Azure name of the resource
func (productApi *ProductApi) AzureName() string {
	return productApi.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (productApi ProductApi) GetAPIVersion() string {
	return "2022-08-01"
}

// GetResourceScope returns the scope of the resource
func (productApi *ProductApi) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (productApi *ProductApi) GetSpec() genruntime.ConvertibleSpec {
	return &productApi.Spec
}

// GetStatus returns the status of this resource
func (productApi *ProductApi) GetStatus() genruntime.ConvertibleStatus {
	return &productApi.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (productApi *ProductApi) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/products/apis"
func (productApi *ProductApi) GetType() string {
	return "Microsoft.ApiManagement/service/products/apis"
}

// NewEmptyStatus returns a new empty (blank) status
func (productApi *ProductApi) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ProductApi_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (productApi *ProductApi) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(productApi.Spec)
	return productApi.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (productApi *ProductApi) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ProductApi_STATUS); ok {
		productApi.Status = *st
		return nil
	}

	// Convert status to required version
	var st ProductApi_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	productApi.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20220801-productapi,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=productapis,verbs=create;update,versions=v1api20220801,name=validate.v1api20220801.productapis.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ProductApi{}

// ValidateCreate validates the creation of the resource
func (productApi *ProductApi) ValidateCreate() (admission.Warnings, error) {
	validations := productApi.createValidations()
	var temp any = productApi
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (productApi *ProductApi) ValidateDelete() (admission.Warnings, error) {
	validations := productApi.deleteValidations()
	var temp any = productApi
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (productApi *ProductApi) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := productApi.updateValidations()
	var temp any = productApi
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (productApi *ProductApi) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){productApi.validateResourceReferences, productApi.validateOwnerReference, productApi.validateSecretDestinations, productApi.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (productApi *ProductApi) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (productApi *ProductApi) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return productApi.validateResourceReferences()
		},
		productApi.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return productApi.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return productApi.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return productApi.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (productApi *ProductApi) validateConfigMapDestinations() (admission.Warnings, error) {
	if productApi.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(productApi, nil, productApi.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (productApi *ProductApi) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(productApi)
}

// validateResourceReferences validates all resource references
func (productApi *ProductApi) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&productApi.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (productApi *ProductApi) validateSecretDestinations() (admission.Warnings, error) {
	if productApi.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(productApi, nil, productApi.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (productApi *ProductApi) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ProductApi)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, productApi)
}

// AssignProperties_From_ProductApi populates our ProductApi from the provided source ProductApi
func (productApi *ProductApi) AssignProperties_From_ProductApi(source *storage.ProductApi) error {

	// ObjectMeta
	productApi.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec ProductApi_Spec
	err := spec.AssignProperties_From_ProductApi_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ProductApi_Spec() to populate field Spec")
	}
	productApi.Spec = spec

	// Status
	var status ProductApi_STATUS
	err = status.AssignProperties_From_ProductApi_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ProductApi_STATUS() to populate field Status")
	}
	productApi.Status = status

	// No error
	return nil
}

// AssignProperties_To_ProductApi populates the provided destination ProductApi from our ProductApi
func (productApi *ProductApi) AssignProperties_To_ProductApi(destination *storage.ProductApi) error {

	// ObjectMeta
	destination.ObjectMeta = *productApi.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.ProductApi_Spec
	err := productApi.Spec.AssignProperties_To_ProductApi_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ProductApi_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.ProductApi_STATUS
	err = productApi.Status.AssignProperties_To_ProductApi_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ProductApi_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (productApi *ProductApi) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: productApi.Spec.OriginalVersion(),
		Kind:    "ProductApi",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis/{apiId}
type ProductApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductApi `json:"items"`
}

type ProductApi_Spec struct {
	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="^[^*#&+:<>?]+$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *ProductApiOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Product resource
	Owner *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Product"`
}

var _ genruntime.ARMTransformer = &ProductApi_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (productApi *ProductApi_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if productApi == nil {
		return nil, nil
	}
	result := &ProductApi_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (productApi *ProductApi_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ProductApi_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (productApi *ProductApi_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ProductApi_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ProductApi_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	productApi.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	productApi.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &ProductApi_Spec{}

// ConvertSpecFrom populates our ProductApi_Spec from the provided source
func (productApi *ProductApi_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.ProductApi_Spec)
	if ok {
		// Populate our instance from source
		return productApi.AssignProperties_From_ProductApi_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.ProductApi_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = productApi.AssignProperties_From_ProductApi_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our ProductApi_Spec
func (productApi *ProductApi_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.ProductApi_Spec)
	if ok {
		// Populate destination from our instance
		return productApi.AssignProperties_To_ProductApi_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.ProductApi_Spec{}
	err := productApi.AssignProperties_To_ProductApi_Spec(dst)
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

// AssignProperties_From_ProductApi_Spec populates our ProductApi_Spec from the provided source ProductApi_Spec
func (productApi *ProductApi_Spec) AssignProperties_From_ProductApi_Spec(source *storage.ProductApi_Spec) error {

	// AzureName
	productApi.AzureName = source.AzureName

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec ProductApiOperatorSpec
		err := operatorSpec.AssignProperties_From_ProductApiOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ProductApiOperatorSpec() to populate field OperatorSpec")
		}
		productApi.OperatorSpec = &operatorSpec
	} else {
		productApi.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		productApi.Owner = &owner
	} else {
		productApi.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_ProductApi_Spec populates the provided destination ProductApi_Spec from our ProductApi_Spec
func (productApi *ProductApi_Spec) AssignProperties_To_ProductApi_Spec(destination *storage.ProductApi_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = productApi.AzureName

	// OperatorSpec
	if productApi.OperatorSpec != nil {
		var operatorSpec storage.ProductApiOperatorSpec
		err := productApi.OperatorSpec.AssignProperties_To_ProductApiOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ProductApiOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = productApi.OriginalVersion()

	// Owner
	if productApi.Owner != nil {
		owner := productApi.Owner.Copy()
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

// Initialize_From_ProductApi_STATUS populates our ProductApi_Spec from the provided source ProductApi_STATUS
func (productApi *ProductApi_Spec) Initialize_From_ProductApi_STATUS(source *ProductApi_STATUS) error {

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (productApi *ProductApi_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (productApi *ProductApi_Spec) SetAzureName(azureName string) { productApi.AzureName = azureName }

type ProductApi_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ProductApi_STATUS{}

// ConvertStatusFrom populates our ProductApi_STATUS from the provided source
func (productApi *ProductApi_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.ProductApi_STATUS)
	if ok {
		// Populate our instance from source
		return productApi.AssignProperties_From_ProductApi_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.ProductApi_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = productApi.AssignProperties_From_ProductApi_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our ProductApi_STATUS
func (productApi *ProductApi_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.ProductApi_STATUS)
	if ok {
		// Populate destination from our instance
		return productApi.AssignProperties_To_ProductApi_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.ProductApi_STATUS{}
	err := productApi.AssignProperties_To_ProductApi_STATUS(dst)
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

var _ genruntime.FromARMConverter = &ProductApi_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (productApi *ProductApi_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ProductApi_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (productApi *ProductApi_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	_, ok := armInput.(ProductApi_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ProductApi_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// No error
	return nil
}

// AssignProperties_From_ProductApi_STATUS populates our ProductApi_STATUS from the provided source ProductApi_STATUS
func (productApi *ProductApi_STATUS) AssignProperties_From_ProductApi_STATUS(source *storage.ProductApi_STATUS) error {

	// Conditions
	productApi.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// No error
	return nil
}

// AssignProperties_To_ProductApi_STATUS populates the provided destination ProductApi_STATUS from our ProductApi_STATUS
func (productApi *ProductApi_STATUS) AssignProperties_To_ProductApi_STATUS(destination *storage.ProductApi_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(productApi.Conditions)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ProductApiOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_ProductApiOperatorSpec populates our ProductApiOperatorSpec from the provided source ProductApiOperatorSpec
func (operator *ProductApiOperatorSpec) AssignProperties_From_ProductApiOperatorSpec(source *storage.ProductApiOperatorSpec) error {

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// No error
	return nil
}

// AssignProperties_To_ProductApiOperatorSpec populates the provided destination ProductApiOperatorSpec from our ProductApiOperatorSpec
func (operator *ProductApiOperatorSpec) AssignProperties_To_ProductApiOperatorSpec(destination *storage.ProductApiOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
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

func init() {
	SchemeBuilder.Register(&ProductApi{}, &ProductApiList{})
}
