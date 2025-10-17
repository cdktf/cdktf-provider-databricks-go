// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappssettingscustomtemplate


type DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/apps_settings_custom_template#permission DataDatabricksAppsSettingsCustomTemplate#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/apps_settings_custom_template#securable_type DataDatabricksAppsSettingsCustomTemplate#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
}

