// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipalfederationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicePrincipalFederationPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/service_principal_federation_policy#description ServicePrincipalFederationPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/service_principal_federation_policy#oidc_policy ServicePrincipalFederationPolicy#oidc_policy}.
	OidcPolicy *ServicePrincipalFederationPolicyOidcPolicy `field:"optional" json:"oidcPolicy" yaml:"oidcPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/service_principal_federation_policy#policy_id ServicePrincipalFederationPolicy#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/service_principal_federation_policy#service_principal_id ServicePrincipalFederationPolicy#service_principal_id}.
	ServicePrincipalId *float64 `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
}

