// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#clients Job#clients}
	Clients *JobNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

