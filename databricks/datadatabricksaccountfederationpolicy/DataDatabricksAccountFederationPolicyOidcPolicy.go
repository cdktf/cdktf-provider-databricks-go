// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountfederationpolicy


type DataDatabricksAccountFederationPolicyOidcPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/account_federation_policy#audiences DataDatabricksAccountFederationPolicy#audiences}.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/account_federation_policy#issuer DataDatabricksAccountFederationPolicy#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/account_federation_policy#jwks_json DataDatabricksAccountFederationPolicy#jwks_json}.
	JwksJson *string `field:"optional" json:"jwksJson" yaml:"jwksJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/account_federation_policy#jwks_uri DataDatabricksAccountFederationPolicy#jwks_uri}.
	JwksUri *string `field:"optional" json:"jwksUri" yaml:"jwksUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/account_federation_policy#subject DataDatabricksAccountFederationPolicy#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/account_federation_policy#subject_claim DataDatabricksAccountFederationPolicy#subject_claim}.
	SubjectClaim *string `field:"optional" json:"subjectClaim" yaml:"subjectClaim"`
}

