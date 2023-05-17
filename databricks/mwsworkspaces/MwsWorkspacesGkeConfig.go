package mwsworkspaces


type MwsWorkspacesGkeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/mws_workspaces#connectivity_type MwsWorkspaces#connectivity_type}.
	ConnectivityType *string `field:"required" json:"connectivityType" yaml:"connectivityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/mws_workspaces#master_ip_range MwsWorkspaces#master_ip_range}.
	MasterIpRange *string `field:"required" json:"masterIpRange" yaml:"masterIpRange"`
}

