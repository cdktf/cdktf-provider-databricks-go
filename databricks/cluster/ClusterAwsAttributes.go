// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterAwsAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/cluster#availability Cluster#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/cluster#ebs_volume_count Cluster#ebs_volume_count}.
	EbsVolumeCount *float64 `field:"optional" json:"ebsVolumeCount" yaml:"ebsVolumeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/cluster#ebs_volume_size Cluster#ebs_volume_size}.
	EbsVolumeSize *float64 `field:"optional" json:"ebsVolumeSize" yaml:"ebsVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/cluster#ebs_volume_type Cluster#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/cluster#first_on_demand Cluster#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/cluster#instance_profile_arn Cluster#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/cluster#spot_bid_price_percent Cluster#spot_bid_price_percent}.
	SpotBidPricePercent *float64 `field:"optional" json:"spotBidPricePercent" yaml:"spotBidPricePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/cluster#zone_id Cluster#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

