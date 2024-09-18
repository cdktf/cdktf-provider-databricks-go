// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskNewClusterClusterMountInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#local_mount_dir_path Job#local_mount_dir_path}.
	LocalMountDirPath *string `field:"required" json:"localMountDirPath" yaml:"localMountDirPath"`
	// network_filesystem_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#network_filesystem_info Job#network_filesystem_info}
	NetworkFilesystemInfo *JobTaskForEachTaskTaskNewClusterClusterMountInfoNetworkFilesystemInfo `field:"required" json:"networkFilesystemInfo" yaml:"networkFilesystemInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#remote_mount_dir_path Job#remote_mount_dir_path}.
	RemoteMountDirPath *string `field:"optional" json:"remoteMountDirPath" yaml:"remoteMountDirPath"`
}

