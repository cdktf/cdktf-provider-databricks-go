// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appssettingscustomtemplate


type AppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/apps_settings_custom_template#permission AppsSettingsCustomTemplate#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/apps_settings_custom_template#securable_type AppsSettingsCustomTemplate#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
}

