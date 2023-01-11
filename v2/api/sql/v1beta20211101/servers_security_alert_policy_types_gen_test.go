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

func Test_ServersSecurityAlertPolicy_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersSecurityAlertPolicy to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersSecurityAlertPolicy, ServersSecurityAlertPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersSecurityAlertPolicy tests if a specific instance of ServersSecurityAlertPolicy round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersSecurityAlertPolicy(subject ServersSecurityAlertPolicy) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ServersSecurityAlertPolicy
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersSecurityAlertPolicy
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

func Test_ServersSecurityAlertPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersSecurityAlertPolicy to ServersSecurityAlertPolicy via AssignProperties_To_ServersSecurityAlertPolicy & AssignProperties_From_ServersSecurityAlertPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersSecurityAlertPolicy, ServersSecurityAlertPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersSecurityAlertPolicy tests if a specific instance of ServersSecurityAlertPolicy can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServersSecurityAlertPolicy(subject ServersSecurityAlertPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ServersSecurityAlertPolicy
	err := copied.AssignProperties_To_ServersSecurityAlertPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersSecurityAlertPolicy
	err = actual.AssignProperties_From_ServersSecurityAlertPolicy(&other)
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

func Test_ServersSecurityAlertPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersSecurityAlertPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersSecurityAlertPolicy, ServersSecurityAlertPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersSecurityAlertPolicy runs a test to see if a specific instance of ServersSecurityAlertPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForServersSecurityAlertPolicy(subject ServersSecurityAlertPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersSecurityAlertPolicy
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

// Generator of ServersSecurityAlertPolicy instances for property testing - lazily instantiated by
// ServersSecurityAlertPolicyGenerator()
var serversSecurityAlertPolicyGenerator gopter.Gen

// ServersSecurityAlertPolicyGenerator returns a generator of ServersSecurityAlertPolicy instances for property testing.
func ServersSecurityAlertPolicyGenerator() gopter.Gen {
	if serversSecurityAlertPolicyGenerator != nil {
		return serversSecurityAlertPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersSecurityAlertPolicy(generators)
	serversSecurityAlertPolicyGenerator = gen.Struct(reflect.TypeOf(ServersSecurityAlertPolicy{}), generators)

	return serversSecurityAlertPolicyGenerator
}

// AddRelatedPropertyGeneratorsForServersSecurityAlertPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersSecurityAlertPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_SecurityAlertPolicy_SpecGenerator()
	gens["Status"] = Servers_SecurityAlertPolicy_STATUSGenerator()
}

func Test_Servers_SecurityAlertPolicy_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_SecurityAlertPolicy_Spec to Servers_SecurityAlertPolicy_Spec via AssignProperties_To_Servers_SecurityAlertPolicy_Spec & AssignProperties_From_Servers_SecurityAlertPolicy_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_SecurityAlertPolicy_Spec, Servers_SecurityAlertPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_SecurityAlertPolicy_Spec tests if a specific instance of Servers_SecurityAlertPolicy_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_SecurityAlertPolicy_Spec(subject Servers_SecurityAlertPolicy_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_SecurityAlertPolicy_Spec
	err := copied.AssignProperties_To_Servers_SecurityAlertPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_SecurityAlertPolicy_Spec
	err = actual.AssignProperties_From_Servers_SecurityAlertPolicy_Spec(&other)
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

func Test_Servers_SecurityAlertPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_SecurityAlertPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_SecurityAlertPolicy_Spec, Servers_SecurityAlertPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_SecurityAlertPolicy_Spec runs a test to see if a specific instance of Servers_SecurityAlertPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_SecurityAlertPolicy_Spec(subject Servers_SecurityAlertPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_SecurityAlertPolicy_Spec
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

// Generator of Servers_SecurityAlertPolicy_Spec instances for property testing - lazily instantiated by
// Servers_SecurityAlertPolicy_SpecGenerator()
var servers_SecurityAlertPolicy_SpecGenerator gopter.Gen

// Servers_SecurityAlertPolicy_SpecGenerator returns a generator of Servers_SecurityAlertPolicy_Spec instances for property testing.
func Servers_SecurityAlertPolicy_SpecGenerator() gopter.Gen {
	if servers_SecurityAlertPolicy_SpecGenerator != nil {
		return servers_SecurityAlertPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_SecurityAlertPolicy_Spec(generators)
	servers_SecurityAlertPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_SecurityAlertPolicy_Spec{}), generators)

	return servers_SecurityAlertPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_SecurityAlertPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_SecurityAlertPolicy_Spec(gens map[string]gopter.Gen) {
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Disabled, ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Enabled))
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_SecurityAlertPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_SecurityAlertPolicy_STATUS to Servers_SecurityAlertPolicy_STATUS via AssignProperties_To_Servers_SecurityAlertPolicy_STATUS & AssignProperties_From_Servers_SecurityAlertPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_SecurityAlertPolicy_STATUS, Servers_SecurityAlertPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_SecurityAlertPolicy_STATUS tests if a specific instance of Servers_SecurityAlertPolicy_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_SecurityAlertPolicy_STATUS(subject Servers_SecurityAlertPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_SecurityAlertPolicy_STATUS
	err := copied.AssignProperties_To_Servers_SecurityAlertPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_SecurityAlertPolicy_STATUS
	err = actual.AssignProperties_From_Servers_SecurityAlertPolicy_STATUS(&other)
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

func Test_Servers_SecurityAlertPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_SecurityAlertPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_SecurityAlertPolicy_STATUS, Servers_SecurityAlertPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_SecurityAlertPolicy_STATUS runs a test to see if a specific instance of Servers_SecurityAlertPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_SecurityAlertPolicy_STATUS(subject Servers_SecurityAlertPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_SecurityAlertPolicy_STATUS
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

// Generator of Servers_SecurityAlertPolicy_STATUS instances for property testing - lazily instantiated by
// Servers_SecurityAlertPolicy_STATUSGenerator()
var servers_SecurityAlertPolicy_STATUSGenerator gopter.Gen

// Servers_SecurityAlertPolicy_STATUSGenerator returns a generator of Servers_SecurityAlertPolicy_STATUS instances for property testing.
// We first initialize servers_SecurityAlertPolicy_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_SecurityAlertPolicy_STATUSGenerator() gopter.Gen {
	if servers_SecurityAlertPolicy_STATUSGenerator != nil {
		return servers_SecurityAlertPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_SecurityAlertPolicy_STATUS(generators)
	servers_SecurityAlertPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_SecurityAlertPolicy_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_SecurityAlertPolicy_STATUS(generators)
	AddRelatedPropertyGeneratorsForServers_SecurityAlertPolicy_STATUS(generators)
	servers_SecurityAlertPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_SecurityAlertPolicy_STATUS{}), generators)

	return servers_SecurityAlertPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_SecurityAlertPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_SecurityAlertPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["CreationTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS_Disabled, ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS_Enabled))
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_SecurityAlertPolicy_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_SecurityAlertPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
