// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobNewClusterClusterLogConfDbfs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#destination Job#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

