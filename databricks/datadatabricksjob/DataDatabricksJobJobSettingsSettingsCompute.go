// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#compute_key DataDatabricksJob#compute_key}.
	ComputeKey *string `field:"optional" json:"computeKey" yaml:"computeKey"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#spec DataDatabricksJob#spec}
	Spec *DataDatabricksJobJobSettingsSettingsComputeSpec `field:"optional" json:"spec" yaml:"spec"`
}

