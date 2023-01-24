package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoDiskSpecDiskType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#azure_disk_volume_type DataDatabricksInstancePool#azure_disk_volume_type}.
	AzureDiskVolumeType *string `field:"optional" json:"azureDiskVolumeType" yaml:"azureDiskVolumeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#ebs_volume_type DataDatabricksInstancePool#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
}

