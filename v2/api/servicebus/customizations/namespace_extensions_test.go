/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
)

func Test_NamespaceSecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &servicebus.NamespaceOperatorSecrets{}
	reflecthelpers.PopulateStruct(secrets)

	obj := &servicebus.Namespace{
		Spec: servicebus.Namespace_Spec{
			OperatorSpec: &servicebus.NamespaceOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := namespaceSecretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(servicebus.NamespaceOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
