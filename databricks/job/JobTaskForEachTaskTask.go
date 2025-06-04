// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#task_key Job#task_key}.
	TaskKey *string `field:"required" json:"taskKey" yaml:"taskKey"`
	// clean_rooms_notebook_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#clean_rooms_notebook_task Job#clean_rooms_notebook_task}
	CleanRoomsNotebookTask *JobTaskForEachTaskTaskCleanRoomsNotebookTask `field:"optional" json:"cleanRoomsNotebookTask" yaml:"cleanRoomsNotebookTask"`
	// condition_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#condition_task Job#condition_task}
	ConditionTask *JobTaskForEachTaskTaskConditionTask `field:"optional" json:"conditionTask" yaml:"conditionTask"`
	// dashboard_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#dashboard_task Job#dashboard_task}
	DashboardTask *JobTaskForEachTaskTaskDashboardTask `field:"optional" json:"dashboardTask" yaml:"dashboardTask"`
	// dbt_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#dbt_task Job#dbt_task}
	DbtTask *JobTaskForEachTaskTaskDbtTask `field:"optional" json:"dbtTask" yaml:"dbtTask"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#depends_on Job#depends_on}
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#description Job#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#disable_auto_optimization Job#disable_auto_optimization}.
	DisableAutoOptimization interface{} `field:"optional" json:"disableAutoOptimization" yaml:"disableAutoOptimization"`
	// email_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#email_notifications Job#email_notifications}
	EmailNotifications *JobTaskForEachTaskTaskEmailNotifications `field:"optional" json:"emailNotifications" yaml:"emailNotifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#environment_key Job#environment_key}.
	EnvironmentKey *string `field:"optional" json:"environmentKey" yaml:"environmentKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#existing_cluster_id Job#existing_cluster_id}.
	ExistingClusterId *string `field:"optional" json:"existingClusterId" yaml:"existingClusterId"`
	// gen_ai_compute_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#gen_ai_compute_task Job#gen_ai_compute_task}
	GenAiComputeTask *JobTaskForEachTaskTaskGenAiComputeTask `field:"optional" json:"genAiComputeTask" yaml:"genAiComputeTask"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#health Job#health}
	Health *JobTaskForEachTaskTaskHealth `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#job_cluster_key Job#job_cluster_key}.
	JobClusterKey *string `field:"optional" json:"jobClusterKey" yaml:"jobClusterKey"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#library Job#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#max_retries Job#max_retries}.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#min_retry_interval_millis Job#min_retry_interval_millis}.
	MinRetryIntervalMillis *float64 `field:"optional" json:"minRetryIntervalMillis" yaml:"minRetryIntervalMillis"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#new_cluster Job#new_cluster}
	NewCluster *JobTaskForEachTaskTaskNewCluster `field:"optional" json:"newCluster" yaml:"newCluster"`
	// notebook_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#notebook_task Job#notebook_task}
	NotebookTask *JobTaskForEachTaskTaskNotebookTask `field:"optional" json:"notebookTask" yaml:"notebookTask"`
	// notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#notification_settings Job#notification_settings}
	NotificationSettings *JobTaskForEachTaskTaskNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// pipeline_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#pipeline_task Job#pipeline_task}
	PipelineTask *JobTaskForEachTaskTaskPipelineTask `field:"optional" json:"pipelineTask" yaml:"pipelineTask"`
	// power_bi_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#power_bi_task Job#power_bi_task}
	PowerBiTask *JobTaskForEachTaskTaskPowerBiTask `field:"optional" json:"powerBiTask" yaml:"powerBiTask"`
	// python_wheel_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#python_wheel_task Job#python_wheel_task}
	PythonWheelTask *JobTaskForEachTaskTaskPythonWheelTask `field:"optional" json:"pythonWheelTask" yaml:"pythonWheelTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#retry_on_timeout Job#retry_on_timeout}.
	RetryOnTimeout interface{} `field:"optional" json:"retryOnTimeout" yaml:"retryOnTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#run_if Job#run_if}.
	RunIf *string `field:"optional" json:"runIf" yaml:"runIf"`
	// run_job_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#run_job_task Job#run_job_task}
	RunJobTask *JobTaskForEachTaskTaskRunJobTask `field:"optional" json:"runJobTask" yaml:"runJobTask"`
	// spark_jar_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#spark_jar_task Job#spark_jar_task}
	SparkJarTask *JobTaskForEachTaskTaskSparkJarTask `field:"optional" json:"sparkJarTask" yaml:"sparkJarTask"`
	// spark_python_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#spark_python_task Job#spark_python_task}
	SparkPythonTask *JobTaskForEachTaskTaskSparkPythonTask `field:"optional" json:"sparkPythonTask" yaml:"sparkPythonTask"`
	// spark_submit_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#spark_submit_task Job#spark_submit_task}
	SparkSubmitTask *JobTaskForEachTaskTaskSparkSubmitTask `field:"optional" json:"sparkSubmitTask" yaml:"sparkSubmitTask"`
	// sql_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#sql_task Job#sql_task}
	SqlTask *JobTaskForEachTaskTaskSqlTask `field:"optional" json:"sqlTask" yaml:"sqlTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#timeout_seconds Job#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// webhook_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job#webhook_notifications Job#webhook_notifications}
	WebhookNotifications *JobTaskForEachTaskTaskWebhookNotifications `field:"optional" json:"webhookNotifications" yaml:"webhookNotifications"`
}

