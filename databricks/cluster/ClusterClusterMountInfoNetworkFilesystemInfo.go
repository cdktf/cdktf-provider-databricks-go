package cluster


type ClusterClusterMountInfoNetworkFilesystemInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#server_address Cluster#server_address}.
	ServerAddress *string `field:"required" json:"serverAddress" yaml:"serverAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#mount_options Cluster#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

