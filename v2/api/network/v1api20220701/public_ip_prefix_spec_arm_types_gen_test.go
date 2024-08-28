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

func Test_IpTag_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpTag_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpTag_ARM, IpTag_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpTag_ARM runs a test to see if a specific instance of IpTag_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpTag_ARM(subject IpTag_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpTag_ARM
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

// Generator of IpTag_ARM instances for property testing - lazily instantiated by IpTag_ARMGenerator()
var ipTag_ARMGenerator gopter.Gen

// IpTag_ARMGenerator returns a generator of IpTag_ARM instances for property testing.
func IpTag_ARMGenerator() gopter.Gen {
	if ipTag_ARMGenerator != nil {
		return ipTag_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpTag_ARM(generators)
	ipTag_ARMGenerator = gen.Struct(reflect.TypeOf(IpTag_ARM{}), generators)

	return ipTag_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIpTag_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpTag_ARM(gens map[string]gopter.Gen) {
	gens["IpTagType"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM, NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM runs a test to see if a specific instance of NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM(subject NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM
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

// Generator of NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator()
var natGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator gopter.Gen

// NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator returns a generator of NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM instances for property testing.
func NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if natGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator != nil {
		return natGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM(generators)
	natGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM{}), generators)

	return natGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPPrefixPropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefixPropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefixPropertiesFormat_ARM, PublicIPPrefixPropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefixPropertiesFormat_ARM runs a test to see if a specific instance of PublicIPPrefixPropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefixPropertiesFormat_ARM(subject PublicIPPrefixPropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefixPropertiesFormat_ARM
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

// Generator of PublicIPPrefixPropertiesFormat_ARM instances for property testing - lazily instantiated by
// PublicIPPrefixPropertiesFormat_ARMGenerator()
var publicIPPrefixPropertiesFormat_ARMGenerator gopter.Gen

// PublicIPPrefixPropertiesFormat_ARMGenerator returns a generator of PublicIPPrefixPropertiesFormat_ARM instances for property testing.
// We first initialize publicIPPrefixPropertiesFormat_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPPrefixPropertiesFormat_ARMGenerator() gopter.Gen {
	if publicIPPrefixPropertiesFormat_ARMGenerator != nil {
		return publicIPPrefixPropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefixPropertiesFormat_ARM(generators)
	publicIPPrefixPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefixPropertiesFormat_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefixPropertiesFormat_ARM(generators)
	AddRelatedPropertyGeneratorsForPublicIPPrefixPropertiesFormat_ARM(generators)
	publicIPPrefixPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefixPropertiesFormat_ARM{}), generators)

	return publicIPPrefixPropertiesFormat_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPPrefixPropertiesFormat_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPPrefixPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["PrefixLength"] = gen.PtrOf(gen.Int())
	gens["PublicIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(IPVersion_ARM_IPv4, IPVersion_ARM_IPv6))
}

// AddRelatedPropertyGeneratorsForPublicIPPrefixPropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPPrefixPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["CustomIPPrefix"] = gen.PtrOf(PublicIpPrefixSubResource_ARMGenerator())
	gens["IpTags"] = gen.SliceOf(IpTag_ARMGenerator())
	gens["NatGateway"] = gen.PtrOf(NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_ARMGenerator())
}

func Test_PublicIPPrefixSku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefixSku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefixSku_ARM, PublicIPPrefixSku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefixSku_ARM runs a test to see if a specific instance of PublicIPPrefixSku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefixSku_ARM(subject PublicIPPrefixSku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefixSku_ARM
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

// Generator of PublicIPPrefixSku_ARM instances for property testing - lazily instantiated by
// PublicIPPrefixSku_ARMGenerator()
var publicIPPrefixSku_ARMGenerator gopter.Gen

// PublicIPPrefixSku_ARMGenerator returns a generator of PublicIPPrefixSku_ARM instances for property testing.
func PublicIPPrefixSku_ARMGenerator() gopter.Gen {
	if publicIPPrefixSku_ARMGenerator != nil {
		return publicIPPrefixSku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefixSku_ARM(generators)
	publicIPPrefixSku_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefixSku_ARM{}), generators)

	return publicIPPrefixSku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPPrefixSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPPrefixSku_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(PublicIPPrefixSku_Name_ARM_Standard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(PublicIPPrefixSku_Tier_ARM_Global, PublicIPPrefixSku_Tier_ARM_Regional))
}

func Test_PublicIPPrefix_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefix_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefix_Spec_ARM, PublicIPPrefix_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefix_Spec_ARM runs a test to see if a specific instance of PublicIPPrefix_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefix_Spec_ARM(subject PublicIPPrefix_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefix_Spec_ARM
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

// Generator of PublicIPPrefix_Spec_ARM instances for property testing - lazily instantiated by
// PublicIPPrefix_Spec_ARMGenerator()
var publicIPPrefix_Spec_ARMGenerator gopter.Gen

// PublicIPPrefix_Spec_ARMGenerator returns a generator of PublicIPPrefix_Spec_ARM instances for property testing.
// We first initialize publicIPPrefix_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPPrefix_Spec_ARMGenerator() gopter.Gen {
	if publicIPPrefix_Spec_ARMGenerator != nil {
		return publicIPPrefix_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefix_Spec_ARM(generators)
	publicIPPrefix_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefix_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefix_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForPublicIPPrefix_Spec_ARM(generators)
	publicIPPrefix_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefix_Spec_ARM{}), generators)

	return publicIPPrefix_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPPrefix_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPPrefix_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPPrefix_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPPrefix_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_ARMGenerator())
	gens["Properties"] = gen.PtrOf(PublicIPPrefixPropertiesFormat_ARMGenerator())
	gens["Sku"] = gen.PtrOf(PublicIPPrefixSku_ARMGenerator())
}

func Test_PublicIpPrefixSubResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIpPrefixSubResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIpPrefixSubResource_ARM, PublicIpPrefixSubResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIpPrefixSubResource_ARM runs a test to see if a specific instance of PublicIpPrefixSubResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIpPrefixSubResource_ARM(subject PublicIpPrefixSubResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIpPrefixSubResource_ARM
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

// Generator of PublicIpPrefixSubResource_ARM instances for property testing - lazily instantiated by
// PublicIpPrefixSubResource_ARMGenerator()
var publicIpPrefixSubResource_ARMGenerator gopter.Gen

// PublicIpPrefixSubResource_ARMGenerator returns a generator of PublicIpPrefixSubResource_ARM instances for property testing.
func PublicIpPrefixSubResource_ARMGenerator() gopter.Gen {
	if publicIpPrefixSubResource_ARMGenerator != nil {
		return publicIpPrefixSubResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIpPrefixSubResource_ARM(generators)
	publicIpPrefixSubResource_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIpPrefixSubResource_ARM{}), generators)

	return publicIpPrefixSubResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIpPrefixSubResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIpPrefixSubResource_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
