// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appssettingscustomtemplate


type AppsSettingsCustomTemplateManifestResourceSpecs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/apps_settings_custom_template#name AppsSettingsCustomTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/apps_settings_custom_template#description AppsSettingsCustomTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/apps_settings_custom_template#job_spec AppsSettingsCustomTemplate#job_spec}.
	JobSpec *AppsSettingsCustomTemplateManifestResourceSpecsJobSpec `field:"optional" json:"jobSpec" yaml:"jobSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/apps_settings_custom_template#secret_spec AppsSettingsCustomTemplate#secret_spec}.
	SecretSpec *AppsSettingsCustomTemplateManifestResourceSpecsSecretSpec `field:"optional" json:"secretSpec" yaml:"secretSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/apps_settings_custom_template#serving_endpoint_spec AppsSettingsCustomTemplate#serving_endpoint_spec}.
	ServingEndpointSpec *AppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpec `field:"optional" json:"servingEndpointSpec" yaml:"servingEndpointSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/apps_settings_custom_template#sql_warehouse_spec AppsSettingsCustomTemplate#sql_warehouse_spec}.
	SqlWarehouseSpec *AppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpec `field:"optional" json:"sqlWarehouseSpec" yaml:"sqlWarehouseSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/apps_settings_custom_template#uc_securable_spec AppsSettingsCustomTemplate#uc_securable_spec}.
	UcSecurableSpec *AppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpec `field:"optional" json:"ucSecurableSpec" yaml:"ucSecurableSpec"`
}

