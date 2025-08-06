// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig


type MwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#connection_state MwsNetworkConnectivityConfig#connection_state}.
	ConnectionState *string `field:"optional" json:"connectionState" yaml:"connectionState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#creation_time MwsNetworkConnectivityConfig#creation_time}.
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#deactivated MwsNetworkConnectivityConfig#deactivated}.
	Deactivated interface{} `field:"optional" json:"deactivated" yaml:"deactivated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#deactivated_at MwsNetworkConnectivityConfig#deactivated_at}.
	DeactivatedAt *float64 `field:"optional" json:"deactivatedAt" yaml:"deactivatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#domain_names MwsNetworkConnectivityConfig#domain_names}.
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#endpoint_name MwsNetworkConnectivityConfig#endpoint_name}.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#group_id MwsNetworkConnectivityConfig#group_id}.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#network_connectivity_config_id MwsNetworkConnectivityConfig#network_connectivity_config_id}.
	NetworkConnectivityConfigId *string `field:"optional" json:"networkConnectivityConfigId" yaml:"networkConnectivityConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#resource_id MwsNetworkConnectivityConfig#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#rule_id MwsNetworkConnectivityConfig#rule_id}.
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_network_connectivity_config#updated_time MwsNetworkConnectivityConfig#updated_time}.
	UpdatedTime *float64 `field:"optional" json:"updatedTime" yaml:"updatedTime"`
}

