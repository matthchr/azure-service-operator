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

func Test_ManagedInstances_ServerTrustCertificate_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedInstances_ServerTrustCertificate to hub returns original",
		prop.ForAll(RunResourceConversionTestForManagedInstances_ServerTrustCertificate, ManagedInstances_ServerTrustCertificateGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForManagedInstances_ServerTrustCertificate tests if a specific instance of ManagedInstances_ServerTrustCertificate round trips to the hub storage version and back losslessly
func RunResourceConversionTestForManagedInstances_ServerTrustCertificate(subject ManagedInstances_ServerTrustCertificate) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ManagedInstances_ServerTrustCertificate
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ManagedInstances_ServerTrustCertificate
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

func Test_ManagedInstances_ServerTrustCertificate_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedInstances_ServerTrustCertificate to ManagedInstances_ServerTrustCertificate via AssignProperties_To_ManagedInstances_ServerTrustCertificate & AssignProperties_From_ManagedInstances_ServerTrustCertificate returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate, ManagedInstances_ServerTrustCertificateGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate tests if a specific instance of ManagedInstances_ServerTrustCertificate can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate(subject ManagedInstances_ServerTrustCertificate) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ManagedInstances_ServerTrustCertificate
	err := copied.AssignProperties_To_ManagedInstances_ServerTrustCertificate(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedInstances_ServerTrustCertificate
	err = actual.AssignProperties_From_ManagedInstances_ServerTrustCertificate(&other)
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

func Test_ManagedInstances_ServerTrustCertificate_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedInstances_ServerTrustCertificate via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedInstances_ServerTrustCertificate, ManagedInstances_ServerTrustCertificateGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedInstances_ServerTrustCertificate runs a test to see if a specific instance of ManagedInstances_ServerTrustCertificate round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedInstances_ServerTrustCertificate(subject ManagedInstances_ServerTrustCertificate) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedInstances_ServerTrustCertificate
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

// Generator of ManagedInstances_ServerTrustCertificate instances for property testing - lazily instantiated by
// ManagedInstances_ServerTrustCertificateGenerator()
var managedInstances_ServerTrustCertificateGenerator gopter.Gen

// ManagedInstances_ServerTrustCertificateGenerator returns a generator of ManagedInstances_ServerTrustCertificate instances for property testing.
func ManagedInstances_ServerTrustCertificateGenerator() gopter.Gen {
	if managedInstances_ServerTrustCertificateGenerator != nil {
		return managedInstances_ServerTrustCertificateGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagedInstances_ServerTrustCertificate(generators)
	managedInstances_ServerTrustCertificateGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_ServerTrustCertificate{}), generators)

	return managedInstances_ServerTrustCertificateGenerator
}

// AddRelatedPropertyGeneratorsForManagedInstances_ServerTrustCertificate is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedInstances_ServerTrustCertificate(gens map[string]gopter.Gen) {
	gens["Spec"] = ManagedInstances_ServerTrustCertificate_SpecGenerator()
	gens["Status"] = ManagedInstances_ServerTrustCertificate_STATUSGenerator()
}

func Test_ManagedInstances_ServerTrustCertificate_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedInstances_ServerTrustCertificate_Spec to ManagedInstances_ServerTrustCertificate_Spec via AssignProperties_To_ManagedInstances_ServerTrustCertificate_Spec & AssignProperties_From_ManagedInstances_ServerTrustCertificate_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate_Spec, ManagedInstances_ServerTrustCertificate_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate_Spec tests if a specific instance of ManagedInstances_ServerTrustCertificate_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate_Spec(subject ManagedInstances_ServerTrustCertificate_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ManagedInstances_ServerTrustCertificate_Spec
	err := copied.AssignProperties_To_ManagedInstances_ServerTrustCertificate_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedInstances_ServerTrustCertificate_Spec
	err = actual.AssignProperties_From_ManagedInstances_ServerTrustCertificate_Spec(&other)
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

func Test_ManagedInstances_ServerTrustCertificate_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedInstances_ServerTrustCertificate_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedInstances_ServerTrustCertificate_Spec, ManagedInstances_ServerTrustCertificate_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedInstances_ServerTrustCertificate_Spec runs a test to see if a specific instance of ManagedInstances_ServerTrustCertificate_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedInstances_ServerTrustCertificate_Spec(subject ManagedInstances_ServerTrustCertificate_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedInstances_ServerTrustCertificate_Spec
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

// Generator of ManagedInstances_ServerTrustCertificate_Spec instances for property testing - lazily instantiated by
// ManagedInstances_ServerTrustCertificate_SpecGenerator()
var managedInstances_ServerTrustCertificate_SpecGenerator gopter.Gen

// ManagedInstances_ServerTrustCertificate_SpecGenerator returns a generator of ManagedInstances_ServerTrustCertificate_Spec instances for property testing.
func ManagedInstances_ServerTrustCertificate_SpecGenerator() gopter.Gen {
	if managedInstances_ServerTrustCertificate_SpecGenerator != nil {
		return managedInstances_ServerTrustCertificate_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedInstances_ServerTrustCertificate_Spec(generators)
	managedInstances_ServerTrustCertificate_SpecGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_ServerTrustCertificate_Spec{}), generators)

	return managedInstances_ServerTrustCertificate_SpecGenerator
}

// AddIndependentPropertyGeneratorsForManagedInstances_ServerTrustCertificate_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedInstances_ServerTrustCertificate_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["PublicBlob"] = gen.PtrOf(gen.AlphaString())
}

func Test_ManagedInstances_ServerTrustCertificate_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedInstances_ServerTrustCertificate_STATUS to ManagedInstances_ServerTrustCertificate_STATUS via AssignProperties_To_ManagedInstances_ServerTrustCertificate_STATUS & AssignProperties_From_ManagedInstances_ServerTrustCertificate_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate_STATUS, ManagedInstances_ServerTrustCertificate_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate_STATUS tests if a specific instance of ManagedInstances_ServerTrustCertificate_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForManagedInstances_ServerTrustCertificate_STATUS(subject ManagedInstances_ServerTrustCertificate_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ManagedInstances_ServerTrustCertificate_STATUS
	err := copied.AssignProperties_To_ManagedInstances_ServerTrustCertificate_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedInstances_ServerTrustCertificate_STATUS
	err = actual.AssignProperties_From_ManagedInstances_ServerTrustCertificate_STATUS(&other)
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

func Test_ManagedInstances_ServerTrustCertificate_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedInstances_ServerTrustCertificate_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedInstances_ServerTrustCertificate_STATUS, ManagedInstances_ServerTrustCertificate_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedInstances_ServerTrustCertificate_STATUS runs a test to see if a specific instance of ManagedInstances_ServerTrustCertificate_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedInstances_ServerTrustCertificate_STATUS(subject ManagedInstances_ServerTrustCertificate_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedInstances_ServerTrustCertificate_STATUS
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

// Generator of ManagedInstances_ServerTrustCertificate_STATUS instances for property testing - lazily instantiated by
// ManagedInstances_ServerTrustCertificate_STATUSGenerator()
var managedInstances_ServerTrustCertificate_STATUSGenerator gopter.Gen

// ManagedInstances_ServerTrustCertificate_STATUSGenerator returns a generator of ManagedInstances_ServerTrustCertificate_STATUS instances for property testing.
func ManagedInstances_ServerTrustCertificate_STATUSGenerator() gopter.Gen {
	if managedInstances_ServerTrustCertificate_STATUSGenerator != nil {
		return managedInstances_ServerTrustCertificate_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedInstances_ServerTrustCertificate_STATUS(generators)
	managedInstances_ServerTrustCertificate_STATUSGenerator = gen.Struct(reflect.TypeOf(ManagedInstances_ServerTrustCertificate_STATUS{}), generators)

	return managedInstances_ServerTrustCertificate_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForManagedInstances_ServerTrustCertificate_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedInstances_ServerTrustCertificate_STATUS(gens map[string]gopter.Gen) {
	gens["CertificateName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PublicBlob"] = gen.PtrOf(gen.AlphaString())
	gens["Thumbprint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
