// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmwsnetworkconnectivityconfig


type DataDatabricksMwsNetworkConnectivityConfigEgressConfig struct {
	// default_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.1/docs/data-sources/mws_network_connectivity_config#default_rules DataDatabricksMwsNetworkConnectivityConfig#default_rules}
	DefaultRules *DataDatabricksMwsNetworkConnectivityConfigEgressConfigDefaultRules `field:"optional" json:"defaultRules" yaml:"defaultRules"`
	// target_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.1/docs/data-sources/mws_network_connectivity_config#target_rules DataDatabricksMwsNetworkConnectivityConfig#target_rules}
	TargetRules *DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRules `field:"optional" json:"targetRules" yaml:"targetRules"`
}

