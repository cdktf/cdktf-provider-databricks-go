// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyEgressNetworkAccessPolicyEnforcement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/account_network_policy#dry_run_mode_product_filter AccountNetworkPolicy#dry_run_mode_product_filter}.
	DryRunModeProductFilter *[]*string `field:"optional" json:"dryRunModeProductFilter" yaml:"dryRunModeProductFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/account_network_policy#enforcement_mode AccountNetworkPolicy#enforcement_mode}.
	EnforcementMode *string `field:"optional" json:"enforcementMode" yaml:"enforcementMode"`
}

