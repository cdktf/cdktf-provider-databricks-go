package mwsworkspaces


type MwsWorkspacesCloudResourceContainer struct {
	// gcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.23.0/docs/resources/mws_workspaces#gcp MwsWorkspaces#gcp}
	Gcp *MwsWorkspacesCloudResourceContainerGcp `field:"required" json:"gcp" yaml:"gcp"`
}

