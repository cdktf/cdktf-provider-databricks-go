package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#local_mount_dir_path DataDatabricksJob#local_mount_dir_path}.
	LocalMountDirPath *string `field:"required" json:"localMountDirPath" yaml:"localMountDirPath"`
	// network_filesystem_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#network_filesystem_info DataDatabricksJob#network_filesystem_info}
	NetworkFilesystemInfo *DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfo `field:"required" json:"networkFilesystemInfo" yaml:"networkFilesystemInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#remote_mount_dir_path DataDatabricksJob#remote_mount_dir_path}.
	RemoteMountDirPath *string `field:"optional" json:"remoteMountDirPath" yaml:"remoteMountDirPath"`
}

