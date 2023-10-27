// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksshare


type DataDatabricksShareObject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#data_object_type DataDatabricksShare#data_object_type}.
	DataObjectType *string `field:"required" json:"dataObjectType" yaml:"dataObjectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#name DataDatabricksShare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#added_at DataDatabricksShare#added_at}.
	AddedAt *float64 `field:"optional" json:"addedAt" yaml:"addedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#added_by DataDatabricksShare#added_by}.
	AddedBy *string `field:"optional" json:"addedBy" yaml:"addedBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#cdf_enabled DataDatabricksShare#cdf_enabled}.
	CdfEnabled interface{} `field:"optional" json:"cdfEnabled" yaml:"cdfEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#comment DataDatabricksShare#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#history_data_sharing_status DataDatabricksShare#history_data_sharing_status}.
	HistoryDataSharingStatus *string `field:"optional" json:"historyDataSharingStatus" yaml:"historyDataSharingStatus"`
	// partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#partition DataDatabricksShare#partition}
	Partition interface{} `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#shared_as DataDatabricksShare#shared_as}.
	SharedAs *string `field:"optional" json:"sharedAs" yaml:"sharedAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#start_version DataDatabricksShare#start_version}.
	StartVersion *float64 `field:"optional" json:"startVersion" yaml:"startVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/share#status DataDatabricksShare#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

