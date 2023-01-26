package job


type JobTaskNewClusterClusterMountInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#local_mount_dir_path Job#local_mount_dir_path}.
	LocalMountDirPath *string `field:"required" json:"localMountDirPath" yaml:"localMountDirPath"`
	// network_filesystem_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#network_filesystem_info Job#network_filesystem_info}
	NetworkFilesystemInfo *JobTaskNewClusterClusterMountInfoNetworkFilesystemInfo `field:"required" json:"networkFilesystemInfo" yaml:"networkFilesystemInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#remote_mount_dir_path Job#remote_mount_dir_path}.
	RemoteMountDirPath *string `field:"optional" json:"remoteMountDirPath" yaml:"remoteMountDirPath"`
}

