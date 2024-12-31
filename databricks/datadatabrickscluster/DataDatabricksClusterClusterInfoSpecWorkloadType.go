// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoSpecWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/cluster#clients DataDatabricksCluster#clients}
	Clients *DataDatabricksClusterClusterInfoSpecWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

