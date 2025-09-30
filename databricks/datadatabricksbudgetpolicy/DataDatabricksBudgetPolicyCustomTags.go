// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksbudgetpolicy


type DataDatabricksBudgetPolicyCustomTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/budget_policy#key DataDatabricksBudgetPolicy#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/budget_policy#value DataDatabricksBudgetPolicy#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

