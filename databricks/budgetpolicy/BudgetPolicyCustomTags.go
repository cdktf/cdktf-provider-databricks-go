// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetpolicy


type BudgetPolicyCustomTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/budget_policy#key BudgetPolicy#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/budget_policy#value BudgetPolicy#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

