// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskNewClusterAzureAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/job#availability Job#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/job#first_on_demand Job#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// log_analytics_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/job#log_analytics_info Job#log_analytics_info}
	LogAnalyticsInfo *JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfo `field:"optional" json:"logAnalyticsInfo" yaml:"logAnalyticsInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/job#spot_bid_max_price Job#spot_bid_max_price}.
	SpotBidMaxPrice *float64 `field:"optional" json:"spotBidMaxPrice" yaml:"spotBidMaxPrice"`
}

