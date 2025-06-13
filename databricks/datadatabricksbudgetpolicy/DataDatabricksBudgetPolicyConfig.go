// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksbudgetpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksBudgetPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/budget_policy#binding_workspace_ids DataDatabricksBudgetPolicy#binding_workspace_ids}.
	BindingWorkspaceIds *[]*float64 `field:"optional" json:"bindingWorkspaceIds" yaml:"bindingWorkspaceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/budget_policy#custom_tags DataDatabricksBudgetPolicy#custom_tags}.
	CustomTags interface{} `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/budget_policy#policy_name DataDatabricksBudgetPolicy#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

