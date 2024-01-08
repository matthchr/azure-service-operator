// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

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

func Test_Service_NamedValue_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_NamedValue_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_NamedValue_Spec_ARM, Service_NamedValue_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_NamedValue_Spec_ARM runs a test to see if a specific instance of Service_NamedValue_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_NamedValue_Spec_ARM(subject Service_NamedValue_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_NamedValue_Spec_ARM
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

// Generator of Service_NamedValue_Spec_ARM instances for property testing - lazily instantiated by
// Service_NamedValue_Spec_ARMGenerator()
var service_NamedValue_Spec_ARMGenerator gopter.Gen

// Service_NamedValue_Spec_ARMGenerator returns a generator of Service_NamedValue_Spec_ARM instances for property testing.
// We first initialize service_NamedValue_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_NamedValue_Spec_ARMGenerator() gopter.Gen {
	if service_NamedValue_Spec_ARMGenerator != nil {
		return service_NamedValue_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_NamedValue_Spec_ARM(generators)
	service_NamedValue_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_NamedValue_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_NamedValue_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForService_NamedValue_Spec_ARM(generators)
	service_NamedValue_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_NamedValue_Spec_ARM{}), generators)

	return service_NamedValue_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_NamedValue_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_NamedValue_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForService_NamedValue_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_NamedValue_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NamedValueCreateContractProperties_ARMGenerator())
}

func Test_NamedValueCreateContractProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamedValueCreateContractProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamedValueCreateContractProperties_ARM, NamedValueCreateContractProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamedValueCreateContractProperties_ARM runs a test to see if a specific instance of NamedValueCreateContractProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamedValueCreateContractProperties_ARM(subject NamedValueCreateContractProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamedValueCreateContractProperties_ARM
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

// Generator of NamedValueCreateContractProperties_ARM instances for property testing - lazily instantiated by
// NamedValueCreateContractProperties_ARMGenerator()
var namedValueCreateContractProperties_ARMGenerator gopter.Gen

// NamedValueCreateContractProperties_ARMGenerator returns a generator of NamedValueCreateContractProperties_ARM instances for property testing.
// We first initialize namedValueCreateContractProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamedValueCreateContractProperties_ARMGenerator() gopter.Gen {
	if namedValueCreateContractProperties_ARMGenerator != nil {
		return namedValueCreateContractProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamedValueCreateContractProperties_ARM(generators)
	namedValueCreateContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(NamedValueCreateContractProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamedValueCreateContractProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForNamedValueCreateContractProperties_ARM(generators)
	namedValueCreateContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(NamedValueCreateContractProperties_ARM{}), generators)

	return namedValueCreateContractProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamedValueCreateContractProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamedValueCreateContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Secret"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.SliceOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamedValueCreateContractProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamedValueCreateContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["KeyVault"] = gen.PtrOf(KeyVaultContractCreateProperties_ARMGenerator())
}

func Test_KeyVaultContractCreateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultContractCreateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultContractCreateProperties_ARM, KeyVaultContractCreateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultContractCreateProperties_ARM runs a test to see if a specific instance of KeyVaultContractCreateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultContractCreateProperties_ARM(subject KeyVaultContractCreateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultContractCreateProperties_ARM
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

// Generator of KeyVaultContractCreateProperties_ARM instances for property testing - lazily instantiated by
// KeyVaultContractCreateProperties_ARMGenerator()
var keyVaultContractCreateProperties_ARMGenerator gopter.Gen

// KeyVaultContractCreateProperties_ARMGenerator returns a generator of KeyVaultContractCreateProperties_ARM instances for property testing.
func KeyVaultContractCreateProperties_ARMGenerator() gopter.Gen {
	if keyVaultContractCreateProperties_ARMGenerator != nil {
		return keyVaultContractCreateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties_ARM(generators)
	keyVaultContractCreateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultContractCreateProperties_ARM{}), generators)

	return keyVaultContractCreateProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties_ARM(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["SecretIdentifier"] = gen.PtrOf(gen.AlphaString())
}
