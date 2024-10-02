/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	containerservice "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &containerservice.ManagedClusterOperatorSecrets{}
	reflecthelpers.PopulateStruct(secrets)

	obj := &containerservice.ManagedCluster{
		Spec: containerservice.ManagedCluster_Spec{
			OperatorSpec: &containerservice.ManagedClusterOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := secretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(containerservice.ManagedClusterOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
