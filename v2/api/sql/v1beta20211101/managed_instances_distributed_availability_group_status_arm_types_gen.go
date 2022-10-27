// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

type ManagedInstances_DistributedAvailabilityGroup_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *DistributedAvailabilityGroupProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The properties of a distributed availability group.
type DistributedAvailabilityGroupProperties_STATUS_ARM struct {
	// DistributedAvailabilityGroupId: The distributed availability group id
	DistributedAvailabilityGroupId *string `json:"distributedAvailabilityGroupId,omitempty"`

	// LastHardenedLsn: The last hardened lsn
	LastHardenedLsn *string `json:"lastHardenedLsn,omitempty"`

	// LinkState: The link state
	LinkState *string `json:"linkState,omitempty"`

	// PrimaryAvailabilityGroupName: The primary availability group name
	PrimaryAvailabilityGroupName *string `json:"primaryAvailabilityGroupName,omitempty"`

	// ReplicationMode: The replication mode of a distributed availability group. Parameter will be ignored during link
	// creation.
	ReplicationMode *DistributedAvailabilityGroupProperties_ReplicationMode_STATUS `json:"replicationMode,omitempty"`

	// SecondaryAvailabilityGroupName: The secondary availability group name
	SecondaryAvailabilityGroupName *string `json:"secondaryAvailabilityGroupName,omitempty"`

	// SourceEndpoint: The source endpoint
	SourceEndpoint *string `json:"sourceEndpoint,omitempty"`

	// SourceReplicaId: The source replica id
	SourceReplicaId *string `json:"sourceReplicaId,omitempty"`

	// TargetDatabase: The name of the target database
	TargetDatabase *string `json:"targetDatabase,omitempty"`

	// TargetReplicaId: The target replica id
	TargetReplicaId *string `json:"targetReplicaId,omitempty"`
}
