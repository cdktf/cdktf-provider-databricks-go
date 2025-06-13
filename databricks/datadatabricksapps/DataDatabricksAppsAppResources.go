// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/apps#name DataDatabricksApps#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/apps#description DataDatabricksApps#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/apps#job DataDatabricksApps#job}.
	Job *DataDatabricksAppsAppResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/apps#secret DataDatabricksApps#secret}.
	Secret *DataDatabricksAppsAppResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/apps#serving_endpoint DataDatabricksApps#serving_endpoint}.
	ServingEndpoint *DataDatabricksAppsAppResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/apps#sql_warehouse DataDatabricksApps#sql_warehouse}.
	SqlWarehouse *DataDatabricksAppsAppResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/apps#uc_securable DataDatabricksApps#uc_securable}.
	UcSecurable *DataDatabricksAppsAppResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

