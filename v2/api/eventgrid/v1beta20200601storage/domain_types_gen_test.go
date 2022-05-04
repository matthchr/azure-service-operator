// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

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

func Test_Domain_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomain, DomainGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomain runs a test to see if a specific instance of Domain round trips to JSON and back losslessly
func RunJSONSerializationTestForDomain(subject Domain) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domain
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

// Generator of Domain instances for property testing - lazily instantiated by DomainGenerator()
var domainGenerator gopter.Gen

// DomainGenerator returns a generator of Domain instances for property testing.
func DomainGenerator() gopter.Gen {
	if domainGenerator != nil {
		return domainGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDomain(generators)
	domainGenerator = gen.Struct(reflect.TypeOf(Domain{}), generators)

	return domainGenerator
}

// AddRelatedPropertyGeneratorsForDomain is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomain(gens map[string]gopter.Gen) {
	gens["Spec"] = DomainsSpecGenerator()
	gens["Status"] = DomainStatusGenerator()
}

func Test_Domain_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainStatus, DomainStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainStatus runs a test to see if a specific instance of Domain_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainStatus(subject Domain_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domain_Status
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

// Generator of Domain_Status instances for property testing - lazily instantiated by DomainStatusGenerator()
var domainStatusGenerator gopter.Gen

// DomainStatusGenerator returns a generator of Domain_Status instances for property testing.
// We first initialize domainStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainStatusGenerator() gopter.Gen {
	if domainStatusGenerator != nil {
		return domainStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainStatus(generators)
	domainStatusGenerator = gen.Struct(reflect.TypeOf(Domain_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainStatus(generators)
	AddRelatedPropertyGeneratorsForDomainStatus(generators)
	domainStatusGenerator = gen.Struct(reflect.TypeOf(Domain_Status{}), generators)

	return domainStatusGenerator
}

// AddIndependentPropertyGeneratorsForDomainStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainStatus(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["InputSchema"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomainStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainStatus(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRuleStatusGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMappingStatusGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_Domains_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domains_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsSpec, DomainsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsSpec runs a test to see if a specific instance of Domains_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsSpec(subject Domains_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domains_Spec
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

// Generator of Domains_Spec instances for property testing - lazily instantiated by DomainsSpecGenerator()
var domainsSpecGenerator gopter.Gen

// DomainsSpecGenerator returns a generator of Domains_Spec instances for property testing.
// We first initialize domainsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainsSpecGenerator() gopter.Gen {
	if domainsSpecGenerator != nil {
		return domainsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsSpec(generators)
	domainsSpecGenerator = gen.Struct(reflect.TypeOf(Domains_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsSpec(generators)
	AddRelatedPropertyGeneratorsForDomainsSpec(generators)
	domainsSpecGenerator = gen.Struct(reflect.TypeOf(Domains_Spec{}), generators)

	return domainsSpecGenerator
}

// AddIndependentPropertyGeneratorsForDomainsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["InputSchema"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomainsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsSpec(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRuleGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(JsonInputSchemaMappingGenerator())
}

func Test_InboundIpRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundIpRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundIpRule, InboundIpRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundIpRule runs a test to see if a specific instance of InboundIpRule round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundIpRule(subject InboundIpRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundIpRule
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

// Generator of InboundIpRule instances for property testing - lazily instantiated by InboundIpRuleGenerator()
var inboundIpRuleGenerator gopter.Gen

// InboundIpRuleGenerator returns a generator of InboundIpRule instances for property testing.
func InboundIpRuleGenerator() gopter.Gen {
	if inboundIpRuleGenerator != nil {
		return inboundIpRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundIpRule(generators)
	inboundIpRuleGenerator = gen.Struct(reflect.TypeOf(InboundIpRule{}), generators)

	return inboundIpRuleGenerator
}

// AddIndependentPropertyGeneratorsForInboundIpRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundIpRule(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.AlphaString())
	gens["IpMask"] = gen.PtrOf(gen.AlphaString())
}

func Test_InboundIpRule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundIpRule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundIpRuleStatus, InboundIpRuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundIpRuleStatus runs a test to see if a specific instance of InboundIpRule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundIpRuleStatus(subject InboundIpRule_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundIpRule_Status
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

// Generator of InboundIpRule_Status instances for property testing - lazily instantiated by
// InboundIpRuleStatusGenerator()
var inboundIpRuleStatusGenerator gopter.Gen

// InboundIpRuleStatusGenerator returns a generator of InboundIpRule_Status instances for property testing.
func InboundIpRuleStatusGenerator() gopter.Gen {
	if inboundIpRuleStatusGenerator != nil {
		return inboundIpRuleStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundIpRuleStatus(generators)
	inboundIpRuleStatusGenerator = gen.Struct(reflect.TypeOf(InboundIpRule_Status{}), generators)

	return inboundIpRuleStatusGenerator
}

// AddIndependentPropertyGeneratorsForInboundIpRuleStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundIpRuleStatus(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.AlphaString())
	gens["IpMask"] = gen.PtrOf(gen.AlphaString())
}

func Test_InputSchemaMapping_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InputSchemaMapping_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInputSchemaMappingStatus, InputSchemaMappingStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInputSchemaMappingStatus runs a test to see if a specific instance of InputSchemaMapping_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForInputSchemaMappingStatus(subject InputSchemaMapping_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InputSchemaMapping_Status
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

// Generator of InputSchemaMapping_Status instances for property testing - lazily instantiated by
// InputSchemaMappingStatusGenerator()
var inputSchemaMappingStatusGenerator gopter.Gen

// InputSchemaMappingStatusGenerator returns a generator of InputSchemaMapping_Status instances for property testing.
func InputSchemaMappingStatusGenerator() gopter.Gen {
	if inputSchemaMappingStatusGenerator != nil {
		return inputSchemaMappingStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInputSchemaMappingStatus(generators)
	inputSchemaMappingStatusGenerator = gen.Struct(reflect.TypeOf(InputSchemaMapping_Status{}), generators)

	return inputSchemaMappingStatusGenerator
}

// AddIndependentPropertyGeneratorsForInputSchemaMappingStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInputSchemaMappingStatus(gens map[string]gopter.Gen) {
	gens["InputSchemaMappingType"] = gen.PtrOf(gen.AlphaString())
}

func Test_JsonInputSchemaMapping_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of JsonInputSchemaMapping via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForJsonInputSchemaMapping, JsonInputSchemaMappingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForJsonInputSchemaMapping runs a test to see if a specific instance of JsonInputSchemaMapping round trips to JSON and back losslessly
func RunJSONSerializationTestForJsonInputSchemaMapping(subject JsonInputSchemaMapping) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual JsonInputSchemaMapping
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

// Generator of JsonInputSchemaMapping instances for property testing - lazily instantiated by
// JsonInputSchemaMappingGenerator()
var jsonInputSchemaMappingGenerator gopter.Gen

// JsonInputSchemaMappingGenerator returns a generator of JsonInputSchemaMapping instances for property testing.
// We first initialize jsonInputSchemaMappingGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func JsonInputSchemaMappingGenerator() gopter.Gen {
	if jsonInputSchemaMappingGenerator != nil {
		return jsonInputSchemaMappingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForJsonInputSchemaMapping(generators)
	jsonInputSchemaMappingGenerator = gen.Struct(reflect.TypeOf(JsonInputSchemaMapping{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForJsonInputSchemaMapping(generators)
	AddRelatedPropertyGeneratorsForJsonInputSchemaMapping(generators)
	jsonInputSchemaMappingGenerator = gen.Struct(reflect.TypeOf(JsonInputSchemaMapping{}), generators)

	return jsonInputSchemaMappingGenerator
}

// AddIndependentPropertyGeneratorsForJsonInputSchemaMapping is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForJsonInputSchemaMapping(gens map[string]gopter.Gen) {
	gens["InputSchemaMappingType"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForJsonInputSchemaMapping is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForJsonInputSchemaMapping(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(JsonInputSchemaMappingPropertiesGenerator())
}

func Test_PrivateEndpointConnection_Status_Domain_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionStatusDomainSubResourceEmbedded, PrivateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionStatusDomainSubResourceEmbedded runs a test to see if a specific instance of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionStatusDomainSubResourceEmbedded(subject PrivateEndpointConnection_Status_Domain_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_Domain_SubResourceEmbedded
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

// Generator of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded instances for property testing - lazily
// instantiated by PrivateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator()
var privateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator gopter.Gen

// PrivateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator returns a generator of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded instances for property testing.
func PrivateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator() gopter.Gen {
	if privateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator != nil {
		return privateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusDomainSubResourceEmbedded(generators)
	privateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_Domain_SubResourceEmbedded{}), generators)

	return privateEndpointConnectionStatusDomainSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusDomainSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusDomainSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatus, SystemDataStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatus runs a test to see if a specific instance of SystemData_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatus(subject SystemData_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_Status
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

// Generator of SystemData_Status instances for property testing - lazily instantiated by SystemDataStatusGenerator()
var systemDataStatusGenerator gopter.Gen

// SystemDataStatusGenerator returns a generator of SystemData_Status instances for property testing.
func SystemDataStatusGenerator() gopter.Gen {
	if systemDataStatusGenerator != nil {
		return systemDataStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatus(generators)
	systemDataStatusGenerator = gen.Struct(reflect.TypeOf(SystemData_Status{}), generators)

	return systemDataStatusGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatus(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}

func Test_JsonInputSchemaMappingProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of JsonInputSchemaMappingProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForJsonInputSchemaMappingProperties, JsonInputSchemaMappingPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForJsonInputSchemaMappingProperties runs a test to see if a specific instance of JsonInputSchemaMappingProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForJsonInputSchemaMappingProperties(subject JsonInputSchemaMappingProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual JsonInputSchemaMappingProperties
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

// Generator of JsonInputSchemaMappingProperties instances for property testing - lazily instantiated by
// JsonInputSchemaMappingPropertiesGenerator()
var jsonInputSchemaMappingPropertiesGenerator gopter.Gen

// JsonInputSchemaMappingPropertiesGenerator returns a generator of JsonInputSchemaMappingProperties instances for property testing.
func JsonInputSchemaMappingPropertiesGenerator() gopter.Gen {
	if jsonInputSchemaMappingPropertiesGenerator != nil {
		return jsonInputSchemaMappingPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForJsonInputSchemaMappingProperties(generators)
	jsonInputSchemaMappingPropertiesGenerator = gen.Struct(reflect.TypeOf(JsonInputSchemaMappingProperties{}), generators)

	return jsonInputSchemaMappingPropertiesGenerator
}

// AddRelatedPropertyGeneratorsForJsonInputSchemaMappingProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForJsonInputSchemaMappingProperties(gens map[string]gopter.Gen) {
	gens["DataVersion"] = gen.PtrOf(JsonFieldWithDefaultGenerator())
	gens["EventTime"] = gen.PtrOf(JsonFieldGenerator())
	gens["EventType"] = gen.PtrOf(JsonFieldWithDefaultGenerator())
	gens["Id"] = gen.PtrOf(JsonFieldGenerator())
	gens["Subject"] = gen.PtrOf(JsonFieldWithDefaultGenerator())
	gens["Topic"] = gen.PtrOf(JsonFieldGenerator())
}

func Test_JsonField_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of JsonField via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForJsonField, JsonFieldGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForJsonField runs a test to see if a specific instance of JsonField round trips to JSON and back losslessly
func RunJSONSerializationTestForJsonField(subject JsonField) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual JsonField
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

// Generator of JsonField instances for property testing - lazily instantiated by JsonFieldGenerator()
var jsonFieldGenerator gopter.Gen

// JsonFieldGenerator returns a generator of JsonField instances for property testing.
func JsonFieldGenerator() gopter.Gen {
	if jsonFieldGenerator != nil {
		return jsonFieldGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForJsonField(generators)
	jsonFieldGenerator = gen.Struct(reflect.TypeOf(JsonField{}), generators)

	return jsonFieldGenerator
}

// AddIndependentPropertyGeneratorsForJsonField is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForJsonField(gens map[string]gopter.Gen) {
	gens["SourceField"] = gen.PtrOf(gen.AlphaString())
}

func Test_JsonFieldWithDefault_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of JsonFieldWithDefault via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForJsonFieldWithDefault, JsonFieldWithDefaultGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForJsonFieldWithDefault runs a test to see if a specific instance of JsonFieldWithDefault round trips to JSON and back losslessly
func RunJSONSerializationTestForJsonFieldWithDefault(subject JsonFieldWithDefault) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual JsonFieldWithDefault
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

// Generator of JsonFieldWithDefault instances for property testing - lazily instantiated by
// JsonFieldWithDefaultGenerator()
var jsonFieldWithDefaultGenerator gopter.Gen

// JsonFieldWithDefaultGenerator returns a generator of JsonFieldWithDefault instances for property testing.
func JsonFieldWithDefaultGenerator() gopter.Gen {
	if jsonFieldWithDefaultGenerator != nil {
		return jsonFieldWithDefaultGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForJsonFieldWithDefault(generators)
	jsonFieldWithDefaultGenerator = gen.Struct(reflect.TypeOf(JsonFieldWithDefault{}), generators)

	return jsonFieldWithDefaultGenerator
}

// AddIndependentPropertyGeneratorsForJsonFieldWithDefault is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForJsonFieldWithDefault(gens map[string]gopter.Gen) {
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["SourceField"] = gen.PtrOf(gen.AlphaString())
}
