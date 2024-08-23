// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

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

func Test_CopyCompletionError_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CopyCompletionError_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCopyCompletionError_ARM, CopyCompletionError_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCopyCompletionError_ARM runs a test to see if a specific instance of CopyCompletionError_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCopyCompletionError_ARM(subject CopyCompletionError_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CopyCompletionError_ARM
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

// Generator of CopyCompletionError_ARM instances for property testing - lazily instantiated by
// CopyCompletionError_ARMGenerator()
var copyCompletionError_ARMGenerator gopter.Gen

// CopyCompletionError_ARMGenerator returns a generator of CopyCompletionError_ARM instances for property testing.
func CopyCompletionError_ARMGenerator() gopter.Gen {
	if copyCompletionError_ARMGenerator != nil {
		return copyCompletionError_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCopyCompletionError_ARM(generators)
	copyCompletionError_ARMGenerator = gen.Struct(reflect.TypeOf(CopyCompletionError_ARM{}), generators)

	return copyCompletionError_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCopyCompletionError_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCopyCompletionError_ARM(gens map[string]gopter.Gen) {
	gens["ErrorCode"] = gen.PtrOf(gen.OneConstOf(CopyCompletionError_ErrorCode_CopySourceNotFound))
	gens["ErrorMessage"] = gen.PtrOf(gen.AlphaString())
}

func Test_SnapshotProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotProperties_ARM, SnapshotProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotProperties_ARM runs a test to see if a specific instance of SnapshotProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotProperties_ARM(subject SnapshotProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotProperties_ARM
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

// Generator of SnapshotProperties_ARM instances for property testing - lazily instantiated by
// SnapshotProperties_ARMGenerator()
var snapshotProperties_ARMGenerator gopter.Gen

// SnapshotProperties_ARMGenerator returns a generator of SnapshotProperties_ARM instances for property testing.
// We first initialize snapshotProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotProperties_ARMGenerator() gopter.Gen {
	if snapshotProperties_ARMGenerator != nil {
		return snapshotProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotProperties_ARM(generators)
	snapshotProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SnapshotProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForSnapshotProperties_ARM(generators)
	snapshotProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SnapshotProperties_ARM{}), generators)

	return snapshotProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotProperties_ARM(gens map[string]gopter.Gen) {
	gens["CompletionPercent"] = gen.PtrOf(gen.Float64())
	gens["DataAccessAuthMode"] = gen.PtrOf(gen.OneConstOf(DataAccessAuthMode_AzureActiveDirectory, DataAccessAuthMode_None))
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.OneConstOf(
		DiskState_ActiveSAS,
		DiskState_ActiveSASFrozen,
		DiskState_ActiveUpload,
		DiskState_Attached,
		DiskState_Frozen,
		DiskState_ReadyToUpload,
		DiskState_Reserved,
		DiskState_Unattached))
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(SnapshotProperties_HyperVGeneration_V1, SnapshotProperties_HyperVGeneration_V2))
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(NetworkAccessPolicy_AllowAll, NetworkAccessPolicy_AllowPrivate, NetworkAccessPolicy_DenyAll))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(SnapshotProperties_OsType_Linux, SnapshotProperties_OsType_Windows))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccess_Disabled, PublicNetworkAccess_Enabled))
	gens["SupportsHibernation"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForSnapshotProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotProperties_ARM(gens map[string]gopter.Gen) {
	gens["CopyCompletionError"] = gen.PtrOf(CopyCompletionError_ARMGenerator())
	gens["CreationData"] = gen.PtrOf(CreationData_ARMGenerator())
	gens["Encryption"] = gen.PtrOf(Encryption_ARMGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollection_ARMGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlan_ARMGenerator())
	gens["SecurityProfile"] = gen.PtrOf(DiskSecurityProfile_ARMGenerator())
	gens["SupportedCapabilities"] = gen.PtrOf(SupportedCapabilities_ARMGenerator())
}

func Test_SnapshotSku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSku_ARM, SnapshotSku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSku_ARM runs a test to see if a specific instance of SnapshotSku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSku_ARM(subject SnapshotSku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku_ARM
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

// Generator of SnapshotSku_ARM instances for property testing - lazily instantiated by SnapshotSku_ARMGenerator()
var snapshotSku_ARMGenerator gopter.Gen

// SnapshotSku_ARMGenerator returns a generator of SnapshotSku_ARM instances for property testing.
func SnapshotSku_ARMGenerator() gopter.Gen {
	if snapshotSku_ARMGenerator != nil {
		return snapshotSku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSku_ARM(generators)
	snapshotSku_ARMGenerator = gen.Struct(reflect.TypeOf(SnapshotSku_ARM{}), generators)

	return snapshotSku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSku_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SnapshotSku_Name_Premium_LRS, SnapshotSku_Name_Standard_LRS, SnapshotSku_Name_Standard_ZRS))
}

func Test_Snapshot_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot_Spec_ARM, Snapshot_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot_Spec_ARM runs a test to see if a specific instance of Snapshot_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot_Spec_ARM(subject Snapshot_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_Spec_ARM
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

// Generator of Snapshot_Spec_ARM instances for property testing - lazily instantiated by Snapshot_Spec_ARMGenerator()
var snapshot_Spec_ARMGenerator gopter.Gen

// Snapshot_Spec_ARMGenerator returns a generator of Snapshot_Spec_ARM instances for property testing.
// We first initialize snapshot_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Snapshot_Spec_ARMGenerator() gopter.Gen {
	if snapshot_Spec_ARMGenerator != nil {
		return snapshot_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_Spec_ARM(generators)
	snapshot_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Snapshot_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForSnapshot_Spec_ARM(generators)
	snapshot_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Snapshot_Spec_ARM{}), generators)

	return snapshot_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshot_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshot_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshot_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_ARMGenerator())
	gens["Properties"] = gen.PtrOf(SnapshotProperties_ARMGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSku_ARMGenerator())
}