// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfos


type DataDatabricksPolicyInfosPoliciesColumnMask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/policy_infos#function_name DataDatabricksPolicyInfos#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/policy_infos#on_column DataDatabricksPolicyInfos#on_column}.
	OnColumn *string `field:"required" json:"onColumn" yaml:"onColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/policy_infos#using DataDatabricksPolicyInfos#using}.
	Using interface{} `field:"optional" json:"using" yaml:"using"`
}

