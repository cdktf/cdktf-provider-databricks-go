// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterClusterMountInfoNetworkFilesystemInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.1/docs/data-sources/job#server_address DataDatabricksJob#server_address}.
	ServerAddress *string `field:"required" json:"serverAddress" yaml:"serverAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.1/docs/data-sources/job#mount_options DataDatabricksJob#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

