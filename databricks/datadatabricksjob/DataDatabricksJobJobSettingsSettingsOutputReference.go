// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Compute() DataDatabricksJobJobSettingsSettingsComputeList
	ComputeInput() interface{}
	Continuous() DataDatabricksJobJobSettingsSettingsContinuousOutputReference
	ContinuousInput() *DataDatabricksJobJobSettingsSettingsContinuous
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DbtTask() DataDatabricksJobJobSettingsSettingsDbtTaskOutputReference
	DbtTaskInput() *DataDatabricksJobJobSettingsSettingsDbtTask
	Deployment() DataDatabricksJobJobSettingsSettingsDeploymentOutputReference
	DeploymentInput() *DataDatabricksJobJobSettingsSettingsDeployment
	EmailNotifications() DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference
	EmailNotificationsInput() *DataDatabricksJobJobSettingsSettingsEmailNotifications
	ExistingClusterId() *string
	SetExistingClusterId(val *string)
	ExistingClusterIdInput() *string
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	GitSource() DataDatabricksJobJobSettingsSettingsGitSourceOutputReference
	GitSourceInput() *DataDatabricksJobJobSettingsSettingsGitSource
	Health() DataDatabricksJobJobSettingsSettingsHealthOutputReference
	HealthInput() *DataDatabricksJobJobSettingsSettingsHealth
	InternalValue() *DataDatabricksJobJobSettingsSettings
	SetInternalValue(val *DataDatabricksJobJobSettingsSettings)
	JobCluster() DataDatabricksJobJobSettingsSettingsJobClusterList
	JobClusterInput() interface{}
	Library() DataDatabricksJobJobSettingsSettingsLibraryList
	LibraryInput() interface{}
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
	NewCluster() DataDatabricksJobJobSettingsSettingsNewClusterOutputReference
	NewClusterInput() *DataDatabricksJobJobSettingsSettingsNewCluster
	NotebookTask() DataDatabricksJobJobSettingsSettingsNotebookTaskOutputReference
	NotebookTaskInput() *DataDatabricksJobJobSettingsSettingsNotebookTask
	NotificationSettings() DataDatabricksJobJobSettingsSettingsNotificationSettingsOutputReference
	NotificationSettingsInput() *DataDatabricksJobJobSettingsSettingsNotificationSettings
	Parameter() DataDatabricksJobJobSettingsSettingsParameterList
	ParameterInput() interface{}
	PipelineTask() DataDatabricksJobJobSettingsSettingsPipelineTaskOutputReference
	PipelineTaskInput() *DataDatabricksJobJobSettingsSettingsPipelineTask
	PythonWheelTask() DataDatabricksJobJobSettingsSettingsPythonWheelTaskOutputReference
	PythonWheelTaskInput() *DataDatabricksJobJobSettingsSettingsPythonWheelTask
	Queue() DataDatabricksJobJobSettingsSettingsQueueOutputReference
	QueueInput() *DataDatabricksJobJobSettingsSettingsQueue
	RetryOnTimeout() interface{}
	SetRetryOnTimeout(val interface{})
	RetryOnTimeoutInput() interface{}
	RunAs() DataDatabricksJobJobSettingsSettingsRunAsOutputReference
	RunAsInput() *DataDatabricksJobJobSettingsSettingsRunAs
	RunJobTask() DataDatabricksJobJobSettingsSettingsRunJobTaskOutputReference
	RunJobTaskInput() *DataDatabricksJobJobSettingsSettingsRunJobTask
	Schedule() DataDatabricksJobJobSettingsSettingsScheduleOutputReference
	ScheduleInput() *DataDatabricksJobJobSettingsSettingsSchedule
	SparkJarTask() DataDatabricksJobJobSettingsSettingsSparkJarTaskOutputReference
	SparkJarTaskInput() *DataDatabricksJobJobSettingsSettingsSparkJarTask
	SparkPythonTask() DataDatabricksJobJobSettingsSettingsSparkPythonTaskOutputReference
	SparkPythonTaskInput() *DataDatabricksJobJobSettingsSettingsSparkPythonTask
	SparkSubmitTask() DataDatabricksJobJobSettingsSettingsSparkSubmitTaskOutputReference
	SparkSubmitTaskInput() *DataDatabricksJobJobSettingsSettingsSparkSubmitTask
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Task() DataDatabricksJobJobSettingsSettingsTaskList
	TaskInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
	Trigger() DataDatabricksJobJobSettingsSettingsTriggerOutputReference
	TriggerInput() *DataDatabricksJobJobSettingsSettingsTrigger
	WebhookNotifications() DataDatabricksJobJobSettingsSettingsWebhookNotificationsOutputReference
	WebhookNotificationsInput() *DataDatabricksJobJobSettingsSettingsWebhookNotifications
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCompute(value interface{})
	PutContinuous(value *DataDatabricksJobJobSettingsSettingsContinuous)
	PutDbtTask(value *DataDatabricksJobJobSettingsSettingsDbtTask)
	PutDeployment(value *DataDatabricksJobJobSettingsSettingsDeployment)
	PutEmailNotifications(value *DataDatabricksJobJobSettingsSettingsEmailNotifications)
	PutGitSource(value *DataDatabricksJobJobSettingsSettingsGitSource)
	PutHealth(value *DataDatabricksJobJobSettingsSettingsHealth)
	PutJobCluster(value interface{})
	PutLibrary(value interface{})
	PutNewCluster(value *DataDatabricksJobJobSettingsSettingsNewCluster)
	PutNotebookTask(value *DataDatabricksJobJobSettingsSettingsNotebookTask)
	PutNotificationSettings(value *DataDatabricksJobJobSettingsSettingsNotificationSettings)
	PutParameter(value interface{})
	PutPipelineTask(value *DataDatabricksJobJobSettingsSettingsPipelineTask)
	PutPythonWheelTask(value *DataDatabricksJobJobSettingsSettingsPythonWheelTask)
	PutQueue(value *DataDatabricksJobJobSettingsSettingsQueue)
	PutRunAs(value *DataDatabricksJobJobSettingsSettingsRunAs)
	PutRunJobTask(value *DataDatabricksJobJobSettingsSettingsRunJobTask)
	PutSchedule(value *DataDatabricksJobJobSettingsSettingsSchedule)
	PutSparkJarTask(value *DataDatabricksJobJobSettingsSettingsSparkJarTask)
	PutSparkPythonTask(value *DataDatabricksJobJobSettingsSettingsSparkPythonTask)
	PutSparkSubmitTask(value *DataDatabricksJobJobSettingsSettingsSparkSubmitTask)
	PutTask(value interface{})
	PutTrigger(value *DataDatabricksJobJobSettingsSettingsTrigger)
	PutWebhookNotifications(value *DataDatabricksJobJobSettingsSettingsWebhookNotifications)
	ResetCompute()
	ResetContinuous()
	ResetDbtTask()
	ResetDeployment()
	ResetEmailNotifications()
	ResetExistingClusterId()
	ResetFormat()
	ResetGitSource()
	ResetHealth()
	ResetJobCluster()
	ResetLibrary()
	ResetMaxConcurrentRuns()
	ResetMaxRetries()
	ResetMinRetryIntervalMillis()
	ResetName()
	ResetNewCluster()
	ResetNotebookTask()
	ResetNotificationSettings()
	ResetParameter()
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
	ResetTimeoutSeconds()
	ResetTrigger()
	ResetWebhookNotifications()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Compute() DataDatabricksJobJobSettingsSettingsComputeList {
	var returns DataDatabricksJobJobSettingsSettingsComputeList
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Continuous() DataDatabricksJobJobSettingsSettingsContinuousOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsContinuousOutputReference
	_jsii_.Get(
		j,
		"continuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ContinuousInput() *DataDatabricksJobJobSettingsSettingsContinuous {
	var returns *DataDatabricksJobJobSettingsSettingsContinuous
	_jsii_.Get(
		j,
		"continuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) DbtTask() DataDatabricksJobJobSettingsSettingsDbtTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsDbtTaskOutputReference
	_jsii_.Get(
		j,
		"dbtTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) DbtTaskInput() *DataDatabricksJobJobSettingsSettingsDbtTask {
	var returns *DataDatabricksJobJobSettingsSettingsDbtTask
	_jsii_.Get(
		j,
		"dbtTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Deployment() DataDatabricksJobJobSettingsSettingsDeploymentOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) DeploymentInput() *DataDatabricksJobJobSettingsSettingsDeployment {
	var returns *DataDatabricksJobJobSettingsSettingsDeployment
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) EmailNotifications() DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference
	_jsii_.Get(
		j,
		"emailNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) EmailNotificationsInput() *DataDatabricksJobJobSettingsSettingsEmailNotifications {
	var returns *DataDatabricksJobJobSettingsSettingsEmailNotifications
	_jsii_.Get(
		j,
		"emailNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ExistingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ExistingClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GitSource() DataDatabricksJobJobSettingsSettingsGitSourceOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsGitSourceOutputReference
	_jsii_.Get(
		j,
		"gitSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GitSourceInput() *DataDatabricksJobJobSettingsSettingsGitSource {
	var returns *DataDatabricksJobJobSettingsSettingsGitSource
	_jsii_.Get(
		j,
		"gitSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Health() DataDatabricksJobJobSettingsSettingsHealthOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) HealthInput() *DataDatabricksJobJobSettingsSettingsHealth {
	var returns *DataDatabricksJobJobSettingsSettingsHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) InternalValue() *DataDatabricksJobJobSettingsSettings {
	var returns *DataDatabricksJobJobSettingsSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) JobCluster() DataDatabricksJobJobSettingsSettingsJobClusterList {
	var returns DataDatabricksJobJobSettingsSettingsJobClusterList
	_jsii_.Get(
		j,
		"jobCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) JobClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Library() DataDatabricksJobJobSettingsSettingsLibraryList {
	var returns DataDatabricksJobJobSettingsSettingsLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) MaxConcurrentRuns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) MaxConcurrentRunsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) MinRetryIntervalMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) MinRetryIntervalMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) NewCluster() DataDatabricksJobJobSettingsSettingsNewClusterOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsNewClusterOutputReference
	_jsii_.Get(
		j,
		"newCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) NewClusterInput() *DataDatabricksJobJobSettingsSettingsNewCluster {
	var returns *DataDatabricksJobJobSettingsSettingsNewCluster
	_jsii_.Get(
		j,
		"newClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) NotebookTask() DataDatabricksJobJobSettingsSettingsNotebookTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsNotebookTaskOutputReference
	_jsii_.Get(
		j,
		"notebookTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) NotebookTaskInput() *DataDatabricksJobJobSettingsSettingsNotebookTask {
	var returns *DataDatabricksJobJobSettingsSettingsNotebookTask
	_jsii_.Get(
		j,
		"notebookTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) NotificationSettings() DataDatabricksJobJobSettingsSettingsNotificationSettingsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) NotificationSettingsInput() *DataDatabricksJobJobSettingsSettingsNotificationSettings {
	var returns *DataDatabricksJobJobSettingsSettingsNotificationSettings
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Parameter() DataDatabricksJobJobSettingsSettingsParameterList {
	var returns DataDatabricksJobJobSettingsSettingsParameterList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PipelineTask() DataDatabricksJobJobSettingsSettingsPipelineTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsPipelineTaskOutputReference
	_jsii_.Get(
		j,
		"pipelineTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PipelineTaskInput() *DataDatabricksJobJobSettingsSettingsPipelineTask {
	var returns *DataDatabricksJobJobSettingsSettingsPipelineTask
	_jsii_.Get(
		j,
		"pipelineTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PythonWheelTask() DataDatabricksJobJobSettingsSettingsPythonWheelTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsPythonWheelTaskOutputReference
	_jsii_.Get(
		j,
		"pythonWheelTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PythonWheelTaskInput() *DataDatabricksJobJobSettingsSettingsPythonWheelTask {
	var returns *DataDatabricksJobJobSettingsSettingsPythonWheelTask
	_jsii_.Get(
		j,
		"pythonWheelTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Queue() DataDatabricksJobJobSettingsSettingsQueueOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsQueueOutputReference
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) QueueInput() *DataDatabricksJobJobSettingsSettingsQueue {
	var returns *DataDatabricksJobJobSettingsSettingsQueue
	_jsii_.Get(
		j,
		"queueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) RetryOnTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) RetryOnTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) RunAs() DataDatabricksJobJobSettingsSettingsRunAsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsRunAsOutputReference
	_jsii_.Get(
		j,
		"runAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) RunAsInput() *DataDatabricksJobJobSettingsSettingsRunAs {
	var returns *DataDatabricksJobJobSettingsSettingsRunAs
	_jsii_.Get(
		j,
		"runAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) RunJobTask() DataDatabricksJobJobSettingsSettingsRunJobTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsRunJobTaskOutputReference
	_jsii_.Get(
		j,
		"runJobTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) RunJobTaskInput() *DataDatabricksJobJobSettingsSettingsRunJobTask {
	var returns *DataDatabricksJobJobSettingsSettingsRunJobTask
	_jsii_.Get(
		j,
		"runJobTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Schedule() DataDatabricksJobJobSettingsSettingsScheduleOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ScheduleInput() *DataDatabricksJobJobSettingsSettingsSchedule {
	var returns *DataDatabricksJobJobSettingsSettingsSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) SparkJarTask() DataDatabricksJobJobSettingsSettingsSparkJarTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsSparkJarTaskOutputReference
	_jsii_.Get(
		j,
		"sparkJarTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) SparkJarTaskInput() *DataDatabricksJobJobSettingsSettingsSparkJarTask {
	var returns *DataDatabricksJobJobSettingsSettingsSparkJarTask
	_jsii_.Get(
		j,
		"sparkJarTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) SparkPythonTask() DataDatabricksJobJobSettingsSettingsSparkPythonTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsSparkPythonTaskOutputReference
	_jsii_.Get(
		j,
		"sparkPythonTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) SparkPythonTaskInput() *DataDatabricksJobJobSettingsSettingsSparkPythonTask {
	var returns *DataDatabricksJobJobSettingsSettingsSparkPythonTask
	_jsii_.Get(
		j,
		"sparkPythonTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) SparkSubmitTask() DataDatabricksJobJobSettingsSettingsSparkSubmitTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsSparkSubmitTaskOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) SparkSubmitTaskInput() *DataDatabricksJobJobSettingsSettingsSparkSubmitTask {
	var returns *DataDatabricksJobJobSettingsSettingsSparkSubmitTask
	_jsii_.Get(
		j,
		"sparkSubmitTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Task() DataDatabricksJobJobSettingsSettingsTaskList {
	var returns DataDatabricksJobJobSettingsSettingsTaskList
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) TaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Trigger() DataDatabricksJobJobSettingsSettingsTriggerOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTriggerOutputReference
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) TriggerInput() *DataDatabricksJobJobSettingsSettingsTrigger {
	var returns *DataDatabricksJobJobSettingsSettingsTrigger
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) WebhookNotifications() DataDatabricksJobJobSettingsSettingsWebhookNotificationsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsWebhookNotificationsOutputReference
	_jsii_.Get(
		j,
		"webhookNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) WebhookNotificationsInput() *DataDatabricksJobJobSettingsSettingsWebhookNotifications {
	var returns *DataDatabricksJobJobSettingsSettingsWebhookNotifications
	_jsii_.Get(
		j,
		"webhookNotificationsInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksJobJobSettingsSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsOutputReference_Override(d DataDatabricksJobJobSettingsSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetExistingClusterId(val *string) {
	if err := j.validateSetExistingClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingClusterId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetInternalValue(val *DataDatabricksJobJobSettingsSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetMaxConcurrentRuns(val *float64) {
	if err := j.validateSetMaxConcurrentRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentRuns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetMinRetryIntervalMillis(val *float64) {
	if err := j.validateSetMinRetryIntervalMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRetryIntervalMillis",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetRetryOnTimeout(val interface{}) {
	if err := j.validateSetRetryOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOnTimeout",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutCompute(value interface{}) {
	if err := d.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutContinuous(value *DataDatabricksJobJobSettingsSettingsContinuous) {
	if err := d.validatePutContinuousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContinuous",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutDbtTask(value *DataDatabricksJobJobSettingsSettingsDbtTask) {
	if err := d.validatePutDbtTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDbtTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutDeployment(value *DataDatabricksJobJobSettingsSettingsDeployment) {
	if err := d.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeployment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutEmailNotifications(value *DataDatabricksJobJobSettingsSettingsEmailNotifications) {
	if err := d.validatePutEmailNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmailNotifications",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutGitSource(value *DataDatabricksJobJobSettingsSettingsGitSource) {
	if err := d.validatePutGitSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutHealth(value *DataDatabricksJobJobSettingsSettingsHealth) {
	if err := d.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHealth",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutJobCluster(value interface{}) {
	if err := d.validatePutJobClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJobCluster",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutLibrary(value interface{}) {
	if err := d.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLibrary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutNewCluster(value *DataDatabricksJobJobSettingsSettingsNewCluster) {
	if err := d.validatePutNewClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNewCluster",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutNotebookTask(value *DataDatabricksJobJobSettingsSettingsNotebookTask) {
	if err := d.validatePutNotebookTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotebookTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutNotificationSettings(value *DataDatabricksJobJobSettingsSettingsNotificationSettings) {
	if err := d.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutParameter(value interface{}) {
	if err := d.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParameter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutPipelineTask(value *DataDatabricksJobJobSettingsSettingsPipelineTask) {
	if err := d.validatePutPipelineTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPipelineTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutPythonWheelTask(value *DataDatabricksJobJobSettingsSettingsPythonWheelTask) {
	if err := d.validatePutPythonWheelTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPythonWheelTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutQueue(value *DataDatabricksJobJobSettingsSettingsQueue) {
	if err := d.validatePutQueueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueue",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutRunAs(value *DataDatabricksJobJobSettingsSettingsRunAs) {
	if err := d.validatePutRunAsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunAs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutRunJobTask(value *DataDatabricksJobJobSettingsSettingsRunJobTask) {
	if err := d.validatePutRunJobTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunJobTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutSchedule(value *DataDatabricksJobJobSettingsSettingsSchedule) {
	if err := d.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutSparkJarTask(value *DataDatabricksJobJobSettingsSettingsSparkJarTask) {
	if err := d.validatePutSparkJarTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkJarTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutSparkPythonTask(value *DataDatabricksJobJobSettingsSettingsSparkPythonTask) {
	if err := d.validatePutSparkPythonTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkPythonTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutSparkSubmitTask(value *DataDatabricksJobJobSettingsSettingsSparkSubmitTask) {
	if err := d.validatePutSparkSubmitTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkSubmitTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutTask(value interface{}) {
	if err := d.validatePutTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutTrigger(value *DataDatabricksJobJobSettingsSettingsTrigger) {
	if err := d.validatePutTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTrigger",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) PutWebhookNotifications(value *DataDatabricksJobJobSettingsSettingsWebhookNotifications) {
	if err := d.validatePutWebhookNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWebhookNotifications",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetCompute() {
	_jsii_.InvokeVoid(
		d,
		"resetCompute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetContinuous() {
	_jsii_.InvokeVoid(
		d,
		"resetContinuous",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetDbtTask() {
	_jsii_.InvokeVoid(
		d,
		"resetDbtTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetDeployment() {
	_jsii_.InvokeVoid(
		d,
		"resetDeployment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetEmailNotifications() {
	_jsii_.InvokeVoid(
		d,
		"resetEmailNotifications",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetExistingClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetExistingClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetGitSource() {
	_jsii_.InvokeVoid(
		d,
		"resetGitSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetHealth() {
	_jsii_.InvokeVoid(
		d,
		"resetHealth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetJobCluster() {
	_jsii_.InvokeVoid(
		d,
		"resetJobCluster",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetLibrary() {
	_jsii_.InvokeVoid(
		d,
		"resetLibrary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetMaxConcurrentRuns() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConcurrentRuns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetMinRetryIntervalMillis() {
	_jsii_.InvokeVoid(
		d,
		"resetMinRetryIntervalMillis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetNewCluster() {
	_jsii_.InvokeVoid(
		d,
		"resetNewCluster",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetNotebookTask() {
	_jsii_.InvokeVoid(
		d,
		"resetNotebookTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		d,
		"resetParameter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetPipelineTask() {
	_jsii_.InvokeVoid(
		d,
		"resetPipelineTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetPythonWheelTask() {
	_jsii_.InvokeVoid(
		d,
		"resetPythonWheelTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetQueue() {
	_jsii_.InvokeVoid(
		d,
		"resetQueue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetRetryOnTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetRetryOnTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetRunAs() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetRunJobTask() {
	_jsii_.InvokeVoid(
		d,
		"resetRunJobTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetSparkJarTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkJarTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetSparkPythonTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkPythonTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetSparkSubmitTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkSubmitTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetTask() {
	_jsii_.InvokeVoid(
		d,
		"resetTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetTrigger() {
	_jsii_.InvokeVoid(
		d,
		"resetTrigger",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ResetWebhookNotifications() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhookNotifications",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

