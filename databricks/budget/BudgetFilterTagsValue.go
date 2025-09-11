// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budget


type BudgetFilterTagsValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/budget#operator Budget#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/budget#values Budget#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

