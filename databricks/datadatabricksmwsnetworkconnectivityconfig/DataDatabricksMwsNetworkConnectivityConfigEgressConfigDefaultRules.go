// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmwsnetworkconnectivityconfig


type DataDatabricksMwsNetworkConnectivityConfigEgressConfigDefaultRules struct {
	// aws_stable_ip_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/mws_network_connectivity_config#aws_stable_ip_rule DataDatabricksMwsNetworkConnectivityConfig#aws_stable_ip_rule}
	AwsStableIpRule *DataDatabricksMwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRule `field:"optional" json:"awsStableIpRule" yaml:"awsStableIpRule"`
	// azure_service_endpoint_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/mws_network_connectivity_config#azure_service_endpoint_rule DataDatabricksMwsNetworkConnectivityConfig#azure_service_endpoint_rule}
	AzureServiceEndpointRule *DataDatabricksMwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRule `field:"optional" json:"azureServiceEndpointRule" yaml:"azureServiceEndpointRule"`
}

