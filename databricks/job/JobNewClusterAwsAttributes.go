// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobNewClusterAwsAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#availability Job#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#ebs_volume_count Job#ebs_volume_count}.
	EbsVolumeCount *float64 `field:"optional" json:"ebsVolumeCount" yaml:"ebsVolumeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#ebs_volume_iops Job#ebs_volume_iops}.
	EbsVolumeIops *float64 `field:"optional" json:"ebsVolumeIops" yaml:"ebsVolumeIops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#ebs_volume_size Job#ebs_volume_size}.
	EbsVolumeSize *float64 `field:"optional" json:"ebsVolumeSize" yaml:"ebsVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#ebs_volume_throughput Job#ebs_volume_throughput}.
	EbsVolumeThroughput *float64 `field:"optional" json:"ebsVolumeThroughput" yaml:"ebsVolumeThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#ebs_volume_type Job#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#first_on_demand Job#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#instance_profile_arn Job#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#spot_bid_price_percent Job#spot_bid_price_percent}.
	SpotBidPricePercent *float64 `field:"optional" json:"spotBidPricePercent" yaml:"spotBidPricePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#zone_id Job#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

