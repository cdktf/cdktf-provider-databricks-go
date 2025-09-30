// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineClusterAzureAttributesLogAnalyticsInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/pipeline#log_analytics_primary_key Pipeline#log_analytics_primary_key}.
	LogAnalyticsPrimaryKey *string `field:"optional" json:"logAnalyticsPrimaryKey" yaml:"logAnalyticsPrimaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/pipeline#log_analytics_workspace_id Pipeline#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"optional" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
}

