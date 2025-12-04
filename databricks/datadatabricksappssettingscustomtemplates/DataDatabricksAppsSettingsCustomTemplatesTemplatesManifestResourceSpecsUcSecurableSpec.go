// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappssettingscustomtemplates


type DataDatabricksAppsSettingsCustomTemplatesTemplatesManifestResourceSpecsUcSecurableSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/apps_settings_custom_templates#permission DataDatabricksAppsSettingsCustomTemplates#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/apps_settings_custom_templates#securable_type DataDatabricksAppsSettingsCustomTemplates#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
}

