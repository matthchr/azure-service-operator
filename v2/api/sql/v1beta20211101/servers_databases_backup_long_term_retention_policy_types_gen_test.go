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

func Test_ServersDatabasesBackupLongTermRetentionPolicy_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabasesBackupLongTermRetentionPolicy to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersDatabasesBackupLongTermRetentionPolicy, ServersDatabasesBackupLongTermRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersDatabasesBackupLongTermRetentionPolicy tests if a specific instance of ServersDatabasesBackupLongTermRetentionPolicy round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersDatabasesBackupLongTermRetentionPolicy(subject ServersDatabasesBackupLongTermRetentionPolicy) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ServersDatabasesBackupLongTermRetentionPolicy
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersDatabasesBackupLongTermRetentionPolicy
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

func Test_ServersDatabasesBackupLongTermRetentionPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabasesBackupLongTermRetentionPolicy to ServersDatabasesBackupLongTermRetentionPolicy via AssignProperties_To_ServersDatabasesBackupLongTermRetentionPolicy & AssignProperties_From_ServersDatabasesBackupLongTermRetentionPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersDatabasesBackupLongTermRetentionPolicy, ServersDatabasesBackupLongTermRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersDatabasesBackupLongTermRetentionPolicy tests if a specific instance of ServersDatabasesBackupLongTermRetentionPolicy can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServersDatabasesBackupLongTermRetentionPolicy(subject ServersDatabasesBackupLongTermRetentionPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ServersDatabasesBackupLongTermRetentionPolicy
	err := copied.AssignProperties_To_ServersDatabasesBackupLongTermRetentionPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersDatabasesBackupLongTermRetentionPolicy
	err = actual.AssignProperties_From_ServersDatabasesBackupLongTermRetentionPolicy(&other)
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

func Test_ServersDatabasesBackupLongTermRetentionPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesBackupLongTermRetentionPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesBackupLongTermRetentionPolicy, ServersDatabasesBackupLongTermRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesBackupLongTermRetentionPolicy runs a test to see if a specific instance of ServersDatabasesBackupLongTermRetentionPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesBackupLongTermRetentionPolicy(subject ServersDatabasesBackupLongTermRetentionPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesBackupLongTermRetentionPolicy
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

// Generator of ServersDatabasesBackupLongTermRetentionPolicy instances for property testing - lazily instantiated by
// ServersDatabasesBackupLongTermRetentionPolicyGenerator()
var serversDatabasesBackupLongTermRetentionPolicyGenerator gopter.Gen

// ServersDatabasesBackupLongTermRetentionPolicyGenerator returns a generator of ServersDatabasesBackupLongTermRetentionPolicy instances for property testing.
func ServersDatabasesBackupLongTermRetentionPolicyGenerator() gopter.Gen {
	if serversDatabasesBackupLongTermRetentionPolicyGenerator != nil {
		return serversDatabasesBackupLongTermRetentionPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersDatabasesBackupLongTermRetentionPolicy(generators)
	serversDatabasesBackupLongTermRetentionPolicyGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesBackupLongTermRetentionPolicy{}), generators)

	return serversDatabasesBackupLongTermRetentionPolicyGenerator
}

// AddRelatedPropertyGeneratorsForServersDatabasesBackupLongTermRetentionPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesBackupLongTermRetentionPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator()
	gens["Status"] = Servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator()
}

func Test_Servers_Databases_BackupLongTermRetentionPolicy_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Databases_BackupLongTermRetentionPolicy_Spec to Servers_Databases_BackupLongTermRetentionPolicy_Spec via AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_Spec & AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Databases_BackupLongTermRetentionPolicy_Spec, Servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Databases_BackupLongTermRetentionPolicy_Spec tests if a specific instance of Servers_Databases_BackupLongTermRetentionPolicy_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_Databases_BackupLongTermRetentionPolicy_Spec(subject Servers_Databases_BackupLongTermRetentionPolicy_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Databases_BackupLongTermRetentionPolicy_Spec
	err := copied.AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Databases_BackupLongTermRetentionPolicy_Spec
	err = actual.AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_Spec(&other)
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

func Test_Servers_Databases_BackupLongTermRetentionPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_BackupLongTermRetentionPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_BackupLongTermRetentionPolicy_Spec, Servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_BackupLongTermRetentionPolicy_Spec runs a test to see if a specific instance of Servers_Databases_BackupLongTermRetentionPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_BackupLongTermRetentionPolicy_Spec(subject Servers_Databases_BackupLongTermRetentionPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_BackupLongTermRetentionPolicy_Spec
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

// Generator of Servers_Databases_BackupLongTermRetentionPolicy_Spec instances for property testing - lazily
// instantiated by Servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator()
var servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator gopter.Gen

// Servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator returns a generator of Servers_Databases_BackupLongTermRetentionPolicy_Spec instances for property testing.
func Servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator() gopter.Gen {
	if servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator != nil {
		return servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_BackupLongTermRetentionPolicy_Spec(generators)
	servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_BackupLongTermRetentionPolicy_Spec{}), generators)

	return servers_Databases_BackupLongTermRetentionPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_BackupLongTermRetentionPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_BackupLongTermRetentionPolicy_Spec(gens map[string]gopter.Gen) {
	gens["MonthlyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["WeekOfYear"] = gen.PtrOf(gen.Int())
	gens["WeeklyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["YearlyRetention"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Databases_BackupLongTermRetentionPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Databases_BackupLongTermRetentionPolicy_STATUS to Servers_Databases_BackupLongTermRetentionPolicy_STATUS via AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_STATUS & AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Databases_BackupLongTermRetentionPolicy_STATUS, Servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Databases_BackupLongTermRetentionPolicy_STATUS tests if a specific instance of Servers_Databases_BackupLongTermRetentionPolicy_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_Databases_BackupLongTermRetentionPolicy_STATUS(subject Servers_Databases_BackupLongTermRetentionPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Databases_BackupLongTermRetentionPolicy_STATUS
	err := copied.AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Databases_BackupLongTermRetentionPolicy_STATUS
	err = actual.AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(&other)
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

func Test_Servers_Databases_BackupLongTermRetentionPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_BackupLongTermRetentionPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_BackupLongTermRetentionPolicy_STATUS, Servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_BackupLongTermRetentionPolicy_STATUS runs a test to see if a specific instance of Servers_Databases_BackupLongTermRetentionPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_BackupLongTermRetentionPolicy_STATUS(subject Servers_Databases_BackupLongTermRetentionPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_BackupLongTermRetentionPolicy_STATUS
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

// Generator of Servers_Databases_BackupLongTermRetentionPolicy_STATUS instances for property testing - lazily
// instantiated by Servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator()
var servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator gopter.Gen

// Servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator returns a generator of Servers_Databases_BackupLongTermRetentionPolicy_STATUS instances for property testing.
func Servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator() gopter.Gen {
	if servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator != nil {
		return servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_BackupLongTermRetentionPolicy_STATUS(generators)
	servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_BackupLongTermRetentionPolicy_STATUS{}), generators)

	return servers_Databases_BackupLongTermRetentionPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_BackupLongTermRetentionPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_BackupLongTermRetentionPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["MonthlyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["WeekOfYear"] = gen.PtrOf(gen.Int())
	gens["WeeklyRetention"] = gen.PtrOf(gen.AlphaString())
	gens["YearlyRetention"] = gen.PtrOf(gen.AlphaString())
}
