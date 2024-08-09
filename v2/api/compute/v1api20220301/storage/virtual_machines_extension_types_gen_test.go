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

func Test_InstanceViewStatus_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InstanceViewStatus via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInstanceViewStatus, InstanceViewStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInstanceViewStatus runs a test to see if a specific instance of InstanceViewStatus round trips to JSON and back losslessly
func RunJSONSerializationTestForInstanceViewStatus(subject InstanceViewStatus) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InstanceViewStatus
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

// Generator of InstanceViewStatus instances for property testing - lazily instantiated by InstanceViewStatusGenerator()
var instanceViewStatusGenerator gopter.Gen

// InstanceViewStatusGenerator returns a generator of InstanceViewStatus instances for property testing.
func InstanceViewStatusGenerator() gopter.Gen {
	if instanceViewStatusGenerator != nil {
		return instanceViewStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInstanceViewStatus(generators)
	instanceViewStatusGenerator = gen.Struct(reflect.TypeOf(InstanceViewStatus{}), generators)

	return instanceViewStatusGenerator
}

// AddIndependentPropertyGeneratorsForInstanceViewStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInstanceViewStatus(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Level"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Time"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualMachineExtensionInstanceView_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineExtensionInstanceView via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineExtensionInstanceView, VirtualMachineExtensionInstanceViewGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineExtensionInstanceView runs a test to see if a specific instance of VirtualMachineExtensionInstanceView round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineExtensionInstanceView(subject VirtualMachineExtensionInstanceView) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineExtensionInstanceView
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

// Generator of VirtualMachineExtensionInstanceView instances for property testing - lazily instantiated by
// VirtualMachineExtensionInstanceViewGenerator()
var virtualMachineExtensionInstanceViewGenerator gopter.Gen

// VirtualMachineExtensionInstanceViewGenerator returns a generator of VirtualMachineExtensionInstanceView instances for property testing.
// We first initialize virtualMachineExtensionInstanceViewGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineExtensionInstanceViewGenerator() gopter.Gen {
	if virtualMachineExtensionInstanceViewGenerator != nil {
		return virtualMachineExtensionInstanceViewGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView(generators)
	virtualMachineExtensionInstanceViewGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView(generators)
	virtualMachineExtensionInstanceViewGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView{}), generators)

	return virtualMachineExtensionInstanceViewGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView(gens map[string]gopter.Gen) {
	gens["Statuses"] = gen.SliceOf(InstanceViewStatusGenerator())
	gens["Substatuses"] = gen.SliceOf(InstanceViewStatusGenerator())
}

func Test_VirtualMachineExtensionInstanceView_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineExtensionInstanceView_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS, VirtualMachineExtensionInstanceView_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS runs a test to see if a specific instance of VirtualMachineExtensionInstanceView_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS(subject VirtualMachineExtensionInstanceView_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineExtensionInstanceView_STATUS
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

// Generator of VirtualMachineExtensionInstanceView_STATUS instances for property testing - lazily instantiated by
// VirtualMachineExtensionInstanceView_STATUSGenerator()
var virtualMachineExtensionInstanceView_STATUSGenerator gopter.Gen

// VirtualMachineExtensionInstanceView_STATUSGenerator returns a generator of VirtualMachineExtensionInstanceView_STATUS instances for property testing.
// We first initialize virtualMachineExtensionInstanceView_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineExtensionInstanceView_STATUSGenerator() gopter.Gen {
	if virtualMachineExtensionInstanceView_STATUSGenerator != nil {
		return virtualMachineExtensionInstanceView_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(generators)
	virtualMachineExtensionInstanceView_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(generators)
	virtualMachineExtensionInstanceView_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView_STATUS{}), generators)

	return virtualMachineExtensionInstanceView_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(gens map[string]gopter.Gen) {
	gens["Statuses"] = gen.SliceOf(InstanceViewStatus_STATUSGenerator())
	gens["Substatuses"] = gen.SliceOf(InstanceViewStatus_STATUSGenerator())
}

func Test_VirtualMachinesExtension_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachinesExtension via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachinesExtension, VirtualMachinesExtensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachinesExtension runs a test to see if a specific instance of VirtualMachinesExtension round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachinesExtension(subject VirtualMachinesExtension) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachinesExtension
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

// Generator of VirtualMachinesExtension instances for property testing - lazily instantiated by
// VirtualMachinesExtensionGenerator()
var virtualMachinesExtensionGenerator gopter.Gen

// VirtualMachinesExtensionGenerator returns a generator of VirtualMachinesExtension instances for property testing.
func VirtualMachinesExtensionGenerator() gopter.Gen {
	if virtualMachinesExtensionGenerator != nil {
		return virtualMachinesExtensionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForVirtualMachinesExtension(generators)
	virtualMachinesExtensionGenerator = gen.Struct(reflect.TypeOf(VirtualMachinesExtension{}), generators)

	return virtualMachinesExtensionGenerator
}

// AddRelatedPropertyGeneratorsForVirtualMachinesExtension is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachinesExtension(gens map[string]gopter.Gen) {
	gens["Spec"] = VirtualMachinesExtension_SpecGenerator()
	gens["Status"] = VirtualMachinesExtension_STATUSGenerator()
}

func Test_VirtualMachinesExtensionOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachinesExtensionOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachinesExtensionOperatorSpec, VirtualMachinesExtensionOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachinesExtensionOperatorSpec runs a test to see if a specific instance of VirtualMachinesExtensionOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachinesExtensionOperatorSpec(subject VirtualMachinesExtensionOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachinesExtensionOperatorSpec
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

// Generator of VirtualMachinesExtensionOperatorSpec instances for property testing - lazily instantiated by
// VirtualMachinesExtensionOperatorSpecGenerator()
var virtualMachinesExtensionOperatorSpecGenerator gopter.Gen

// VirtualMachinesExtensionOperatorSpecGenerator returns a generator of VirtualMachinesExtensionOperatorSpec instances for property testing.
func VirtualMachinesExtensionOperatorSpecGenerator() gopter.Gen {
	if virtualMachinesExtensionOperatorSpecGenerator != nil {
		return virtualMachinesExtensionOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	virtualMachinesExtensionOperatorSpecGenerator = gen.Struct(reflect.TypeOf(VirtualMachinesExtensionOperatorSpec{}), generators)

	return virtualMachinesExtensionOperatorSpecGenerator
}

func Test_VirtualMachinesExtension_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachinesExtension_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachinesExtension_STATUS, VirtualMachinesExtension_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachinesExtension_STATUS runs a test to see if a specific instance of VirtualMachinesExtension_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachinesExtension_STATUS(subject VirtualMachinesExtension_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachinesExtension_STATUS
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

// Generator of VirtualMachinesExtension_STATUS instances for property testing - lazily instantiated by
// VirtualMachinesExtension_STATUSGenerator()
var virtualMachinesExtension_STATUSGenerator gopter.Gen

// VirtualMachinesExtension_STATUSGenerator returns a generator of VirtualMachinesExtension_STATUS instances for property testing.
// We first initialize virtualMachinesExtension_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachinesExtension_STATUSGenerator() gopter.Gen {
	if virtualMachinesExtension_STATUSGenerator != nil {
		return virtualMachinesExtension_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachinesExtension_STATUS(generators)
	virtualMachinesExtension_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachinesExtension_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachinesExtension_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualMachinesExtension_STATUS(generators)
	virtualMachinesExtension_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachinesExtension_STATUS{}), generators)

	return virtualMachinesExtension_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachinesExtension_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachinesExtension_STATUS(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["SuppressFailures"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachinesExtension_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachinesExtension_STATUS(gens map[string]gopter.Gen) {
	gens["InstanceView"] = gen.PtrOf(VirtualMachineExtensionInstanceView_STATUSGenerator())
	gens["ProtectedSettingsFromKeyVault"] = gen.PtrOf(KeyVaultSecretReference_STATUSGenerator())
}

func Test_VirtualMachinesExtension_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachinesExtension_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachinesExtension_Spec, VirtualMachinesExtension_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachinesExtension_Spec runs a test to see if a specific instance of VirtualMachinesExtension_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachinesExtension_Spec(subject VirtualMachinesExtension_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachinesExtension_Spec
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

// Generator of VirtualMachinesExtension_Spec instances for property testing - lazily instantiated by
// VirtualMachinesExtension_SpecGenerator()
var virtualMachinesExtension_SpecGenerator gopter.Gen

// VirtualMachinesExtension_SpecGenerator returns a generator of VirtualMachinesExtension_Spec instances for property testing.
// We first initialize virtualMachinesExtension_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachinesExtension_SpecGenerator() gopter.Gen {
	if virtualMachinesExtension_SpecGenerator != nil {
		return virtualMachinesExtension_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachinesExtension_Spec(generators)
	virtualMachinesExtension_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualMachinesExtension_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachinesExtension_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualMachinesExtension_Spec(generators)
	virtualMachinesExtension_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualMachinesExtension_Spec{}), generators)

	return virtualMachinesExtension_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachinesExtension_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachinesExtension_Spec(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["SuppressFailures"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachinesExtension_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachinesExtension_Spec(gens map[string]gopter.Gen) {
	gens["InstanceView"] = gen.PtrOf(VirtualMachineExtensionInstanceViewGenerator())
	gens["OperatorSpec"] = gen.PtrOf(VirtualMachinesExtensionOperatorSpecGenerator())
	gens["ProtectedSettingsFromKeyVault"] = gen.PtrOf(KeyVaultSecretReferenceGenerator())
}
