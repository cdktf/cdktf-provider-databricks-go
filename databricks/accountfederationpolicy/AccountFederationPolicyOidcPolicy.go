// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountfederationpolicy


type AccountFederationPolicyOidcPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/account_federation_policy#audiences AccountFederationPolicy#audiences}.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/account_federation_policy#issuer AccountFederationPolicy#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/account_federation_policy#jwks_json AccountFederationPolicy#jwks_json}.
	JwksJson *string `field:"optional" json:"jwksJson" yaml:"jwksJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/account_federation_policy#jwks_uri AccountFederationPolicy#jwks_uri}.
	JwksUri *string `field:"optional" json:"jwksUri" yaml:"jwksUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/account_federation_policy#subject AccountFederationPolicy#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/account_federation_policy#subject_claim AccountFederationPolicy#subject_claim}.
	SubjectClaim *string `field:"optional" json:"subjectClaim" yaml:"subjectClaim"`
}

