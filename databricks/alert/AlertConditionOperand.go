// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alert


type AlertConditionOperand struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.1/docs/resources/alert#column Alert#column}
	Column *AlertConditionOperandColumn `field:"required" json:"column" yaml:"column"`
}

