// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/cluster#clients Cluster#clients}
	Clients *ClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

