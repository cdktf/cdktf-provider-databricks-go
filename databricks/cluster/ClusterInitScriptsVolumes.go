// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterInitScriptsVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.28.0/docs/resources/cluster#destination Cluster#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}
