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

func Test_Service_PolicyFragment_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_PolicyFragment_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_PolicyFragment_Spec_ARM, Service_PolicyFragment_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_PolicyFragment_Spec_ARM runs a test to see if a specific instance of Service_PolicyFragment_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_PolicyFragment_Spec_ARM(subject Service_PolicyFragment_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_PolicyFragment_Spec_ARM
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

// Generator of Service_PolicyFragment_Spec_ARM instances for property testing - lazily instantiated by
// Service_PolicyFragment_Spec_ARMGenerator()
var service_PolicyFragment_Spec_ARMGenerator gopter.Gen

// Service_PolicyFragment_Spec_ARMGenerator returns a generator of Service_PolicyFragment_Spec_ARM instances for property testing.
// We first initialize service_PolicyFragment_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_PolicyFragment_Spec_ARMGenerator() gopter.Gen {
	if service_PolicyFragment_Spec_ARMGenerator != nil {
		return service_PolicyFragment_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_PolicyFragment_Spec_ARM(generators)
	service_PolicyFragment_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_PolicyFragment_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_PolicyFragment_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForService_PolicyFragment_Spec_ARM(generators)
	service_PolicyFragment_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_PolicyFragment_Spec_ARM{}), generators)

	return service_PolicyFragment_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_PolicyFragment_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_PolicyFragment_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForService_PolicyFragment_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_PolicyFragment_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PolicyFragmentContractProperties_ARMGenerator())
}

func Test_PolicyFragmentContractProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PolicyFragmentContractProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicyFragmentContractProperties_ARM, PolicyFragmentContractProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicyFragmentContractProperties_ARM runs a test to see if a specific instance of PolicyFragmentContractProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicyFragmentContractProperties_ARM(subject PolicyFragmentContractProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PolicyFragmentContractProperties_ARM
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

// Generator of PolicyFragmentContractProperties_ARM instances for property testing - lazily instantiated by
// PolicyFragmentContractProperties_ARMGenerator()
var policyFragmentContractProperties_ARMGenerator gopter.Gen

// PolicyFragmentContractProperties_ARMGenerator returns a generator of PolicyFragmentContractProperties_ARM instances for property testing.
func PolicyFragmentContractProperties_ARMGenerator() gopter.Gen {
	if policyFragmentContractProperties_ARMGenerator != nil {
		return policyFragmentContractProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicyFragmentContractProperties_ARM(generators)
	policyFragmentContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(PolicyFragmentContractProperties_ARM{}), generators)

	return policyFragmentContractProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPolicyFragmentContractProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPolicyFragmentContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Format"] = gen.PtrOf(gen.OneConstOf(PolicyFragmentContractProperties_Format_Rawxml, PolicyFragmentContractProperties_Format_Xml))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
