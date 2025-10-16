// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2


type DataDatabricksAlertsV2Alerts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#id DataDatabricksAlertsV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#custom_description DataDatabricksAlertsV2#custom_description}.
	CustomDescription *string `field:"optional" json:"customDescription" yaml:"customDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#custom_summary DataDatabricksAlertsV2#custom_summary}.
	CustomSummary *string `field:"optional" json:"customSummary" yaml:"customSummary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#display_name DataDatabricksAlertsV2#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#evaluation DataDatabricksAlertsV2#evaluation}.
	Evaluation *DataDatabricksAlertsV2AlertsEvaluation `field:"optional" json:"evaluation" yaml:"evaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#parent_path DataDatabricksAlertsV2#parent_path}.
	ParentPath *string `field:"optional" json:"parentPath" yaml:"parentPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#query_text DataDatabricksAlertsV2#query_text}.
	QueryText *string `field:"optional" json:"queryText" yaml:"queryText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#run_as DataDatabricksAlertsV2#run_as}.
	RunAs *DataDatabricksAlertsV2AlertsRunAs `field:"optional" json:"runAs" yaml:"runAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#run_as_user_name DataDatabricksAlertsV2#run_as_user_name}.
	RunAsUserName *string `field:"optional" json:"runAsUserName" yaml:"runAsUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#schedule DataDatabricksAlertsV2#schedule}.
	Schedule *DataDatabricksAlertsV2AlertsSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/alerts_v2#warehouse_id DataDatabricksAlertsV2#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

