// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountfederationpolicies


type DataDatabricksAccountFederationPoliciesPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/account_federation_policies#description DataDatabricksAccountFederationPolicies#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/account_federation_policies#oidc_policy DataDatabricksAccountFederationPolicies#oidc_policy}.
	OidcPolicy *DataDatabricksAccountFederationPoliciesPoliciesOidcPolicy `field:"optional" json:"oidcPolicy" yaml:"oidcPolicy"`
}

