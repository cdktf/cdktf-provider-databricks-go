package job


type JobNewClusterClusterMountInfoNetworkFilesystemInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#server_address Job#server_address}.
	ServerAddress *string `field:"required" json:"serverAddress" yaml:"serverAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#mount_options Job#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

