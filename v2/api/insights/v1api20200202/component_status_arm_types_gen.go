// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200202

type Component_STATUS_ARM struct {
	// Etag: Resource etag
	Etag *string `json:"etag,omitempty"`

	// Id: Azure resource Id
	Id *string `json:"id,omitempty"`

	// Kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string,
	// values should typically be one of the following: web, ios, other, store, java, phone.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// Name: Azure resource name
	Name *string `json:"name,omitempty"`

	// Properties: Properties that define an Application Insights component resource.
	Properties *ApplicationInsightsComponentProperties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Azure resource type
	Type *string `json:"type,omitempty"`
}

// Properties that define an Application Insights component resource.
type ApplicationInsightsComponentProperties_STATUS_ARM struct {
	// AppId: Application Insights Unique ID for your Application.
	AppId *string `json:"AppId,omitempty"`

	// ApplicationId: The unique ID of your application. This field mirrors the 'Name' field and cannot be changed.
	ApplicationId *string `json:"ApplicationId,omitempty"`

	// Application_Type: Type of application being monitored.
	Application_Type *ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM `json:"Application_Type,omitempty"`

	// ConnectionString: Application Insights component connection string.
	ConnectionString *string `json:"ConnectionString,omitempty"`

	// CreationDate: Creation Date for the Application Insights component, in ISO 8601 format.
	CreationDate *string `json:"CreationDate,omitempty"`

	// DisableIpMasking: Disable IP masking.
	DisableIpMasking *bool `json:"DisableIpMasking,omitempty"`

	// DisableLocalAuth: Disable Non-AAD based Auth.
	DisableLocalAuth *bool `json:"DisableLocalAuth,omitempty"`

	// Flow_Type: Used by the Application Insights system to determine what kind of flow this component was created by. This is
	// to be set to 'Bluefield' when creating/updating a component via the REST API.
	Flow_Type *ApplicationInsightsComponentProperties_Flow_Type_STATUS_ARM `json:"Flow_Type,omitempty"`

	// ForceCustomerStorageForProfiler: Force users to create their own storage account for profiler and debugger.
	ForceCustomerStorageForProfiler *bool `json:"ForceCustomerStorageForProfiler,omitempty"`

	// HockeyAppId: The unique application ID created when a new application is added to HockeyApp, used for communications
	// with HockeyApp.
	HockeyAppId *string `json:"HockeyAppId,omitempty"`

	// HockeyAppToken: Token used to authenticate communications with between Application Insights and HockeyApp.
	HockeyAppToken *string `json:"HockeyAppToken,omitempty"`

	// ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.
	ImmediatePurgeDataOn30Days *bool `json:"ImmediatePurgeDataOn30Days,omitempty"`

	// IngestionMode: Indicates the flow of the ingestion.
	IngestionMode *ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM `json:"IngestionMode,omitempty"`

	// InstrumentationKey: Application Insights Instrumentation key. A read-only value that applications can use to identify
	// the destination for all telemetry sent to Azure Application Insights. This value will be supplied upon construction of
	// each new Application Insights component.
	InstrumentationKey *string `json:"InstrumentationKey,omitempty"`

	// LaMigrationDate: The date which the component got migrated to LA, in ISO 8601 format.
	LaMigrationDate *string `json:"LaMigrationDate,omitempty"`

	// Name: Application name.
	Name *string `json:"Name,omitempty"`

	// PrivateLinkScopedResources: List of linked private link scope resources.
	PrivateLinkScopedResources []PrivateLinkScopedResource_STATUS_ARM `json:"PrivateLinkScopedResources,omitempty"`

	// ProvisioningState: Current state of this component: whether or not is has been provisioned within the resource group it
	// is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying,
	// Canceled, and Failed.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PublicNetworkAccessForIngestion: The network access type for accessing Application Insights ingestion.
	PublicNetworkAccessForIngestion *PublicNetworkAccessType_STATUS_ARM `json:"publicNetworkAccessForIngestion,omitempty"`

	// PublicNetworkAccessForQuery: The network access type for accessing Application Insights query.
	PublicNetworkAccessForQuery *PublicNetworkAccessType_STATUS_ARM `json:"publicNetworkAccessForQuery,omitempty"`

	// Request_Source: Describes what tool created this Application Insights component. Customers using this API should set
	// this to the default 'rest'.
	Request_Source *ApplicationInsightsComponentProperties_Request_Source_STATUS_ARM `json:"Request_Source,omitempty"`

	// RetentionInDays: Retention period in days.
	RetentionInDays *int `json:"RetentionInDays,omitempty"`

	// SamplingPercentage: Percentage of the data produced by the application being monitored that is being sampled for
	// Application Insights telemetry.
	SamplingPercentage *float64 `json:"SamplingPercentage,omitempty"`

	// TenantId: Azure Tenant Id.
	TenantId *string `json:"TenantId,omitempty"`

	// WorkspaceResourceId: Resource Id of the log analytics workspace which the data will be ingested to. This property is
	// required to create an application with this API version. Applications from older versions will not have this property.
	WorkspaceResourceId *string `json:"WorkspaceResourceId,omitempty"`
}

type ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM string

const (
	ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM_Other = ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM("other")
	ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM_Web   = ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM("web")
)

// Mapping from string to ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM
var applicationInsightsComponentProperties_Application_Type_STATUS_ARM_Values = map[string]ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM{
	"other": ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM_Other,
	"web":   ApplicationInsightsComponentProperties_Application_Type_STATUS_ARM_Web,
}

type ApplicationInsightsComponentProperties_Flow_Type_STATUS_ARM string

const ApplicationInsightsComponentProperties_Flow_Type_STATUS_ARM_Bluefield = ApplicationInsightsComponentProperties_Flow_Type_STATUS_ARM("Bluefield")

// Mapping from string to ApplicationInsightsComponentProperties_Flow_Type_STATUS_ARM
var applicationInsightsComponentProperties_Flow_Type_STATUS_ARM_Values = map[string]ApplicationInsightsComponentProperties_Flow_Type_STATUS_ARM{
	"bluefield": ApplicationInsightsComponentProperties_Flow_Type_STATUS_ARM_Bluefield,
}

type ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM string

const (
	ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM_ApplicationInsights                       = ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM("ApplicationInsights")
	ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM_ApplicationInsightsWithDiagnosticSettings = ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM("ApplicationInsightsWithDiagnosticSettings")
	ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM_LogAnalytics                              = ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM("LogAnalytics")
)

// Mapping from string to ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM
var applicationInsightsComponentProperties_IngestionMode_STATUS_ARM_Values = map[string]ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM{
	"applicationinsights":                       ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM_ApplicationInsights,
	"applicationinsightswithdiagnosticsettings": ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM_ApplicationInsightsWithDiagnosticSettings,
	"loganalytics":                              ApplicationInsightsComponentProperties_IngestionMode_STATUS_ARM_LogAnalytics,
}

type ApplicationInsightsComponentProperties_Request_Source_STATUS_ARM string

const ApplicationInsightsComponentProperties_Request_Source_STATUS_ARM_Rest = ApplicationInsightsComponentProperties_Request_Source_STATUS_ARM("rest")

// Mapping from string to ApplicationInsightsComponentProperties_Request_Source_STATUS_ARM
var applicationInsightsComponentProperties_Request_Source_STATUS_ARM_Values = map[string]ApplicationInsightsComponentProperties_Request_Source_STATUS_ARM{
	"rest": ApplicationInsightsComponentProperties_Request_Source_STATUS_ARM_Rest,
}

// The private link scope resource reference.
type PrivateLinkScopedResource_STATUS_ARM struct {
	// ResourceId: The full resource Id of the private link scope resource.
	ResourceId *string `json:"ResourceId,omitempty"`

	// ScopeId: The private link scope unique Identifier.
	ScopeId *string `json:"ScopeId,omitempty"`
}

// The network access type for operating on the Application Insights Component. By default it is Enabled
type PublicNetworkAccessType_STATUS_ARM string

const (
	PublicNetworkAccessType_STATUS_ARM_Disabled = PublicNetworkAccessType_STATUS_ARM("Disabled")
	PublicNetworkAccessType_STATUS_ARM_Enabled  = PublicNetworkAccessType_STATUS_ARM("Enabled")
)

// Mapping from string to PublicNetworkAccessType_STATUS_ARM
var publicNetworkAccessType_STATUS_ARM_Values = map[string]PublicNetworkAccessType_STATUS_ARM{
	"disabled": PublicNetworkAccessType_STATUS_ARM_Disabled,
	"enabled":  PublicNetworkAccessType_STATUS_ARM_Enabled,
}
