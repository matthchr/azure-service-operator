// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

type ConsumerGroup_StatusARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Properties: Single item in List or Get Consumer group operation
	Properties *ConsumerGroup_Status_PropertiesARM `json:"properties,omitempty"`

	//SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or
	//"Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

type ConsumerGroup_Status_PropertiesARM struct {
	//CreatedAt: Exact time the message was created.
	CreatedAt *string `json:"createdAt,omitempty"`

	//UpdatedAt: The exact time the message was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`

	//UserMetadata: User Metadata is a placeholder to store user-defined string data
	//with maximum length 1024. e.g. it can be used to store descriptive data, such as
	//list of teams and their contact information also user-defined configuration
	//settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty"`
}