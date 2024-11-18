/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package generic

import (
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

type Builder struct {
	options Options
	//schemer Schemer,
	fieldIndexer         client.FieldIndexer
	armConnectionFactory arm.ARMConnectionFactory
	credentialProvider   identity.CredentialProvider
	kubeClient           kubeclient.Client
	positiveConditions   *conditions.PositiveConditionBuilder
	expressionEvaluator  asocel.ExpressionEvaluator

	objs []*registration.StorageType
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Options() Options {
	return b.options
}

func (b *Builder) FieldIndexer() client.FieldIndexer {
	return b.fieldIndexer
}

func (b *Builder) ARMConnectionFactory() client.FieldIndexer {
	return b.fieldIndexer
}

func (b *Builder) CredentialProvider() identity.CredentialProvider {
	return b.credentialProvider
}

func (b *Builder) KubeClient() kubeclient.Client {
	return b.kubeClient
}

func (b *Builder) ConditionBuilder() *conditions.PositiveConditionBuilder {
	return b.positiveConditions
}

func (b *Builder) ExpressionEvaluator() asocel.ExpressionEvaluator {
	return b.expressionEvaluator
}

func (b *Builder) WithOptions(options Options) *Builder {
	b.options = options
	return b
}

func (b *Builder) WithIndexer(indexer client.FieldIndexer) *Builder {
	b.fieldIndexer = indexer
	return b
}

func (b *Builder) WithARMConnectionFactory(armConnectionFactory arm.ARMConnectionFactory) *Builder {
	b.armConnectionFactory = armConnectionFactory
	return b
}

func (b *Builder) WithCredentialProvider(credentialProvider identity.CredentialProvider) *Builder {
	b.credentialProvider = credentialProvider
	return b
}

func (b *Builder) WithKubeClient(kubeClient kubeclient.Client) *Builder {
	b.kubeClient = kubeClient
	return b
}

func (b *Builder) WithConditionBuilder(conditionBuilder *conditions.PositiveConditionBuilder) *Builder {
	b.positiveConditions = conditionBuilder
	return b
}

func (b *Builder) WithExpressionEvaluator(expressionEvaluator asocel.ExpressionEvaluator) *Builder {
	b.expressionEvaluator = expressionEvaluator
	return b
}

func (b *Builder) Build() {

}
