// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2EvaluationThresholdColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/alert_v2#aggregation AlertV2#aggregation}.
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/alert_v2#display AlertV2#display}.
	Display *string `field:"optional" json:"display" yaml:"display"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/alert_v2#name AlertV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

