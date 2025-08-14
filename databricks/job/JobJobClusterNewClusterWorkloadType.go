// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobJobClusterNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/job#clients Job#clients}
	Clients *JobJobClusterNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

