package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoAzureAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#availability DataDatabricksInstancePool#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#spot_bid_max_price DataDatabricksInstancePool#spot_bid_max_price}.
	SpotBidMaxPrice *float64 `field:"optional" json:"spotBidMaxPrice" yaml:"spotBidMaxPrice"`
}

