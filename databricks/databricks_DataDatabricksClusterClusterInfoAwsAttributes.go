// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksClusterClusterInfoAwsAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#availability DataDatabricksCluster#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#ebs_volume_count DataDatabricksCluster#ebs_volume_count}.
	EbsVolumeCount *float64 `field:"optional" json:"ebsVolumeCount" yaml:"ebsVolumeCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#ebs_volume_size DataDatabricksCluster#ebs_volume_size}.
	EbsVolumeSize *float64 `field:"optional" json:"ebsVolumeSize" yaml:"ebsVolumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#ebs_volume_type DataDatabricksCluster#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#first_on_demand DataDatabricksCluster#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#instance_profile_arn DataDatabricksCluster#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#spot_bid_price_percent DataDatabricksCluster#spot_bid_price_percent}.
	SpotBidPricePercent *float64 `field:"optional" json:"spotBidPricePercent" yaml:"spotBidPricePercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#zone_id DataDatabricksCluster#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}
