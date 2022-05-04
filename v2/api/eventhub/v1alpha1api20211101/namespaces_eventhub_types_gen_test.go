// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"encoding/json"
	alpha20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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

func Test_NamespacesEventhub_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhub to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesEventhub tests if a specific instance of NamespacesEventhub round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.NamespacesEventhub
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesEventhub
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

func Test_NamespacesEventhub_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhub to NamespacesEventhub via AssignPropertiesToNamespacesEventhub & AssignPropertiesFromNamespacesEventhub returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhub tests if a specific instance of NamespacesEventhub can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.NamespacesEventhub
	err := copied.AssignPropertiesToNamespacesEventhub(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhub
	err = actual.AssignPropertiesFromNamespacesEventhub(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhub_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhub via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhub runs a test to see if a specific instance of NamespacesEventhub round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhub
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

// Generator of NamespacesEventhub instances for property testing - lazily instantiated by NamespacesEventhubGenerator()
var namespacesEventhubGenerator gopter.Gen

// NamespacesEventhubGenerator returns a generator of NamespacesEventhub instances for property testing.
func NamespacesEventhubGenerator() gopter.Gen {
	if namespacesEventhubGenerator != nil {
		return namespacesEventhubGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhub(generators)
	namespacesEventhubGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhub{}), generators)

	return namespacesEventhubGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhub is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhub(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesEventhubsSpecGenerator()
	gens["Status"] = EventhubStatusGenerator()
}

func Test_Eventhub_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Eventhub_Status to Eventhub_Status via AssignPropertiesToEventhubStatus & AssignPropertiesFromEventhubStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForEventhubStatus, EventhubStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForEventhubStatus tests if a specific instance of Eventhub_Status can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForEventhubStatus(subject Eventhub_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.Eventhub_Status
	err := copied.AssignPropertiesToEventhubStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Eventhub_Status
	err = actual.AssignPropertiesFromEventhubStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Eventhub_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Eventhub_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventhubStatus, EventhubStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventhubStatus runs a test to see if a specific instance of Eventhub_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForEventhubStatus(subject Eventhub_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Eventhub_Status
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

// Generator of Eventhub_Status instances for property testing - lazily instantiated by EventhubStatusGenerator()
var eventhubStatusGenerator gopter.Gen

// EventhubStatusGenerator returns a generator of Eventhub_Status instances for property testing.
// We first initialize eventhubStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventhubStatusGenerator() gopter.Gen {
	if eventhubStatusGenerator != nil {
		return eventhubStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhubStatus(generators)
	eventhubStatusGenerator = gen.Struct(reflect.TypeOf(Eventhub_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhubStatus(generators)
	AddRelatedPropertyGeneratorsForEventhubStatus(generators)
	eventhubStatusGenerator = gen.Struct(reflect.TypeOf(Eventhub_Status{}), generators)

	return eventhubStatusGenerator
}

// AddIndependentPropertyGeneratorsForEventhubStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventhubStatus(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PartitionIds"] = gen.SliceOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EventhubStatusPropertiesStatusActive,
		EventhubStatusPropertiesStatusCreating,
		EventhubStatusPropertiesStatusDeleting,
		EventhubStatusPropertiesStatusDisabled,
		EventhubStatusPropertiesStatusReceiveDisabled,
		EventhubStatusPropertiesStatusRenaming,
		EventhubStatusPropertiesStatusRestoring,
		EventhubStatusPropertiesStatusSendDisabled,
		EventhubStatusPropertiesStatusUnknown))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventhubStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventhubStatus(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescriptionStatusGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_NamespacesEventhubs_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubs_Spec to NamespacesEventhubs_Spec via AssignPropertiesToNamespacesEventhubsSpec & AssignPropertiesFromNamespacesEventhubsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsSpec, NamespacesEventhubsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsSpec tests if a specific instance of NamespacesEventhubs_Spec can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsSpec(subject NamespacesEventhubs_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.NamespacesEventhubs_Spec
	err := copied.AssignPropertiesToNamespacesEventhubsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubs_Spec
	err = actual.AssignPropertiesFromNamespacesEventhubsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubs_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubs_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsSpec, NamespacesEventhubsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsSpec runs a test to see if a specific instance of NamespacesEventhubs_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsSpec(subject NamespacesEventhubs_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubs_Spec
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

// Generator of NamespacesEventhubs_Spec instances for property testing - lazily instantiated by
// NamespacesEventhubsSpecGenerator()
var namespacesEventhubsSpecGenerator gopter.Gen

// NamespacesEventhubsSpecGenerator returns a generator of NamespacesEventhubs_Spec instances for property testing.
// We first initialize namespacesEventhubsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsSpecGenerator() gopter.Gen {
	if namespacesEventhubsSpecGenerator != nil {
		return namespacesEventhubsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpec(generators)
	namespacesEventhubsSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpec(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsSpec(generators)
	namespacesEventhubsSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec{}), generators)

	return namespacesEventhubsSpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsSpec(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(NamespacesEventhubsSpecPropertiesCaptureDescriptionGenerator())
}

func Test_CaptureDescription_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from CaptureDescription_Status to CaptureDescription_Status via AssignPropertiesToCaptureDescriptionStatus & AssignPropertiesFromCaptureDescriptionStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForCaptureDescriptionStatus, CaptureDescriptionStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForCaptureDescriptionStatus tests if a specific instance of CaptureDescription_Status can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForCaptureDescriptionStatus(subject CaptureDescription_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.CaptureDescription_Status
	err := copied.AssignPropertiesToCaptureDescriptionStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual CaptureDescription_Status
	err = actual.AssignPropertiesFromCaptureDescriptionStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_CaptureDescription_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescriptionStatus, CaptureDescriptionStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescriptionStatus runs a test to see if a specific instance of CaptureDescription_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescriptionStatus(subject CaptureDescription_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription_Status
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

// Generator of CaptureDescription_Status instances for property testing - lazily instantiated by
// CaptureDescriptionStatusGenerator()
var captureDescriptionStatusGenerator gopter.Gen

// CaptureDescriptionStatusGenerator returns a generator of CaptureDescription_Status instances for property testing.
// We first initialize captureDescriptionStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescriptionStatusGenerator() gopter.Gen {
	if captureDescriptionStatusGenerator != nil {
		return captureDescriptionStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescriptionStatus(generators)
	captureDescriptionStatusGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescriptionStatus(generators)
	AddRelatedPropertyGeneratorsForCaptureDescriptionStatus(generators)
	captureDescriptionStatusGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_Status{}), generators)

	return captureDescriptionStatusGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescriptionStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescriptionStatus(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(CaptureDescriptionStatusEncodingAvro, CaptureDescriptionStatusEncodingAvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescriptionStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescriptionStatus(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(DestinationStatusGenerator())
}

func Test_NamespacesEventhubs_Spec_Properties_CaptureDescription_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubs_Spec_Properties_CaptureDescription to NamespacesEventhubs_Spec_Properties_CaptureDescription via AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescription & AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescription returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsSpecPropertiesCaptureDescription, NamespacesEventhubsSpecPropertiesCaptureDescriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsSpecPropertiesCaptureDescription tests if a specific instance of NamespacesEventhubs_Spec_Properties_CaptureDescription can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsSpecPropertiesCaptureDescription(subject NamespacesEventhubs_Spec_Properties_CaptureDescription) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.NamespacesEventhubs_Spec_Properties_CaptureDescription
	err := copied.AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescription(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubs_Spec_Properties_CaptureDescription
	err = actual.AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescription(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubs_Spec_Properties_CaptureDescription_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubs_Spec_Properties_CaptureDescription via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescription, NamespacesEventhubsSpecPropertiesCaptureDescriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescription runs a test to see if a specific instance of NamespacesEventhubs_Spec_Properties_CaptureDescription round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescription(subject NamespacesEventhubs_Spec_Properties_CaptureDescription) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubs_Spec_Properties_CaptureDescription
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

// Generator of NamespacesEventhubs_Spec_Properties_CaptureDescription instances for property testing - lazily
// instantiated by NamespacesEventhubsSpecPropertiesCaptureDescriptionGenerator()
var namespacesEventhubsSpecPropertiesCaptureDescriptionGenerator gopter.Gen

// NamespacesEventhubsSpecPropertiesCaptureDescriptionGenerator returns a generator of NamespacesEventhubs_Spec_Properties_CaptureDescription instances for property testing.
// We first initialize namespacesEventhubsSpecPropertiesCaptureDescriptionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsSpecPropertiesCaptureDescriptionGenerator() gopter.Gen {
	if namespacesEventhubsSpecPropertiesCaptureDescriptionGenerator != nil {
		return namespacesEventhubsSpecPropertiesCaptureDescriptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescription(generators)
	namespacesEventhubsSpecPropertiesCaptureDescriptionGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_Properties_CaptureDescription{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescription(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescription(generators)
	namespacesEventhubsSpecPropertiesCaptureDescriptionGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_Properties_CaptureDescription{}), generators)

	return namespacesEventhubsSpecPropertiesCaptureDescriptionGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescription is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescription(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(NamespacesEventhubsSpecPropertiesCaptureDescriptionEncodingAvro, NamespacesEventhubsSpecPropertiesCaptureDescriptionEncodingAvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescription is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescription(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator())
}

func Test_Destination_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Destination_Status to Destination_Status via AssignPropertiesToDestinationStatus & AssignPropertiesFromDestinationStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForDestinationStatus, DestinationStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDestinationStatus tests if a specific instance of Destination_Status can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForDestinationStatus(subject Destination_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.Destination_Status
	err := copied.AssignPropertiesToDestinationStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Destination_Status
	err = actual.AssignPropertiesFromDestinationStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Destination_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestinationStatus, DestinationStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestinationStatus runs a test to see if a specific instance of Destination_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDestinationStatus(subject Destination_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_Status
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

// Generator of Destination_Status instances for property testing - lazily instantiated by DestinationStatusGenerator()
var destinationStatusGenerator gopter.Gen

// DestinationStatusGenerator returns a generator of Destination_Status instances for property testing.
func DestinationStatusGenerator() gopter.Gen {
	if destinationStatusGenerator != nil {
		return destinationStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestinationStatus(generators)
	destinationStatusGenerator = gen.Struct(reflect.TypeOf(Destination_Status{}), generators)

	return destinationStatusGenerator
}

// AddIndependentPropertyGeneratorsForDestinationStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestinationStatus(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination to NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination via AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination & AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination, NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination tests if a specific instance of NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(subject NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
	err := copied.AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
	err = actual.AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination, NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination runs a test to see if a specific instance of NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(subject NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
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

// Generator of NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination instances for property testing -
// lazily instantiated by NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator()
var namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator gopter.Gen

// NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator returns a generator of NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination instances for property testing.
func NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator() gopter.Gen {
	if namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator != nil {
		return namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(generators)
	namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination{}), generators)

	return namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}
