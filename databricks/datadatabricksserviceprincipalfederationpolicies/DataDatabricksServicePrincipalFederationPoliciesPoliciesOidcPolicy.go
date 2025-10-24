// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipalfederationpolicies


type DataDatabricksServicePrincipalFederationPoliciesPoliciesOidcPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/service_principal_federation_policies#audiences DataDatabricksServicePrincipalFederationPolicies#audiences}.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/service_principal_federation_policies#issuer DataDatabricksServicePrincipalFederationPolicies#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/service_principal_federation_policies#jwks_json DataDatabricksServicePrincipalFederationPolicies#jwks_json}.
	JwksJson *string `field:"optional" json:"jwksJson" yaml:"jwksJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/service_principal_federation_policies#jwks_uri DataDatabricksServicePrincipalFederationPolicies#jwks_uri}.
	JwksUri *string `field:"optional" json:"jwksUri" yaml:"jwksUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/service_principal_federation_policies#subject DataDatabricksServicePrincipalFederationPolicies#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/service_principal_federation_policies#subject_claim DataDatabricksServicePrincipalFederationPolicies#subject_claim}.
	SubjectClaim *string `field:"optional" json:"subjectClaim" yaml:"subjectClaim"`
}

