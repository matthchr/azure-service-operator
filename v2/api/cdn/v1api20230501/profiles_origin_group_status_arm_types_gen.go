// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

type Profiles_OriginGroup_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the origin group.
	Properties *AFDOriginGroupProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Read only system data
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The JSON object that contains the properties of the origin group.
type AFDOriginGroupProperties_STATUS_ARM struct {
	DeploymentStatus *AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM `json:"deploymentStatus,omitempty"`

	// HealthProbeSettings: Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings *HealthProbeParameters_STATUS_ARM `json:"healthProbeSettings,omitempty"`

	// LoadBalancingSettings: Load balancing settings for a backend pool
	LoadBalancingSettings *LoadBalancingSettingsParameters_STATUS_ARM `json:"loadBalancingSettings,omitempty"`

	// ProfileName: The name of the profile which holds the origin group.
	ProfileName *string `json:"profileName,omitempty"`

	// ProvisioningState: Provisioning status
	ProvisioningState *AFDOriginGroupProperties_ProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`

	// SessionAffinityState: Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
	SessionAffinityState *AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM `json:"sessionAffinityState,omitempty"`

	// TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: Time in minutes to shift the traffic to the endpoint gradually
	// when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not
	// supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int `json:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes,omitempty"`
}

type AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM string

const (
	AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_Failed     = AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM("Failed")
	AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_InProgress = AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM("InProgress")
	AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_NotStarted = AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM("NotStarted")
	AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_Succeeded  = AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM("Succeeded")
)

// Mapping from string to AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM
var aFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_Values = map[string]AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM{
	"failed":     AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_Failed,
	"inprogress": AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_InProgress,
	"notstarted": AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_NotStarted,
	"succeeded":  AFDOriginGroupProperties_DeploymentStatus_STATUS_ARM_Succeeded,
}

type AFDOriginGroupProperties_ProvisioningState_STATUS_ARM string

const (
	AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Creating  = AFDOriginGroupProperties_ProvisioningState_STATUS_ARM("Creating")
	AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Deleting  = AFDOriginGroupProperties_ProvisioningState_STATUS_ARM("Deleting")
	AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Failed    = AFDOriginGroupProperties_ProvisioningState_STATUS_ARM("Failed")
	AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Succeeded = AFDOriginGroupProperties_ProvisioningState_STATUS_ARM("Succeeded")
	AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Updating  = AFDOriginGroupProperties_ProvisioningState_STATUS_ARM("Updating")
)

// Mapping from string to AFDOriginGroupProperties_ProvisioningState_STATUS_ARM
var aFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Values = map[string]AFDOriginGroupProperties_ProvisioningState_STATUS_ARM{
	"creating":  AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Creating,
	"deleting":  AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Deleting,
	"failed":    AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Failed,
	"succeeded": AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Succeeded,
	"updating":  AFDOriginGroupProperties_ProvisioningState_STATUS_ARM_Updating,
}

type AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM string

const (
	AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM_Disabled = AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM("Disabled")
	AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM_Enabled  = AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM("Enabled")
)

// Mapping from string to AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM
var aFDOriginGroupProperties_SessionAffinityState_STATUS_ARM_Values = map[string]AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM{
	"disabled": AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM_Disabled,
	"enabled":  AFDOriginGroupProperties_SessionAffinityState_STATUS_ARM_Enabled,
}

// The JSON object that contains the properties to send health probes to origin.
type HealthProbeParameters_STATUS_ARM struct {
	// ProbeIntervalInSeconds: The number of seconds between health probes.Default is 240sec.
	ProbeIntervalInSeconds *int `json:"probeIntervalInSeconds,omitempty"`

	// ProbePath: The path relative to the origin that is used to determine the health of the origin.
	ProbePath *string `json:"probePath,omitempty"`

	// ProbeProtocol: Protocol to use for health probe.
	ProbeProtocol *HealthProbeParameters_ProbeProtocol_STATUS_ARM `json:"probeProtocol,omitempty"`

	// ProbeRequestType: The type of health probe request that is made.
	ProbeRequestType *HealthProbeParameters_ProbeRequestType_STATUS_ARM `json:"probeRequestType,omitempty"`
}

// Round-Robin load balancing settings for a backend pool
type LoadBalancingSettingsParameters_STATUS_ARM struct {
	// AdditionalLatencyInMilliseconds: The additional latency in milliseconds for probes to fall into the lowest latency bucket
	AdditionalLatencyInMilliseconds *int `json:"additionalLatencyInMilliseconds,omitempty"`

	// SampleSize: The number of samples to consider for load balancing decisions
	SampleSize *int `json:"sampleSize,omitempty"`

	// SuccessfulSamplesRequired: The number of samples within the sample period that must succeed
	SuccessfulSamplesRequired *int `json:"successfulSamplesRequired,omitempty"`
}

type HealthProbeParameters_ProbeProtocol_STATUS_ARM string

const (
	HealthProbeParameters_ProbeProtocol_STATUS_ARM_Http   = HealthProbeParameters_ProbeProtocol_STATUS_ARM("Http")
	HealthProbeParameters_ProbeProtocol_STATUS_ARM_Https  = HealthProbeParameters_ProbeProtocol_STATUS_ARM("Https")
	HealthProbeParameters_ProbeProtocol_STATUS_ARM_NotSet = HealthProbeParameters_ProbeProtocol_STATUS_ARM("NotSet")
)

// Mapping from string to HealthProbeParameters_ProbeProtocol_STATUS_ARM
var healthProbeParameters_ProbeProtocol_STATUS_ARM_Values = map[string]HealthProbeParameters_ProbeProtocol_STATUS_ARM{
	"http":   HealthProbeParameters_ProbeProtocol_STATUS_ARM_Http,
	"https":  HealthProbeParameters_ProbeProtocol_STATUS_ARM_Https,
	"notset": HealthProbeParameters_ProbeProtocol_STATUS_ARM_NotSet,
}

type HealthProbeParameters_ProbeRequestType_STATUS_ARM string

const (
	HealthProbeParameters_ProbeRequestType_STATUS_ARM_GET    = HealthProbeParameters_ProbeRequestType_STATUS_ARM("GET")
	HealthProbeParameters_ProbeRequestType_STATUS_ARM_HEAD   = HealthProbeParameters_ProbeRequestType_STATUS_ARM("HEAD")
	HealthProbeParameters_ProbeRequestType_STATUS_ARM_NotSet = HealthProbeParameters_ProbeRequestType_STATUS_ARM("NotSet")
)

// Mapping from string to HealthProbeParameters_ProbeRequestType_STATUS_ARM
var healthProbeParameters_ProbeRequestType_STATUS_ARM_Values = map[string]HealthProbeParameters_ProbeRequestType_STATUS_ARM{
	"get":    HealthProbeParameters_ProbeRequestType_STATUS_ARM_GET,
	"head":   HealthProbeParameters_ProbeRequestType_STATUS_ARM_HEAD,
	"notset": HealthProbeParameters_ProbeRequestType_STATUS_ARM_NotSet,
}
