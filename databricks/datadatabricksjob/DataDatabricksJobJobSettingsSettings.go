// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettings struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#compute DataDatabricksJob#compute}
	Compute interface{} `field:"optional" json:"compute" yaml:"compute"`
	// continuous block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#continuous DataDatabricksJob#continuous}
	Continuous *DataDatabricksJobJobSettingsSettingsContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// dbt_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#dbt_task DataDatabricksJob#dbt_task}
	DbtTask *DataDatabricksJobJobSettingsSettingsDbtTask `field:"optional" json:"dbtTask" yaml:"dbtTask"`
	// email_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#email_notifications DataDatabricksJob#email_notifications}
	EmailNotifications *DataDatabricksJobJobSettingsSettingsEmailNotifications `field:"optional" json:"emailNotifications" yaml:"emailNotifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#existing_cluster_id DataDatabricksJob#existing_cluster_id}.
	ExistingClusterId *string `field:"optional" json:"existingClusterId" yaml:"existingClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#format DataDatabricksJob#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// git_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#git_source DataDatabricksJob#git_source}
	GitSource *DataDatabricksJobJobSettingsSettingsGitSource `field:"optional" json:"gitSource" yaml:"gitSource"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#health DataDatabricksJob#health}
	Health *DataDatabricksJobJobSettingsSettingsHealth `field:"optional" json:"health" yaml:"health"`
	// job_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#job_cluster DataDatabricksJob#job_cluster}
	JobCluster interface{} `field:"optional" json:"jobCluster" yaml:"jobCluster"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#library DataDatabricksJob#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#max_concurrent_runs DataDatabricksJob#max_concurrent_runs}.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#max_retries DataDatabricksJob#max_retries}.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#min_retry_interval_millis DataDatabricksJob#min_retry_interval_millis}.
	MinRetryIntervalMillis *float64 `field:"optional" json:"minRetryIntervalMillis" yaml:"minRetryIntervalMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#name DataDatabricksJob#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#new_cluster DataDatabricksJob#new_cluster}
	NewCluster *DataDatabricksJobJobSettingsSettingsNewCluster `field:"optional" json:"newCluster" yaml:"newCluster"`
	// notebook_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#notebook_task DataDatabricksJob#notebook_task}
	NotebookTask *DataDatabricksJobJobSettingsSettingsNotebookTask `field:"optional" json:"notebookTask" yaml:"notebookTask"`
	// notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#notification_settings DataDatabricksJob#notification_settings}
	NotificationSettings *DataDatabricksJobJobSettingsSettingsNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#parameter DataDatabricksJob#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// pipeline_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#pipeline_task DataDatabricksJob#pipeline_task}
	PipelineTask *DataDatabricksJobJobSettingsSettingsPipelineTask `field:"optional" json:"pipelineTask" yaml:"pipelineTask"`
	// python_wheel_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#python_wheel_task DataDatabricksJob#python_wheel_task}
	PythonWheelTask *DataDatabricksJobJobSettingsSettingsPythonWheelTask `field:"optional" json:"pythonWheelTask" yaml:"pythonWheelTask"`
	// queue block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#queue DataDatabricksJob#queue}
	Queue *DataDatabricksJobJobSettingsSettingsQueue `field:"optional" json:"queue" yaml:"queue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#retry_on_timeout DataDatabricksJob#retry_on_timeout}.
	RetryOnTimeout interface{} `field:"optional" json:"retryOnTimeout" yaml:"retryOnTimeout"`
	// run_as block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#run_as DataDatabricksJob#run_as}
	RunAs *DataDatabricksJobJobSettingsSettingsRunAs `field:"optional" json:"runAs" yaml:"runAs"`
	// run_job_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#run_job_task DataDatabricksJob#run_job_task}
	RunJobTask *DataDatabricksJobJobSettingsSettingsRunJobTask `field:"optional" json:"runJobTask" yaml:"runJobTask"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#schedule DataDatabricksJob#schedule}
	Schedule *DataDatabricksJobJobSettingsSettingsSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// spark_jar_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#spark_jar_task DataDatabricksJob#spark_jar_task}
	SparkJarTask *DataDatabricksJobJobSettingsSettingsSparkJarTask `field:"optional" json:"sparkJarTask" yaml:"sparkJarTask"`
	// spark_python_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#spark_python_task DataDatabricksJob#spark_python_task}
	SparkPythonTask *DataDatabricksJobJobSettingsSettingsSparkPythonTask `field:"optional" json:"sparkPythonTask" yaml:"sparkPythonTask"`
	// spark_submit_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#spark_submit_task DataDatabricksJob#spark_submit_task}
	SparkSubmitTask *DataDatabricksJobJobSettingsSettingsSparkSubmitTask `field:"optional" json:"sparkSubmitTask" yaml:"sparkSubmitTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#tags DataDatabricksJob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#task DataDatabricksJob#task}
	Task interface{} `field:"optional" json:"task" yaml:"task"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#timeout_seconds DataDatabricksJob#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#trigger DataDatabricksJob#trigger}
	Trigger *DataDatabricksJobJobSettingsSettingsTrigger `field:"optional" json:"trigger" yaml:"trigger"`
	// webhook_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.26.0/docs/data-sources/job#webhook_notifications DataDatabricksJob#webhook_notifications}
	WebhookNotifications *DataDatabricksJobJobSettingsSettingsWebhookNotifications `field:"optional" json:"webhookNotifications" yaml:"webhookNotifications"`
}

