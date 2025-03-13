// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoAutoscale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/cluster#max_workers DataDatabricksCluster#max_workers}.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/cluster#min_workers DataDatabricksCluster#min_workers}.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
}

