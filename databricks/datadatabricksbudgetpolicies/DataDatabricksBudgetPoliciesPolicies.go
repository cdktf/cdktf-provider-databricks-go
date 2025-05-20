// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksbudgetpolicies


type DataDatabricksBudgetPoliciesPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/budget_policies#binding_workspace_ids DataDatabricksBudgetPolicies#binding_workspace_ids}.
	BindingWorkspaceIds *[]*float64 `field:"optional" json:"bindingWorkspaceIds" yaml:"bindingWorkspaceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/budget_policies#custom_tags DataDatabricksBudgetPolicies#custom_tags}.
	CustomTags interface{} `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/budget_policies#policy_name DataDatabricksBudgetPolicies#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

