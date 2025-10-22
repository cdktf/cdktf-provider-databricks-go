// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipalfederationpolicy


type ServicePrincipalFederationPolicyOidcPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/service_principal_federation_policy#audiences ServicePrincipalFederationPolicy#audiences}.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/service_principal_federation_policy#issuer ServicePrincipalFederationPolicy#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/service_principal_federation_policy#jwks_json ServicePrincipalFederationPolicy#jwks_json}.
	JwksJson *string `field:"optional" json:"jwksJson" yaml:"jwksJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/service_principal_federation_policy#jwks_uri ServicePrincipalFederationPolicy#jwks_uri}.
	JwksUri *string `field:"optional" json:"jwksUri" yaml:"jwksUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/service_principal_federation_policy#subject ServicePrincipalFederationPolicy#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/service_principal_federation_policy#subject_claim ServicePrincipalFederationPolicy#subject_claim}.
	SubjectClaim *string `field:"optional" json:"subjectClaim" yaml:"subjectClaim"`
}

