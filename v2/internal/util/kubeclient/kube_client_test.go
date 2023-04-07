/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package kubeclient_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-logr/logr"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

const testNamespace = "testnamespace"

func Test_ListChunked_OneChunk(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	kubeClient := NewKubeClient(newTestScheme())
	g.Expect(addSecrets(ctx, kubeClient, 100)).To(Succeed())

	list := v1.SecretList{}
	g.Expect(kubeClient.ListChunked(ctx, logr.Discard(), &list, client.Limit(200))).To(Succeed())
	g.Expect(list.Items).To(HaveLen(100))
}

// TODO: No point in testing smaller chunk sizes here because the fake client doesn't actually emulate
// TODO: chunking

func NewKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}

func makeSecrets(count int) []v1.Secret {
	results := make([]v1.Secret, 0, count)

	for i := 0; i < count; i++ {
		secret := v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("foo%d", i),
				Namespace: testNamespace,
			},
		}

		results = append(results, secret)
	}

	return results
}

func addSecrets(ctx context.Context, client kubeclient.Client, count int) error {
	secrets := makeSecrets(count)

	for _, secret := range secrets {
		err := client.Create(ctx, &secret)
		if err != nil {
			return err
		}
	}

	return nil
}

func newTestScheme() *runtime.Scheme {
	s := runtime.NewScheme()
	_ = apiextensions.AddToScheme(s)
	_ = v1.AddToScheme(s)

	return s
}
