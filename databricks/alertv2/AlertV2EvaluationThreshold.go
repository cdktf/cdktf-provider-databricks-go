// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2EvaluationThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/alert_v2#column AlertV2#column}.
	Column *AlertV2EvaluationThresholdColumn `field:"optional" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/alert_v2#value AlertV2#value}.
	Value *AlertV2EvaluationThresholdValue `field:"optional" json:"value" yaml:"value"`
}

