// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskGenAiComputeTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#dl_runtime_image Job#dl_runtime_image}.
	DlRuntimeImage *string `field:"required" json:"dlRuntimeImage" yaml:"dlRuntimeImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#command Job#command}.
	Command *string `field:"optional" json:"command" yaml:"command"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#compute Job#compute}
	Compute *JobTaskGenAiComputeTaskCompute `field:"optional" json:"compute" yaml:"compute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#mlflow_experiment_name Job#mlflow_experiment_name}.
	MlflowExperimentName *string `field:"optional" json:"mlflowExperimentName" yaml:"mlflowExperimentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#source Job#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#training_script_path Job#training_script_path}.
	TrainingScriptPath *string `field:"optional" json:"trainingScriptPath" yaml:"trainingScriptPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#yaml_parameters Job#yaml_parameters}.
	YamlParameters *string `field:"optional" json:"yamlParameters" yaml:"yamlParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/job#yaml_parameters_file_path Job#yaml_parameters_file_path}.
	YamlParametersFilePath *string `field:"optional" json:"yamlParametersFilePath" yaml:"yamlParametersFilePath"`
}

