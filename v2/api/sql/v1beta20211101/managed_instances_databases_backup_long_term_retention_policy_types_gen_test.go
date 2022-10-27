// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101storage"
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

func Test_ManagedInstances_Databases_BackupLongTermRetentionPolicy_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedInstances_Databases_BackupLongTermRetentionPolicy to hub returns original",
		prop.ForAll(RunResourceConversionTestForManagedInstances_Databases_BackupLongTermRetentionPolicy, ManagedInstances_Databases_BackupLongTermRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForManagedInstances_Databases_BackupLongTermRetentionPolicy tests if a specific instance of ManagedInstances_Databases_BackupLongTermRetentionPolicy round trips to the hub storage version and back losslessly
func RunResourceConversionTestForManagedInstances_Databases_BackupLongTermRetentionPolicy(subject ManagedInstances_Databases_BackupLongTermRetentionPolicy) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ManagedInstances_Databases_BackupLongTermRetentionPolicy
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ManagedInstances_Databases_BackupLongTermRetentionPolicy
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

func Test_ManagedInstances_Databases_BackupLongTermRetentionPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedInstances_Databases_BackupLongTermRetentionPolicy to ManagedInstances_Databases_BackupLongTermRetentionPolicy via AssignProperties_To_ManagedInstances_Databases_BackupLongTermRetentionPolicy & AssignProperties_From_ManagedInstances_Databases_BackupLongTermRetentionPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy, ManagedInstances_Databases_BackupLongTermRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy tests if a specific instance of ManagedInstances_Databases_BackupLongTermRetentionPolicy can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy(subject ManagedInstances_Databases_BackupLongTermRetentionPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ManagedInstances_Databases_BackupLongTermRetentionPolicy
	err := copied.AssignProperties_To_ManagedInstances_Databases_BackupLongTermRetentionPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedInstances_Databases_BackupLongTermRetentionPolicy
	err = actual.AssignProperties_From_ManagedInstances_Databases_BackupLongTermRetentionPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ManagedInstances_Databases_BackupLongTermRetentionPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedInstances_Databases_BackupLongTermRetentionPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy, ManagedInstances_Databases_BackupLongTermRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy runs a test to see if a specific instance of ManagedInstances_Databases_BackupLongTermRetentionPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy(subject ManagedInstances_Databases_BackupLongTermRetentionPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedInstances_Databases_BackupLongTermRetentionPolicy
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

// Generator of ManagedInstances_Databases_BackupLongTermRetentionPolicy instances for property testing - lazily
// instantiated by ManagedInstances_Databases_BackupLongTermRetentionPolicyGenerator()
var managedInstances_Databases_BackupLongTermRetentionPolicyGenerator gopter.Gen

// ManagedInstances_Databases_BackupLongTermRetentionPolicyGenerator returns a generator of ManagedInstances_Databases_BackupLongTermRetentionPolicy instances for property testing.
func ManagedInstances_Databases_BackupLongTermRetentionPolicyGenerator() gopter.Gen {
	if managedInstances_Databases_BackupLongTermRetentionPolicyGenerator != nil {
		return managedInstances_Databases_BackupLongTermRetentionPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy(generators)
	managedInstances_Databases_BackupLongTermRetentionPolicyGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_Databases_BackupLongTermRetentionPolicy{}), generators)

	return managedInstances_Databases_BackupLongTermRetentionPolicyGenerator
}

// AddRelatedPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = ManagedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator()
	gens["Status"] = ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator()
}

func Test_ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec to ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec via AssignProperties_To_ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec & AssignProperties_From_ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec, ManagedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec tests if a specific instance of ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec(subject ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec
	err := copied.AssignProperties_To_ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec
	err = actual.AssignProperties_From_ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec, ManagedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec runs a test to see if a specific instance of ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec(subject ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec
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

// Generator of ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec instances for property testing - lazily
// instantiated by ManagedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator()
var managedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator gopter.Gen

// ManagedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator returns a generator of ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec instances for property testing.
func ManagedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator() gopter.Gen {
	if managedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator != nil {
		return managedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec(generators)
	managedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec{}), generators)

	return managedInstances_Databases_BackupLongTermRetentionPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec(gens map[string]gopter.Gen) {
	gens["MonthlyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["WeekOfYear"] = gen.PtrOf(gen.Int())
	gens["WeeklyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["YearlyRetention"] = gen.PtrOf(gen.AlphaString())
}

func Test_ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS to ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS via AssignProperties_To_ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS & AssignProperties_From_ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS, ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS tests if a specific instance of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS(subject ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS
	err := copied.AssignProperties_To_ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS
	err = actual.AssignProperties_From_ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS, ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS runs a test to see if a specific instance of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS(subject ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS
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

// Generator of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS instances for property testing - lazily
// instantiated by ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator()
var managedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator gopter.Gen

// ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator returns a generator of ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS instances for property testing.
func ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator() gopter.Gen {
	if managedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator != nil {
		return managedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS(generators)
	managedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS{}), generators)

	return managedInstances_Databases_BackupLongTermRetentionPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedInstances_Databases_BackupLongTermRetentionPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["MonthlyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["WeekOfYear"] = gen.PtrOf(gen.Int())
	gens["WeeklyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["YearlyRetention"] = gen.PtrOf(gen.AlphaString())
}
