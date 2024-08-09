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

func Test_ProductApi_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductApi via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductApi, ProductApiGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductApi runs a test to see if a specific instance of ProductApi round trips to JSON and back losslessly
func RunJSONSerializationTestForProductApi(subject ProductApi) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductApi
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

// Generator of ProductApi instances for property testing - lazily instantiated by ProductApiGenerator()
var productApiGenerator gopter.Gen

// ProductApiGenerator returns a generator of ProductApi instances for property testing.
func ProductApiGenerator() gopter.Gen {
	if productApiGenerator != nil {
		return productApiGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProductApi(generators)
	productApiGenerator = gen.Struct(reflect.TypeOf(ProductApi{}), generators)

	return productApiGenerator
}

// AddRelatedPropertyGeneratorsForProductApi is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProductApi(gens map[string]gopter.Gen) {
	gens["Spec"] = ProductApi_SpecGenerator()
	gens["Status"] = ProductApi_STATUSGenerator()
}

func Test_ProductApiOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductApiOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductApiOperatorSpec, ProductApiOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductApiOperatorSpec runs a test to see if a specific instance of ProductApiOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForProductApiOperatorSpec(subject ProductApiOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductApiOperatorSpec
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

// Generator of ProductApiOperatorSpec instances for property testing - lazily instantiated by
// ProductApiOperatorSpecGenerator()
var productApiOperatorSpecGenerator gopter.Gen

// ProductApiOperatorSpecGenerator returns a generator of ProductApiOperatorSpec instances for property testing.
func ProductApiOperatorSpecGenerator() gopter.Gen {
	if productApiOperatorSpecGenerator != nil {
		return productApiOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	productApiOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ProductApiOperatorSpec{}), generators)

	return productApiOperatorSpecGenerator
}

func Test_ProductApi_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductApi_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductApi_STATUS, ProductApi_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductApi_STATUS runs a test to see if a specific instance of ProductApi_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProductApi_STATUS(subject ProductApi_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductApi_STATUS
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

// Generator of ProductApi_STATUS instances for property testing - lazily instantiated by ProductApi_STATUSGenerator()
var productApi_STATUSGenerator gopter.Gen

// ProductApi_STATUSGenerator returns a generator of ProductApi_STATUS instances for property testing.
func ProductApi_STATUSGenerator() gopter.Gen {
	if productApi_STATUSGenerator != nil {
		return productApi_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	productApi_STATUSGenerator = gen.Struct(reflect.TypeOf(ProductApi_STATUS{}), generators)

	return productApi_STATUSGenerator
}

func Test_ProductApi_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductApi_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductApi_Spec, ProductApi_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductApi_Spec runs a test to see if a specific instance of ProductApi_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForProductApi_Spec(subject ProductApi_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductApi_Spec
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

// Generator of ProductApi_Spec instances for property testing - lazily instantiated by ProductApi_SpecGenerator()
var productApi_SpecGenerator gopter.Gen

// ProductApi_SpecGenerator returns a generator of ProductApi_Spec instances for property testing.
// We first initialize productApi_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProductApi_SpecGenerator() gopter.Gen {
	if productApi_SpecGenerator != nil {
		return productApi_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductApi_Spec(generators)
	productApi_SpecGenerator = gen.Struct(reflect.TypeOf(ProductApi_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductApi_Spec(generators)
	AddRelatedPropertyGeneratorsForProductApi_Spec(generators)
	productApi_SpecGenerator = gen.Struct(reflect.TypeOf(ProductApi_Spec{}), generators)

	return productApi_SpecGenerator
}

// AddIndependentPropertyGeneratorsForProductApi_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductApi_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForProductApi_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProductApi_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ProductApiOperatorSpecGenerator())
}
