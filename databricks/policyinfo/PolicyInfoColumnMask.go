// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyinfo


type PolicyInfoColumnMask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/policy_info#function_name PolicyInfo#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/policy_info#on_column PolicyInfo#on_column}.
	OnColumn *string `field:"required" json:"onColumn" yaml:"onColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/policy_info#using PolicyInfo#using}.
	Using interface{} `field:"optional" json:"using" yaml:"using"`
}

