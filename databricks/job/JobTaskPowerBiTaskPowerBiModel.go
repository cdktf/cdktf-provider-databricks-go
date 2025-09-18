// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskPowerBiTaskPowerBiModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#authentication_method Job#authentication_method}.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#model_name Job#model_name}.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#overwrite_existing Job#overwrite_existing}.
	OverwriteExisting interface{} `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#storage_mode Job#storage_mode}.
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#workspace_name Job#workspace_name}.
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

