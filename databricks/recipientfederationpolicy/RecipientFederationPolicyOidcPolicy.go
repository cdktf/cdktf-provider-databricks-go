// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recipientfederationpolicy


type RecipientFederationPolicyOidcPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/recipient_federation_policy#issuer RecipientFederationPolicy#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/recipient_federation_policy#subject RecipientFederationPolicy#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/recipient_federation_policy#subject_claim RecipientFederationPolicy#subject_claim}.
	SubjectClaim *string `field:"required" json:"subjectClaim" yaml:"subjectClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/recipient_federation_policy#audiences RecipientFederationPolicy#audiences}.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
}

