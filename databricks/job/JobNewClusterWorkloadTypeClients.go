// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobNewClusterWorkloadTypeClients struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#jobs Job#jobs}.
	Jobs interface{} `field:"optional" json:"jobs" yaml:"jobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#notebooks Job#notebooks}.
	Notebooks interface{} `field:"optional" json:"notebooks" yaml:"notebooks"`
}

