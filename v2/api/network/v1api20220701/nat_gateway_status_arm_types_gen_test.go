// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_NatGatewayPropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewayPropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewayPropertiesFormat_STATUS_ARM, NatGatewayPropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewayPropertiesFormat_STATUS_ARM runs a test to see if a specific instance of NatGatewayPropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewayPropertiesFormat_STATUS_ARM(subject NatGatewayPropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewayPropertiesFormat_STATUS_ARM
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

// Generator of NatGatewayPropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated by
// NatGatewayPropertiesFormat_STATUS_ARMGenerator()
var natGatewayPropertiesFormat_STATUS_ARMGenerator gopter.Gen

// NatGatewayPropertiesFormat_STATUS_ARMGenerator returns a generator of NatGatewayPropertiesFormat_STATUS_ARM instances for property testing.
// We first initialize natGatewayPropertiesFormat_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NatGatewayPropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if natGatewayPropertiesFormat_STATUS_ARMGenerator != nil {
		return natGatewayPropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewayPropertiesFormat_STATUS_ARM(generators)
	natGatewayPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NatGatewayPropertiesFormat_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewayPropertiesFormat_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForNatGatewayPropertiesFormat_STATUS_ARM(generators)
	natGatewayPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NatGatewayPropertiesFormat_STATUS_ARM{}), generators)

	return natGatewayPropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewayPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewayPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ApplicationGatewayProvisioningState_STATUS_ARM_Deleting,
		ApplicationGatewayProvisioningState_STATUS_ARM_Failed,
		ApplicationGatewayProvisioningState_STATUS_ARM_Succeeded,
		ApplicationGatewayProvisioningState_STATUS_ARM_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNatGatewayPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNatGatewayPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PublicIpAddresses"] = gen.SliceOf(ApplicationGatewaySubResource_STATUS_ARMGenerator())
	gens["PublicIpPrefixes"] = gen.SliceOf(ApplicationGatewaySubResource_STATUS_ARMGenerator())
	gens["Subnets"] = gen.SliceOf(ApplicationGatewaySubResource_STATUS_ARMGenerator())
}

func Test_NatGatewaySku_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewaySku_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewaySku_STATUS_ARM, NatGatewaySku_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewaySku_STATUS_ARM runs a test to see if a specific instance of NatGatewaySku_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewaySku_STATUS_ARM(subject NatGatewaySku_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewaySku_STATUS_ARM
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

// Generator of NatGatewaySku_STATUS_ARM instances for property testing - lazily instantiated by
// NatGatewaySku_STATUS_ARMGenerator()
var natGatewaySku_STATUS_ARMGenerator gopter.Gen

// NatGatewaySku_STATUS_ARMGenerator returns a generator of NatGatewaySku_STATUS_ARM instances for property testing.
func NatGatewaySku_STATUS_ARMGenerator() gopter.Gen {
	if natGatewaySku_STATUS_ARMGenerator != nil {
		return natGatewaySku_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewaySku_STATUS_ARM(generators)
	natGatewaySku_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NatGatewaySku_STATUS_ARM{}), generators)

	return natGatewaySku_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewaySku_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewaySku_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(NatGatewaySku_Name_STATUS_ARM_Standard))
}

func Test_NatGateway_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGateway_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGateway_STATUS_ARM, NatGateway_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGateway_STATUS_ARM runs a test to see if a specific instance of NatGateway_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGateway_STATUS_ARM(subject NatGateway_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGateway_STATUS_ARM
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

// Generator of NatGateway_STATUS_ARM instances for property testing - lazily instantiated by
// NatGateway_STATUS_ARMGenerator()
var natGateway_STATUS_ARMGenerator gopter.Gen

// NatGateway_STATUS_ARMGenerator returns a generator of NatGateway_STATUS_ARM instances for property testing.
// We first initialize natGateway_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NatGateway_STATUS_ARMGenerator() gopter.Gen {
	if natGateway_STATUS_ARMGenerator != nil {
		return natGateway_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_STATUS_ARM(generators)
	natGateway_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NatGateway_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForNatGateway_STATUS_ARM(generators)
	natGateway_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NatGateway_STATUS_ARM{}), generators)

	return natGateway_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGateway_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGateway_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNatGateway_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNatGateway_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NatGatewayPropertiesFormat_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(NatGatewaySku_STATUS_ARMGenerator())
}
