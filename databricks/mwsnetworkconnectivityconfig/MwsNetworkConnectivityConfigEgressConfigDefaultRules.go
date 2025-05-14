// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig


type MwsNetworkConnectivityConfigEgressConfigDefaultRules struct {
	// aws_stable_ip_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/resources/mws_network_connectivity_config#aws_stable_ip_rule MwsNetworkConnectivityConfig#aws_stable_ip_rule}
	AwsStableIpRule *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRule `field:"optional" json:"awsStableIpRule" yaml:"awsStableIpRule"`
	// azure_service_endpoint_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/resources/mws_network_connectivity_config#azure_service_endpoint_rule MwsNetworkConnectivityConfig#azure_service_endpoint_rule}
	AzureServiceEndpointRule *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRule `field:"optional" json:"azureServiceEndpointRule" yaml:"azureServiceEndpointRule"`
}

