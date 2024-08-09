// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_AfdOrigin_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdOrigin to hub returns original",
		prop.ForAll(RunResourceConversionTestForAfdOrigin, AfdOriginGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForAfdOrigin tests if a specific instance of AfdOrigin round trips to the hub storage version and back losslessly
func RunResourceConversionTestForAfdOrigin(subject AfdOrigin) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.AfdOrigin
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual AfdOrigin
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AfdOrigin_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdOrigin to AfdOrigin via AssignProperties_To_AfdOrigin & AssignProperties_From_AfdOrigin returns original",
		prop.ForAll(RunPropertyAssignmentTestForAfdOrigin, AfdOriginGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAfdOrigin tests if a specific instance of AfdOrigin can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAfdOrigin(subject AfdOrigin) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AfdOrigin
	err := copied.AssignProperties_To_AfdOrigin(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AfdOrigin
	err = actual.AssignProperties_From_AfdOrigin(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AfdOrigin_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdOrigin via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdOrigin, AfdOriginGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdOrigin runs a test to see if a specific instance of AfdOrigin round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdOrigin(subject AfdOrigin) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdOrigin
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AfdOrigin instances for property testing - lazily instantiated by AfdOriginGenerator()
var afdOriginGenerator gopter.Gen

// AfdOriginGenerator returns a generator of AfdOrigin instances for property testing.
func AfdOriginGenerator() gopter.Gen {
	if afdOriginGenerator != nil {
		return afdOriginGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAfdOrigin(generators)
	afdOriginGenerator = gen.Struct(reflect.TypeOf(AfdOrigin{}), generators)

	return afdOriginGenerator
}

// AddRelatedPropertyGeneratorsForAfdOrigin is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdOrigin(gens map[string]gopter.Gen) {
	gens["Spec"] = AfdOrigin_SpecGenerator()
	gens["Status"] = AfdOrigin_STATUSGenerator()
}

func Test_AfdOriginOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdOriginOperatorSpec to AfdOriginOperatorSpec via AssignProperties_To_AfdOriginOperatorSpec & AssignProperties_From_AfdOriginOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForAfdOriginOperatorSpec, AfdOriginOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAfdOriginOperatorSpec tests if a specific instance of AfdOriginOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAfdOriginOperatorSpec(subject AfdOriginOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AfdOriginOperatorSpec
	err := copied.AssignProperties_To_AfdOriginOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AfdOriginOperatorSpec
	err = actual.AssignProperties_From_AfdOriginOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AfdOriginOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdOriginOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdOriginOperatorSpec, AfdOriginOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdOriginOperatorSpec runs a test to see if a specific instance of AfdOriginOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdOriginOperatorSpec(subject AfdOriginOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdOriginOperatorSpec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AfdOriginOperatorSpec instances for property testing - lazily instantiated by
// AfdOriginOperatorSpecGenerator()
var afdOriginOperatorSpecGenerator gopter.Gen

// AfdOriginOperatorSpecGenerator returns a generator of AfdOriginOperatorSpec instances for property testing.
func AfdOriginOperatorSpecGenerator() gopter.Gen {
	if afdOriginOperatorSpecGenerator != nil {
		return afdOriginOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	afdOriginOperatorSpecGenerator = gen.Struct(reflect.TypeOf(AfdOriginOperatorSpec{}), generators)

	return afdOriginOperatorSpecGenerator
}

func Test_AfdOrigin_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdOrigin_STATUS to AfdOrigin_STATUS via AssignProperties_To_AfdOrigin_STATUS & AssignProperties_From_AfdOrigin_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForAfdOrigin_STATUS, AfdOrigin_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAfdOrigin_STATUS tests if a specific instance of AfdOrigin_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAfdOrigin_STATUS(subject AfdOrigin_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AfdOrigin_STATUS
	err := copied.AssignProperties_To_AfdOrigin_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AfdOrigin_STATUS
	err = actual.AssignProperties_From_AfdOrigin_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AfdOrigin_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdOrigin_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdOrigin_STATUS, AfdOrigin_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdOrigin_STATUS runs a test to see if a specific instance of AfdOrigin_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdOrigin_STATUS(subject AfdOrigin_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdOrigin_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AfdOrigin_STATUS instances for property testing - lazily instantiated by AfdOrigin_STATUSGenerator()
var afdOrigin_STATUSGenerator gopter.Gen

// AfdOrigin_STATUSGenerator returns a generator of AfdOrigin_STATUS instances for property testing.
// We first initialize afdOrigin_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AfdOrigin_STATUSGenerator() gopter.Gen {
	if afdOrigin_STATUSGenerator != nil {
		return afdOrigin_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdOrigin_STATUS(generators)
	afdOrigin_STATUSGenerator = gen.Struct(reflect.TypeOf(AfdOrigin_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdOrigin_STATUS(generators)
	AddRelatedPropertyGeneratorsForAfdOrigin_STATUS(generators)
	afdOrigin_STATUSGenerator = gen.Struct(reflect.TypeOf(AfdOrigin_STATUS{}), generators)

	return afdOrigin_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAfdOrigin_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdOrigin_STATUS(gens map[string]gopter.Gen) {
	gens["DeploymentStatus"] = gen.PtrOf(gen.OneConstOf(
		AFDOriginProperties_DeploymentStatus_STATUS_Failed,
		AFDOriginProperties_DeploymentStatus_STATUS_InProgress,
		AFDOriginProperties_DeploymentStatus_STATUS_NotStarted,
		AFDOriginProperties_DeploymentStatus_STATUS_Succeeded))
	gens["EnabledState"] = gen.PtrOf(gen.OneConstOf(AFDOriginProperties_EnabledState_STATUS_Disabled, AFDOriginProperties_EnabledState_STATUS_Enabled))
	gens["EnforceCertificateNameCheck"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["HttpPort"] = gen.PtrOf(gen.Int())
	gens["HttpsPort"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["OriginGroupName"] = gen.PtrOf(gen.AlphaString())
	gens["OriginHostHeader"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		AFDOriginProperties_ProvisioningState_STATUS_Creating,
		AFDOriginProperties_ProvisioningState_STATUS_Deleting,
		AFDOriginProperties_ProvisioningState_STATUS_Failed,
		AFDOriginProperties_ProvisioningState_STATUS_Succeeded,
		AFDOriginProperties_ProvisioningState_STATUS_Updating))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAfdOrigin_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdOrigin_STATUS(gens map[string]gopter.Gen) {
	gens["AzureOrigin"] = gen.PtrOf(ResourceReference_STATUSGenerator())
	gens["SharedPrivateLinkResource"] = gen.PtrOf(SharedPrivateLinkResourceProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_AfdOrigin_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdOrigin_Spec to AfdOrigin_Spec via AssignProperties_To_AfdOrigin_Spec & AssignProperties_From_AfdOrigin_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForAfdOrigin_Spec, AfdOrigin_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAfdOrigin_Spec tests if a specific instance of AfdOrigin_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAfdOrigin_Spec(subject AfdOrigin_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AfdOrigin_Spec
	err := copied.AssignProperties_To_AfdOrigin_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AfdOrigin_Spec
	err = actual.AssignProperties_From_AfdOrigin_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AfdOrigin_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdOrigin_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdOrigin_Spec, AfdOrigin_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdOrigin_Spec runs a test to see if a specific instance of AfdOrigin_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdOrigin_Spec(subject AfdOrigin_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdOrigin_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AfdOrigin_Spec instances for property testing - lazily instantiated by AfdOrigin_SpecGenerator()
var afdOrigin_SpecGenerator gopter.Gen

// AfdOrigin_SpecGenerator returns a generator of AfdOrigin_Spec instances for property testing.
// We first initialize afdOrigin_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AfdOrigin_SpecGenerator() gopter.Gen {
	if afdOrigin_SpecGenerator != nil {
		return afdOrigin_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdOrigin_Spec(generators)
	afdOrigin_SpecGenerator = gen.Struct(reflect.TypeOf(AfdOrigin_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdOrigin_Spec(generators)
	AddRelatedPropertyGeneratorsForAfdOrigin_Spec(generators)
	afdOrigin_SpecGenerator = gen.Struct(reflect.TypeOf(AfdOrigin_Spec{}), generators)

	return afdOrigin_SpecGenerator
}

// AddIndependentPropertyGeneratorsForAfdOrigin_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdOrigin_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EnabledState"] = gen.PtrOf(gen.OneConstOf(AFDOriginProperties_EnabledState_Disabled, AFDOriginProperties_EnabledState_Enabled))
	gens["EnforceCertificateNameCheck"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["HttpPort"] = gen.PtrOf(gen.Int())
	gens["HttpsPort"] = gen.PtrOf(gen.Int())
	gens["OriginHostHeader"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAfdOrigin_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdOrigin_Spec(gens map[string]gopter.Gen) {
	gens["AzureOrigin"] = gen.PtrOf(ResourceReferenceGenerator())
	gens["OperatorSpec"] = gen.PtrOf(AfdOriginOperatorSpecGenerator())
	gens["SharedPrivateLinkResource"] = gen.PtrOf(SharedPrivateLinkResourcePropertiesGenerator())
}

func Test_SharedPrivateLinkResourceProperties_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SharedPrivateLinkResourceProperties to SharedPrivateLinkResourceProperties via AssignProperties_To_SharedPrivateLinkResourceProperties & AssignProperties_From_SharedPrivateLinkResourceProperties returns original",
		prop.ForAll(RunPropertyAssignmentTestForSharedPrivateLinkResourceProperties, SharedPrivateLinkResourcePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSharedPrivateLinkResourceProperties tests if a specific instance of SharedPrivateLinkResourceProperties can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSharedPrivateLinkResourceProperties(subject SharedPrivateLinkResourceProperties) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SharedPrivateLinkResourceProperties
	err := copied.AssignProperties_To_SharedPrivateLinkResourceProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SharedPrivateLinkResourceProperties
	err = actual.AssignProperties_From_SharedPrivateLinkResourceProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SharedPrivateLinkResourceProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourceProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourceProperties, SharedPrivateLinkResourcePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourceProperties runs a test to see if a specific instance of SharedPrivateLinkResourceProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourceProperties(subject SharedPrivateLinkResourceProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourceProperties
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SharedPrivateLinkResourceProperties instances for property testing - lazily instantiated by
// SharedPrivateLinkResourcePropertiesGenerator()
var sharedPrivateLinkResourcePropertiesGenerator gopter.Gen

// SharedPrivateLinkResourcePropertiesGenerator returns a generator of SharedPrivateLinkResourceProperties instances for property testing.
// We first initialize sharedPrivateLinkResourcePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SharedPrivateLinkResourcePropertiesGenerator() gopter.Gen {
	if sharedPrivateLinkResourcePropertiesGenerator != nil {
		return sharedPrivateLinkResourcePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties(generators)
	sharedPrivateLinkResourcePropertiesGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties(generators)
	AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties(generators)
	sharedPrivateLinkResourcePropertiesGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperties{}), generators)

	return sharedPrivateLinkResourcePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkLocation"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		SharedPrivateLinkResourceProperties_Status_Approved,
		SharedPrivateLinkResourceProperties_Status_Disconnected,
		SharedPrivateLinkResourceProperties_Status_Pending,
		SharedPrivateLinkResourceProperties_Status_Rejected,
		SharedPrivateLinkResourceProperties_Status_Timeout))
}

// AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties(gens map[string]gopter.Gen) {
	gens["PrivateLink"] = gen.PtrOf(ResourceReferenceGenerator())
}

func Test_SharedPrivateLinkResourceProperties_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SharedPrivateLinkResourceProperties_STATUS to SharedPrivateLinkResourceProperties_STATUS via AssignProperties_To_SharedPrivateLinkResourceProperties_STATUS & AssignProperties_From_SharedPrivateLinkResourceProperties_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSharedPrivateLinkResourceProperties_STATUS, SharedPrivateLinkResourceProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSharedPrivateLinkResourceProperties_STATUS tests if a specific instance of SharedPrivateLinkResourceProperties_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSharedPrivateLinkResourceProperties_STATUS(subject SharedPrivateLinkResourceProperties_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SharedPrivateLinkResourceProperties_STATUS
	err := copied.AssignProperties_To_SharedPrivateLinkResourceProperties_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SharedPrivateLinkResourceProperties_STATUS
	err = actual.AssignProperties_From_SharedPrivateLinkResourceProperties_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SharedPrivateLinkResourceProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourceProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourceProperties_STATUS, SharedPrivateLinkResourceProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourceProperties_STATUS runs a test to see if a specific instance of SharedPrivateLinkResourceProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourceProperties_STATUS(subject SharedPrivateLinkResourceProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourceProperties_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SharedPrivateLinkResourceProperties_STATUS instances for property testing - lazily instantiated by
// SharedPrivateLinkResourceProperties_STATUSGenerator()
var sharedPrivateLinkResourceProperties_STATUSGenerator gopter.Gen

// SharedPrivateLinkResourceProperties_STATUSGenerator returns a generator of SharedPrivateLinkResourceProperties_STATUS instances for property testing.
// We first initialize sharedPrivateLinkResourceProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SharedPrivateLinkResourceProperties_STATUSGenerator() gopter.Gen {
	if sharedPrivateLinkResourceProperties_STATUSGenerator != nil {
		return sharedPrivateLinkResourceProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties_STATUS(generators)
	sharedPrivateLinkResourceProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties_STATUS(generators)
	sharedPrivateLinkResourceProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperties_STATUS{}), generators)

	return sharedPrivateLinkResourceProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties_STATUS(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkLocation"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		SharedPrivateLinkResourceProperties_Status_STATUS_Approved,
		SharedPrivateLinkResourceProperties_Status_STATUS_Disconnected,
		SharedPrivateLinkResourceProperties_Status_STATUS_Pending,
		SharedPrivateLinkResourceProperties_Status_STATUS_Rejected,
		SharedPrivateLinkResourceProperties_Status_STATUS_Timeout))
}

// AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties_STATUS(gens map[string]gopter.Gen) {
	gens["PrivateLink"] = gen.PtrOf(ResourceReference_STATUSGenerator())
}
