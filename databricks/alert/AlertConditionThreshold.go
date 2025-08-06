// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alert


type AlertConditionThreshold struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/alert#value Alert#value}
	Value *AlertConditionThresholdValue `field:"required" json:"value" yaml:"value"`
}

