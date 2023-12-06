// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobJobClusterNewClusterAutoscale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/job#max_workers Job#max_workers}.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/job#min_workers Job#min_workers}.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
}

