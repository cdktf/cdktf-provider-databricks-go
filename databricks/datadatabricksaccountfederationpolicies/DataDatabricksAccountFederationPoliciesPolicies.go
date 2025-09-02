// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountfederationpolicies


type DataDatabricksAccountFederationPoliciesPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/account_federation_policies#description DataDatabricksAccountFederationPolicies#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/account_federation_policies#oidc_policy DataDatabricksAccountFederationPolicies#oidc_policy}.
	OidcPolicy *DataDatabricksAccountFederationPoliciesPoliciesOidcPolicy `field:"optional" json:"oidcPolicy" yaml:"oidcPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/account_federation_policies#policy_id DataDatabricksAccountFederationPolicies#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/account_federation_policies#service_principal_id DataDatabricksAccountFederationPolicies#service_principal_id}.
	ServicePrincipalId *float64 `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
}

