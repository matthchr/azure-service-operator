// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

type Servers_Ipv6FirewallRule_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *IPv6ServerFirewallRuleProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The properties of an IPv6 server firewall rule.
type IPv6ServerFirewallRuleProperties_STATUS_ARM struct {
	// EndIPv6Address: The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to
	// startIpAddress.
	EndIPv6Address *string `json:"endIPv6Address,omitempty"`

	// StartIPv6Address: The start IP address of the firewall rule. Must be IPv6 format.
	StartIPv6Address *string `json:"startIPv6Address,omitempty"`
}
