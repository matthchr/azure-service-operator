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

func Test_ServersIPV6FirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersIPV6FirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersIPV6FirewallRule, ServersIPV6FirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersIPV6FirewallRule runs a test to see if a specific instance of ServersIPV6FirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForServersIPV6FirewallRule(subject ServersIPV6FirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersIPV6FirewallRule
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

// Generator of ServersIPV6FirewallRule instances for property testing - lazily instantiated by
// ServersIPV6FirewallRuleGenerator()
var serversIPV6FirewallRuleGenerator gopter.Gen

// ServersIPV6FirewallRuleGenerator returns a generator of ServersIPV6FirewallRule instances for property testing.
func ServersIPV6FirewallRuleGenerator() gopter.Gen {
	if serversIPV6FirewallRuleGenerator != nil {
		return serversIPV6FirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersIPV6FirewallRule(generators)
	serversIPV6FirewallRuleGenerator = gen.Struct(reflect.TypeOf(ServersIPV6FirewallRule{}), generators)

	return serversIPV6FirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForServersIPV6FirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersIPV6FirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersIPV6FirewallRule_SpecGenerator()
	gens["Status"] = ServersIPV6FirewallRule_STATUSGenerator()
}

func Test_ServersIPV6FirewallRuleOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersIPV6FirewallRuleOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersIPV6FirewallRuleOperatorSpec, ServersIPV6FirewallRuleOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersIPV6FirewallRuleOperatorSpec runs a test to see if a specific instance of ServersIPV6FirewallRuleOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersIPV6FirewallRuleOperatorSpec(subject ServersIPV6FirewallRuleOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersIPV6FirewallRuleOperatorSpec
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

// Generator of ServersIPV6FirewallRuleOperatorSpec instances for property testing - lazily instantiated by
// ServersIPV6FirewallRuleOperatorSpecGenerator()
var serversIPV6FirewallRuleOperatorSpecGenerator gopter.Gen

// ServersIPV6FirewallRuleOperatorSpecGenerator returns a generator of ServersIPV6FirewallRuleOperatorSpec instances for property testing.
func ServersIPV6FirewallRuleOperatorSpecGenerator() gopter.Gen {
	if serversIPV6FirewallRuleOperatorSpecGenerator != nil {
		return serversIPV6FirewallRuleOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	serversIPV6FirewallRuleOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ServersIPV6FirewallRuleOperatorSpec{}), generators)

	return serversIPV6FirewallRuleOperatorSpecGenerator
}

func Test_ServersIPV6FirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersIPV6FirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersIPV6FirewallRule_STATUS, ServersIPV6FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersIPV6FirewallRule_STATUS runs a test to see if a specific instance of ServersIPV6FirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersIPV6FirewallRule_STATUS(subject ServersIPV6FirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersIPV6FirewallRule_STATUS
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

// Generator of ServersIPV6FirewallRule_STATUS instances for property testing - lazily instantiated by
// ServersIPV6FirewallRule_STATUSGenerator()
var serversIPV6FirewallRule_STATUSGenerator gopter.Gen

// ServersIPV6FirewallRule_STATUSGenerator returns a generator of ServersIPV6FirewallRule_STATUS instances for property testing.
func ServersIPV6FirewallRule_STATUSGenerator() gopter.Gen {
	if serversIPV6FirewallRule_STATUSGenerator != nil {
		return serversIPV6FirewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_STATUS(generators)
	serversIPV6FirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersIPV6FirewallRule_STATUS{}), generators)

	return serversIPV6FirewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["EndIPv6Address"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIPv6Address"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersIPV6FirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersIPV6FirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersIPV6FirewallRule_Spec, ServersIPV6FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersIPV6FirewallRule_Spec runs a test to see if a specific instance of ServersIPV6FirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersIPV6FirewallRule_Spec(subject ServersIPV6FirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersIPV6FirewallRule_Spec
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

// Generator of ServersIPV6FirewallRule_Spec instances for property testing - lazily instantiated by
// ServersIPV6FirewallRule_SpecGenerator()
var serversIPV6FirewallRule_SpecGenerator gopter.Gen

// ServersIPV6FirewallRule_SpecGenerator returns a generator of ServersIPV6FirewallRule_Spec instances for property testing.
// We first initialize serversIPV6FirewallRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersIPV6FirewallRule_SpecGenerator() gopter.Gen {
	if serversIPV6FirewallRule_SpecGenerator != nil {
		return serversIPV6FirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_Spec(generators)
	serversIPV6FirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(ServersIPV6FirewallRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_Spec(generators)
	AddRelatedPropertyGeneratorsForServersIPV6FirewallRule_Spec(generators)
	serversIPV6FirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(ServersIPV6FirewallRule_Spec{}), generators)

	return serversIPV6FirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIPv6Address"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["StartIPv6Address"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersIPV6FirewallRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersIPV6FirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ServersIPV6FirewallRuleOperatorSpecGenerator())
}
