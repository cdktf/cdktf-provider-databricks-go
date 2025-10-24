// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskPowerBiTaskPowerBiModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#authentication_method DataDatabricksJob#authentication_method}.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#model_name DataDatabricksJob#model_name}.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#overwrite_existing DataDatabricksJob#overwrite_existing}.
	OverwriteExisting interface{} `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#storage_mode DataDatabricksJob#storage_mode}.
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#workspace_name DataDatabricksJob#workspace_name}.
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

