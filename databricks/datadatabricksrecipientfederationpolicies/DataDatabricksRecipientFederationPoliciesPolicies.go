// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksrecipientfederationpolicies


type DataDatabricksRecipientFederationPoliciesPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/recipient_federation_policies#comment DataDatabricksRecipientFederationPolicies#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/recipient_federation_policies#name DataDatabricksRecipientFederationPolicies#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/recipient_federation_policies#oidc_policy DataDatabricksRecipientFederationPolicies#oidc_policy}.
	OidcPolicy *DataDatabricksRecipientFederationPoliciesPoliciesOidcPolicy `field:"optional" json:"oidcPolicy" yaml:"oidcPolicy"`
}

