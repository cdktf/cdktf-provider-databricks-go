// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig


type MwsNetworkConnectivityConfigEgressConfig struct {
	// default_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.71.0/docs/resources/mws_network_connectivity_config#default_rules MwsNetworkConnectivityConfig#default_rules}
	DefaultRules *MwsNetworkConnectivityConfigEgressConfigDefaultRules `field:"optional" json:"defaultRules" yaml:"defaultRules"`
	// target_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.71.0/docs/resources/mws_network_connectivity_config#target_rules MwsNetworkConnectivityConfig#target_rules}
	TargetRules *MwsNetworkConnectivityConfigEgressConfigTargetRules `field:"optional" json:"targetRules" yaml:"targetRules"`
}

