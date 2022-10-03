package mwsworkspaces


type MwsWorkspacesNetworkGcpCommonNetworkConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#gke_cluster_master_ip_range MwsWorkspaces#gke_cluster_master_ip_range}.
	GkeClusterMasterIpRange *string `field:"required" json:"gkeClusterMasterIpRange" yaml:"gkeClusterMasterIpRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#gke_connectivity_type MwsWorkspaces#gke_connectivity_type}.
	GkeConnectivityType *string `field:"required" json:"gkeConnectivityType" yaml:"gkeConnectivityType"`
}

