// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job databricks_job}.
type Job interface {
	cdktf.TerraformResource
	AlwaysRunning() interface{}
	SetAlwaysRunning(val interface{})
	AlwaysRunningInput() interface{}
	BudgetPolicyId() *string
	SetBudgetPolicyId(val *string)
	BudgetPolicyIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Continuous() JobContinuousOutputReference
	ContinuousInput() *JobContinuous
	ControlRunState() interface{}
	SetControlRunState(val interface{})
	ControlRunStateInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DbtTask() JobDbtTaskOutputReference
	DbtTaskInput() *JobDbtTask
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() JobDeploymentOutputReference
	DeploymentInput() *JobDeployment
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EditMode() *string
	SetEditMode(val *string)
	EditModeInput() *string
	EmailNotifications() JobEmailNotificationsOutputReference
	EmailNotificationsInput() *JobEmailNotifications
	Environment() JobEnvironmentList
	EnvironmentInput() interface{}
	ExistingClusterId() *string
	SetExistingClusterId(val *string)
	ExistingClusterIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GitSource() JobGitSourceOutputReference
	GitSourceInput() *JobGitSource
	Health() JobHealthOutputReference
	HealthInput() *JobHealth
	Id() *string
	SetId(val *string)
	IdInput() *string
	JobCluster() JobJobClusterList
	JobClusterInput() interface{}
	Library() JobLibraryList
	LibraryInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxConcurrentRuns() *float64
	SetMaxConcurrentRuns(val *float64)
	MaxConcurrentRunsInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MinRetryIntervalMillis() *float64
	SetMinRetryIntervalMillis(val *float64)
	MinRetryIntervalMillisInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewCluster() JobNewClusterOutputReference
	NewClusterInput() *JobNewCluster
	// The tree node.
	Node() constructs.Node
	NotebookTask() JobNotebookTaskOutputReference
	NotebookTaskInput() *JobNotebookTask
	NotificationSettings() JobNotificationSettingsOutputReference
	NotificationSettingsInput() *JobNotificationSettings
	Parameter() JobParameterList
	ParameterInput() interface{}
	PerformanceTarget() *string
	SetPerformanceTarget(val *string)
	PerformanceTargetInput() *string
	PipelineTask() JobPipelineTaskOutputReference
	PipelineTaskInput() *JobPipelineTask
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PythonWheelTask() JobPythonWheelTaskOutputReference
	PythonWheelTaskInput() *JobPythonWheelTask
	Queue() JobQueueOutputReference
	QueueInput() *JobQueue
	// Experimental.
	RawOverrides() interface{}
	RetryOnTimeout() interface{}
	SetRetryOnTimeout(val interface{})
	RetryOnTimeoutInput() interface{}
	RunAs() JobRunAsOutputReference
	RunAsInput() *JobRunAs
	RunJobTask() JobRunJobTaskOutputReference
	RunJobTaskInput() *JobRunJobTask
	Schedule() JobScheduleOutputReference
	ScheduleInput() *JobSchedule
	SparkJarTask() JobSparkJarTaskOutputReference
	SparkJarTaskInput() *JobSparkJarTask
	SparkPythonTask() JobSparkPythonTaskOutputReference
	SparkPythonTaskInput() *JobSparkPythonTask
	SparkSubmitTask() JobSparkSubmitTaskOutputReference
	SparkSubmitTaskInput() *JobSparkSubmitTask
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Task() JobTaskList
	TaskInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() JobTimeoutsOutputReference
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
	TimeoutsInput() interface{}
	Trigger() JobTriggerOutputReference
	TriggerInput() *JobTrigger
	Url() *string
	WebhookNotifications() JobWebhookNotificationsOutputReference
	WebhookNotificationsInput() *JobWebhookNotifications
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutContinuous(value *JobContinuous)
	PutDbtTask(value *JobDbtTask)
	PutDeployment(value *JobDeployment)
	PutEmailNotifications(value *JobEmailNotifications)
	PutEnvironment(value interface{})
	PutGitSource(value *JobGitSource)
	PutHealth(value *JobHealth)
	PutJobCluster(value interface{})
	PutLibrary(value interface{})
	PutNewCluster(value *JobNewCluster)
	PutNotebookTask(value *JobNotebookTask)
	PutNotificationSettings(value *JobNotificationSettings)
	PutParameter(value interface{})
	PutPipelineTask(value *JobPipelineTask)
	PutPythonWheelTask(value *JobPythonWheelTask)
	PutQueue(value *JobQueue)
	PutRunAs(value *JobRunAs)
	PutRunJobTask(value *JobRunJobTask)
	PutSchedule(value *JobSchedule)
	PutSparkJarTask(value *JobSparkJarTask)
	PutSparkPythonTask(value *JobSparkPythonTask)
	PutSparkSubmitTask(value *JobSparkSubmitTask)
	PutTask(value interface{})
	PutTimeouts(value *JobTimeouts)
	PutTrigger(value *JobTrigger)
	PutWebhookNotifications(value *JobWebhookNotifications)
	ResetAlwaysRunning()
	ResetBudgetPolicyId()
	ResetContinuous()
	ResetControlRunState()
	ResetDbtTask()
	ResetDeployment()
	ResetDescription()
	ResetEditMode()
	ResetEmailNotifications()
	ResetEnvironment()
	ResetExistingClusterId()
	ResetFormat()
	ResetGitSource()
	ResetHealth()
	ResetId()
	ResetJobCluster()
	ResetLibrary()
	ResetMaxConcurrentRuns()
	ResetMaxRetries()
	ResetMinRetryIntervalMillis()
	ResetName()
	ResetNewCluster()
	ResetNotebookTask()
	ResetNotificationSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameter()
	ResetPerformanceTarget()
	ResetPipelineTask()
	ResetPythonWheelTask()
	ResetQueue()
	ResetRetryOnTimeout()
	ResetRunAs()
	ResetRunJobTask()
	ResetSchedule()
	ResetSparkJarTask()
	ResetSparkPythonTask()
	ResetSparkSubmitTask()
	ResetTags()
	ResetTask()
	ResetTimeouts()
	ResetTimeoutSeconds()
	ResetTrigger()
	ResetWebhookNotifications()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Job
type jsiiProxy_Job struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Job) AlwaysRunning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysRunning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) AlwaysRunningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysRunningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) BudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) BudgetPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Continuous() JobContinuousOutputReference {
	var returns JobContinuousOutputReference
	_jsii_.Get(
		j,
		"continuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ContinuousInput() *JobContinuous {
	var returns *JobContinuous
	_jsii_.Get(
		j,
		"continuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ControlRunState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlRunState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ControlRunStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlRunStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DbtTask() JobDbtTaskOutputReference {
	var returns JobDbtTaskOutputReference
	_jsii_.Get(
		j,
		"dbtTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DbtTaskInput() *JobDbtTask {
	var returns *JobDbtTask
	_jsii_.Get(
		j,
		"dbtTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Deployment() JobDeploymentOutputReference {
	var returns JobDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DeploymentInput() *JobDeployment {
	var returns *JobDeployment
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) EditMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) EditModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) EmailNotifications() JobEmailNotificationsOutputReference {
	var returns JobEmailNotificationsOutputReference
	_jsii_.Get(
		j,
		"emailNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) EmailNotificationsInput() *JobEmailNotifications {
	var returns *JobEmailNotifications
	_jsii_.Get(
		j,
		"emailNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Environment() JobEnvironmentList {
	var returns JobEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ExistingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ExistingClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) GitSource() JobGitSourceOutputReference {
	var returns JobGitSourceOutputReference
	_jsii_.Get(
		j,
		"gitSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) GitSourceInput() *JobGitSource {
	var returns *JobGitSource
	_jsii_.Get(
		j,
		"gitSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Health() JobHealthOutputReference {
	var returns JobHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) HealthInput() *JobHealth {
	var returns *JobHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) JobCluster() JobJobClusterList {
	var returns JobJobClusterList
	_jsii_.Get(
		j,
		"jobCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) JobClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Library() JobLibraryList {
	var returns JobLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) MaxConcurrentRuns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) MaxConcurrentRunsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) MinRetryIntervalMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) MinRetryIntervalMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) NewCluster() JobNewClusterOutputReference {
	var returns JobNewClusterOutputReference
	_jsii_.Get(
		j,
		"newCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) NewClusterInput() *JobNewCluster {
	var returns *JobNewCluster
	_jsii_.Get(
		j,
		"newClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) NotebookTask() JobNotebookTaskOutputReference {
	var returns JobNotebookTaskOutputReference
	_jsii_.Get(
		j,
		"notebookTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) NotebookTaskInput() *JobNotebookTask {
	var returns *JobNotebookTask
	_jsii_.Get(
		j,
		"notebookTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) NotificationSettings() JobNotificationSettingsOutputReference {
	var returns JobNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) NotificationSettingsInput() *JobNotificationSettings {
	var returns *JobNotificationSettings
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Parameter() JobParameterList {
	var returns JobParameterList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PerformanceTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PerformanceTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PipelineTask() JobPipelineTaskOutputReference {
	var returns JobPipelineTaskOutputReference
	_jsii_.Get(
		j,
		"pipelineTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PipelineTaskInput() *JobPipelineTask {
	var returns *JobPipelineTask
	_jsii_.Get(
		j,
		"pipelineTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PythonWheelTask() JobPythonWheelTaskOutputReference {
	var returns JobPythonWheelTaskOutputReference
	_jsii_.Get(
		j,
		"pythonWheelTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PythonWheelTaskInput() *JobPythonWheelTask {
	var returns *JobPythonWheelTask
	_jsii_.Get(
		j,
		"pythonWheelTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Queue() JobQueueOutputReference {
	var returns JobQueueOutputReference
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) QueueInput() *JobQueue {
	var returns *JobQueue
	_jsii_.Get(
		j,
		"queueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RetryOnTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RetryOnTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RunAs() JobRunAsOutputReference {
	var returns JobRunAsOutputReference
	_jsii_.Get(
		j,
		"runAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RunAsInput() *JobRunAs {
	var returns *JobRunAs
	_jsii_.Get(
		j,
		"runAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RunJobTask() JobRunJobTaskOutputReference {
	var returns JobRunJobTaskOutputReference
	_jsii_.Get(
		j,
		"runJobTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RunJobTaskInput() *JobRunJobTask {
	var returns *JobRunJobTask
	_jsii_.Get(
		j,
		"runJobTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Schedule() JobScheduleOutputReference {
	var returns JobScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ScheduleInput() *JobSchedule {
	var returns *JobSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SparkJarTask() JobSparkJarTaskOutputReference {
	var returns JobSparkJarTaskOutputReference
	_jsii_.Get(
		j,
		"sparkJarTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SparkJarTaskInput() *JobSparkJarTask {
	var returns *JobSparkJarTask
	_jsii_.Get(
		j,
		"sparkJarTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SparkPythonTask() JobSparkPythonTaskOutputReference {
	var returns JobSparkPythonTaskOutputReference
	_jsii_.Get(
		j,
		"sparkPythonTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SparkPythonTaskInput() *JobSparkPythonTask {
	var returns *JobSparkPythonTask
	_jsii_.Get(
		j,
		"sparkPythonTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SparkSubmitTask() JobSparkSubmitTaskOutputReference {
	var returns JobSparkSubmitTaskOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SparkSubmitTaskInput() *JobSparkSubmitTask {
	var returns *JobSparkSubmitTask
	_jsii_.Get(
		j,
		"sparkSubmitTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Task() JobTaskList {
	var returns JobTaskList
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Timeouts() JobTimeoutsOutputReference {
	var returns JobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Trigger() JobTriggerOutputReference {
	var returns JobTriggerOutputReference
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TriggerInput() *JobTrigger {
	var returns *JobTrigger
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) WebhookNotifications() JobWebhookNotificationsOutputReference {
	var returns JobWebhookNotificationsOutputReference
	_jsii_.Get(
		j,
		"webhookNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) WebhookNotificationsInput() *JobWebhookNotifications {
	var returns *JobWebhookNotifications
	_jsii_.Get(
		j,
		"webhookNotificationsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job databricks_job} Resource.
func NewJob(scope constructs.Construct, id *string, config *JobConfig) Job {
	_init_.Initialize()

	if err := validateNewJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Job{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.Job",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/job databricks_job} Resource.
func NewJob_Override(j Job, scope constructs.Construct, id *string, config *JobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.Job",
		[]interface{}{scope, id, config},
		j,
	)
}

func (j *jsiiProxy_Job)SetAlwaysRunning(val interface{}) {
	if err := j.validateSetAlwaysRunningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysRunning",
		val,
	)
}

func (j *jsiiProxy_Job)SetBudgetPolicyId(val *string) {
	if err := j.validateSetBudgetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetPolicyId",
		val,
	)
}

func (j *jsiiProxy_Job)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Job)SetControlRunState(val interface{}) {
	if err := j.validateSetControlRunStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlRunState",
		val,
	)
}

func (j *jsiiProxy_Job)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Job)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Job)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Job)SetEditMode(val *string) {
	if err := j.validateSetEditModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"editMode",
		val,
	)
}

func (j *jsiiProxy_Job)SetExistingClusterId(val *string) {
	if err := j.validateSetExistingClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingClusterId",
		val,
	)
}

func (j *jsiiProxy_Job)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Job)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_Job)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Job)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Job)SetMaxConcurrentRuns(val *float64) {
	if err := j.validateSetMaxConcurrentRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentRuns",
		val,
	)
}

func (j *jsiiProxy_Job)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_Job)SetMinRetryIntervalMillis(val *float64) {
	if err := j.validateSetMinRetryIntervalMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRetryIntervalMillis",
		val,
	)
}

func (j *jsiiProxy_Job)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Job)SetPerformanceTarget(val *string) {
	if err := j.validateSetPerformanceTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceTarget",
		val,
	)
}

func (j *jsiiProxy_Job)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Job)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Job)SetRetryOnTimeout(val interface{}) {
	if err := j.validateSetRetryOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOnTimeout",
		val,
	)
}

func (j *jsiiProxy_Job)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Job)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

// Generates CDKTF code for importing a Job resource upon running "cdktf plan <stack-name>".
func Job_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.job.Job",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Job_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.job.Job",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Job_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.job.Job",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Job_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.job.Job",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Job_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.job.Job",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) AddMoveTarget(moveTarget *string) {
	if err := j.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (j *jsiiProxy_Job) AddOverride(path *string, value interface{}) {
	if err := j.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (j *jsiiProxy_Job) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := j.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (j *jsiiProxy_Job) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) MoveFromId(id *string) {
	if err := j.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveFromId",
		[]interface{}{id},
	)
}

func (j *jsiiProxy_Job) MoveTo(moveTarget *string, index interface{}) {
	if err := j.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (j *jsiiProxy_Job) MoveToId(id *string) {
	if err := j.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveToId",
		[]interface{}{id},
	)
}

func (j *jsiiProxy_Job) OverrideLogicalId(newLogicalId *string) {
	if err := j.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (j *jsiiProxy_Job) PutContinuous(value *JobContinuous) {
	if err := j.validatePutContinuousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putContinuous",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutDbtTask(value *JobDbtTask) {
	if err := j.validatePutDbtTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDbtTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutDeployment(value *JobDeployment) {
	if err := j.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDeployment",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutEmailNotifications(value *JobEmailNotifications) {
	if err := j.validatePutEmailNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEmailNotifications",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutEnvironment(value interface{}) {
	if err := j.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutGitSource(value *JobGitSource) {
	if err := j.validatePutGitSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGitSource",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutHealth(value *JobHealth) {
	if err := j.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHealth",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutJobCluster(value interface{}) {
	if err := j.validatePutJobClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putJobCluster",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutLibrary(value interface{}) {
	if err := j.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLibrary",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutNewCluster(value *JobNewCluster) {
	if err := j.validatePutNewClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNewCluster",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutNotebookTask(value *JobNotebookTask) {
	if err := j.validatePutNotebookTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNotebookTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutNotificationSettings(value *JobNotificationSettings) {
	if err := j.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutParameter(value interface{}) {
	if err := j.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putParameter",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutPipelineTask(value *JobPipelineTask) {
	if err := j.validatePutPipelineTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPipelineTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutPythonWheelTask(value *JobPythonWheelTask) {
	if err := j.validatePutPythonWheelTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPythonWheelTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutQueue(value *JobQueue) {
	if err := j.validatePutQueueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putQueue",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutRunAs(value *JobRunAs) {
	if err := j.validatePutRunAsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putRunAs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutRunJobTask(value *JobRunJobTask) {
	if err := j.validatePutRunJobTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putRunJobTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutSchedule(value *JobSchedule) {
	if err := j.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSchedule",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutSparkJarTask(value *JobSparkJarTask) {
	if err := j.validatePutSparkJarTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkJarTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutSparkPythonTask(value *JobSparkPythonTask) {
	if err := j.validatePutSparkPythonTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkPythonTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutSparkSubmitTask(value *JobSparkSubmitTask) {
	if err := j.validatePutSparkSubmitTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkSubmitTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutTask(value interface{}) {
	if err := j.validatePutTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutTimeouts(value *JobTimeouts) {
	if err := j.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutTrigger(value *JobTrigger) {
	if err := j.validatePutTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTrigger",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutWebhookNotifications(value *JobWebhookNotifications) {
	if err := j.validatePutWebhookNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWebhookNotifications",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) ResetAlwaysRunning() {
	_jsii_.InvokeVoid(
		j,
		"resetAlwaysRunning",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetBudgetPolicyId() {
	_jsii_.InvokeVoid(
		j,
		"resetBudgetPolicyId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetContinuous() {
	_jsii_.InvokeVoid(
		j,
		"resetContinuous",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetControlRunState() {
	_jsii_.InvokeVoid(
		j,
		"resetControlRunState",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetDbtTask() {
	_jsii_.InvokeVoid(
		j,
		"resetDbtTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetDeployment() {
	_jsii_.InvokeVoid(
		j,
		"resetDeployment",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetDescription() {
	_jsii_.InvokeVoid(
		j,
		"resetDescription",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetEditMode() {
	_jsii_.InvokeVoid(
		j,
		"resetEditMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetEmailNotifications() {
	_jsii_.InvokeVoid(
		j,
		"resetEmailNotifications",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetEnvironment() {
	_jsii_.InvokeVoid(
		j,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetExistingClusterId() {
	_jsii_.InvokeVoid(
		j,
		"resetExistingClusterId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetFormat() {
	_jsii_.InvokeVoid(
		j,
		"resetFormat",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetGitSource() {
	_jsii_.InvokeVoid(
		j,
		"resetGitSource",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetHealth() {
	_jsii_.InvokeVoid(
		j,
		"resetHealth",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetId() {
	_jsii_.InvokeVoid(
		j,
		"resetId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetJobCluster() {
	_jsii_.InvokeVoid(
		j,
		"resetJobCluster",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetLibrary() {
	_jsii_.InvokeVoid(
		j,
		"resetLibrary",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetMaxConcurrentRuns() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxConcurrentRuns",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetMinRetryIntervalMillis() {
	_jsii_.InvokeVoid(
		j,
		"resetMinRetryIntervalMillis",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetName() {
	_jsii_.InvokeVoid(
		j,
		"resetName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetNewCluster() {
	_jsii_.InvokeVoid(
		j,
		"resetNewCluster",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetNotebookTask() {
	_jsii_.InvokeVoid(
		j,
		"resetNotebookTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		j,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		j,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetParameter() {
	_jsii_.InvokeVoid(
		j,
		"resetParameter",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetPerformanceTarget() {
	_jsii_.InvokeVoid(
		j,
		"resetPerformanceTarget",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetPipelineTask() {
	_jsii_.InvokeVoid(
		j,
		"resetPipelineTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetPythonWheelTask() {
	_jsii_.InvokeVoid(
		j,
		"resetPythonWheelTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetQueue() {
	_jsii_.InvokeVoid(
		j,
		"resetQueue",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetRetryOnTimeout() {
	_jsii_.InvokeVoid(
		j,
		"resetRetryOnTimeout",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetRunAs() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetRunJobTask() {
	_jsii_.InvokeVoid(
		j,
		"resetRunJobTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetSchedule() {
	_jsii_.InvokeVoid(
		j,
		"resetSchedule",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetSparkJarTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkJarTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetSparkPythonTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkPythonTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetSparkSubmitTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkSubmitTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetTags() {
	_jsii_.InvokeVoid(
		j,
		"resetTags",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetTask() {
	_jsii_.InvokeVoid(
		j,
		"resetTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetTimeouts() {
	_jsii_.InvokeVoid(
		j,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetTrigger() {
	_jsii_.InvokeVoid(
		j,
		"resetTrigger",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetWebhookNotifications() {
	_jsii_.InvokeVoid(
		j,
		"resetWebhookNotifications",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

