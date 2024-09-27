/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

// TODO: Webhook tests?
// - Ensure that we don't have multiple overlapping writes?
// - Ensure CEL expression is parsable?
// TODO: Different version tests
// TODO: Secret tests

func Test_ExportCELConfigMap(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	principalIdKey := "principalId"

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMapExpressions: []*core.DestinationExpression{
					{
						Name:  configMapName,
						Key:   principalIdKey,
						Value: "self.status.principalId",
					},
					{
						Name:  configMapName,
						Value: `{"greeting": "hello", "classification": "friend"}`,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	// The ConfigMap should exist with the expected value
	tc.ExpectConfigMapHasKeysAndValues(
		configMapName,
		"principalId", *mi.Status.PrincipalId,
		"greeting", "hello",
		"classification", "friend")
}

func Test_ExportCELSecret(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	secretName := "my-configmap"
	secretKey := "key1"

	// Create a storage account and export the secret
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			Sku: &storage.Sku{
				Name: to.Ptr(storage.SkuName_Standard_LRS),
			},
			AccessTier: to.Ptr(storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot),
			OperatorSpec: &storage.StorageAccountOperatorSpec{
				Secrets: &storage.StorageAccountOperatorSecrets{
					Key1: &genruntime.SecretDestination{
						Name: secretName,
						Key:  secretKey,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(acct)

	// The secret should exist with the expected value
	tc.ExpectSecretHasKeys("key1")
}
