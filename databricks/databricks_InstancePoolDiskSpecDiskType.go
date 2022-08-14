// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type InstancePoolDiskSpecDiskType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#azure_disk_volume_type InstancePool#azure_disk_volume_type}.
	AzureDiskVolumeType *string `field:"optional" json:"azureDiskVolumeType" yaml:"azureDiskVolumeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#ebs_volume_type InstancePool#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
}

