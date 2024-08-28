// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

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

func Test_BackupShortTermRetentionPolicyProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupShortTermRetentionPolicyProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupShortTermRetentionPolicyProperties_STATUS_ARM, BackupShortTermRetentionPolicyProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupShortTermRetentionPolicyProperties_STATUS_ARM runs a test to see if a specific instance of BackupShortTermRetentionPolicyProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupShortTermRetentionPolicyProperties_STATUS_ARM(subject BackupShortTermRetentionPolicyProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupShortTermRetentionPolicyProperties_STATUS_ARM
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

// Generator of BackupShortTermRetentionPolicyProperties_STATUS_ARM instances for property testing - lazily instantiated
// by BackupShortTermRetentionPolicyProperties_STATUS_ARMGenerator()
var backupShortTermRetentionPolicyProperties_STATUS_ARMGenerator gopter.Gen

// BackupShortTermRetentionPolicyProperties_STATUS_ARMGenerator returns a generator of BackupShortTermRetentionPolicyProperties_STATUS_ARM instances for property testing.
func BackupShortTermRetentionPolicyProperties_STATUS_ARMGenerator() gopter.Gen {
	if backupShortTermRetentionPolicyProperties_STATUS_ARMGenerator != nil {
		return backupShortTermRetentionPolicyProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupShortTermRetentionPolicyProperties_STATUS_ARM(generators)
	backupShortTermRetentionPolicyProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(BackupShortTermRetentionPolicyProperties_STATUS_ARM{}), generators)

	return backupShortTermRetentionPolicyProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackupShortTermRetentionPolicyProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupShortTermRetentionPolicyProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DiffBackupIntervalInHours"] = gen.PtrOf(gen.OneConstOf(BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS_ARM_12, BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS_ARM_24))
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
}

func Test_Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM, Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM runs a test to see if a specific instance of Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM(subject Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM
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

// Generator of Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM instances for property testing - lazily
// instantiated by Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator()
var servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator gopter.Gen

// Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator returns a generator of Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM instances for property testing.
// We first initialize servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator() gopter.Gen {
	if servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator != nil {
		return servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM(generators)
	servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM(generators)
	servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM{}), generators)

	return servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BackupShortTermRetentionPolicyProperties_STATUS_ARMGenerator())
}
