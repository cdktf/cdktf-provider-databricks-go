package mwsworkspaces


type MwsWorkspacesCloudResourceBucket struct {
	// gcp block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#gcp MwsWorkspaces#gcp}
	Gcp *MwsWorkspacesCloudResourceBucketGcp `field:"required" json:"gcp" yaml:"gcp"`
}

