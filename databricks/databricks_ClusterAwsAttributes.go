// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type ClusterAwsAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#availability Cluster#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#ebs_volume_count Cluster#ebs_volume_count}.
	EbsVolumeCount *float64 `field:"optional" json:"ebsVolumeCount" yaml:"ebsVolumeCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#ebs_volume_size Cluster#ebs_volume_size}.
	EbsVolumeSize *float64 `field:"optional" json:"ebsVolumeSize" yaml:"ebsVolumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#ebs_volume_type Cluster#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#first_on_demand Cluster#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#instance_profile_arn Cluster#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#spot_bid_price_percent Cluster#spot_bid_price_percent}.
	SpotBidPricePercent *float64 `field:"optional" json:"spotBidPricePercent" yaml:"spotBidPricePercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#zone_id Cluster#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}
