// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssharepluginframework


type DataDatabricksSharePluginframeworkObject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#name DataDatabricksSharePluginframework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#added_at DataDatabricksSharePluginframework#added_at}.
	AddedAt *float64 `field:"optional" json:"addedAt" yaml:"addedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#added_by DataDatabricksSharePluginframework#added_by}.
	AddedBy *string `field:"optional" json:"addedBy" yaml:"addedBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#cdf_enabled DataDatabricksSharePluginframework#cdf_enabled}.
	CdfEnabled interface{} `field:"optional" json:"cdfEnabled" yaml:"cdfEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#comment DataDatabricksSharePluginframework#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#content DataDatabricksSharePluginframework#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#data_object_type DataDatabricksSharePluginframework#data_object_type}.
	DataObjectType *string `field:"optional" json:"dataObjectType" yaml:"dataObjectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#effective_cdf_enabled DataDatabricksSharePluginframework#effective_cdf_enabled}.
	EffectiveCdfEnabled interface{} `field:"optional" json:"effectiveCdfEnabled" yaml:"effectiveCdfEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#effective_history_data_sharing_status DataDatabricksSharePluginframework#effective_history_data_sharing_status}.
	EffectiveHistoryDataSharingStatus *string `field:"optional" json:"effectiveHistoryDataSharingStatus" yaml:"effectiveHistoryDataSharingStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#effective_shared_as DataDatabricksSharePluginframework#effective_shared_as}.
	EffectiveSharedAs *string `field:"optional" json:"effectiveSharedAs" yaml:"effectiveSharedAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#effective_start_version DataDatabricksSharePluginframework#effective_start_version}.
	EffectiveStartVersion *float64 `field:"optional" json:"effectiveStartVersion" yaml:"effectiveStartVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#history_data_sharing_status DataDatabricksSharePluginframework#history_data_sharing_status}.
	HistoryDataSharingStatus *string `field:"optional" json:"historyDataSharingStatus" yaml:"historyDataSharingStatus"`
	// partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#partition DataDatabricksSharePluginframework#partition}
	Partition interface{} `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#shared_as DataDatabricksSharePluginframework#shared_as}.
	SharedAs *string `field:"optional" json:"sharedAs" yaml:"sharedAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#start_version DataDatabricksSharePluginframework#start_version}.
	StartVersion *float64 `field:"optional" json:"startVersion" yaml:"startVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#status DataDatabricksSharePluginframework#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/share_pluginframework#string_shared_as DataDatabricksSharePluginframework#string_shared_as}.
	StringSharedAs *string `field:"optional" json:"stringSharedAs" yaml:"stringSharedAs"`
}
