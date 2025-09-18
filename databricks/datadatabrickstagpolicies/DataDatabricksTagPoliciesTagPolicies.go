// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstagpolicies


type DataDatabricksTagPoliciesTagPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/tag_policies#tag_key DataDatabricksTagPolicies#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/tag_policies#description DataDatabricksTagPolicies#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/tag_policies#values DataDatabricksTagPolicies#values}.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

