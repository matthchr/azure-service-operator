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

func Test_ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM, ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM runs a test to see if a specific instance of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM(subject ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM
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

// Generator of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM instances for property testing -
// lazily instantiated by ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator()
var managedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator gopter.Gen

// ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator returns a generator of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM instances for property testing.
// We first initialize managedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator() gopter.Gen {
	if managedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator != nil {
		return managedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM(generators)
	managedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM(generators)
	managedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM{}), generators)

	return managedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BaseLongTermRetentionPolicyProperties_STATUS_ARMGenerator())
}

func Test_BaseLongTermRetentionPolicyProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BaseLongTermRetentionPolicyProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBaseLongTermRetentionPolicyProperties_STATUS_ARM, BaseLongTermRetentionPolicyProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBaseLongTermRetentionPolicyProperties_STATUS_ARM runs a test to see if a specific instance of BaseLongTermRetentionPolicyProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBaseLongTermRetentionPolicyProperties_STATUS_ARM(subject BaseLongTermRetentionPolicyProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BaseLongTermRetentionPolicyProperties_STATUS_ARM
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

// Generator of BaseLongTermRetentionPolicyProperties_STATUS_ARM instances for property testing - lazily instantiated by
// BaseLongTermRetentionPolicyProperties_STATUS_ARMGenerator()
var baseLongTermRetentionPolicyProperties_STATUS_ARMGenerator gopter.Gen

// BaseLongTermRetentionPolicyProperties_STATUS_ARMGenerator returns a generator of BaseLongTermRetentionPolicyProperties_STATUS_ARM instances for property testing.
func BaseLongTermRetentionPolicyProperties_STATUS_ARMGenerator() gopter.Gen {
	if baseLongTermRetentionPolicyProperties_STATUS_ARMGenerator != nil {
		return baseLongTermRetentionPolicyProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBaseLongTermRetentionPolicyProperties_STATUS_ARM(generators)
	baseLongTermRetentionPolicyProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(BaseLongTermRetentionPolicyProperties_STATUS_ARM{}), generators)

	return baseLongTermRetentionPolicyProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBaseLongTermRetentionPolicyProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBaseLongTermRetentionPolicyProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["MonthlyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["WeekOfYear"] = gen.PtrOf(gen.Int())
	gens["WeeklyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["YearlyRetention"] = gen.PtrOf(gen.AlphaString())
}
