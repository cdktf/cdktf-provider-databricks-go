// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoDiskSpecDiskType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/instance_pool#azure_disk_volume_type DataDatabricksInstancePool#azure_disk_volume_type}.
	AzureDiskVolumeType *string `field:"optional" json:"azureDiskVolumeType" yaml:"azureDiskVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/instance_pool#ebs_volume_type DataDatabricksInstancePool#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
}

