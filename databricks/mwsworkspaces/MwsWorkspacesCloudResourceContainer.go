package mwsworkspaces


type MwsWorkspacesCloudResourceContainer struct {
	// gcp block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#gcp MwsWorkspaces#gcp}
	Gcp *MwsWorkspacesCloudResourceContainerGcp `field:"required" json:"gcp" yaml:"gcp"`
}

