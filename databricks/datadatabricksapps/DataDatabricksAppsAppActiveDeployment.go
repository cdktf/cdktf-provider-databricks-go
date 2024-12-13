// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppActiveDeployment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#create_time DataDatabricksApps#create_time}.
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#creator DataDatabricksApps#creator}.
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#deployment_artifacts DataDatabricksApps#deployment_artifacts}.
	DeploymentArtifacts *DataDatabricksAppsAppActiveDeploymentDeploymentArtifacts `field:"optional" json:"deploymentArtifacts" yaml:"deploymentArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#deployment_id DataDatabricksApps#deployment_id}.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#mode DataDatabricksApps#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#source_code_path DataDatabricksApps#source_code_path}.
	SourceCodePath *string `field:"optional" json:"sourceCodePath" yaml:"sourceCodePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#status DataDatabricksApps#status}.
	Status *DataDatabricksAppsAppActiveDeploymentStatus `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#update_time DataDatabricksApps#update_time}.
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
}

