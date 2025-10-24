// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyEgressNetworkAccessPolicyEnforcement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/account_network_policy#dry_run_mode_product_filter DataDatabricksAccountNetworkPolicy#dry_run_mode_product_filter}.
	DryRunModeProductFilter *[]*string `field:"optional" json:"dryRunModeProductFilter" yaml:"dryRunModeProductFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/account_network_policy#enforcement_mode DataDatabricksAccountNetworkPolicy#enforcement_mode}.
	EnforcementMode *string `field:"optional" json:"enforcementMode" yaml:"enforcementMode"`
}

