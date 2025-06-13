// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskRunJobTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#job_id Job#job_id}.
	JobId *float64 `field:"required" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#dbt_commands Job#dbt_commands}.
	DbtCommands *[]*string `field:"optional" json:"dbtCommands" yaml:"dbtCommands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#jar_params Job#jar_params}.
	JarParams *[]*string `field:"optional" json:"jarParams" yaml:"jarParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#job_parameters Job#job_parameters}.
	JobParameters *map[string]*string `field:"optional" json:"jobParameters" yaml:"jobParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#notebook_params Job#notebook_params}.
	NotebookParams *map[string]*string `field:"optional" json:"notebookParams" yaml:"notebookParams"`
	// pipeline_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#pipeline_params Job#pipeline_params}
	PipelineParams *JobTaskForEachTaskTaskRunJobTaskPipelineParams `field:"optional" json:"pipelineParams" yaml:"pipelineParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#python_named_params Job#python_named_params}.
	PythonNamedParams *map[string]*string `field:"optional" json:"pythonNamedParams" yaml:"pythonNamedParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#python_params Job#python_params}.
	PythonParams *[]*string `field:"optional" json:"pythonParams" yaml:"pythonParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#spark_submit_params Job#spark_submit_params}.
	SparkSubmitParams *[]*string `field:"optional" json:"sparkSubmitParams" yaml:"sparkSubmitParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/job#sql_params Job#sql_params}.
	SqlParams *map[string]*string `field:"optional" json:"sqlParams" yaml:"sqlParams"`
}

