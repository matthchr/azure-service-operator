// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

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

func Test_Servers_DnsAliase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_DnsAliase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_DnsAliase, Servers_DnsAliaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_DnsAliase runs a test to see if a specific instance of Servers_DnsAliase round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_DnsAliase(subject Servers_DnsAliase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_DnsAliase
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

// Generator of Servers_DnsAliase instances for property testing - lazily instantiated by Servers_DnsAliaseGenerator()
var servers_DnsAliaseGenerator gopter.Gen

// Servers_DnsAliaseGenerator returns a generator of Servers_DnsAliase instances for property testing.
func Servers_DnsAliaseGenerator() gopter.Gen {
	if servers_DnsAliaseGenerator != nil {
		return servers_DnsAliaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServers_DnsAliase(generators)
	servers_DnsAliaseGenerator = gen.Struct(reflect.TypeOf(Servers_DnsAliase{}), generators)

	return servers_DnsAliaseGenerator
}

// AddRelatedPropertyGeneratorsForServers_DnsAliase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_DnsAliase(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_DnsAliase_SpecGenerator()
	gens["Status"] = Servers_DnsAliase_STATUSGenerator()
}

func Test_Servers_DnsAliase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_DnsAliase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_DnsAliase_Spec, Servers_DnsAliase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_DnsAliase_Spec runs a test to see if a specific instance of Servers_DnsAliase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_DnsAliase_Spec(subject Servers_DnsAliase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_DnsAliase_Spec
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

// Generator of Servers_DnsAliase_Spec instances for property testing - lazily instantiated by
// Servers_DnsAliase_SpecGenerator()
var servers_DnsAliase_SpecGenerator gopter.Gen

// Servers_DnsAliase_SpecGenerator returns a generator of Servers_DnsAliase_Spec instances for property testing.
func Servers_DnsAliase_SpecGenerator() gopter.Gen {
	if servers_DnsAliase_SpecGenerator != nil {
		return servers_DnsAliase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_DnsAliase_Spec(generators)
	servers_DnsAliase_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_DnsAliase_Spec{}), generators)

	return servers_DnsAliase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_DnsAliase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_DnsAliase_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
}

func Test_Servers_DnsAliase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_DnsAliase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_DnsAliase_STATUS, Servers_DnsAliase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_DnsAliase_STATUS runs a test to see if a specific instance of Servers_DnsAliase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_DnsAliase_STATUS(subject Servers_DnsAliase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_DnsAliase_STATUS
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

// Generator of Servers_DnsAliase_STATUS instances for property testing - lazily instantiated by
// Servers_DnsAliase_STATUSGenerator()
var servers_DnsAliase_STATUSGenerator gopter.Gen

// Servers_DnsAliase_STATUSGenerator returns a generator of Servers_DnsAliase_STATUS instances for property testing.
func Servers_DnsAliase_STATUSGenerator() gopter.Gen {
	if servers_DnsAliase_STATUSGenerator != nil {
		return servers_DnsAliase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_DnsAliase_STATUS(generators)
	servers_DnsAliase_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_DnsAliase_STATUS{}), generators)

	return servers_DnsAliase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_DnsAliase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_DnsAliase_STATUS(gens map[string]gopter.Gen) {
	gens["AzureDnsRecord"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
