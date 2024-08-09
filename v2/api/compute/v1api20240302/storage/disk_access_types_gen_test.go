// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
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

func Test_DiskAccess_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskAccess via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskAccess, DiskAccessGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskAccess runs a test to see if a specific instance of DiskAccess round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskAccess(subject DiskAccess) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskAccess
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

// Generator of DiskAccess instances for property testing - lazily instantiated by DiskAccessGenerator()
var diskAccessGenerator gopter.Gen

// DiskAccessGenerator returns a generator of DiskAccess instances for property testing.
func DiskAccessGenerator() gopter.Gen {
	if diskAccessGenerator != nil {
		return diskAccessGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDiskAccess(generators)
	diskAccessGenerator = gen.Struct(reflect.TypeOf(DiskAccess{}), generators)

	return diskAccessGenerator
}

// AddRelatedPropertyGeneratorsForDiskAccess is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskAccess(gens map[string]gopter.Gen) {
	gens["Spec"] = DiskAccess_SpecGenerator()
	gens["Status"] = DiskAccess_STATUSGenerator()
}

func Test_DiskAccessOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskAccessOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskAccessOperatorSpec, DiskAccessOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskAccessOperatorSpec runs a test to see if a specific instance of DiskAccessOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskAccessOperatorSpec(subject DiskAccessOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskAccessOperatorSpec
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

// Generator of DiskAccessOperatorSpec instances for property testing - lazily instantiated by
// DiskAccessOperatorSpecGenerator()
var diskAccessOperatorSpecGenerator gopter.Gen

// DiskAccessOperatorSpecGenerator returns a generator of DiskAccessOperatorSpec instances for property testing.
func DiskAccessOperatorSpecGenerator() gopter.Gen {
	if diskAccessOperatorSpecGenerator != nil {
		return diskAccessOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	diskAccessOperatorSpecGenerator = gen.Struct(reflect.TypeOf(DiskAccessOperatorSpec{}), generators)

	return diskAccessOperatorSpecGenerator
}

func Test_DiskAccess_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskAccess_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskAccess_STATUS, DiskAccess_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskAccess_STATUS runs a test to see if a specific instance of DiskAccess_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskAccess_STATUS(subject DiskAccess_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskAccess_STATUS
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

// Generator of DiskAccess_STATUS instances for property testing - lazily instantiated by DiskAccess_STATUSGenerator()
var diskAccess_STATUSGenerator gopter.Gen

// DiskAccess_STATUSGenerator returns a generator of DiskAccess_STATUS instances for property testing.
// We first initialize diskAccess_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskAccess_STATUSGenerator() gopter.Gen {
	if diskAccess_STATUSGenerator != nil {
		return diskAccess_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_STATUS(generators)
	diskAccess_STATUSGenerator = gen.Struct(reflect.TypeOf(DiskAccess_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_STATUS(generators)
	AddRelatedPropertyGeneratorsForDiskAccess_STATUS(generators)
	diskAccess_STATUSGenerator = gen.Struct(reflect.TypeOf(DiskAccess_STATUS{}), generators)

	return diskAccess_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDiskAccess_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskAccess_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TimeCreated"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskAccess_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskAccess_STATUS(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUSGenerator())
}

func Test_DiskAccess_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskAccess_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskAccess_Spec, DiskAccess_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskAccess_Spec runs a test to see if a specific instance of DiskAccess_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskAccess_Spec(subject DiskAccess_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskAccess_Spec
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

// Generator of DiskAccess_Spec instances for property testing - lazily instantiated by DiskAccess_SpecGenerator()
var diskAccess_SpecGenerator gopter.Gen

// DiskAccess_SpecGenerator returns a generator of DiskAccess_Spec instances for property testing.
// We first initialize diskAccess_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskAccess_SpecGenerator() gopter.Gen {
	if diskAccess_SpecGenerator != nil {
		return diskAccess_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_Spec(generators)
	diskAccess_SpecGenerator = gen.Struct(reflect.TypeOf(DiskAccess_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_Spec(generators)
	AddRelatedPropertyGeneratorsForDiskAccess_Spec(generators)
	diskAccess_SpecGenerator = gen.Struct(reflect.TypeOf(DiskAccess_Spec{}), generators)

	return diskAccess_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDiskAccess_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskAccess_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskAccess_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskAccess_Spec(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["OperatorSpec"] = gen.PtrOf(DiskAccessOperatorSpecGenerator())
}

func Test_PrivateEndpointConnection_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS, PrivateEndpointConnection_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS runs a test to see if a specific instance of PrivateEndpointConnection_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS(subject PrivateEndpointConnection_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS
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

// Generator of PrivateEndpointConnection_STATUS instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUSGenerator()
var privateEndpointConnection_STATUSGenerator gopter.Gen

// PrivateEndpointConnection_STATUSGenerator returns a generator of PrivateEndpointConnection_STATUS instances for property testing.
func PrivateEndpointConnection_STATUSGenerator() gopter.Gen {
	if privateEndpointConnection_STATUSGenerator != nil {
		return privateEndpointConnection_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(generators)
	privateEndpointConnection_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS{}), generators)

	return privateEndpointConnection_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
