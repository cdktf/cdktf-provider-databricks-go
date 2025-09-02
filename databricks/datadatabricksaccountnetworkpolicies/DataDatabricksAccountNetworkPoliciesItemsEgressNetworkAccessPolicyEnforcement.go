// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessPolicyEnforcement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/account_network_policies#dry_run_mode_product_filter DataDatabricksAccountNetworkPolicies#dry_run_mode_product_filter}.
	DryRunModeProductFilter *[]*string `field:"optional" json:"dryRunModeProductFilter" yaml:"dryRunModeProductFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/account_network_policies#enforcement_mode DataDatabricksAccountNetworkPolicies#enforcement_mode}.
	EnforcementMode *string `field:"optional" json:"enforcementMode" yaml:"enforcementMode"`
}

