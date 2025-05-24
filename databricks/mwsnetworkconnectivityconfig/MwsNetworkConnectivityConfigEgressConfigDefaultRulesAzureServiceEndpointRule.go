// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig


type MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/mws_network_connectivity_config#subnets MwsNetworkConnectivityConfig#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/mws_network_connectivity_config#target_region MwsNetworkConnectivityConfig#target_region}.
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/mws_network_connectivity_config#target_services MwsNetworkConnectivityConfig#target_services}.
	TargetServices *[]*string `field:"optional" json:"targetServices" yaml:"targetServices"`
}

