// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoClusterLogStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/cluster_pluginframework#last_attempted DataDatabricksClusterPluginframework#last_attempted}.
	LastAttempted *float64 `field:"optional" json:"lastAttempted" yaml:"lastAttempted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/cluster_pluginframework#last_exception DataDatabricksClusterPluginframework#last_exception}.
	LastException *string `field:"optional" json:"lastException" yaml:"lastException"`
}

