// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsNewClusterClusterMountInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#local_mount_dir_path DataDatabricksJob#local_mount_dir_path}.
	LocalMountDirPath *string `field:"required" json:"localMountDirPath" yaml:"localMountDirPath"`
	// network_filesystem_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#network_filesystem_info DataDatabricksJob#network_filesystem_info}
	NetworkFilesystemInfo *DataDatabricksJobJobSettingsSettingsNewClusterClusterMountInfoNetworkFilesystemInfo `field:"required" json:"networkFilesystemInfo" yaml:"networkFilesystemInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#remote_mount_dir_path DataDatabricksJob#remote_mount_dir_path}.
	RemoteMountDirPath *string `field:"optional" json:"remoteMountDirPath" yaml:"remoteMountDirPath"`
}

