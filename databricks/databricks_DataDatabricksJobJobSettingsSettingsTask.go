// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsTask struct {
	// dbt_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#dbt_task DataDatabricksJob#dbt_task}
	DbtTask *DataDatabricksJobJobSettingsSettingsTaskDbtTask `field:"optional" json:"dbtTask" yaml:"dbtTask"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#depends_on DataDatabricksJob#depends_on}
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#description DataDatabricksJob#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// email_notifications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#email_notifications DataDatabricksJob#email_notifications}
	EmailNotifications *DataDatabricksJobJobSettingsSettingsTaskEmailNotifications `field:"optional" json:"emailNotifications" yaml:"emailNotifications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#existing_cluster_id DataDatabricksJob#existing_cluster_id}.
	ExistingClusterId *string `field:"optional" json:"existingClusterId" yaml:"existingClusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#job_cluster_key DataDatabricksJob#job_cluster_key}.
	JobClusterKey *string `field:"optional" json:"jobClusterKey" yaml:"jobClusterKey"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#library DataDatabricksJob#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#max_retries DataDatabricksJob#max_retries}.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#min_retry_interval_millis DataDatabricksJob#min_retry_interval_millis}.
	MinRetryIntervalMillis *float64 `field:"optional" json:"minRetryIntervalMillis" yaml:"minRetryIntervalMillis"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#new_cluster DataDatabricksJob#new_cluster}
	NewCluster *DataDatabricksJobJobSettingsSettingsTaskNewCluster `field:"optional" json:"newCluster" yaml:"newCluster"`
	// notebook_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#notebook_task DataDatabricksJob#notebook_task}
	NotebookTask *DataDatabricksJobJobSettingsSettingsTaskNotebookTask `field:"optional" json:"notebookTask" yaml:"notebookTask"`
	// pipeline_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#pipeline_task DataDatabricksJob#pipeline_task}
	PipelineTask *DataDatabricksJobJobSettingsSettingsTaskPipelineTask `field:"optional" json:"pipelineTask" yaml:"pipelineTask"`
	// python_wheel_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#python_wheel_task DataDatabricksJob#python_wheel_task}
	PythonWheelTask *DataDatabricksJobJobSettingsSettingsTaskPythonWheelTask `field:"optional" json:"pythonWheelTask" yaml:"pythonWheelTask"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#retry_on_timeout DataDatabricksJob#retry_on_timeout}.
	RetryOnTimeout interface{} `field:"optional" json:"retryOnTimeout" yaml:"retryOnTimeout"`
	// spark_jar_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#spark_jar_task DataDatabricksJob#spark_jar_task}
	SparkJarTask *DataDatabricksJobJobSettingsSettingsTaskSparkJarTask `field:"optional" json:"sparkJarTask" yaml:"sparkJarTask"`
	// spark_python_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#spark_python_task DataDatabricksJob#spark_python_task}
	SparkPythonTask *DataDatabricksJobJobSettingsSettingsTaskSparkPythonTask `field:"optional" json:"sparkPythonTask" yaml:"sparkPythonTask"`
	// spark_submit_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#spark_submit_task DataDatabricksJob#spark_submit_task}
	SparkSubmitTask *DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTask `field:"optional" json:"sparkSubmitTask" yaml:"sparkSubmitTask"`
	// sql_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#sql_task DataDatabricksJob#sql_task}
	SqlTask *DataDatabricksJobJobSettingsSettingsTaskSqlTask `field:"optional" json:"sqlTask" yaml:"sqlTask"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#task_key DataDatabricksJob#task_key}.
	TaskKey *string `field:"optional" json:"taskKey" yaml:"taskKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#timeout_seconds DataDatabricksJob#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

