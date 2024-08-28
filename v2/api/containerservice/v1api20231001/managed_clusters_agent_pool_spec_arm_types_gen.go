// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231001

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ManagedClusters_AgentPool_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of an agent pool.
	Properties *ManagedClusterAgentPoolProfileProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ManagedClusters_AgentPool_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-10-01"
func (pool ManagedClusters_AgentPool_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (pool *ManagedClusters_AgentPool_Spec_ARM) GetName() string {
	return pool.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/agentPools"
func (pool *ManagedClusters_AgentPool_Spec_ARM) GetType() string {
	return "Microsoft.ContainerService/managedClusters/agentPools"
}

// Properties for the container service agent pool profile.
type ManagedClusterAgentPoolProfileProperties_ARM struct {
	// AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
	// property is 'VirtualMachineScaleSets'.
	AvailabilityZones          []string `json:"availabilityZones"`
	CapacityReservationGroupID *string  `json:"capacityReservationGroupID,omitempty"`

	// Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
	// for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count *int `json:"count,omitempty"`

	// CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
	// a snapshot.
	CreationData *CreationData_ARM `json:"creationData,omitempty"`

	// EnableAutoScaling: Whether to enable auto-scaler
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty"`

	// EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
	// see: https://docs.microsoft.com/azure/aks/enable-host-encryption
	EnableEncryptionAtHost *bool `json:"enableEncryptionAtHost,omitempty"`

	// EnableFIPS: See [Add a FIPS-enabled node
	// pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more
	// details.
	EnableFIPS *bool `json:"enableFIPS,omitempty"`

	// EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
	// A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
	// to minimize hops. For more information see [assigning a public IP per
	// node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The
	// default is false.
	EnableNodePublicIP *bool `json:"enableNodePublicIP,omitempty"`

	// EnableUltraSSD: Whether to enable UltraSSD
	EnableUltraSSD *bool `json:"enableUltraSSD,omitempty"`

	// GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
	GpuInstanceProfile *GPUInstanceProfile_ARM `json:"gpuInstanceProfile,omitempty"`
	HostGroupID        *string                 `json:"hostGroupID,omitempty"`

	// KubeletConfig: The Kubelet configuration on the agent pool nodes.
	KubeletConfig *KubeletConfig_ARM `json:"kubeletConfig,omitempty"`

	// KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
	// storage.
	KubeletDiskType *KubeletDiskType_ARM `json:"kubeletDiskType,omitempty"`

	// LinuxOSConfig: The OS configuration of Linux agent nodes.
	LinuxOSConfig *LinuxOSConfig_ARM `json:"linuxOSConfig,omitempty"`

	// MaxCount: The maximum number of nodes for auto-scaling
	MaxCount *int `json:"maxCount,omitempty"`

	// MaxPods: The maximum number of pods that can run on a node.
	MaxPods *int `json:"maxPods,omitempty"`

	// MinCount: The minimum number of nodes for auto-scaling
	MinCount *int `json:"minCount,omitempty"`

	// Mode: A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool
	// restrictions  and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
	Mode *AgentPoolMode_ARM `json:"mode,omitempty"`

	// NetworkProfile: Network-related settings of an agent pool.
	NetworkProfile *AgentPoolNetworkProfile_ARM `json:"networkProfile,omitempty"`

	// NodeLabels: The node labels to be persisted across all nodes in agent pool.
	NodeLabels           map[string]string `json:"nodeLabels"`
	NodePublicIPPrefixID *string           `json:"nodePublicIPPrefixID,omitempty"`

	// NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `json:"nodeTaints"`

	// OrchestratorVersion: Both patch version <major.minor.patch> (e.g. 1.20.13) and <major.minor> (e.g. 1.20) are supported.
	// When <major.minor> is specified, the latest supported GA patch version is chosen automatically. Updating the cluster
	// with the same <major.minor> once it has been created (e.g. 1.14.x -> 1.14) will not trigger an upgrade, even if a newer
	// patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same
	// Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor
	// version must be within two minor versions of the control plane version. The node pool version cannot be greater than the
	// control plane version. For more information see [upgrading a node
	// pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion *string `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB        *int    `json:"osDiskSizeGB,omitempty"`

	// OsDiskType: The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested
	// OSDiskSizeGB. Otherwise,  defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral
	// OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
	OsDiskType *OSDiskType_ARM `json:"osDiskType,omitempty"`

	// OsSKU: Specifies the OS SKU used by the agent pool. The default is Ubuntu if OSType is Linux. The default is Windows2019
	// when  Kubernetes <= 1.24 or Windows2022 when Kubernetes >= 1.25 if OSType is Windows.
	OsSKU *OSSKU_ARM `json:"osSKU,omitempty"`

	// OsType: The operating system type. The default is Linux.
	OsType      *OSType_ARM `json:"osType,omitempty"`
	PodSubnetID *string     `json:"podSubnetID,omitempty"`

	// PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
	// field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
	// be stopped if it is Running and provisioning state is Succeeded
	PowerState                *PowerState_ARM `json:"powerState,omitempty"`
	ProximityPlacementGroupID *string         `json:"proximityPlacementGroupID,omitempty"`

	// ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
	ScaleDownMode *ScaleDownMode_ARM `json:"scaleDownMode,omitempty"`

	// ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is
	// 'Delete'.
	ScaleSetEvictionPolicy *ScaleSetEvictionPolicy_ARM `json:"scaleSetEvictionPolicy,omitempty"`

	// ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
	ScaleSetPriority *ScaleSetPriority_ARM `json:"scaleSetPriority,omitempty"`

	// SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
	// on-demand price. For more details on spot pricing, see [spot VMs
	// pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
	SpotMaxPrice *float64 `json:"spotMaxPrice,omitempty"`

	// Tags: The tags to be persisted on the agent pool virtual machine scale set.
	Tags map[string]string `json:"tags"`

	// Type: The type of Agent Pool.
	Type *AgentPoolType_ARM `json:"type,omitempty"`

	// UpgradeSettings: Settings for upgrading the agentpool
	UpgradeSettings *AgentPoolUpgradeSettings_ARM `json:"upgradeSettings,omitempty"`

	// VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
	// might fail to run correctly. For more details on restricted VM sizes, see:
	// https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize       *string `json:"vmSize,omitempty"`
	VnetSubnetID *string `json:"vnetSubnetID,omitempty"`

	// WorkloadRuntime: Determines the type of workload a node can run.
	WorkloadRuntime *WorkloadRuntime_ARM `json:"workloadRuntime,omitempty"`
}

// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions
// and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
// +kubebuilder:validation:Enum={"System","User"}
type AgentPoolMode_ARM string

const (
	AgentPoolMode_ARM_System = AgentPoolMode_ARM("System")
	AgentPoolMode_ARM_User   = AgentPoolMode_ARM("User")
)

// Mapping from string to AgentPoolMode_ARM
var agentPoolMode_ARM_Values = map[string]AgentPoolMode_ARM{
	"system": AgentPoolMode_ARM_System,
	"user":   AgentPoolMode_ARM_User,
}

// Network settings of an agent pool.
type AgentPoolNetworkProfile_ARM struct {
	// AllowedHostPorts: The port ranges that are allowed to access. The specified ranges are allowed to overlap.
	AllowedHostPorts          []PortRange_ARM `json:"allowedHostPorts"`
	ApplicationSecurityGroups []string        `json:"applicationSecurityGroups,omitempty"`

	// NodePublicIPTags: IPTags of instance-level public IPs.
	NodePublicIPTags []IPTag_ARM `json:"nodePublicIPTags"`
}

// The type of Agent Pool.
// +kubebuilder:validation:Enum={"AvailabilitySet","VirtualMachineScaleSets"}
type AgentPoolType_ARM string

const (
	AgentPoolType_ARM_AvailabilitySet         = AgentPoolType_ARM("AvailabilitySet")
	AgentPoolType_ARM_VirtualMachineScaleSets = AgentPoolType_ARM("VirtualMachineScaleSets")
)

// Mapping from string to AgentPoolType_ARM
var agentPoolType_ARM_Values = map[string]AgentPoolType_ARM{
	"availabilityset":         AgentPoolType_ARM_AvailabilitySet,
	"virtualmachinescalesets": AgentPoolType_ARM_VirtualMachineScaleSets,
}

// Settings for upgrading an agentpool
type AgentPoolUpgradeSettings_ARM struct {
	// DrainTimeoutInMinutes: The amount of time (in minutes) to wait on eviction of pods and graceful termination per node.
	// This eviction wait time honors waiting on pod disruption budgets. If this time is exceeded, the upgrade fails. If not
	// specified, the default is 30 minutes.
	DrainTimeoutInMinutes *int `json:"drainTimeoutInMinutes,omitempty"`

	// MaxSurge: This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it
	// is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded
	// up. If not specified, the default is 1. For more information, including best practices, see:
	// https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade
	MaxSurge *string `json:"maxSurge,omitempty"`
}

// Data used when creating a target resource from a source resource.
type CreationData_ARM struct {
	SourceResourceId *string `json:"sourceResourceId,omitempty"`
}

// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
// +kubebuilder:validation:Enum={"MIG1g","MIG2g","MIG3g","MIG4g","MIG7g"}
type GPUInstanceProfile_ARM string

const (
	GPUInstanceProfile_ARM_MIG1G = GPUInstanceProfile_ARM("MIG1g")
	GPUInstanceProfile_ARM_MIG2G = GPUInstanceProfile_ARM("MIG2g")
	GPUInstanceProfile_ARM_MIG3G = GPUInstanceProfile_ARM("MIG3g")
	GPUInstanceProfile_ARM_MIG4G = GPUInstanceProfile_ARM("MIG4g")
	GPUInstanceProfile_ARM_MIG7G = GPUInstanceProfile_ARM("MIG7g")
)

// Mapping from string to GPUInstanceProfile_ARM
var gPUInstanceProfile_ARM_Values = map[string]GPUInstanceProfile_ARM{
	"mig1g": GPUInstanceProfile_ARM_MIG1G,
	"mig2g": GPUInstanceProfile_ARM_MIG2G,
	"mig3g": GPUInstanceProfile_ARM_MIG3G,
	"mig4g": GPUInstanceProfile_ARM_MIG4G,
	"mig7g": GPUInstanceProfile_ARM_MIG7G,
}

// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type KubeletConfig_ARM struct {
	// AllowedUnsafeSysctls: Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).
	AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls"`

	// ContainerLogMaxFiles: The maximum number of container log files that can be present for a container. The number must be
	// ≥ 2.
	ContainerLogMaxFiles *int `json:"containerLogMaxFiles,omitempty"`

	// ContainerLogMaxSizeMB: The maximum size (e.g. 10Mi) of container log file before it is rotated.
	ContainerLogMaxSizeMB *int `json:"containerLogMaxSizeMB,omitempty"`

	// CpuCfsQuota: The default is true.
	CpuCfsQuota *bool `json:"cpuCfsQuota,omitempty"`

	// CpuCfsQuotaPeriod: The default is '100ms.' Valid values are a sequence of decimal numbers with an optional fraction and
	// a unit suffix. For example: '300ms', '2h45m'. Supported units are 'ns', 'us', 'ms', 's', 'm', and 'h'.
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`

	// CpuManagerPolicy: The default is 'none'. See [Kubernetes CPU management
	// policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more
	// information. Allowed values are 'none' and 'static'.
	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty"`

	// FailSwapOn: If set to true it will make the Kubelet fail to start if swap is enabled on the node.
	FailSwapOn *bool `json:"failSwapOn,omitempty"`

	// ImageGcHighThreshold: To disable image garbage collection, set to 100. The default is 85%
	ImageGcHighThreshold *int `json:"imageGcHighThreshold,omitempty"`

	// ImageGcLowThreshold: This cannot be set higher than imageGcHighThreshold. The default is 80%
	ImageGcLowThreshold *int `json:"imageGcLowThreshold,omitempty"`

	// PodMaxPids: The maximum number of processes per pod.
	PodMaxPids *int `json:"podMaxPids,omitempty"`

	// TopologyManagerPolicy: For more information see [Kubernetes Topology
	// Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is 'none'. Allowed values
	// are 'none', 'best-effort', 'restricted', and 'single-numa-node'.
	TopologyManagerPolicy *string `json:"topologyManagerPolicy,omitempty"`
}

// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
// +kubebuilder:validation:Enum={"OS","Temporary"}
type KubeletDiskType_ARM string

const (
	KubeletDiskType_ARM_OS        = KubeletDiskType_ARM("OS")
	KubeletDiskType_ARM_Temporary = KubeletDiskType_ARM("Temporary")
)

// Mapping from string to KubeletDiskType_ARM
var kubeletDiskType_ARM_Values = map[string]KubeletDiskType_ARM{
	"os":        KubeletDiskType_ARM_OS,
	"temporary": KubeletDiskType_ARM_Temporary,
}

// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type LinuxOSConfig_ARM struct {
	// SwapFileSizeMB: The size in MB of a swap file that will be created on each node.
	SwapFileSizeMB *int `json:"swapFileSizeMB,omitempty"`

	// Sysctls: Sysctl settings for Linux agent nodes.
	Sysctls *SysctlConfig_ARM `json:"sysctls,omitempty"`

	// TransparentHugePageDefrag: Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is
	// 'madvise'. For more information see [Transparent
	// Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).
	TransparentHugePageDefrag *string `json:"transparentHugePageDefrag,omitempty"`

	// TransparentHugePageEnabled: Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more
	// information see [Transparent
	// Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).
	TransparentHugePageEnabled *string `json:"transparentHugePageEnabled,omitempty"`
}

// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise,
// defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral
// OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
// +kubebuilder:validation:Enum={"Ephemeral","Managed"}
type OSDiskType_ARM string

const (
	OSDiskType_ARM_Ephemeral = OSDiskType_ARM("Ephemeral")
	OSDiskType_ARM_Managed   = OSDiskType_ARM("Managed")
)

// Mapping from string to OSDiskType_ARM
var oSDiskType_ARM_Values = map[string]OSDiskType_ARM{
	"ephemeral": OSDiskType_ARM_Ephemeral,
	"managed":   OSDiskType_ARM_Managed,
}

// Specifies the OS SKU used by the agent pool. The default is Ubuntu if OSType is Linux. The default is Windows2019 when
// Kubernetes <= 1.24 or Windows2022 when Kubernetes >= 1.25 if OSType is Windows.
// +kubebuilder:validation:Enum={"AzureLinux","CBLMariner","Ubuntu","Windows2019","Windows2022"}
type OSSKU_ARM string

const (
	OSSKU_ARM_AzureLinux  = OSSKU_ARM("AzureLinux")
	OSSKU_ARM_CBLMariner  = OSSKU_ARM("CBLMariner")
	OSSKU_ARM_Ubuntu      = OSSKU_ARM("Ubuntu")
	OSSKU_ARM_Windows2019 = OSSKU_ARM("Windows2019")
	OSSKU_ARM_Windows2022 = OSSKU_ARM("Windows2022")
)

// Mapping from string to OSSKU_ARM
var oSSKU_ARM_Values = map[string]OSSKU_ARM{
	"azurelinux":  OSSKU_ARM_AzureLinux,
	"cblmariner":  OSSKU_ARM_CBLMariner,
	"ubuntu":      OSSKU_ARM_Ubuntu,
	"windows2019": OSSKU_ARM_Windows2019,
	"windows2022": OSSKU_ARM_Windows2022,
}

// The operating system type. The default is Linux.
// +kubebuilder:validation:Enum={"Linux","Windows"}
type OSType_ARM string

const (
	OSType_ARM_Linux   = OSType_ARM("Linux")
	OSType_ARM_Windows = OSType_ARM("Windows")
)

// Mapping from string to OSType_ARM
var oSType_ARM_Values = map[string]OSType_ARM{
	"linux":   OSType_ARM_Linux,
	"windows": OSType_ARM_Windows,
}

// Describes the Power State of the cluster
type PowerState_ARM struct {
	// Code: Tells whether the cluster is Running or Stopped
	Code *PowerState_Code_ARM `json:"code,omitempty"`
}

// Describes how VMs are added to or removed from Agent Pools. See [billing
// states](https://docs.microsoft.com/azure/virtual-machines/states-billing).
// +kubebuilder:validation:Enum={"Deallocate","Delete"}
type ScaleDownMode_ARM string

const (
	ScaleDownMode_ARM_Deallocate = ScaleDownMode_ARM("Deallocate")
	ScaleDownMode_ARM_Delete     = ScaleDownMode_ARM("Delete")
)

// Mapping from string to ScaleDownMode_ARM
var scaleDownMode_ARM_Values = map[string]ScaleDownMode_ARM{
	"deallocate": ScaleDownMode_ARM_Deallocate,
	"delete":     ScaleDownMode_ARM_Delete,
}

// The eviction policy specifies what to do with the VM when it is evicted. The default is Delete. For more information
// about eviction see [spot VMs](https://docs.microsoft.com/azure/virtual-machines/spot-vms)
// +kubebuilder:validation:Enum={"Deallocate","Delete"}
type ScaleSetEvictionPolicy_ARM string

const (
	ScaleSetEvictionPolicy_ARM_Deallocate = ScaleSetEvictionPolicy_ARM("Deallocate")
	ScaleSetEvictionPolicy_ARM_Delete     = ScaleSetEvictionPolicy_ARM("Delete")
)

// Mapping from string to ScaleSetEvictionPolicy_ARM
var scaleSetEvictionPolicy_ARM_Values = map[string]ScaleSetEvictionPolicy_ARM{
	"deallocate": ScaleSetEvictionPolicy_ARM_Deallocate,
	"delete":     ScaleSetEvictionPolicy_ARM_Delete,
}

// The Virtual Machine Scale Set priority.
// +kubebuilder:validation:Enum={"Regular","Spot"}
type ScaleSetPriority_ARM string

const (
	ScaleSetPriority_ARM_Regular = ScaleSetPriority_ARM("Regular")
	ScaleSetPriority_ARM_Spot    = ScaleSetPriority_ARM("Spot")
)

// Mapping from string to ScaleSetPriority_ARM
var scaleSetPriority_ARM_Values = map[string]ScaleSetPriority_ARM{
	"regular": ScaleSetPriority_ARM_Regular,
	"spot":    ScaleSetPriority_ARM_Spot,
}

// Determines the type of workload a node can run.
// +kubebuilder:validation:Enum={"OCIContainer","WasmWasi"}
type WorkloadRuntime_ARM string

const (
	WorkloadRuntime_ARM_OCIContainer = WorkloadRuntime_ARM("OCIContainer")
	WorkloadRuntime_ARM_WasmWasi     = WorkloadRuntime_ARM("WasmWasi")
)

// Mapping from string to WorkloadRuntime_ARM
var workloadRuntime_ARM_Values = map[string]WorkloadRuntime_ARM{
	"ocicontainer": WorkloadRuntime_ARM_OCIContainer,
	"wasmwasi":     WorkloadRuntime_ARM_WasmWasi,
}

// Contains the IPTag associated with the object.
type IPTag_ARM struct {
	// IpTagType: The IP tag type. Example: RoutingPreference.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: The value of the IP tag associated with the public IP. Example: Internet.
	Tag *string `json:"tag,omitempty"`
}

// The port range.
type PortRange_ARM struct {
	// PortEnd: The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or
	// equal to portStart.
	PortEnd *int `json:"portEnd,omitempty"`

	// PortStart: The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or
	// equal to portEnd.
	PortStart *int `json:"portStart,omitempty"`

	// Protocol: The network protocol of the port.
	Protocol *PortRange_Protocol_ARM `json:"protocol,omitempty"`
}

// +kubebuilder:validation:Enum={"Running","Stopped"}
type PowerState_Code_ARM string

const (
	PowerState_Code_ARM_Running = PowerState_Code_ARM("Running")
	PowerState_Code_ARM_Stopped = PowerState_Code_ARM("Stopped")
)

// Mapping from string to PowerState_Code_ARM
var powerState_Code_ARM_Values = map[string]PowerState_Code_ARM{
	"running": PowerState_Code_ARM_Running,
	"stopped": PowerState_Code_ARM_Stopped,
}

// Sysctl settings for Linux agent nodes.
type SysctlConfig_ARM struct {
	// FsAioMaxNr: Sysctl setting fs.aio-max-nr.
	FsAioMaxNr *int `json:"fsAioMaxNr,omitempty"`

	// FsFileMax: Sysctl setting fs.file-max.
	FsFileMax *int `json:"fsFileMax,omitempty"`

	// FsInotifyMaxUserWatches: Sysctl setting fs.inotify.max_user_watches.
	FsInotifyMaxUserWatches *int `json:"fsInotifyMaxUserWatches,omitempty"`

	// FsNrOpen: Sysctl setting fs.nr_open.
	FsNrOpen *int `json:"fsNrOpen,omitempty"`

	// KernelThreadsMax: Sysctl setting kernel.threads-max.
	KernelThreadsMax *int `json:"kernelThreadsMax,omitempty"`

	// NetCoreNetdevMaxBacklog: Sysctl setting net.core.netdev_max_backlog.
	NetCoreNetdevMaxBacklog *int `json:"netCoreNetdevMaxBacklog,omitempty"`

	// NetCoreOptmemMax: Sysctl setting net.core.optmem_max.
	NetCoreOptmemMax *int `json:"netCoreOptmemMax,omitempty"`

	// NetCoreRmemDefault: Sysctl setting net.core.rmem_default.
	NetCoreRmemDefault *int `json:"netCoreRmemDefault,omitempty"`

	// NetCoreRmemMax: Sysctl setting net.core.rmem_max.
	NetCoreRmemMax *int `json:"netCoreRmemMax,omitempty"`

	// NetCoreSomaxconn: Sysctl setting net.core.somaxconn.
	NetCoreSomaxconn *int `json:"netCoreSomaxconn,omitempty"`

	// NetCoreWmemDefault: Sysctl setting net.core.wmem_default.
	NetCoreWmemDefault *int `json:"netCoreWmemDefault,omitempty"`

	// NetCoreWmemMax: Sysctl setting net.core.wmem_max.
	NetCoreWmemMax *int `json:"netCoreWmemMax,omitempty"`

	// NetIpv4IpLocalPortRange: Sysctl setting net.ipv4.ip_local_port_range.
	NetIpv4IpLocalPortRange *string `json:"netIpv4IpLocalPortRange,omitempty"`

	// NetIpv4NeighDefaultGcThresh1: Sysctl setting net.ipv4.neigh.default.gc_thresh1.
	NetIpv4NeighDefaultGcThresh1 *int `json:"netIpv4NeighDefaultGcThresh1,omitempty"`

	// NetIpv4NeighDefaultGcThresh2: Sysctl setting net.ipv4.neigh.default.gc_thresh2.
	NetIpv4NeighDefaultGcThresh2 *int `json:"netIpv4NeighDefaultGcThresh2,omitempty"`

	// NetIpv4NeighDefaultGcThresh3: Sysctl setting net.ipv4.neigh.default.gc_thresh3.
	NetIpv4NeighDefaultGcThresh3 *int `json:"netIpv4NeighDefaultGcThresh3,omitempty"`

	// NetIpv4TcpFinTimeout: Sysctl setting net.ipv4.tcp_fin_timeout.
	NetIpv4TcpFinTimeout *int `json:"netIpv4TcpFinTimeout,omitempty"`

	// NetIpv4TcpKeepaliveProbes: Sysctl setting net.ipv4.tcp_keepalive_probes.
	NetIpv4TcpKeepaliveProbes *int `json:"netIpv4TcpKeepaliveProbes,omitempty"`

	// NetIpv4TcpKeepaliveTime: Sysctl setting net.ipv4.tcp_keepalive_time.
	NetIpv4TcpKeepaliveTime *int `json:"netIpv4TcpKeepaliveTime,omitempty"`

	// NetIpv4TcpMaxSynBacklog: Sysctl setting net.ipv4.tcp_max_syn_backlog.
	NetIpv4TcpMaxSynBacklog *int `json:"netIpv4TcpMaxSynBacklog,omitempty"`

	// NetIpv4TcpMaxTwBuckets: Sysctl setting net.ipv4.tcp_max_tw_buckets.
	NetIpv4TcpMaxTwBuckets *int `json:"netIpv4TcpMaxTwBuckets,omitempty"`

	// NetIpv4TcpTwReuse: Sysctl setting net.ipv4.tcp_tw_reuse.
	NetIpv4TcpTwReuse *bool `json:"netIpv4TcpTwReuse,omitempty"`

	// NetIpv4TcpkeepaliveIntvl: Sysctl setting net.ipv4.tcp_keepalive_intvl.
	NetIpv4TcpkeepaliveIntvl *int `json:"netIpv4TcpkeepaliveIntvl,omitempty"`

	// NetNetfilterNfConntrackBuckets: Sysctl setting net.netfilter.nf_conntrack_buckets.
	NetNetfilterNfConntrackBuckets *int `json:"netNetfilterNfConntrackBuckets,omitempty"`

	// NetNetfilterNfConntrackMax: Sysctl setting net.netfilter.nf_conntrack_max.
	NetNetfilterNfConntrackMax *int `json:"netNetfilterNfConntrackMax,omitempty"`

	// VmMaxMapCount: Sysctl setting vm.max_map_count.
	VmMaxMapCount *int `json:"vmMaxMapCount,omitempty"`

	// VmSwappiness: Sysctl setting vm.swappiness.
	VmSwappiness *int `json:"vmSwappiness,omitempty"`

	// VmVfsCachePressure: Sysctl setting vm.vfs_cache_pressure.
	VmVfsCachePressure *int `json:"vmVfsCachePressure,omitempty"`
}

// +kubebuilder:validation:Enum={"TCP","UDP"}
type PortRange_Protocol_ARM string

const (
	PortRange_Protocol_ARM_TCP = PortRange_Protocol_ARM("TCP")
	PortRange_Protocol_ARM_UDP = PortRange_Protocol_ARM("UDP")
)

// Mapping from string to PortRange_Protocol_ARM
var portRange_Protocol_ARM_Values = map[string]PortRange_Protocol_ARM{
	"tcp": PortRange_Protocol_ARM_TCP,
	"udp": PortRange_Protocol_ARM_UDP,
}
