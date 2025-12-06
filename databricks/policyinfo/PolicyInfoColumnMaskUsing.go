// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyinfo


type PolicyInfoColumnMaskUsing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/policy_info#alias PolicyInfo#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/policy_info#constant PolicyInfo#constant}.
	Constant *string `field:"optional" json:"constant" yaml:"constant"`
}

