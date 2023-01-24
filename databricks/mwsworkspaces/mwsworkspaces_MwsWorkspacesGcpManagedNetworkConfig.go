package mwsworkspaces


type MwsWorkspacesGcpManagedNetworkConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#gke_cluster_pod_ip_range MwsWorkspaces#gke_cluster_pod_ip_range}.
	GkeClusterPodIpRange *string `field:"required" json:"gkeClusterPodIpRange" yaml:"gkeClusterPodIpRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#gke_cluster_service_ip_range MwsWorkspaces#gke_cluster_service_ip_range}.
	GkeClusterServiceIpRange *string `field:"required" json:"gkeClusterServiceIpRange" yaml:"gkeClusterServiceIpRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#subnet_cidr MwsWorkspaces#subnet_cidr}.
	SubnetCidr *string `field:"required" json:"subnetCidr" yaml:"subnetCidr"`
}

