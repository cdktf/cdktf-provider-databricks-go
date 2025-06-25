// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#clients Job#clients}
	Clients *JobTaskForEachTaskTaskNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

