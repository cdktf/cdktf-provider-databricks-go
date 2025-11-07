// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappssettingscustomtemplate


type DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/apps_settings_custom_template#name DataDatabricksAppsSettingsCustomTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/apps_settings_custom_template#description DataDatabricksAppsSettingsCustomTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/apps_settings_custom_template#job_spec DataDatabricksAppsSettingsCustomTemplate#job_spec}.
	JobSpec *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsJobSpec `field:"optional" json:"jobSpec" yaml:"jobSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/apps_settings_custom_template#secret_spec DataDatabricksAppsSettingsCustomTemplate#secret_spec}.
	SecretSpec *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSecretSpec `field:"optional" json:"secretSpec" yaml:"secretSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/apps_settings_custom_template#serving_endpoint_spec DataDatabricksAppsSettingsCustomTemplate#serving_endpoint_spec}.
	ServingEndpointSpec *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpec `field:"optional" json:"servingEndpointSpec" yaml:"servingEndpointSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/apps_settings_custom_template#sql_warehouse_spec DataDatabricksAppsSettingsCustomTemplate#sql_warehouse_spec}.
	SqlWarehouseSpec *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpec `field:"optional" json:"sqlWarehouseSpec" yaml:"sqlWarehouseSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/apps_settings_custom_template#uc_securable_spec DataDatabricksAppsSettingsCustomTemplate#uc_securable_spec}.
	UcSecurableSpec *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpec `field:"optional" json:"ucSecurableSpec" yaml:"ucSecurableSpec"`
}

