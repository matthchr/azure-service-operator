/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	machinelearning "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20240401/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &machinelearning.WorkspaceOperatorSecrets{}
	reflecthelpers.PopulateStruct(secrets)

	obj := &machinelearning.Workspace{
		Spec: machinelearning.Workspace_Spec{
			OperatorSpec: &machinelearning.WorkspaceOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := secretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(machinelearning.WorkspaceOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
