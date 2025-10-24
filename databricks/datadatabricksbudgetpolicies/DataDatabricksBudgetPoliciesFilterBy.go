// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksbudgetpolicies


type DataDatabricksBudgetPoliciesFilterBy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/budget_policies#creator_user_id DataDatabricksBudgetPolicies#creator_user_id}.
	CreatorUserId *float64 `field:"optional" json:"creatorUserId" yaml:"creatorUserId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/budget_policies#creator_user_name DataDatabricksBudgetPolicies#creator_user_name}.
	CreatorUserName *string `field:"optional" json:"creatorUserName" yaml:"creatorUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/budget_policies#policy_name DataDatabricksBudgetPolicies#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

