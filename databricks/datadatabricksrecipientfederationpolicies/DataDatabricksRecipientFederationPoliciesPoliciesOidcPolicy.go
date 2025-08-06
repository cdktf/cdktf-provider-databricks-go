// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksrecipientfederationpolicies


type DataDatabricksRecipientFederationPoliciesPoliciesOidcPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/recipient_federation_policies#issuer DataDatabricksRecipientFederationPolicies#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/recipient_federation_policies#subject DataDatabricksRecipientFederationPolicies#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/recipient_federation_policies#subject_claim DataDatabricksRecipientFederationPolicies#subject_claim}.
	SubjectClaim *string `field:"required" json:"subjectClaim" yaml:"subjectClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/recipient_federation_policies#audiences DataDatabricksRecipientFederationPolicies#audiences}.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
}

