// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItems struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/account_network_policies#account_id DataDatabricksAccountNetworkPolicies#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/account_network_policies#egress DataDatabricksAccountNetworkPolicies#egress}.
	Egress *DataDatabricksAccountNetworkPoliciesItemsEgress `field:"optional" json:"egress" yaml:"egress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/account_network_policies#network_policy_id DataDatabricksAccountNetworkPolicies#network_policy_id}.
	NetworkPolicyId *string `field:"optional" json:"networkPolicyId" yaml:"networkPolicyId"`
}

