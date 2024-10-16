// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/storage"
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

func Test_DnsForwardingRuleSetsVirtualNetworkLink_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsForwardingRuleSetsVirtualNetworkLink to hub returns original",
		prop.ForAll(RunResourceConversionTestForDnsForwardingRuleSetsVirtualNetworkLink, DnsForwardingRuleSetsVirtualNetworkLinkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDnsForwardingRuleSetsVirtualNetworkLink tests if a specific instance of DnsForwardingRuleSetsVirtualNetworkLink round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDnsForwardingRuleSetsVirtualNetworkLink(subject DnsForwardingRuleSetsVirtualNetworkLink) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.DnsForwardingRuleSetsVirtualNetworkLink
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DnsForwardingRuleSetsVirtualNetworkLink
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

func Test_DnsForwardingRuleSetsVirtualNetworkLink_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsForwardingRuleSetsVirtualNetworkLink to DnsForwardingRuleSetsVirtualNetworkLink via AssignProperties_To_DnsForwardingRuleSetsVirtualNetworkLink & AssignProperties_From_DnsForwardingRuleSetsVirtualNetworkLink returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsForwardingRuleSetsVirtualNetworkLink, DnsForwardingRuleSetsVirtualNetworkLinkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsForwardingRuleSetsVirtualNetworkLink tests if a specific instance of DnsForwardingRuleSetsVirtualNetworkLink can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsForwardingRuleSetsVirtualNetworkLink(subject DnsForwardingRuleSetsVirtualNetworkLink) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsForwardingRuleSetsVirtualNetworkLink
	err := copied.AssignProperties_To_DnsForwardingRuleSetsVirtualNetworkLink(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsForwardingRuleSetsVirtualNetworkLink
	err = actual.AssignProperties_From_DnsForwardingRuleSetsVirtualNetworkLink(&other)
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

func Test_DnsForwardingRuleSetsVirtualNetworkLink_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleSetsVirtualNetworkLink via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink, DnsForwardingRuleSetsVirtualNetworkLinkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink runs a test to see if a specific instance of DnsForwardingRuleSetsVirtualNetworkLink round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink(subject DnsForwardingRuleSetsVirtualNetworkLink) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleSetsVirtualNetworkLink
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

// Generator of DnsForwardingRuleSetsVirtualNetworkLink instances for property testing - lazily instantiated by
// DnsForwardingRuleSetsVirtualNetworkLinkGenerator()
var dnsForwardingRuleSetsVirtualNetworkLinkGenerator gopter.Gen

// DnsForwardingRuleSetsVirtualNetworkLinkGenerator returns a generator of DnsForwardingRuleSetsVirtualNetworkLink instances for property testing.
func DnsForwardingRuleSetsVirtualNetworkLinkGenerator() gopter.Gen {
	if dnsForwardingRuleSetsVirtualNetworkLinkGenerator != nil {
		return dnsForwardingRuleSetsVirtualNetworkLinkGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink(generators)
	dnsForwardingRuleSetsVirtualNetworkLinkGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleSetsVirtualNetworkLink{}), generators)

	return dnsForwardingRuleSetsVirtualNetworkLinkGenerator
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsForwardingRulesets_VirtualNetworkLink_SpecGenerator()
	gens["Status"] = DnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator()
}

func Test_DnsForwardingRulesets_VirtualNetworkLink_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsForwardingRulesets_VirtualNetworkLink_STATUS to DnsForwardingRulesets_VirtualNetworkLink_STATUS via AssignProperties_To_DnsForwardingRulesets_VirtualNetworkLink_STATUS & AssignProperties_From_DnsForwardingRulesets_VirtualNetworkLink_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsForwardingRulesets_VirtualNetworkLink_STATUS, DnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsForwardingRulesets_VirtualNetworkLink_STATUS tests if a specific instance of DnsForwardingRulesets_VirtualNetworkLink_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsForwardingRulesets_VirtualNetworkLink_STATUS(subject DnsForwardingRulesets_VirtualNetworkLink_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsForwardingRulesets_VirtualNetworkLink_STATUS
	err := copied.AssignProperties_To_DnsForwardingRulesets_VirtualNetworkLink_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsForwardingRulesets_VirtualNetworkLink_STATUS
	err = actual.AssignProperties_From_DnsForwardingRulesets_VirtualNetworkLink_STATUS(&other)
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

func Test_DnsForwardingRulesets_VirtualNetworkLink_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRulesets_VirtualNetworkLink_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRulesets_VirtualNetworkLink_STATUS, DnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRulesets_VirtualNetworkLink_STATUS runs a test to see if a specific instance of DnsForwardingRulesets_VirtualNetworkLink_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRulesets_VirtualNetworkLink_STATUS(subject DnsForwardingRulesets_VirtualNetworkLink_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRulesets_VirtualNetworkLink_STATUS
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

// Generator of DnsForwardingRulesets_VirtualNetworkLink_STATUS instances for property testing - lazily instantiated by
// DnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator()
var dnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator gopter.Gen

// DnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator returns a generator of DnsForwardingRulesets_VirtualNetworkLink_STATUS instances for property testing.
// We first initialize dnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator() gopter.Gen {
	if dnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator != nil {
		return dnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_STATUS(generators)
	dnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesets_VirtualNetworkLink_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_STATUS(generators)
	dnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesets_VirtualNetworkLink_STATUS{}), generators)

	return dnsForwardingRulesets_VirtualNetworkLink_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_Canceled,
		DnsresolverProvisioningState_STATUS_Creating,
		DnsresolverProvisioningState_STATUS_Deleting,
		DnsresolverProvisioningState_STATUS_Failed,
		DnsresolverProvisioningState_STATUS_Succeeded,
		DnsresolverProvisioningState_STATUS_Updating))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
	gens["VirtualNetwork"] = gen.PtrOf(DnsresolverSubResource_STATUSGenerator())
}

func Test_DnsForwardingRulesets_VirtualNetworkLink_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsForwardingRulesets_VirtualNetworkLink_Spec to DnsForwardingRulesets_VirtualNetworkLink_Spec via AssignProperties_To_DnsForwardingRulesets_VirtualNetworkLink_Spec & AssignProperties_From_DnsForwardingRulesets_VirtualNetworkLink_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsForwardingRulesets_VirtualNetworkLink_Spec, DnsForwardingRulesets_VirtualNetworkLink_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsForwardingRulesets_VirtualNetworkLink_Spec tests if a specific instance of DnsForwardingRulesets_VirtualNetworkLink_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsForwardingRulesets_VirtualNetworkLink_Spec(subject DnsForwardingRulesets_VirtualNetworkLink_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsForwardingRulesets_VirtualNetworkLink_Spec
	err := copied.AssignProperties_To_DnsForwardingRulesets_VirtualNetworkLink_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsForwardingRulesets_VirtualNetworkLink_Spec
	err = actual.AssignProperties_From_DnsForwardingRulesets_VirtualNetworkLink_Spec(&other)
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

func Test_DnsForwardingRulesets_VirtualNetworkLink_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRulesets_VirtualNetworkLink_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRulesets_VirtualNetworkLink_Spec, DnsForwardingRulesets_VirtualNetworkLink_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRulesets_VirtualNetworkLink_Spec runs a test to see if a specific instance of DnsForwardingRulesets_VirtualNetworkLink_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRulesets_VirtualNetworkLink_Spec(subject DnsForwardingRulesets_VirtualNetworkLink_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRulesets_VirtualNetworkLink_Spec
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

// Generator of DnsForwardingRulesets_VirtualNetworkLink_Spec instances for property testing - lazily instantiated by
// DnsForwardingRulesets_VirtualNetworkLink_SpecGenerator()
var dnsForwardingRulesets_VirtualNetworkLink_SpecGenerator gopter.Gen

// DnsForwardingRulesets_VirtualNetworkLink_SpecGenerator returns a generator of DnsForwardingRulesets_VirtualNetworkLink_Spec instances for property testing.
// We first initialize dnsForwardingRulesets_VirtualNetworkLink_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRulesets_VirtualNetworkLink_SpecGenerator() gopter.Gen {
	if dnsForwardingRulesets_VirtualNetworkLink_SpecGenerator != nil {
		return dnsForwardingRulesets_VirtualNetworkLink_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_Spec(generators)
	dnsForwardingRulesets_VirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesets_VirtualNetworkLink_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_Spec(generators)
	dnsForwardingRulesets_VirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesets_VirtualNetworkLink_Spec{}), generators)

	return dnsForwardingRulesets_VirtualNetworkLink_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRulesets_VirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(DnsresolverSubResourceGenerator())
}

func Test_DnsresolverSubResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsresolverSubResource to DnsresolverSubResource via AssignProperties_To_DnsresolverSubResource & AssignProperties_From_DnsresolverSubResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsresolverSubResource, DnsresolverSubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsresolverSubResource tests if a specific instance of DnsresolverSubResource can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsresolverSubResource(subject DnsresolverSubResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsresolverSubResource
	err := copied.AssignProperties_To_DnsresolverSubResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsresolverSubResource
	err = actual.AssignProperties_From_DnsresolverSubResource(&other)
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

func Test_DnsresolverSubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsresolverSubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsresolverSubResource, DnsresolverSubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsresolverSubResource runs a test to see if a specific instance of DnsresolverSubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsresolverSubResource(subject DnsresolverSubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsresolverSubResource
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

// Generator of DnsresolverSubResource instances for property testing - lazily instantiated by
// DnsresolverSubResourceGenerator()
var dnsresolverSubResourceGenerator gopter.Gen

// DnsresolverSubResourceGenerator returns a generator of DnsresolverSubResource instances for property testing.
func DnsresolverSubResourceGenerator() gopter.Gen {
	if dnsresolverSubResourceGenerator != nil {
		return dnsresolverSubResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	dnsresolverSubResourceGenerator = gen.Struct(reflect.TypeOf(DnsresolverSubResource{}), generators)

	return dnsresolverSubResourceGenerator
}

func Test_DnsresolverSubResource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsresolverSubResource_STATUS to DnsresolverSubResource_STATUS via AssignProperties_To_DnsresolverSubResource_STATUS & AssignProperties_From_DnsresolverSubResource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsresolverSubResource_STATUS, DnsresolverSubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsresolverSubResource_STATUS tests if a specific instance of DnsresolverSubResource_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsresolverSubResource_STATUS(subject DnsresolverSubResource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsresolverSubResource_STATUS
	err := copied.AssignProperties_To_DnsresolverSubResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsresolverSubResource_STATUS
	err = actual.AssignProperties_From_DnsresolverSubResource_STATUS(&other)
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

func Test_DnsresolverSubResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsresolverSubResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsresolverSubResource_STATUS, DnsresolverSubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsresolverSubResource_STATUS runs a test to see if a specific instance of DnsresolverSubResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsresolverSubResource_STATUS(subject DnsresolverSubResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsresolverSubResource_STATUS
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

// Generator of DnsresolverSubResource_STATUS instances for property testing - lazily instantiated by
// DnsresolverSubResource_STATUSGenerator()
var dnsresolverSubResource_STATUSGenerator gopter.Gen

// DnsresolverSubResource_STATUSGenerator returns a generator of DnsresolverSubResource_STATUS instances for property testing.
func DnsresolverSubResource_STATUSGenerator() gopter.Gen {
	if dnsresolverSubResource_STATUSGenerator != nil {
		return dnsresolverSubResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS(generators)
	dnsresolverSubResource_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsresolverSubResource_STATUS{}), generators)

	return dnsresolverSubResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
