// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmlflowmodel


type DataDatabricksMlflowModelLatestVersions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#creation_timestamp DataDatabricksMlflowModel#creation_timestamp}.
	CreationTimestamp *float64 `field:"optional" json:"creationTimestamp" yaml:"creationTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#current_stage DataDatabricksMlflowModel#current_stage}.
	CurrentStage *string `field:"optional" json:"currentStage" yaml:"currentStage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#description DataDatabricksMlflowModel#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#last_updated_timestamp DataDatabricksMlflowModel#last_updated_timestamp}.
	LastUpdatedTimestamp *float64 `field:"optional" json:"lastUpdatedTimestamp" yaml:"lastUpdatedTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#name DataDatabricksMlflowModel#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#run_id DataDatabricksMlflowModel#run_id}.
	RunId *string `field:"optional" json:"runId" yaml:"runId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#run_link DataDatabricksMlflowModel#run_link}.
	RunLink *string `field:"optional" json:"runLink" yaml:"runLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#source DataDatabricksMlflowModel#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#status DataDatabricksMlflowModel#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#status_message DataDatabricksMlflowModel#status_message}.
	StatusMessage *string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#tags DataDatabricksMlflowModel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#user_id DataDatabricksMlflowModel#user_id}.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/mlflow_model#version DataDatabricksMlflowModel#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

