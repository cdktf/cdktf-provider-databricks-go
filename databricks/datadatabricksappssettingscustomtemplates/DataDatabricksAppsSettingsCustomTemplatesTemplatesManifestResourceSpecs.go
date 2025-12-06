// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappssettingscustomtemplates


type DataDatabricksAppsSettingsCustomTemplatesTemplatesManifestResourceSpecs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_templates#name DataDatabricksAppsSettingsCustomTemplates#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_templates#description DataDatabricksAppsSettingsCustomTemplates#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_templates#job_spec DataDatabricksAppsSettingsCustomTemplates#job_spec}.
	JobSpec *DataDatabricksAppsSettingsCustomTemplatesTemplatesManifestResourceSpecsJobSpec `field:"optional" json:"jobSpec" yaml:"jobSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_templates#secret_spec DataDatabricksAppsSettingsCustomTemplates#secret_spec}.
	SecretSpec *DataDatabricksAppsSettingsCustomTemplatesTemplatesManifestResourceSpecsSecretSpec `field:"optional" json:"secretSpec" yaml:"secretSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_templates#serving_endpoint_spec DataDatabricksAppsSettingsCustomTemplates#serving_endpoint_spec}.
	ServingEndpointSpec *DataDatabricksAppsSettingsCustomTemplatesTemplatesManifestResourceSpecsServingEndpointSpec `field:"optional" json:"servingEndpointSpec" yaml:"servingEndpointSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_templates#sql_warehouse_spec DataDatabricksAppsSettingsCustomTemplates#sql_warehouse_spec}.
	SqlWarehouseSpec *DataDatabricksAppsSettingsCustomTemplatesTemplatesManifestResourceSpecsSqlWarehouseSpec `field:"optional" json:"sqlWarehouseSpec" yaml:"sqlWarehouseSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_templates#uc_securable_spec DataDatabricksAppsSettingsCustomTemplates#uc_securable_spec}.
	UcSecurableSpec *DataDatabricksAppsSettingsCustomTemplatesTemplatesManifestResourceSpecsUcSecurableSpec `field:"optional" json:"ucSecurableSpec" yaml:"ucSecurableSpec"`
}

