// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budget


type BudgetFilterWorkspaceId struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/budget#operator Budget#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/budget#values Budget#values}.
	Values *[]*float64 `field:"optional" json:"values" yaml:"values"`
}

