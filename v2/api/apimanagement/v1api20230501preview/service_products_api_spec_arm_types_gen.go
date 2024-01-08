// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Service_Products_Api_Spec_ARM struct {
	Name string `json:"name,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Service_Products_Api_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (productsApi Service_Products_Api_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (productsApi *Service_Products_Api_Spec_ARM) GetName() string {
	return productsApi.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/products/apis"
func (productsApi *Service_Products_Api_Spec_ARM) GetType() string {
	return "Microsoft.ApiManagement/service/products/apis"
}
