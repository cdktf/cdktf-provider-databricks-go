// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterAzureAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/cluster#availability Cluster#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/cluster#first_on_demand Cluster#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// log_analytics_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/cluster#log_analytics_info Cluster#log_analytics_info}
	LogAnalyticsInfo *ClusterAzureAttributesLogAnalyticsInfo `field:"optional" json:"logAnalyticsInfo" yaml:"logAnalyticsInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/cluster#spot_bid_max_price Cluster#spot_bid_max_price}.
	SpotBidMaxPrice *float64 `field:"optional" json:"spotBidMaxPrice" yaml:"spotBidMaxPrice"`
}

