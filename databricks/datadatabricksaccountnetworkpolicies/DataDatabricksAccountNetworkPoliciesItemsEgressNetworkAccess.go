// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/account_network_policies#restriction_mode DataDatabricksAccountNetworkPolicies#restriction_mode}.
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/account_network_policies#allowed_internet_destinations DataDatabricksAccountNetworkPolicies#allowed_internet_destinations}.
	AllowedInternetDestinations interface{} `field:"optional" json:"allowedInternetDestinations" yaml:"allowedInternetDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/account_network_policies#allowed_storage_destinations DataDatabricksAccountNetworkPolicies#allowed_storage_destinations}.
	AllowedStorageDestinations interface{} `field:"optional" json:"allowedStorageDestinations" yaml:"allowedStorageDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/account_network_policies#policy_enforcement DataDatabricksAccountNetworkPolicies#policy_enforcement}.
	PolicyEnforcement *DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessPolicyEnforcement `field:"optional" json:"policyEnforcement" yaml:"policyEnforcement"`
}

