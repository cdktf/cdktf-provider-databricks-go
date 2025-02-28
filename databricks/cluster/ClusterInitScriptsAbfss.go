// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterInitScriptsAbfss struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/resources/cluster#destination Cluster#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

