// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppAppActiveDeployment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#create_time DataDatabricksApp#create_time}.
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#creator DataDatabricksApp#creator}.
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#deployment_artifacts DataDatabricksApp#deployment_artifacts}.
	DeploymentArtifacts *DataDatabricksAppAppActiveDeploymentDeploymentArtifacts `field:"optional" json:"deploymentArtifacts" yaml:"deploymentArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#deployment_id DataDatabricksApp#deployment_id}.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#mode DataDatabricksApp#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#source_code_path DataDatabricksApp#source_code_path}.
	SourceCodePath *string `field:"optional" json:"sourceCodePath" yaml:"sourceCodePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#status DataDatabricksApp#status}.
	Status *DataDatabricksAppAppActiveDeploymentStatus `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#update_time DataDatabricksApp#update_time}.
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
}

