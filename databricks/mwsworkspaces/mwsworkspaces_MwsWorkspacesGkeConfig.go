package mwsworkspaces


type MwsWorkspacesGkeConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#connectivity_type MwsWorkspaces#connectivity_type}.
	ConnectivityType *string `field:"required" json:"connectivityType" yaml:"connectivityType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#master_ip_range MwsWorkspaces#master_ip_range}.
	MasterIpRange *string `field:"required" json:"masterIpRange" yaml:"masterIpRange"`
}

