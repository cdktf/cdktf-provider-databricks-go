// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancepool


type InstancePoolDiskSpecDiskType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#azure_disk_volume_type InstancePool#azure_disk_volume_type}.
	AzureDiskVolumeType *string `field:"optional" json:"azureDiskVolumeType" yaml:"azureDiskVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#ebs_volume_type InstancePool#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
}

