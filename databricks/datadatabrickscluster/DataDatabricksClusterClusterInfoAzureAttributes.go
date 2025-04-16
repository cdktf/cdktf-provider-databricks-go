// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoAzureAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/cluster#availability DataDatabricksCluster#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/cluster#first_on_demand DataDatabricksCluster#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// log_analytics_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/cluster#log_analytics_info DataDatabricksCluster#log_analytics_info}
	LogAnalyticsInfo *DataDatabricksClusterClusterInfoAzureAttributesLogAnalyticsInfo `field:"optional" json:"logAnalyticsInfo" yaml:"logAnalyticsInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/cluster#spot_bid_max_price DataDatabricksCluster#spot_bid_max_price}.
	SpotBidMaxPrice *float64 `field:"optional" json:"spotBidMaxPrice" yaml:"spotBidMaxPrice"`
}

