// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipalfederationpolicy


type DataDatabricksServicePrincipalFederationPolicyOidcPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/service_principal_federation_policy#audiences DataDatabricksServicePrincipalFederationPolicy#audiences}.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/service_principal_federation_policy#issuer DataDatabricksServicePrincipalFederationPolicy#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/service_principal_federation_policy#jwks_json DataDatabricksServicePrincipalFederationPolicy#jwks_json}.
	JwksJson *string `field:"optional" json:"jwksJson" yaml:"jwksJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/service_principal_federation_policy#jwks_uri DataDatabricksServicePrincipalFederationPolicy#jwks_uri}.
	JwksUri *string `field:"optional" json:"jwksUri" yaml:"jwksUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/service_principal_federation_policy#subject DataDatabricksServicePrincipalFederationPolicy#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/service_principal_federation_policy#subject_claim DataDatabricksServicePrincipalFederationPolicy#subject_claim}.
	SubjectClaim *string `field:"optional" json:"subjectClaim" yaml:"subjectClaim"`
}

