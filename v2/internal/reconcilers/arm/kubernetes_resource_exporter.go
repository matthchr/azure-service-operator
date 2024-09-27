/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

type kubernetesResourceExporter interface {
	Export(ctx context.Context) ([]client.Object, error)
}

var _ kubernetesResourceExporter = &configMapExpressionExporter{}

type configMapExpressionExporter struct {
	obj                 genruntime.ARMMetaObject
	versionedObj        genruntime.ARMMetaObject
	log                 logr.Logger
	expressionEvaluator asocel.ExpressionEvaluator
}

func (c *configMapExpressionExporter) Export(ctx context.Context) ([]client.Object, error) {
	cmExporter, ok := c.obj.(configmaps.Exporter)
	if !ok {
		return nil, nil
	}

	cmExpressions := cmExporter.ConfigMapDestinationExpressions()
	if cmExpressions == nil {
		return nil, nil
	}

	resources, err := c.parseConfigMaps(cmExpressions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse configmap expressions for export")
	}
	return resources, err
}

func (c *configMapExpressionExporter) parseConfigMaps(expressions []*core.DestinationExpression) ([]client.Object, error) {
	collector := configmaps.NewCollector(c.obj.GetNamespace())

	for _, expression := range expressions {
		value, err := c.expressionEvaluator.CompileAndRun(expression.Value, c.versionedObj, nil)
		if err != nil {
			return nil, err
		}

		if value.Value != "" {
			collector.AddValue(
				&genruntime.ConfigMapDestination{
					Name: expression.Name,
					Key:  expression.Key,
				}, value.Value)
		} else if len(value.Values) > 0 {
			for k, v := range value.Values {
				collector.AddValue(
					&genruntime.ConfigMapDestination{
						Name: expression.Name,
						Key:  k,
					}, v)
			}
		} else {
			return nil, errors.Errorf("unexpected expression output")
		}
	}

	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return configmaps.SliceToClientObjectSlice(result), nil
}

var _ kubernetesResourceExporter = &secretExpressionExporter{}

type secretExpressionExporter struct {
	obj                 genruntime.ARMMetaObject
	versionedObj        genruntime.ARMMetaObject
	log                 logr.Logger
	expressionEvaluator asocel.ExpressionEvaluator
}

func (s *secretExpressionExporter) Export(ctx context.Context) ([]client.Object, error) {
	secretExporter, ok := s.obj.(secrets.Exporter)
	if !ok {
		return nil, nil
	}

	secretExpressions := secretExporter.SecretDestinationExpressions()

	if secretExpressions == nil {
		return nil, nil
	}
	resources, err := s.parseSecrets(secretExpressions, s.versionedObj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse secret expressions for export")
	}

	return resources, nil
}

func (s *secretExpressionExporter) parseSecrets(expressions []*core.DestinationExpression, versionedObj genruntime.ARMMetaObject) ([]client.Object, error) {
	collector := secrets.NewCollector(versionedObj.GetNamespace())

	for _, expression := range expressions {
		// TODO: Secret needs to be specified here. Specifically, it needs to be a map[string]string where the keys are
		// TODO: the "name" of the secret and the values are the secret values themselves
		value, err := s.expressionEvaluator.CompileAndRun(expression.Value, versionedObj, nil) // TODO: Secret should not be nil here
		if err != nil {
			return nil, err
		}

		if value.Value != "" {
			collector.AddValue(
				&genruntime.SecretDestination{
					Name: expression.Name,
					Key:  expression.Key,
				}, value.Value)
		} else if len(value.Values) > 0 {
			for k, v := range value.Values {
				collector.AddValue(
					&genruntime.SecretDestination{
						Name: expression.Name,
						Key:  k,
					}, v)
			}
		} else {
			return nil, errors.Errorf("unexpected expression output")
		}
	}

	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return secrets.SliceToClientObjectSlice(result), nil
}

var _ kubernetesResourceExporter = &handcraftedKubernetesResourceExporter{}

type handcraftedKubernetesResourceExporter struct {
	obj        genruntime.ARMMetaObject
	log        logr.Logger
	extension  genruntime.ResourceExtension
	connection Connection
}

func (h *handcraftedKubernetesResourceExporter) Export(ctx context.Context) ([]client.Object, error) {
	retriever := extensions.CreateKubernetesExporter(ctx, h.extension, h.connection.Client(), h.log)
	resources, err := retriever(h.obj)
	if err != nil {
		return nil, errors.Wrap(err, "extension failed to produce resources for export")
	}

	return resources, nil
}

var _ kubernetesResourceExporter = &autogeneratedKubernetesResourceExporter{}

type autogeneratedKubernetesResourceExporter struct {
	obj          genruntime.ARMMetaObject
	versionedObj genruntime.ARMMetaObject
	log          logr.Logger
	connection   Connection
}

func (a *autogeneratedKubernetesResourceExporter) Export(ctx context.Context) ([]client.Object, error) {
	exporter, ok := a.versionedObj.(genruntime.KubernetesExporter)
	if !ok {
		return nil, nil
	}

	resources, err := exporter.ExportKubernetesResources(ctx, a.obj, a.connection.Client(), a.log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to produce resources for export")
	}

	return resources, nil
}
