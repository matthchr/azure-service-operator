/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KubernetesConfigExporter defines a resource which can create configmaps in Kubernetes
type KubernetesConfigExporter interface {
	// ExportKubernetesConfigMaps provides a list of Kubernetes ConfigMaps for the operator to create once the resource which
	// implements this interface is successfully provisioned. This method is invoked once a resource has been
	// successfully created in Azure, but before the Ready condition has been marked successful.
	ExportKubernetesConfigMaps(ctx context.Context) ([]client.Object, error)
}
