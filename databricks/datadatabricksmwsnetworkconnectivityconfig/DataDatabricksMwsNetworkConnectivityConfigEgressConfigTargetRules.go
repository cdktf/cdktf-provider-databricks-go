// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmwsnetworkconnectivityconfig


type DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRules struct {
	// aws_private_endpoint_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/mws_network_connectivity_config#aws_private_endpoint_rules DataDatabricksMwsNetworkConnectivityConfig#aws_private_endpoint_rules}
	AwsPrivateEndpointRules interface{} `field:"optional" json:"awsPrivateEndpointRules" yaml:"awsPrivateEndpointRules"`
	// azure_private_endpoint_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/mws_network_connectivity_config#azure_private_endpoint_rules DataDatabricksMwsNetworkConnectivityConfig#azure_private_endpoint_rules}
	AzurePrivateEndpointRules interface{} `field:"optional" json:"azurePrivateEndpointRules" yaml:"azurePrivateEndpointRules"`
}

