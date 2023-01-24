package pipeline


type PipelineClusterAzureAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#availability Pipeline#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#first_on_demand Pipeline#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#spot_bid_max_price Pipeline#spot_bid_max_price}.
	SpotBidMaxPrice *float64 `field:"optional" json:"spotBidMaxPrice" yaml:"spotBidMaxPrice"`
}

