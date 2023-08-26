// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobComputeSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.0/docs/resources/job#kind Job#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

