// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappssettingscustomtemplates


type DataDatabricksAppsSettingsCustomTemplatesTemplatesManifest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/apps_settings_custom_templates#name DataDatabricksAppsSettingsCustomTemplates#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/apps_settings_custom_templates#version DataDatabricksAppsSettingsCustomTemplates#version}.
	Version *float64 `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/apps_settings_custom_templates#description DataDatabricksAppsSettingsCustomTemplates#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/apps_settings_custom_templates#resource_specs DataDatabricksAppsSettingsCustomTemplates#resource_specs}.
	ResourceSpecs interface{} `field:"optional" json:"resourceSpecs" yaml:"resourceSpecs"`
}

