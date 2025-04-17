// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alert


type AlertConditionOperandColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/alert#name Alert#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

