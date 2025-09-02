// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksrecipientfederationpolicy


type DataDatabricksRecipientFederationPolicyOidcPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/recipient_federation_policy#issuer DataDatabricksRecipientFederationPolicy#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/recipient_federation_policy#subject DataDatabricksRecipientFederationPolicy#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/recipient_federation_policy#subject_claim DataDatabricksRecipientFederationPolicy#subject_claim}.
	SubjectClaim *string `field:"required" json:"subjectClaim" yaml:"subjectClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/recipient_federation_policy#audiences DataDatabricksRecipientFederationPolicy#audiences}.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
}

