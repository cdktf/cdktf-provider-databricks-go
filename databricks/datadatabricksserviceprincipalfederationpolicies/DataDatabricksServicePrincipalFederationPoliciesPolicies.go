// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipalfederationpolicies


type DataDatabricksServicePrincipalFederationPoliciesPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/service_principal_federation_policies#policy_id DataDatabricksServicePrincipalFederationPolicies#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/service_principal_federation_policies#service_principal_id DataDatabricksServicePrincipalFederationPolicies#service_principal_id}.
	ServicePrincipalId *float64 `field:"required" json:"servicePrincipalId" yaml:"servicePrincipalId"`
}

