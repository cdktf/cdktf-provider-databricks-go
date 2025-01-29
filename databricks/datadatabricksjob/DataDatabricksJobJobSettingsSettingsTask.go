// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#task_key DataDatabricksJob#task_key}.
	TaskKey *string `field:"required" json:"taskKey" yaml:"taskKey"`
	// condition_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#condition_task DataDatabricksJob#condition_task}
	ConditionTask *DataDatabricksJobJobSettingsSettingsTaskConditionTask `field:"optional" json:"conditionTask" yaml:"conditionTask"`
	// dbt_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#dbt_task DataDatabricksJob#dbt_task}
	DbtTask *DataDatabricksJobJobSettingsSettingsTaskDbtTask `field:"optional" json:"dbtTask" yaml:"dbtTask"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#depends_on DataDatabricksJob#depends_on}
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#description DataDatabricksJob#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// email_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#email_notifications DataDatabricksJob#email_notifications}
	EmailNotifications *DataDatabricksJobJobSettingsSettingsTaskEmailNotifications `field:"optional" json:"emailNotifications" yaml:"emailNotifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#environment_key DataDatabricksJob#environment_key}.
	EnvironmentKey *string `field:"optional" json:"environmentKey" yaml:"environmentKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#existing_cluster_id DataDatabricksJob#existing_cluster_id}.
	ExistingClusterId *string `field:"optional" json:"existingClusterId" yaml:"existingClusterId"`
	// for_each_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#for_each_task DataDatabricksJob#for_each_task}
	ForEachTask *DataDatabricksJobJobSettingsSettingsTaskForEachTask `field:"optional" json:"forEachTask" yaml:"forEachTask"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#health DataDatabricksJob#health}
	Health *DataDatabricksJobJobSettingsSettingsTaskHealth `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#job_cluster_key DataDatabricksJob#job_cluster_key}.
	JobClusterKey *string `field:"optional" json:"jobClusterKey" yaml:"jobClusterKey"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#library DataDatabricksJob#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#max_retries DataDatabricksJob#max_retries}.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#min_retry_interval_millis DataDatabricksJob#min_retry_interval_millis}.
	MinRetryIntervalMillis *float64 `field:"optional" json:"minRetryIntervalMillis" yaml:"minRetryIntervalMillis"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#new_cluster DataDatabricksJob#new_cluster}
	NewCluster *DataDatabricksJobJobSettingsSettingsTaskNewCluster `field:"optional" json:"newCluster" yaml:"newCluster"`
	// notebook_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#notebook_task DataDatabricksJob#notebook_task}
	NotebookTask *DataDatabricksJobJobSettingsSettingsTaskNotebookTask `field:"optional" json:"notebookTask" yaml:"notebookTask"`
	// notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#notification_settings DataDatabricksJob#notification_settings}
	NotificationSettings *DataDatabricksJobJobSettingsSettingsTaskNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// pipeline_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#pipeline_task DataDatabricksJob#pipeline_task}
	PipelineTask *DataDatabricksJobJobSettingsSettingsTaskPipelineTask `field:"optional" json:"pipelineTask" yaml:"pipelineTask"`
	// python_wheel_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#python_wheel_task DataDatabricksJob#python_wheel_task}
	PythonWheelTask *DataDatabricksJobJobSettingsSettingsTaskPythonWheelTask `field:"optional" json:"pythonWheelTask" yaml:"pythonWheelTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#retry_on_timeout DataDatabricksJob#retry_on_timeout}.
	RetryOnTimeout interface{} `field:"optional" json:"retryOnTimeout" yaml:"retryOnTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#run_if DataDatabricksJob#run_if}.
	RunIf *string `field:"optional" json:"runIf" yaml:"runIf"`
	// run_job_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#run_job_task DataDatabricksJob#run_job_task}
	RunJobTask *DataDatabricksJobJobSettingsSettingsTaskRunJobTask `field:"optional" json:"runJobTask" yaml:"runJobTask"`
	// spark_jar_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#spark_jar_task DataDatabricksJob#spark_jar_task}
	SparkJarTask *DataDatabricksJobJobSettingsSettingsTaskSparkJarTask `field:"optional" json:"sparkJarTask" yaml:"sparkJarTask"`
	// spark_python_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#spark_python_task DataDatabricksJob#spark_python_task}
	SparkPythonTask *DataDatabricksJobJobSettingsSettingsTaskSparkPythonTask `field:"optional" json:"sparkPythonTask" yaml:"sparkPythonTask"`
	// spark_submit_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#spark_submit_task DataDatabricksJob#spark_submit_task}
	SparkSubmitTask *DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTask `field:"optional" json:"sparkSubmitTask" yaml:"sparkSubmitTask"`
	// sql_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#sql_task DataDatabricksJob#sql_task}
	SqlTask *DataDatabricksJobJobSettingsSettingsTaskSqlTask `field:"optional" json:"sqlTask" yaml:"sqlTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#timeout_seconds DataDatabricksJob#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// webhook_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/job#webhook_notifications DataDatabricksJob#webhook_notifications}
	WebhookNotifications *DataDatabricksJobJobSettingsSettingsTaskWebhookNotifications `field:"optional" json:"webhookNotifications" yaml:"webhookNotifications"`
}

