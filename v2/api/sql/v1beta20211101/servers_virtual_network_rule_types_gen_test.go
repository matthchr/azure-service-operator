// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101storage"
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

func Test_Servers_VirtualNetworkRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_VirtualNetworkRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForServers_VirtualNetworkRule, Servers_VirtualNetworkRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServers_VirtualNetworkRule tests if a specific instance of Servers_VirtualNetworkRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServers_VirtualNetworkRule(subject Servers_VirtualNetworkRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.Servers_VirtualNetworkRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Servers_VirtualNetworkRule
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

func Test_Servers_VirtualNetworkRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_VirtualNetworkRule to Servers_VirtualNetworkRule via AssignProperties_To_Servers_VirtualNetworkRule & AssignProperties_From_Servers_VirtualNetworkRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_VirtualNetworkRule, Servers_VirtualNetworkRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_VirtualNetworkRule tests if a specific instance of Servers_VirtualNetworkRule can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_VirtualNetworkRule(subject Servers_VirtualNetworkRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_VirtualNetworkRule
	err := copied.AssignProperties_To_Servers_VirtualNetworkRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_VirtualNetworkRule
	err = actual.AssignProperties_From_Servers_VirtualNetworkRule(&other)
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

func Test_Servers_VirtualNetworkRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_VirtualNetworkRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_VirtualNetworkRule, Servers_VirtualNetworkRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_VirtualNetworkRule runs a test to see if a specific instance of Servers_VirtualNetworkRule round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_VirtualNetworkRule(subject Servers_VirtualNetworkRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_VirtualNetworkRule
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

// Generator of Servers_VirtualNetworkRule instances for property testing - lazily instantiated by
// Servers_VirtualNetworkRuleGenerator()
var servers_VirtualNetworkRuleGenerator gopter.Gen

// Servers_VirtualNetworkRuleGenerator returns a generator of Servers_VirtualNetworkRule instances for property testing.
func Servers_VirtualNetworkRuleGenerator() gopter.Gen {
	if servers_VirtualNetworkRuleGenerator != nil {
		return servers_VirtualNetworkRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServers_VirtualNetworkRule(generators)
	servers_VirtualNetworkRuleGenerator = gen.Struct(reflect.TypeOf(Servers_VirtualNetworkRule{}), generators)

	return servers_VirtualNetworkRuleGenerator
}

// AddRelatedPropertyGeneratorsForServers_VirtualNetworkRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_VirtualNetworkRule(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_VirtualNetworkRule_SpecGenerator()
	gens["Status"] = Servers_VirtualNetworkRule_STATUSGenerator()
}

func Test_Servers_VirtualNetworkRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_VirtualNetworkRule_Spec to Servers_VirtualNetworkRule_Spec via AssignProperties_To_Servers_VirtualNetworkRule_Spec & AssignProperties_From_Servers_VirtualNetworkRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_VirtualNetworkRule_Spec, Servers_VirtualNetworkRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_VirtualNetworkRule_Spec tests if a specific instance of Servers_VirtualNetworkRule_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_VirtualNetworkRule_Spec(subject Servers_VirtualNetworkRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_VirtualNetworkRule_Spec
	err := copied.AssignProperties_To_Servers_VirtualNetworkRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_VirtualNetworkRule_Spec
	err = actual.AssignProperties_From_Servers_VirtualNetworkRule_Spec(&other)
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

func Test_Servers_VirtualNetworkRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_VirtualNetworkRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_VirtualNetworkRule_Spec, Servers_VirtualNetworkRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_VirtualNetworkRule_Spec runs a test to see if a specific instance of Servers_VirtualNetworkRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_VirtualNetworkRule_Spec(subject Servers_VirtualNetworkRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_VirtualNetworkRule_Spec
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

// Generator of Servers_VirtualNetworkRule_Spec instances for property testing - lazily instantiated by
// Servers_VirtualNetworkRule_SpecGenerator()
var servers_VirtualNetworkRule_SpecGenerator gopter.Gen

// Servers_VirtualNetworkRule_SpecGenerator returns a generator of Servers_VirtualNetworkRule_Spec instances for property testing.
func Servers_VirtualNetworkRule_SpecGenerator() gopter.Gen {
	if servers_VirtualNetworkRule_SpecGenerator != nil {
		return servers_VirtualNetworkRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_Spec(generators)
	servers_VirtualNetworkRule_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_VirtualNetworkRule_Spec{}), generators)

	return servers_VirtualNetworkRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
}

func Test_Servers_VirtualNetworkRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_VirtualNetworkRule_STATUS to Servers_VirtualNetworkRule_STATUS via AssignProperties_To_Servers_VirtualNetworkRule_STATUS & AssignProperties_From_Servers_VirtualNetworkRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_VirtualNetworkRule_STATUS, Servers_VirtualNetworkRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_VirtualNetworkRule_STATUS tests if a specific instance of Servers_VirtualNetworkRule_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_VirtualNetworkRule_STATUS(subject Servers_VirtualNetworkRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_VirtualNetworkRule_STATUS
	err := copied.AssignProperties_To_Servers_VirtualNetworkRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_VirtualNetworkRule_STATUS
	err = actual.AssignProperties_From_Servers_VirtualNetworkRule_STATUS(&other)
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

func Test_Servers_VirtualNetworkRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_VirtualNetworkRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_VirtualNetworkRule_STATUS, Servers_VirtualNetworkRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_VirtualNetworkRule_STATUS runs a test to see if a specific instance of Servers_VirtualNetworkRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_VirtualNetworkRule_STATUS(subject Servers_VirtualNetworkRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_VirtualNetworkRule_STATUS
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

// Generator of Servers_VirtualNetworkRule_STATUS instances for property testing - lazily instantiated by
// Servers_VirtualNetworkRule_STATUSGenerator()
var servers_VirtualNetworkRule_STATUSGenerator gopter.Gen

// Servers_VirtualNetworkRule_STATUSGenerator returns a generator of Servers_VirtualNetworkRule_STATUS instances for property testing.
func Servers_VirtualNetworkRule_STATUSGenerator() gopter.Gen {
	if servers_VirtualNetworkRule_STATUSGenerator != nil {
		return servers_VirtualNetworkRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_STATUS(generators)
	servers_VirtualNetworkRule_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_VirtualNetworkRule_STATUS{}), generators)

	return servers_VirtualNetworkRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		VirtualNetworkRuleProperties_State_STATUS_Deleting,
		VirtualNetworkRuleProperties_State_STATUS_Failed,
		VirtualNetworkRuleProperties_State_STATUS_InProgress,
		VirtualNetworkRuleProperties_State_STATUS_Initializing,
		VirtualNetworkRuleProperties_State_STATUS_Ready,
		VirtualNetworkRuleProperties_State_STATUS_Unknown))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkSubnetId"] = gen.PtrOf(gen.AlphaString())
}
