// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/job#log_analytics_primary_key Job#log_analytics_primary_key}.
	LogAnalyticsPrimaryKey *string `field:"optional" json:"logAnalyticsPrimaryKey" yaml:"logAnalyticsPrimaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/job#log_analytics_workspace_id Job#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"optional" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
}

