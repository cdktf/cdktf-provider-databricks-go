// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alert


type AlertCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/alert#op Alert#op}.
	Op *string `field:"required" json:"op" yaml:"op"`
	// operand block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/alert#operand Alert#operand}
	Operand *AlertConditionOperand `field:"required" json:"operand" yaml:"operand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/alert#empty_result_state Alert#empty_result_state}.
	EmptyResultState *string `field:"optional" json:"emptyResultState" yaml:"emptyResultState"`
	// threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/alert#threshold Alert#threshold}
	Threshold *AlertConditionThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}

