// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksshare


type DataDatabricksShareObject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#name DataDatabricksShare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#cdf_enabled DataDatabricksShare#cdf_enabled}.
	CdfEnabled interface{} `field:"optional" json:"cdfEnabled" yaml:"cdfEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#comment DataDatabricksShare#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#content DataDatabricksShare#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#data_object_type DataDatabricksShare#data_object_type}.
	DataObjectType *string `field:"optional" json:"dataObjectType" yaml:"dataObjectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#history_data_sharing_status DataDatabricksShare#history_data_sharing_status}.
	HistoryDataSharingStatus *string `field:"optional" json:"historyDataSharingStatus" yaml:"historyDataSharingStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#partition DataDatabricksShare#partition}.
	Partition interface{} `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#shared_as DataDatabricksShare#shared_as}.
	SharedAs *string `field:"optional" json:"sharedAs" yaml:"sharedAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#start_version DataDatabricksShare#start_version}.
	StartVersion *float64 `field:"optional" json:"startVersion" yaml:"startVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#string_shared_as DataDatabricksShare#string_shared_as}.
	StringSharedAs *string `field:"optional" json:"stringSharedAs" yaml:"stringSharedAs"`
}

