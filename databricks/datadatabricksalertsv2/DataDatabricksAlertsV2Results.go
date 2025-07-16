// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2


type DataDatabricksAlertsV2Results struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#custom_description DataDatabricksAlertsV2#custom_description}.
	CustomDescription *string `field:"optional" json:"customDescription" yaml:"customDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#custom_summary DataDatabricksAlertsV2#custom_summary}.
	CustomSummary *string `field:"optional" json:"customSummary" yaml:"customSummary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#display_name DataDatabricksAlertsV2#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#evaluation DataDatabricksAlertsV2#evaluation}.
	Evaluation *DataDatabricksAlertsV2ResultsEvaluation `field:"optional" json:"evaluation" yaml:"evaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#parent_path DataDatabricksAlertsV2#parent_path}.
	ParentPath *string `field:"optional" json:"parentPath" yaml:"parentPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#query_text DataDatabricksAlertsV2#query_text}.
	QueryText *string `field:"optional" json:"queryText" yaml:"queryText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#run_as_user_name DataDatabricksAlertsV2#run_as_user_name}.
	RunAsUserName *string `field:"optional" json:"runAsUserName" yaml:"runAsUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#schedule DataDatabricksAlertsV2#schedule}.
	Schedule *DataDatabricksAlertsV2ResultsSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/alerts_v2#warehouse_id DataDatabricksAlertsV2#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

