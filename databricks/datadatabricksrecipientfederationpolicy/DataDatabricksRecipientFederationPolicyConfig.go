// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksrecipientfederationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksRecipientFederationPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/recipient_federation_policy#comment DataDatabricksRecipientFederationPolicy#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/recipient_federation_policy#name DataDatabricksRecipientFederationPolicy#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/recipient_federation_policy#oidc_policy DataDatabricksRecipientFederationPolicy#oidc_policy}.
	OidcPolicy *DataDatabricksRecipientFederationPolicyOidcPolicy `field:"optional" json:"oidcPolicy" yaml:"oidcPolicy"`
}

