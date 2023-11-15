// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterAzureAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.30.0/docs/resources/cluster#availability Cluster#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.30.0/docs/resources/cluster#first_on_demand Cluster#first_on_demand}.
	FirstOnDemand *float64 `field:"optional" json:"firstOnDemand" yaml:"firstOnDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.30.0/docs/resources/cluster#spot_bid_max_price Cluster#spot_bid_max_price}.
	SpotBidMaxPrice *float64 `field:"optional" json:"spotBidMaxPrice" yaml:"spotBidMaxPrice"`
}

