// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfo


type DataDatabricksPolicyInfoColumnMask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/policy_info#function_name DataDatabricksPolicyInfo#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/policy_info#on_column DataDatabricksPolicyInfo#on_column}.
	OnColumn *string `field:"required" json:"onColumn" yaml:"onColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/policy_info#using DataDatabricksPolicyInfo#using}.
	Using interface{} `field:"optional" json:"using" yaml:"using"`
}

