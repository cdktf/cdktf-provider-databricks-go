// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig


type MwsNetworkConnectivityConfigEgressConfigTargetRules struct {
	// azure_private_endpoint_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/mws_network_connectivity_config#azure_private_endpoint_rules MwsNetworkConnectivityConfig#azure_private_endpoint_rules}
	AzurePrivateEndpointRules interface{} `field:"optional" json:"azurePrivateEndpointRules" yaml:"azurePrivateEndpointRules"`
}

