// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppActiveDeployment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#create_time App#create_time}.
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#creator App#creator}.
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#deployment_artifacts App#deployment_artifacts}.
	DeploymentArtifacts *AppActiveDeploymentDeploymentArtifacts `field:"optional" json:"deploymentArtifacts" yaml:"deploymentArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#deployment_id App#deployment_id}.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#mode App#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#source_code_path App#source_code_path}.
	SourceCodePath *string `field:"optional" json:"sourceCodePath" yaml:"sourceCodePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#status App#status}.
	Status *AppActiveDeploymentStatus `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#update_time App#update_time}.
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
}

