package mwsworkspaces


type MwsWorkspacesNetwork struct {
	// gcp_common_network_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#gcp_common_network_config MwsWorkspaces#gcp_common_network_config}
	GcpCommonNetworkConfig *MwsWorkspacesNetworkGcpCommonNetworkConfig `field:"required" json:"gcpCommonNetworkConfig" yaml:"gcpCommonNetworkConfig"`
	// gcp_managed_network_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#gcp_managed_network_config MwsWorkspaces#gcp_managed_network_config}
	GcpManagedNetworkConfig *MwsWorkspacesNetworkGcpManagedNetworkConfig `field:"optional" json:"gcpManagedNetworkConfig" yaml:"gcpManagedNetworkConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#network_id MwsWorkspaces#network_id}.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
}

