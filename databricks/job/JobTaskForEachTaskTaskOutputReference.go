// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskOutputReference interface {
	cdktf.ComplexObject
	CleanRoomsNotebookTask() JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference
	CleanRoomsNotebookTaskInput() *JobTaskForEachTaskTaskCleanRoomsNotebookTask
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
	ConditionTask() JobTaskForEachTaskTaskConditionTaskOutputReference
	ConditionTaskInput() *JobTaskForEachTaskTaskConditionTask
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DashboardTask() JobTaskForEachTaskTaskDashboardTaskOutputReference
	DashboardTaskInput() *JobTaskForEachTaskTaskDashboardTask
	DbtCloudTask() JobTaskForEachTaskTaskDbtCloudTaskOutputReference
	DbtCloudTaskInput() *JobTaskForEachTaskTaskDbtCloudTask
	DbtTask() JobTaskForEachTaskTaskDbtTaskOutputReference
	DbtTaskInput() *JobTaskForEachTaskTaskDbtTask
	DependsOn() JobTaskForEachTaskTaskDependsOnList
	DependsOnInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableAutoOptimization() interface{}
	SetDisableAutoOptimization(val interface{})
	DisableAutoOptimizationInput() interface{}
	EmailNotifications() JobTaskForEachTaskTaskEmailNotificationsOutputReference
	EmailNotificationsInput() *JobTaskForEachTaskTaskEmailNotifications
	EnvironmentKey() *string
	SetEnvironmentKey(val *string)
	EnvironmentKeyInput() *string
	ExistingClusterId() *string
	SetExistingClusterId(val *string)
	ExistingClusterIdInput() *string
	// Experimental.
	Fqn() *string
	GenAiComputeTask() JobTaskForEachTaskTaskGenAiComputeTaskOutputReference
	GenAiComputeTaskInput() *JobTaskForEachTaskTaskGenAiComputeTask
	Health() JobTaskForEachTaskTaskHealthOutputReference
	HealthInput() *JobTaskForEachTaskTaskHealth
	InternalValue() *JobTaskForEachTaskTask
	SetInternalValue(val *JobTaskForEachTaskTask)
	JobClusterKey() *string
	SetJobClusterKey(val *string)
	JobClusterKeyInput() *string
	Library() JobTaskForEachTaskTaskLibraryList
	LibraryInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MinRetryIntervalMillis() *float64
	SetMinRetryIntervalMillis(val *float64)
	MinRetryIntervalMillisInput() *float64
	NewCluster() JobTaskForEachTaskTaskNewClusterOutputReference
	NewClusterInput() *JobTaskForEachTaskTaskNewCluster
	NotebookTask() JobTaskForEachTaskTaskNotebookTaskOutputReference
	NotebookTaskInput() *JobTaskForEachTaskTaskNotebookTask
	NotificationSettings() JobTaskForEachTaskTaskNotificationSettingsOutputReference
	NotificationSettingsInput() *JobTaskForEachTaskTaskNotificationSettings
	PipelineTask() JobTaskForEachTaskTaskPipelineTaskOutputReference
	PipelineTaskInput() *JobTaskForEachTaskTaskPipelineTask
	PowerBiTask() JobTaskForEachTaskTaskPowerBiTaskOutputReference
	PowerBiTaskInput() *JobTaskForEachTaskTaskPowerBiTask
	PythonWheelTask() JobTaskForEachTaskTaskPythonWheelTaskOutputReference
	PythonWheelTaskInput() *JobTaskForEachTaskTaskPythonWheelTask
	RetryOnTimeout() interface{}
	SetRetryOnTimeout(val interface{})
	RetryOnTimeoutInput() interface{}
	RunIf() *string
	SetRunIf(val *string)
	RunIfInput() *string
	RunJobTask() JobTaskForEachTaskTaskRunJobTaskOutputReference
	RunJobTaskInput() *JobTaskForEachTaskTaskRunJobTask
	SparkJarTask() JobTaskForEachTaskTaskSparkJarTaskOutputReference
	SparkJarTaskInput() *JobTaskForEachTaskTaskSparkJarTask
	SparkPythonTask() JobTaskForEachTaskTaskSparkPythonTaskOutputReference
	SparkPythonTaskInput() *JobTaskForEachTaskTaskSparkPythonTask
	SparkSubmitTask() JobTaskForEachTaskTaskSparkSubmitTaskOutputReference
	SparkSubmitTaskInput() *JobTaskForEachTaskTaskSparkSubmitTask
	SqlTask() JobTaskForEachTaskTaskSqlTaskOutputReference
	SqlTaskInput() *JobTaskForEachTaskTaskSqlTask
	TaskKey() *string
	SetTaskKey(val *string)
	TaskKeyInput() *string
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
	WebhookNotifications() JobTaskForEachTaskTaskWebhookNotificationsOutputReference
	WebhookNotificationsInput() *JobTaskForEachTaskTaskWebhookNotifications
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
	PutCleanRoomsNotebookTask(value *JobTaskForEachTaskTaskCleanRoomsNotebookTask)
	PutConditionTask(value *JobTaskForEachTaskTaskConditionTask)
	PutDashboardTask(value *JobTaskForEachTaskTaskDashboardTask)
	PutDbtCloudTask(value *JobTaskForEachTaskTaskDbtCloudTask)
	PutDbtTask(value *JobTaskForEachTaskTaskDbtTask)
	PutDependsOn(value interface{})
	PutEmailNotifications(value *JobTaskForEachTaskTaskEmailNotifications)
	PutGenAiComputeTask(value *JobTaskForEachTaskTaskGenAiComputeTask)
	PutHealth(value *JobTaskForEachTaskTaskHealth)
	PutLibrary(value interface{})
	PutNewCluster(value *JobTaskForEachTaskTaskNewCluster)
	PutNotebookTask(value *JobTaskForEachTaskTaskNotebookTask)
	PutNotificationSettings(value *JobTaskForEachTaskTaskNotificationSettings)
	PutPipelineTask(value *JobTaskForEachTaskTaskPipelineTask)
	PutPowerBiTask(value *JobTaskForEachTaskTaskPowerBiTask)
	PutPythonWheelTask(value *JobTaskForEachTaskTaskPythonWheelTask)
	PutRunJobTask(value *JobTaskForEachTaskTaskRunJobTask)
	PutSparkJarTask(value *JobTaskForEachTaskTaskSparkJarTask)
	PutSparkPythonTask(value *JobTaskForEachTaskTaskSparkPythonTask)
	PutSparkSubmitTask(value *JobTaskForEachTaskTaskSparkSubmitTask)
	PutSqlTask(value *JobTaskForEachTaskTaskSqlTask)
	PutWebhookNotifications(value *JobTaskForEachTaskTaskWebhookNotifications)
	ResetCleanRoomsNotebookTask()
	ResetConditionTask()
	ResetDashboardTask()
	ResetDbtCloudTask()
	ResetDbtTask()
	ResetDependsOn()
	ResetDescription()
	ResetDisableAutoOptimization()
	ResetEmailNotifications()
	ResetEnvironmentKey()
	ResetExistingClusterId()
	ResetGenAiComputeTask()
	ResetHealth()
	ResetJobClusterKey()
	ResetLibrary()
	ResetMaxRetries()
	ResetMinRetryIntervalMillis()
	ResetNewCluster()
	ResetNotebookTask()
	ResetNotificationSettings()
	ResetPipelineTask()
	ResetPowerBiTask()
	ResetPythonWheelTask()
	ResetRetryOnTimeout()
	ResetRunIf()
	ResetRunJobTask()
	ResetSparkJarTask()
	ResetSparkPythonTask()
	ResetSparkSubmitTask()
	ResetSqlTask()
	ResetTimeoutSeconds()
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

// The jsii proxy struct for JobTaskForEachTaskTaskOutputReference
type jsiiProxy_JobTaskForEachTaskTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) CleanRoomsNotebookTask() JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference {
	var returns JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference
	_jsii_.Get(
		j,
		"cleanRoomsNotebookTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) CleanRoomsNotebookTaskInput() *JobTaskForEachTaskTaskCleanRoomsNotebookTask {
	var returns *JobTaskForEachTaskTaskCleanRoomsNotebookTask
	_jsii_.Get(
		j,
		"cleanRoomsNotebookTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ConditionTask() JobTaskForEachTaskTaskConditionTaskOutputReference {
	var returns JobTaskForEachTaskTaskConditionTaskOutputReference
	_jsii_.Get(
		j,
		"conditionTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ConditionTaskInput() *JobTaskForEachTaskTaskConditionTask {
	var returns *JobTaskForEachTaskTaskConditionTask
	_jsii_.Get(
		j,
		"conditionTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DashboardTask() JobTaskForEachTaskTaskDashboardTaskOutputReference {
	var returns JobTaskForEachTaskTaskDashboardTaskOutputReference
	_jsii_.Get(
		j,
		"dashboardTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DashboardTaskInput() *JobTaskForEachTaskTaskDashboardTask {
	var returns *JobTaskForEachTaskTaskDashboardTask
	_jsii_.Get(
		j,
		"dashboardTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DbtCloudTask() JobTaskForEachTaskTaskDbtCloudTaskOutputReference {
	var returns JobTaskForEachTaskTaskDbtCloudTaskOutputReference
	_jsii_.Get(
		j,
		"dbtCloudTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DbtCloudTaskInput() *JobTaskForEachTaskTaskDbtCloudTask {
	var returns *JobTaskForEachTaskTaskDbtCloudTask
	_jsii_.Get(
		j,
		"dbtCloudTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DbtTask() JobTaskForEachTaskTaskDbtTaskOutputReference {
	var returns JobTaskForEachTaskTaskDbtTaskOutputReference
	_jsii_.Get(
		j,
		"dbtTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DbtTaskInput() *JobTaskForEachTaskTaskDbtTask {
	var returns *JobTaskForEachTaskTaskDbtTask
	_jsii_.Get(
		j,
		"dbtTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DependsOn() JobTaskForEachTaskTaskDependsOnList {
	var returns JobTaskForEachTaskTaskDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DisableAutoOptimization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) DisableAutoOptimizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) EmailNotifications() JobTaskForEachTaskTaskEmailNotificationsOutputReference {
	var returns JobTaskForEachTaskTaskEmailNotificationsOutputReference
	_jsii_.Get(
		j,
		"emailNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) EmailNotificationsInput() *JobTaskForEachTaskTaskEmailNotifications {
	var returns *JobTaskForEachTaskTaskEmailNotifications
	_jsii_.Get(
		j,
		"emailNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) EnvironmentKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) EnvironmentKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ExistingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ExistingClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GenAiComputeTask() JobTaskForEachTaskTaskGenAiComputeTaskOutputReference {
	var returns JobTaskForEachTaskTaskGenAiComputeTaskOutputReference
	_jsii_.Get(
		j,
		"genAiComputeTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GenAiComputeTaskInput() *JobTaskForEachTaskTaskGenAiComputeTask {
	var returns *JobTaskForEachTaskTaskGenAiComputeTask
	_jsii_.Get(
		j,
		"genAiComputeTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) Health() JobTaskForEachTaskTaskHealthOutputReference {
	var returns JobTaskForEachTaskTaskHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) HealthInput() *JobTaskForEachTaskTaskHealth {
	var returns *JobTaskForEachTaskTaskHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) InternalValue() *JobTaskForEachTaskTask {
	var returns *JobTaskForEachTaskTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) JobClusterKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobClusterKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) JobClusterKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobClusterKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) Library() JobTaskForEachTaskTaskLibraryList {
	var returns JobTaskForEachTaskTaskLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) MinRetryIntervalMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) MinRetryIntervalMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) NewCluster() JobTaskForEachTaskTaskNewClusterOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterOutputReference
	_jsii_.Get(
		j,
		"newCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) NewClusterInput() *JobTaskForEachTaskTaskNewCluster {
	var returns *JobTaskForEachTaskTaskNewCluster
	_jsii_.Get(
		j,
		"newClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) NotebookTask() JobTaskForEachTaskTaskNotebookTaskOutputReference {
	var returns JobTaskForEachTaskTaskNotebookTaskOutputReference
	_jsii_.Get(
		j,
		"notebookTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) NotebookTaskInput() *JobTaskForEachTaskTaskNotebookTask {
	var returns *JobTaskForEachTaskTaskNotebookTask
	_jsii_.Get(
		j,
		"notebookTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) NotificationSettings() JobTaskForEachTaskTaskNotificationSettingsOutputReference {
	var returns JobTaskForEachTaskTaskNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) NotificationSettingsInput() *JobTaskForEachTaskTaskNotificationSettings {
	var returns *JobTaskForEachTaskTaskNotificationSettings
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PipelineTask() JobTaskForEachTaskTaskPipelineTaskOutputReference {
	var returns JobTaskForEachTaskTaskPipelineTaskOutputReference
	_jsii_.Get(
		j,
		"pipelineTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PipelineTaskInput() *JobTaskForEachTaskTaskPipelineTask {
	var returns *JobTaskForEachTaskTaskPipelineTask
	_jsii_.Get(
		j,
		"pipelineTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PowerBiTask() JobTaskForEachTaskTaskPowerBiTaskOutputReference {
	var returns JobTaskForEachTaskTaskPowerBiTaskOutputReference
	_jsii_.Get(
		j,
		"powerBiTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PowerBiTaskInput() *JobTaskForEachTaskTaskPowerBiTask {
	var returns *JobTaskForEachTaskTaskPowerBiTask
	_jsii_.Get(
		j,
		"powerBiTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PythonWheelTask() JobTaskForEachTaskTaskPythonWheelTaskOutputReference {
	var returns JobTaskForEachTaskTaskPythonWheelTaskOutputReference
	_jsii_.Get(
		j,
		"pythonWheelTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PythonWheelTaskInput() *JobTaskForEachTaskTaskPythonWheelTask {
	var returns *JobTaskForEachTaskTaskPythonWheelTask
	_jsii_.Get(
		j,
		"pythonWheelTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) RetryOnTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) RetryOnTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) RunIf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) RunIfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) RunJobTask() JobTaskForEachTaskTaskRunJobTaskOutputReference {
	var returns JobTaskForEachTaskTaskRunJobTaskOutputReference
	_jsii_.Get(
		j,
		"runJobTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) RunJobTaskInput() *JobTaskForEachTaskTaskRunJobTask {
	var returns *JobTaskForEachTaskTaskRunJobTask
	_jsii_.Get(
		j,
		"runJobTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) SparkJarTask() JobTaskForEachTaskTaskSparkJarTaskOutputReference {
	var returns JobTaskForEachTaskTaskSparkJarTaskOutputReference
	_jsii_.Get(
		j,
		"sparkJarTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) SparkJarTaskInput() *JobTaskForEachTaskTaskSparkJarTask {
	var returns *JobTaskForEachTaskTaskSparkJarTask
	_jsii_.Get(
		j,
		"sparkJarTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) SparkPythonTask() JobTaskForEachTaskTaskSparkPythonTaskOutputReference {
	var returns JobTaskForEachTaskTaskSparkPythonTaskOutputReference
	_jsii_.Get(
		j,
		"sparkPythonTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) SparkPythonTaskInput() *JobTaskForEachTaskTaskSparkPythonTask {
	var returns *JobTaskForEachTaskTaskSparkPythonTask
	_jsii_.Get(
		j,
		"sparkPythonTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) SparkSubmitTask() JobTaskForEachTaskTaskSparkSubmitTaskOutputReference {
	var returns JobTaskForEachTaskTaskSparkSubmitTaskOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) SparkSubmitTaskInput() *JobTaskForEachTaskTaskSparkSubmitTask {
	var returns *JobTaskForEachTaskTaskSparkSubmitTask
	_jsii_.Get(
		j,
		"sparkSubmitTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) SqlTask() JobTaskForEachTaskTaskSqlTaskOutputReference {
	var returns JobTaskForEachTaskTaskSqlTaskOutputReference
	_jsii_.Get(
		j,
		"sqlTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) SqlTaskInput() *JobTaskForEachTaskTaskSqlTask {
	var returns *JobTaskForEachTaskTaskSqlTask
	_jsii_.Get(
		j,
		"sqlTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) TaskKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) TaskKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) WebhookNotifications() JobTaskForEachTaskTaskWebhookNotificationsOutputReference {
	var returns JobTaskForEachTaskTaskWebhookNotificationsOutputReference
	_jsii_.Get(
		j,
		"webhookNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) WebhookNotificationsInput() *JobTaskForEachTaskTaskWebhookNotifications {
	var returns *JobTaskForEachTaskTaskWebhookNotifications
	_jsii_.Get(
		j,
		"webhookNotificationsInput",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskOutputReference_Override(j JobTaskForEachTaskTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetDisableAutoOptimization(val interface{}) {
	if err := j.validateSetDisableAutoOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoOptimization",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetEnvironmentKey(val *string) {
	if err := j.validateSetEnvironmentKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentKey",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetExistingClusterId(val *string) {
	if err := j.validateSetExistingClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingClusterId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetInternalValue(val *JobTaskForEachTaskTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetJobClusterKey(val *string) {
	if err := j.validateSetJobClusterKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobClusterKey",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetMinRetryIntervalMillis(val *float64) {
	if err := j.validateSetMinRetryIntervalMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRetryIntervalMillis",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetRetryOnTimeout(val interface{}) {
	if err := j.validateSetRetryOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOnTimeout",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetRunIf(val *string) {
	if err := j.validateSetRunIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runIf",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetTaskKey(val *string) {
	if err := j.validateSetTaskKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskKey",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutCleanRoomsNotebookTask(value *JobTaskForEachTaskTaskCleanRoomsNotebookTask) {
	if err := j.validatePutCleanRoomsNotebookTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCleanRoomsNotebookTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutConditionTask(value *JobTaskForEachTaskTaskConditionTask) {
	if err := j.validatePutConditionTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putConditionTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutDashboardTask(value *JobTaskForEachTaskTaskDashboardTask) {
	if err := j.validatePutDashboardTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDashboardTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutDbtCloudTask(value *JobTaskForEachTaskTaskDbtCloudTask) {
	if err := j.validatePutDbtCloudTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDbtCloudTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutDbtTask(value *JobTaskForEachTaskTaskDbtTask) {
	if err := j.validatePutDbtTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDbtTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutDependsOn(value interface{}) {
	if err := j.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutEmailNotifications(value *JobTaskForEachTaskTaskEmailNotifications) {
	if err := j.validatePutEmailNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEmailNotifications",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutGenAiComputeTask(value *JobTaskForEachTaskTaskGenAiComputeTask) {
	if err := j.validatePutGenAiComputeTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGenAiComputeTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutHealth(value *JobTaskForEachTaskTaskHealth) {
	if err := j.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHealth",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutLibrary(value interface{}) {
	if err := j.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLibrary",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutNewCluster(value *JobTaskForEachTaskTaskNewCluster) {
	if err := j.validatePutNewClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNewCluster",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutNotebookTask(value *JobTaskForEachTaskTaskNotebookTask) {
	if err := j.validatePutNotebookTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNotebookTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutNotificationSettings(value *JobTaskForEachTaskTaskNotificationSettings) {
	if err := j.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutPipelineTask(value *JobTaskForEachTaskTaskPipelineTask) {
	if err := j.validatePutPipelineTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPipelineTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutPowerBiTask(value *JobTaskForEachTaskTaskPowerBiTask) {
	if err := j.validatePutPowerBiTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPowerBiTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutPythonWheelTask(value *JobTaskForEachTaskTaskPythonWheelTask) {
	if err := j.validatePutPythonWheelTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPythonWheelTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutRunJobTask(value *JobTaskForEachTaskTaskRunJobTask) {
	if err := j.validatePutRunJobTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putRunJobTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutSparkJarTask(value *JobTaskForEachTaskTaskSparkJarTask) {
	if err := j.validatePutSparkJarTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkJarTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutSparkPythonTask(value *JobTaskForEachTaskTaskSparkPythonTask) {
	if err := j.validatePutSparkPythonTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkPythonTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutSparkSubmitTask(value *JobTaskForEachTaskTaskSparkSubmitTask) {
	if err := j.validatePutSparkSubmitTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkSubmitTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutSqlTask(value *JobTaskForEachTaskTaskSqlTask) {
	if err := j.validatePutSqlTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSqlTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) PutWebhookNotifications(value *JobTaskForEachTaskTaskWebhookNotifications) {
	if err := j.validatePutWebhookNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWebhookNotifications",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetCleanRoomsNotebookTask() {
	_jsii_.InvokeVoid(
		j,
		"resetCleanRoomsNotebookTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetConditionTask() {
	_jsii_.InvokeVoid(
		j,
		"resetConditionTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetDashboardTask() {
	_jsii_.InvokeVoid(
		j,
		"resetDashboardTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetDbtCloudTask() {
	_jsii_.InvokeVoid(
		j,
		"resetDbtCloudTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetDbtTask() {
	_jsii_.InvokeVoid(
		j,
		"resetDbtTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		j,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		j,
		"resetDescription",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetDisableAutoOptimization() {
	_jsii_.InvokeVoid(
		j,
		"resetDisableAutoOptimization",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetEmailNotifications() {
	_jsii_.InvokeVoid(
		j,
		"resetEmailNotifications",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetEnvironmentKey() {
	_jsii_.InvokeVoid(
		j,
		"resetEnvironmentKey",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetExistingClusterId() {
	_jsii_.InvokeVoid(
		j,
		"resetExistingClusterId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetGenAiComputeTask() {
	_jsii_.InvokeVoid(
		j,
		"resetGenAiComputeTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetHealth() {
	_jsii_.InvokeVoid(
		j,
		"resetHealth",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetJobClusterKey() {
	_jsii_.InvokeVoid(
		j,
		"resetJobClusterKey",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetLibrary() {
	_jsii_.InvokeVoid(
		j,
		"resetLibrary",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetMinRetryIntervalMillis() {
	_jsii_.InvokeVoid(
		j,
		"resetMinRetryIntervalMillis",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetNewCluster() {
	_jsii_.InvokeVoid(
		j,
		"resetNewCluster",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetNotebookTask() {
	_jsii_.InvokeVoid(
		j,
		"resetNotebookTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		j,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetPipelineTask() {
	_jsii_.InvokeVoid(
		j,
		"resetPipelineTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetPowerBiTask() {
	_jsii_.InvokeVoid(
		j,
		"resetPowerBiTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetPythonWheelTask() {
	_jsii_.InvokeVoid(
		j,
		"resetPythonWheelTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetRetryOnTimeout() {
	_jsii_.InvokeVoid(
		j,
		"resetRetryOnTimeout",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetRunIf() {
	_jsii_.InvokeVoid(
		j,
		"resetRunIf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetRunJobTask() {
	_jsii_.InvokeVoid(
		j,
		"resetRunJobTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetSparkJarTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkJarTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetSparkPythonTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkPythonTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetSparkSubmitTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkSubmitTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetSqlTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSqlTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ResetWebhookNotifications() {
	_jsii_.InvokeVoid(
		j,
		"resetWebhookNotifications",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

