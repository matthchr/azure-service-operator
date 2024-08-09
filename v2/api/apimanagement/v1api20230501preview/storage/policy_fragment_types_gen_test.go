// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
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

func Test_PolicyFragment_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PolicyFragment to hub returns original",
		prop.ForAll(RunResourceConversionTestForPolicyFragment, PolicyFragmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPolicyFragment tests if a specific instance of PolicyFragment round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPolicyFragment(subject PolicyFragment) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.PolicyFragment
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PolicyFragment
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

func Test_PolicyFragment_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PolicyFragment to PolicyFragment via AssignProperties_To_PolicyFragment & AssignProperties_From_PolicyFragment returns original",
		prop.ForAll(RunPropertyAssignmentTestForPolicyFragment, PolicyFragmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPolicyFragment tests if a specific instance of PolicyFragment can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPolicyFragment(subject PolicyFragment) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PolicyFragment
	err := copied.AssignProperties_To_PolicyFragment(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PolicyFragment
	err = actual.AssignProperties_From_PolicyFragment(&other)
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

func Test_PolicyFragment_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PolicyFragment via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicyFragment, PolicyFragmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicyFragment runs a test to see if a specific instance of PolicyFragment round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicyFragment(subject PolicyFragment) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PolicyFragment
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

// Generator of PolicyFragment instances for property testing - lazily instantiated by PolicyFragmentGenerator()
var policyFragmentGenerator gopter.Gen

// PolicyFragmentGenerator returns a generator of PolicyFragment instances for property testing.
func PolicyFragmentGenerator() gopter.Gen {
	if policyFragmentGenerator != nil {
		return policyFragmentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPolicyFragment(generators)
	policyFragmentGenerator = gen.Struct(reflect.TypeOf(PolicyFragment{}), generators)

	return policyFragmentGenerator
}

// AddRelatedPropertyGeneratorsForPolicyFragment is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPolicyFragment(gens map[string]gopter.Gen) {
	gens["Spec"] = PolicyFragment_SpecGenerator()
	gens["Status"] = PolicyFragment_STATUSGenerator()
}

func Test_PolicyFragmentOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PolicyFragmentOperatorSpec to PolicyFragmentOperatorSpec via AssignProperties_To_PolicyFragmentOperatorSpec & AssignProperties_From_PolicyFragmentOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPolicyFragmentOperatorSpec, PolicyFragmentOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPolicyFragmentOperatorSpec tests if a specific instance of PolicyFragmentOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPolicyFragmentOperatorSpec(subject PolicyFragmentOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PolicyFragmentOperatorSpec
	err := copied.AssignProperties_To_PolicyFragmentOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PolicyFragmentOperatorSpec
	err = actual.AssignProperties_From_PolicyFragmentOperatorSpec(&other)
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

func Test_PolicyFragmentOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PolicyFragmentOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicyFragmentOperatorSpec, PolicyFragmentOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicyFragmentOperatorSpec runs a test to see if a specific instance of PolicyFragmentOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicyFragmentOperatorSpec(subject PolicyFragmentOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PolicyFragmentOperatorSpec
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

// Generator of PolicyFragmentOperatorSpec instances for property testing - lazily instantiated by
// PolicyFragmentOperatorSpecGenerator()
var policyFragmentOperatorSpecGenerator gopter.Gen

// PolicyFragmentOperatorSpecGenerator returns a generator of PolicyFragmentOperatorSpec instances for property testing.
func PolicyFragmentOperatorSpecGenerator() gopter.Gen {
	if policyFragmentOperatorSpecGenerator != nil {
		return policyFragmentOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	policyFragmentOperatorSpecGenerator = gen.Struct(reflect.TypeOf(PolicyFragmentOperatorSpec{}), generators)

	return policyFragmentOperatorSpecGenerator
}

func Test_PolicyFragment_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PolicyFragment_STATUS to PolicyFragment_STATUS via AssignProperties_To_PolicyFragment_STATUS & AssignProperties_From_PolicyFragment_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPolicyFragment_STATUS, PolicyFragment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPolicyFragment_STATUS tests if a specific instance of PolicyFragment_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPolicyFragment_STATUS(subject PolicyFragment_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PolicyFragment_STATUS
	err := copied.AssignProperties_To_PolicyFragment_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PolicyFragment_STATUS
	err = actual.AssignProperties_From_PolicyFragment_STATUS(&other)
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

func Test_PolicyFragment_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PolicyFragment_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicyFragment_STATUS, PolicyFragment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicyFragment_STATUS runs a test to see if a specific instance of PolicyFragment_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicyFragment_STATUS(subject PolicyFragment_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PolicyFragment_STATUS
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

// Generator of PolicyFragment_STATUS instances for property testing - lazily instantiated by
// PolicyFragment_STATUSGenerator()
var policyFragment_STATUSGenerator gopter.Gen

// PolicyFragment_STATUSGenerator returns a generator of PolicyFragment_STATUS instances for property testing.
func PolicyFragment_STATUSGenerator() gopter.Gen {
	if policyFragment_STATUSGenerator != nil {
		return policyFragment_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicyFragment_STATUS(generators)
	policyFragment_STATUSGenerator = gen.Struct(reflect.TypeOf(PolicyFragment_STATUS{}), generators)

	return policyFragment_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPolicyFragment_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPolicyFragment_STATUS(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Format"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_PolicyFragment_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PolicyFragment_Spec to PolicyFragment_Spec via AssignProperties_To_PolicyFragment_Spec & AssignProperties_From_PolicyFragment_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPolicyFragment_Spec, PolicyFragment_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPolicyFragment_Spec tests if a specific instance of PolicyFragment_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPolicyFragment_Spec(subject PolicyFragment_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PolicyFragment_Spec
	err := copied.AssignProperties_To_PolicyFragment_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PolicyFragment_Spec
	err = actual.AssignProperties_From_PolicyFragment_Spec(&other)
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

func Test_PolicyFragment_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PolicyFragment_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicyFragment_Spec, PolicyFragment_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicyFragment_Spec runs a test to see if a specific instance of PolicyFragment_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicyFragment_Spec(subject PolicyFragment_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PolicyFragment_Spec
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

// Generator of PolicyFragment_Spec instances for property testing - lazily instantiated by
// PolicyFragment_SpecGenerator()
var policyFragment_SpecGenerator gopter.Gen

// PolicyFragment_SpecGenerator returns a generator of PolicyFragment_Spec instances for property testing.
// We first initialize policyFragment_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PolicyFragment_SpecGenerator() gopter.Gen {
	if policyFragment_SpecGenerator != nil {
		return policyFragment_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicyFragment_Spec(generators)
	policyFragment_SpecGenerator = gen.Struct(reflect.TypeOf(PolicyFragment_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicyFragment_Spec(generators)
	AddRelatedPropertyGeneratorsForPolicyFragment_Spec(generators)
	policyFragment_SpecGenerator = gen.Struct(reflect.TypeOf(PolicyFragment_Spec{}), generators)

	return policyFragment_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPolicyFragment_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPolicyFragment_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Format"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPolicyFragment_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPolicyFragment_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(PolicyFragmentOperatorSpecGenerator())
}
