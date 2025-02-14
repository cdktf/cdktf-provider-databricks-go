// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterClusterMountInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/cluster#local_mount_dir_path Cluster#local_mount_dir_path}.
	LocalMountDirPath *string `field:"required" json:"localMountDirPath" yaml:"localMountDirPath"`
	// network_filesystem_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/cluster#network_filesystem_info Cluster#network_filesystem_info}
	NetworkFilesystemInfo *ClusterClusterMountInfoNetworkFilesystemInfo `field:"required" json:"networkFilesystemInfo" yaml:"networkFilesystemInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/cluster#remote_mount_dir_path Cluster#remote_mount_dir_path}.
	RemoteMountDirPath *string `field:"optional" json:"remoteMountDirPath" yaml:"remoteMountDirPath"`
}

