// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoSpecAzureAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster_pluginframework#availability DataDatabricksClusterPluginframework#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster_pluginframework#first_on_demand DataDatabricksClusterPluginframework#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster_pluginframework#log_analytics_info DataDatabricksClusterPluginframework#log_analytics_info}.
	LogAnalyticsInfo interface{} `field:"optional" json:"logAnalyticsInfo" yaml:"logAnalyticsInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster_pluginframework#spot_bid_max_price DataDatabricksClusterPluginframework#spot_bid_max_price}.
	SpotBidMaxPrice *float64 `field:"optional" json:"spotBidMaxPrice" yaml:"spotBidMaxPrice"`
}

