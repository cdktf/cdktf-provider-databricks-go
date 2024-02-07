// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.0/docs/resources/job#compute_key Job#compute_key}.
	ComputeKey *string `field:"optional" json:"computeKey" yaml:"computeKey"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.0/docs/resources/job#spec Job#spec}
	Spec *JobComputeSpec `field:"optional" json:"spec" yaml:"spec"`
}

