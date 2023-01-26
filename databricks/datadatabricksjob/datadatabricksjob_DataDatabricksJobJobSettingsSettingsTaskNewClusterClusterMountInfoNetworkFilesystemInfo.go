package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#server_address DataDatabricksJob#server_address}.
	ServerAddress *string `field:"required" json:"serverAddress" yaml:"serverAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#mount_options DataDatabricksJob#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

