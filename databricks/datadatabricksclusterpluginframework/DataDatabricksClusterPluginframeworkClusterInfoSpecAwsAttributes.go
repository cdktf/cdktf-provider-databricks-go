// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoSpecAwsAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#availability DataDatabricksClusterPluginframework#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#ebs_volume_count DataDatabricksClusterPluginframework#ebs_volume_count}.
	EbsVolumeCount *float64 `field:"optional" json:"ebsVolumeCount" yaml:"ebsVolumeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#ebs_volume_iops DataDatabricksClusterPluginframework#ebs_volume_iops}.
	EbsVolumeIops *float64 `field:"optional" json:"ebsVolumeIops" yaml:"ebsVolumeIops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#ebs_volume_size DataDatabricksClusterPluginframework#ebs_volume_size}.
	EbsVolumeSize *float64 `field:"optional" json:"ebsVolumeSize" yaml:"ebsVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#ebs_volume_throughput DataDatabricksClusterPluginframework#ebs_volume_throughput}.
	EbsVolumeThroughput *float64 `field:"optional" json:"ebsVolumeThroughput" yaml:"ebsVolumeThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#ebs_volume_type DataDatabricksClusterPluginframework#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#first_on_demand DataDatabricksClusterPluginframework#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#instance_profile_arn DataDatabricksClusterPluginframework#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#spot_bid_price_percent DataDatabricksClusterPluginframework#spot_bid_price_percent}.
	SpotBidPricePercent *float64 `field:"optional" json:"spotBidPricePercent" yaml:"spotBidPricePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#zone_id DataDatabricksClusterPluginframework#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

