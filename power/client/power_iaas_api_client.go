// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/client/authentication"
	"github.com/IBM-Cloud/power-go-client/power/client/bluemix_service_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/catalog"
	"github.com/IBM-Cloud/power-go-client/power/client/hardware_platforms"
	"github.com/IBM-Cloud/power-go-client/power/client/iaas_service_broker"
	"github.com/IBM-Cloud/power-go-client/power/client/internal_power_v_s_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/internal_power_v_s_locations"
	"github.com/IBM-Cloud/power-go-client/power/client/internal_storage_regions"
	"github.com/IBM-Cloud/power-go-client/power/client/internal_transit_gateway"
	"github.com/IBM-Cloud/power-go-client/power/client/open_stacks"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_cloud_connections"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_disaster_recovery"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_events"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_jobs"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_placement_groups"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_pod_capacity"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_a_p"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_p_p_placement_groups"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_service_d_h_c_p"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_shared_processor_pools"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_snapshots"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_storage_capacity"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_storage_tiers"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_system_pools"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tasks"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants_ssh_keys"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_connections"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_policies"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volume_groups"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volume_onboarding"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
	"github.com/IBM-Cloud/power-go-client/power/client/service_bindings"
	"github.com/IBM-Cloud/power-go-client/power/client/service_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/storage_types"
	"github.com/IBM-Cloud/power-go-client/power/client/swagger_spec"
	"github.com/IBM-Cloud/power-go-client/power/client/workspaces"
)

// Default power iaas API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new power iaas API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PowerIaasAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new power iaas API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PowerIaasAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new power iaas API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PowerIaasAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PowerIaasAPI)
	cli.Transport = transport
	cli.Authentication = authentication.New(transport, formats)
	cli.BluemixServiceInstances = bluemix_service_instances.New(transport, formats)
	cli.Catalog = catalog.New(transport, formats)
	cli.HardwarePlatforms = hardware_platforms.New(transport, formats)
	cli.IaasServiceBroker = iaas_service_broker.New(transport, formats)
	cli.InternalPowervsInstances = internal_power_v_s_instances.New(transport, formats)
	cli.InternalPowervsLocations = internal_power_v_s_locations.New(transport, formats)
	cli.InternalStorageRegions = internal_storage_regions.New(transport, formats)
	cli.InternalTransitGateway = internal_transit_gateway.New(transport, formats)
	cli.OpenStacks = open_stacks.New(transport, formats)
	cli.PCloudCloudConnections = p_cloud_cloud_connections.New(transport, formats)
	cli.PCloudDisasterRecovery = p_cloud_disaster_recovery.New(transport, formats)
	cli.PCloudEvents = p_cloud_events.New(transport, formats)
	cli.PCloudImages = p_cloud_images.New(transport, formats)
	cli.PCloudInstances = p_cloud_instances.New(transport, formats)
	cli.PCloudJobs = p_cloud_jobs.New(transport, formats)
	cli.PCloudNetworks = p_cloud_networks.New(transport, formats)
	cli.PCloudpVMInstances = p_cloud_p_vm_instances.New(transport, formats)
	cli.PCloudPlacementGroups = p_cloud_placement_groups.New(transport, formats)
	cli.PCloudPodCapacity = p_cloud_pod_capacity.New(transport, formats)
	cli.PCloudsap = p_cloud_s_a_p.New(transport, formats)
	cli.PCloudsppPlacementGroups = p_cloud_s_p_p_placement_groups.New(transport, formats)
	cli.PCloudServicedhcp = p_cloud_service_d_h_c_p.New(transport, formats)
	cli.PCloudSharedProcessorPools = p_cloud_shared_processor_pools.New(transport, formats)
	cli.PCloudSnapshots = p_cloud_snapshots.New(transport, formats)
	cli.PCloudStorageCapacity = p_cloud_storage_capacity.New(transport, formats)
	cli.PCloudStorageTiers = p_cloud_storage_tiers.New(transport, formats)
	cli.PCloudSystemPools = p_cloud_system_pools.New(transport, formats)
	cli.PCloudTasks = p_cloud_tasks.New(transport, formats)
	cli.PCloudTenants = p_cloud_tenants.New(transport, formats)
	cli.PCloudTenantsSSHKeys = p_cloud_tenants_ssh_keys.New(transport, formats)
	cli.PCloudvpnConnections = p_cloud_v_p_n_connections.New(transport, formats)
	cli.PCloudvpnPolicies = p_cloud_v_p_n_policies.New(transport, formats)
	cli.PCloudVolumeGroups = p_cloud_volume_groups.New(transport, formats)
	cli.PCloudVolumeOnboarding = p_cloud_volume_onboarding.New(transport, formats)
	cli.PCloudVolumes = p_cloud_volumes.New(transport, formats)
	cli.ServiceBindings = service_bindings.New(transport, formats)
	cli.ServiceInstances = service_instances.New(transport, formats)
	cli.StorageTypes = storage_types.New(transport, formats)
	cli.SwaggerSpec = swagger_spec.New(transport, formats)
	cli.Workspaces = workspaces.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PowerIaasAPI is a client for power iaas API
type PowerIaasAPI struct {
	Authentication authentication.ClientService

	BluemixServiceInstances bluemix_service_instances.ClientService

	Catalog catalog.ClientService

	HardwarePlatforms hardware_platforms.ClientService

	IaasServiceBroker iaas_service_broker.ClientService

	InternalPowervsInstances internal_power_v_s_instances.ClientService

	InternalPowervsLocations internal_power_v_s_locations.ClientService

	InternalStorageRegions internal_storage_regions.ClientService

	InternalTransitGateway internal_transit_gateway.ClientService

	OpenStacks open_stacks.ClientService

	PCloudCloudConnections p_cloud_cloud_connections.ClientService

	PCloudDisasterRecovery p_cloud_disaster_recovery.ClientService

	PCloudEvents p_cloud_events.ClientService

	PCloudImages p_cloud_images.ClientService

	PCloudInstances p_cloud_instances.ClientService

	PCloudJobs p_cloud_jobs.ClientService

	PCloudNetworks p_cloud_networks.ClientService

	PCloudpVMInstances p_cloud_p_vm_instances.ClientService

	PCloudPlacementGroups p_cloud_placement_groups.ClientService

	PCloudPodCapacity p_cloud_pod_capacity.ClientService

	PCloudsap p_cloud_s_a_p.ClientService

	PCloudsppPlacementGroups p_cloud_s_p_p_placement_groups.ClientService

	PCloudServicedhcp p_cloud_service_d_h_c_p.ClientService

	PCloudSharedProcessorPools p_cloud_shared_processor_pools.ClientService

	PCloudSnapshots p_cloud_snapshots.ClientService

	PCloudStorageCapacity p_cloud_storage_capacity.ClientService

	PCloudStorageTiers p_cloud_storage_tiers.ClientService

	PCloudSystemPools p_cloud_system_pools.ClientService

	PCloudTasks p_cloud_tasks.ClientService

	PCloudTenants p_cloud_tenants.ClientService

	PCloudTenantsSSHKeys p_cloud_tenants_ssh_keys.ClientService

	PCloudvpnConnections p_cloud_v_p_n_connections.ClientService

	PCloudvpnPolicies p_cloud_v_p_n_policies.ClientService

	PCloudVolumeGroups p_cloud_volume_groups.ClientService

	PCloudVolumeOnboarding p_cloud_volume_onboarding.ClientService

	PCloudVolumes p_cloud_volumes.ClientService

	ServiceBindings service_bindings.ClientService

	ServiceInstances service_instances.ClientService

	StorageTypes storage_types.ClientService

	SwaggerSpec swagger_spec.ClientService

	Workspaces workspaces.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PowerIaasAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Authentication.SetTransport(transport)
	c.BluemixServiceInstances.SetTransport(transport)
	c.Catalog.SetTransport(transport)
	c.HardwarePlatforms.SetTransport(transport)
	c.IaasServiceBroker.SetTransport(transport)
	c.InternalPowervsInstances.SetTransport(transport)
	c.InternalPowervsLocations.SetTransport(transport)
	c.InternalStorageRegions.SetTransport(transport)
	c.InternalTransitGateway.SetTransport(transport)
	c.OpenStacks.SetTransport(transport)
	c.PCloudCloudConnections.SetTransport(transport)
	c.PCloudDisasterRecovery.SetTransport(transport)
	c.PCloudEvents.SetTransport(transport)
	c.PCloudImages.SetTransport(transport)
	c.PCloudInstances.SetTransport(transport)
	c.PCloudJobs.SetTransport(transport)
	c.PCloudNetworks.SetTransport(transport)
	c.PCloudpVMInstances.SetTransport(transport)
	c.PCloudPlacementGroups.SetTransport(transport)
	c.PCloudPodCapacity.SetTransport(transport)
	c.PCloudsap.SetTransport(transport)
	c.PCloudsppPlacementGroups.SetTransport(transport)
	c.PCloudServicedhcp.SetTransport(transport)
	c.PCloudSharedProcessorPools.SetTransport(transport)
	c.PCloudSnapshots.SetTransport(transport)
	c.PCloudStorageCapacity.SetTransport(transport)
	c.PCloudStorageTiers.SetTransport(transport)
	c.PCloudSystemPools.SetTransport(transport)
	c.PCloudTasks.SetTransport(transport)
	c.PCloudTenants.SetTransport(transport)
	c.PCloudTenantsSSHKeys.SetTransport(transport)
	c.PCloudvpnConnections.SetTransport(transport)
	c.PCloudvpnPolicies.SetTransport(transport)
	c.PCloudVolumeGroups.SetTransport(transport)
	c.PCloudVolumeOnboarding.SetTransport(transport)
	c.PCloudVolumes.SetTransport(transport)
	c.ServiceBindings.SetTransport(transport)
	c.ServiceInstances.SetTransport(transport)
	c.StorageTypes.SetTransport(transport)
	c.SwaggerSpec.SetTransport(transport)
	c.Workspaces.SetTransport(transport)
}
