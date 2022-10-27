// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_ManagedInstances_DnsAliase_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedInstances_DnsAliase_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedInstances_DnsAliase_STATUS_ARM, ManagedInstances_DnsAliase_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedInstances_DnsAliase_STATUS_ARM runs a test to see if a specific instance of ManagedInstances_DnsAliase_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedInstances_DnsAliase_STATUS_ARM(subject ManagedInstances_DnsAliase_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedInstances_DnsAliase_STATUS_ARM
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

// Generator of ManagedInstances_DnsAliase_STATUS_ARM instances for property testing - lazily instantiated by
// ManagedInstances_DnsAliase_STATUS_ARMGenerator()
var managedInstances_DnsAliase_STATUS_ARMGenerator gopter.Gen

// ManagedInstances_DnsAliase_STATUS_ARMGenerator returns a generator of ManagedInstances_DnsAliase_STATUS_ARM instances for property testing.
// We first initialize managedInstances_DnsAliase_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedInstances_DnsAliase_STATUS_ARMGenerator() gopter.Gen {
	if managedInstances_DnsAliase_STATUS_ARMGenerator != nil {
		return managedInstances_DnsAliase_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedInstances_DnsAliase_STATUS_ARM(generators)
	managedInstances_DnsAliase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_DnsAliase_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedInstances_DnsAliase_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagedInstances_DnsAliase_STATUS_ARM(generators)
	managedInstances_DnsAliase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_DnsAliase_STATUS_ARM{}), generators)

	return managedInstances_DnsAliase_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedInstances_DnsAliase_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedInstances_DnsAliase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedInstances_DnsAliase_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedInstances_DnsAliase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagedServerDnsAliasProperties_STATUS_ARMGenerator())
}

func Test_ManagedServerDnsAliasProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedServerDnsAliasProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedServerDnsAliasProperties_STATUS_ARM, ManagedServerDnsAliasProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedServerDnsAliasProperties_STATUS_ARM runs a test to see if a specific instance of ManagedServerDnsAliasProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedServerDnsAliasProperties_STATUS_ARM(subject ManagedServerDnsAliasProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedServerDnsAliasProperties_STATUS_ARM
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

// Generator of ManagedServerDnsAliasProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ManagedServerDnsAliasProperties_STATUS_ARMGenerator()
var managedServerDnsAliasProperties_STATUS_ARMGenerator gopter.Gen

// ManagedServerDnsAliasProperties_STATUS_ARMGenerator returns a generator of ManagedServerDnsAliasProperties_STATUS_ARM instances for property testing.
func ManagedServerDnsAliasProperties_STATUS_ARMGenerator() gopter.Gen {
	if managedServerDnsAliasProperties_STATUS_ARMGenerator != nil {
		return managedServerDnsAliasProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedServerDnsAliasProperties_STATUS_ARM(generators)
	managedServerDnsAliasProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedServerDnsAliasProperties_STATUS_ARM{}), generators)

	return managedServerDnsAliasProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedServerDnsAliasProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedServerDnsAliasProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AzureDnsRecord"] = gen.PtrOf(gen.AlphaString())
	gens["PublicAzureDnsRecord"] = gen.PtrOf(gen.AlphaString())
}
