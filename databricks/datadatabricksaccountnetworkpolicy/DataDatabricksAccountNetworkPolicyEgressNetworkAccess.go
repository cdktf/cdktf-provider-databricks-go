// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyEgressNetworkAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/account_network_policy#restriction_mode DataDatabricksAccountNetworkPolicy#restriction_mode}.
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/account_network_policy#allowed_internet_destinations DataDatabricksAccountNetworkPolicy#allowed_internet_destinations}.
	AllowedInternetDestinations interface{} `field:"optional" json:"allowedInternetDestinations" yaml:"allowedInternetDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/account_network_policy#allowed_storage_destinations DataDatabricksAccountNetworkPolicy#allowed_storage_destinations}.
	AllowedStorageDestinations interface{} `field:"optional" json:"allowedStorageDestinations" yaml:"allowedStorageDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/account_network_policy#policy_enforcement DataDatabricksAccountNetworkPolicy#policy_enforcement}.
	PolicyEnforcement *DataDatabricksAccountNetworkPolicyEgressNetworkAccessPolicyEnforcement `field:"optional" json:"policyEnforcement" yaml:"policyEnforcement"`
}

