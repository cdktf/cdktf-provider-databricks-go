// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappssettingscustomtemplates


type DataDatabricksAppsSettingsCustomTemplatesTemplates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/apps_settings_custom_templates#git_provider DataDatabricksAppsSettingsCustomTemplates#git_provider}.
	GitProvider *string `field:"required" json:"gitProvider" yaml:"gitProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/apps_settings_custom_templates#git_repo DataDatabricksAppsSettingsCustomTemplates#git_repo}.
	GitRepo *string `field:"required" json:"gitRepo" yaml:"gitRepo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/apps_settings_custom_templates#manifest DataDatabricksAppsSettingsCustomTemplates#manifest}.
	Manifest *DataDatabricksAppsSettingsCustomTemplatesTemplatesManifest `field:"required" json:"manifest" yaml:"manifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/apps_settings_custom_templates#name DataDatabricksAppsSettingsCustomTemplates#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/apps_settings_custom_templates#path DataDatabricksAppsSettingsCustomTemplates#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/apps_settings_custom_templates#description DataDatabricksAppsSettingsCustomTemplates#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

