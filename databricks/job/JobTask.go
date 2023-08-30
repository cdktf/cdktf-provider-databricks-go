// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#compute_key Job#compute_key}.
	ComputeKey *string `field:"optional" json:"computeKey" yaml:"computeKey"`
	// condition_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#condition_task Job#condition_task}
	ConditionTask *JobTaskConditionTask `field:"optional" json:"conditionTask" yaml:"conditionTask"`
	// dbt_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#dbt_task Job#dbt_task}
	DbtTask *JobTaskDbtTask `field:"optional" json:"dbtTask" yaml:"dbtTask"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#depends_on Job#depends_on}
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#description Job#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// email_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#email_notifications Job#email_notifications}
	EmailNotifications *JobTaskEmailNotifications `field:"optional" json:"emailNotifications" yaml:"emailNotifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#existing_cluster_id Job#existing_cluster_id}.
	ExistingClusterId *string `field:"optional" json:"existingClusterId" yaml:"existingClusterId"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#health Job#health}
	Health *JobTaskHealth `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#job_cluster_key Job#job_cluster_key}.
	JobClusterKey *string `field:"optional" json:"jobClusterKey" yaml:"jobClusterKey"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#library Job#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#max_retries Job#max_retries}.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#min_retry_interval_millis Job#min_retry_interval_millis}.
	MinRetryIntervalMillis *float64 `field:"optional" json:"minRetryIntervalMillis" yaml:"minRetryIntervalMillis"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#new_cluster Job#new_cluster}
	NewCluster *JobTaskNewCluster `field:"optional" json:"newCluster" yaml:"newCluster"`
	// notebook_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#notebook_task Job#notebook_task}
	NotebookTask *JobTaskNotebookTask `field:"optional" json:"notebookTask" yaml:"notebookTask"`
	// notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#notification_settings Job#notification_settings}
	NotificationSettings *JobTaskNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// pipeline_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#pipeline_task Job#pipeline_task}
	PipelineTask *JobTaskPipelineTask `field:"optional" json:"pipelineTask" yaml:"pipelineTask"`
	// python_wheel_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#python_wheel_task Job#python_wheel_task}
	PythonWheelTask *JobTaskPythonWheelTask `field:"optional" json:"pythonWheelTask" yaml:"pythonWheelTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#retry_on_timeout Job#retry_on_timeout}.
	RetryOnTimeout interface{} `field:"optional" json:"retryOnTimeout" yaml:"retryOnTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#run_if Job#run_if}.
	RunIf *string `field:"optional" json:"runIf" yaml:"runIf"`
	// run_job_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#run_job_task Job#run_job_task}
	RunJobTask *JobTaskRunJobTask `field:"optional" json:"runJobTask" yaml:"runJobTask"`
	// spark_jar_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#spark_jar_task Job#spark_jar_task}
	SparkJarTask *JobTaskSparkJarTask `field:"optional" json:"sparkJarTask" yaml:"sparkJarTask"`
	// spark_python_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#spark_python_task Job#spark_python_task}
	SparkPythonTask *JobTaskSparkPythonTask `field:"optional" json:"sparkPythonTask" yaml:"sparkPythonTask"`
	// spark_submit_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#spark_submit_task Job#spark_submit_task}
	SparkSubmitTask *JobTaskSparkSubmitTask `field:"optional" json:"sparkSubmitTask" yaml:"sparkSubmitTask"`
	// sql_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#sql_task Job#sql_task}
	SqlTask *JobTaskSqlTask `field:"optional" json:"sqlTask" yaml:"sqlTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#task_key Job#task_key}.
	TaskKey *string `field:"optional" json:"taskKey" yaml:"taskKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#timeout_seconds Job#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

