// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharepluginframework


type SharePluginframeworkObject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#data_object_type SharePluginframework#data_object_type}.
	DataObjectType *string `field:"required" json:"dataObjectType" yaml:"dataObjectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#name SharePluginframework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#added_at SharePluginframework#added_at}.
	AddedAt *float64 `field:"optional" json:"addedAt" yaml:"addedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#added_by SharePluginframework#added_by}.
	AddedBy *string `field:"optional" json:"addedBy" yaml:"addedBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#cdf_enabled SharePluginframework#cdf_enabled}.
	CdfEnabled interface{} `field:"optional" json:"cdfEnabled" yaml:"cdfEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#comment SharePluginframework#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#content SharePluginframework#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#effective_cdf_enabled SharePluginframework#effective_cdf_enabled}.
	EffectiveCdfEnabled interface{} `field:"optional" json:"effectiveCdfEnabled" yaml:"effectiveCdfEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#effective_history_data_sharing_status SharePluginframework#effective_history_data_sharing_status}.
	EffectiveHistoryDataSharingStatus *string `field:"optional" json:"effectiveHistoryDataSharingStatus" yaml:"effectiveHistoryDataSharingStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#effective_shared_as SharePluginframework#effective_shared_as}.
	EffectiveSharedAs *string `field:"optional" json:"effectiveSharedAs" yaml:"effectiveSharedAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#effective_start_version SharePluginframework#effective_start_version}.
	EffectiveStartVersion *float64 `field:"optional" json:"effectiveStartVersion" yaml:"effectiveStartVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#history_data_sharing_status SharePluginframework#history_data_sharing_status}.
	HistoryDataSharingStatus *string `field:"optional" json:"historyDataSharingStatus" yaml:"historyDataSharingStatus"`
	// partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#partition SharePluginframework#partition}
	Partition interface{} `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#shared_as SharePluginframework#shared_as}.
	SharedAs *string `field:"optional" json:"sharedAs" yaml:"sharedAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#start_version SharePluginframework#start_version}.
	StartVersion *float64 `field:"optional" json:"startVersion" yaml:"startVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#status SharePluginframework#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/share_pluginframework#string_shared_as SharePluginframework#string_shared_as}.
	StringSharedAs *string `field:"optional" json:"stringSharedAs" yaml:"stringSharedAs"`
}

