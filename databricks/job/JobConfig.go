package job

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#always_running Job#always_running}.
	AlwaysRunning interface{} `field:"optional" json:"alwaysRunning" yaml:"alwaysRunning"`
	// continuous block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#continuous Job#continuous}
	Continuous *JobContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// dbt_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#dbt_task Job#dbt_task}
	DbtTask *JobDbtTask `field:"optional" json:"dbtTask" yaml:"dbtTask"`
	// email_notifications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#email_notifications Job#email_notifications}
	EmailNotifications *JobEmailNotifications `field:"optional" json:"emailNotifications" yaml:"emailNotifications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#existing_cluster_id Job#existing_cluster_id}.
	ExistingClusterId *string `field:"optional" json:"existingClusterId" yaml:"existingClusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#format Job#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// git_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#git_source Job#git_source}
	GitSource *JobGitSource `field:"optional" json:"gitSource" yaml:"gitSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#id Job#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// job_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#job_cluster Job#job_cluster}
	JobCluster interface{} `field:"optional" json:"jobCluster" yaml:"jobCluster"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#library Job#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#max_concurrent_runs Job#max_concurrent_runs}.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#max_retries Job#max_retries}.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#min_retry_interval_millis Job#min_retry_interval_millis}.
	MinRetryIntervalMillis *float64 `field:"optional" json:"minRetryIntervalMillis" yaml:"minRetryIntervalMillis"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#name Job#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#new_cluster Job#new_cluster}
	NewCluster *JobNewCluster `field:"optional" json:"newCluster" yaml:"newCluster"`
	// notebook_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#notebook_task Job#notebook_task}
	NotebookTask *JobNotebookTask `field:"optional" json:"notebookTask" yaml:"notebookTask"`
	// pipeline_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#pipeline_task Job#pipeline_task}
	PipelineTask *JobPipelineTask `field:"optional" json:"pipelineTask" yaml:"pipelineTask"`
	// python_wheel_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#python_wheel_task Job#python_wheel_task}
	PythonWheelTask *JobPythonWheelTask `field:"optional" json:"pythonWheelTask" yaml:"pythonWheelTask"`
	// queue block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#queue Job#queue}
	Queue *JobQueue `field:"optional" json:"queue" yaml:"queue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#retry_on_timeout Job#retry_on_timeout}.
	RetryOnTimeout interface{} `field:"optional" json:"retryOnTimeout" yaml:"retryOnTimeout"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#schedule Job#schedule}
	Schedule *JobSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// spark_jar_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#spark_jar_task Job#spark_jar_task}
	SparkJarTask *JobSparkJarTask `field:"optional" json:"sparkJarTask" yaml:"sparkJarTask"`
	// spark_python_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#spark_python_task Job#spark_python_task}
	SparkPythonTask *JobSparkPythonTask `field:"optional" json:"sparkPythonTask" yaml:"sparkPythonTask"`
	// spark_submit_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#spark_submit_task Job#spark_submit_task}
	SparkSubmitTask *JobSparkSubmitTask `field:"optional" json:"sparkSubmitTask" yaml:"sparkSubmitTask"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#tags Job#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#task Job#task}
	Task interface{} `field:"optional" json:"task" yaml:"task"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#timeouts Job#timeouts}
	Timeouts *JobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#timeout_seconds Job#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#trigger Job#trigger}
	Trigger *JobTrigger `field:"optional" json:"trigger" yaml:"trigger"`
	// webhook_notifications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#webhook_notifications Job#webhook_notifications}
	WebhookNotifications *JobWebhookNotifications `field:"optional" json:"webhookNotifications" yaml:"webhookNotifications"`
}

