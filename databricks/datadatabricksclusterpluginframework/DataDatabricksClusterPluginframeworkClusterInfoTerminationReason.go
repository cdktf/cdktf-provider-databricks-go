// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoTerminationReason struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/data-sources/cluster_pluginframework#code DataDatabricksClusterPluginframework#code}.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/data-sources/cluster_pluginframework#parameters DataDatabricksClusterPluginframework#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/data-sources/cluster_pluginframework#type DataDatabricksClusterPluginframework#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

