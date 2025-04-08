// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterAzureAttributesLogAnalyticsInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/resources/cluster#log_analytics_primary_key Cluster#log_analytics_primary_key}.
	LogAnalyticsPrimaryKey *string `field:"optional" json:"logAnalyticsPrimaryKey" yaml:"logAnalyticsPrimaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/resources/cluster#log_analytics_workspace_id Cluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"optional" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
}

