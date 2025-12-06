// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappssettingscustomtemplate


type DataDatabricksAppsSettingsCustomTemplateManifest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_template#name DataDatabricksAppsSettingsCustomTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_template#version DataDatabricksAppsSettingsCustomTemplate#version}.
	Version *float64 `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_template#description DataDatabricksAppsSettingsCustomTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps_settings_custom_template#resource_specs DataDatabricksAppsSettingsCustomTemplate#resource_specs}.
	ResourceSpecs interface{} `field:"optional" json:"resourceSpecs" yaml:"resourceSpecs"`
}

