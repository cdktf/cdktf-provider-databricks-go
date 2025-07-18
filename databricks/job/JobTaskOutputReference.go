// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskOutputReference interface {
	cdktf.ComplexObject
	CleanRoomsNotebookTask() JobTaskCleanRoomsNotebookTaskOutputReference
	CleanRoomsNotebookTaskInput() *JobTaskCleanRoomsNotebookTask
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
	ConditionTask() JobTaskConditionTaskOutputReference
	ConditionTaskInput() *JobTaskConditionTask
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DashboardTask() JobTaskDashboardTaskOutputReference
	DashboardTaskInput() *JobTaskDashboardTask
	DbtCloudTask() JobTaskDbtCloudTaskOutputReference
	DbtCloudTaskInput() *JobTaskDbtCloudTask
	DbtPlatformTask() JobTaskDbtPlatformTaskOutputReference
	DbtPlatformTaskInput() *JobTaskDbtPlatformTask
	DbtTask() JobTaskDbtTaskOutputReference
	DbtTaskInput() *JobTaskDbtTask
	DependsOn() JobTaskDependsOnList
	DependsOnInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableAutoOptimization() interface{}
	SetDisableAutoOptimization(val interface{})
	DisableAutoOptimizationInput() interface{}
	EmailNotifications() JobTaskEmailNotificationsOutputReference
	EmailNotificationsInput() *JobTaskEmailNotifications
	EnvironmentKey() *string
	SetEnvironmentKey(val *string)
	EnvironmentKeyInput() *string
	ExistingClusterId() *string
	SetExistingClusterId(val *string)
	ExistingClusterIdInput() *string
	ForEachTask() JobTaskForEachTaskOutputReference
	ForEachTaskInput() *JobTaskForEachTask
	// Experimental.
	Fqn() *string
	GenAiComputeTask() JobTaskGenAiComputeTaskOutputReference
	GenAiComputeTaskInput() *JobTaskGenAiComputeTask
	Health() JobTaskHealthOutputReference
	HealthInput() *JobTaskHealth
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobClusterKey() *string
	SetJobClusterKey(val *string)
	JobClusterKeyInput() *string
	Library() JobTaskLibraryList
	LibraryInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MinRetryIntervalMillis() *float64
	SetMinRetryIntervalMillis(val *float64)
	MinRetryIntervalMillisInput() *float64
	NewCluster() JobTaskNewClusterOutputReference
	NewClusterInput() *JobTaskNewCluster
	NotebookTask() JobTaskNotebookTaskOutputReference
	NotebookTaskInput() *JobTaskNotebookTask
	NotificationSettings() JobTaskNotificationSettingsOutputReference
	NotificationSettingsInput() *JobTaskNotificationSettings
	PipelineTask() JobTaskPipelineTaskOutputReference
	PipelineTaskInput() *JobTaskPipelineTask
	PowerBiTask() JobTaskPowerBiTaskOutputReference
	PowerBiTaskInput() *JobTaskPowerBiTask
	PythonWheelTask() JobTaskPythonWheelTaskOutputReference
	PythonWheelTaskInput() *JobTaskPythonWheelTask
	RetryOnTimeout() interface{}
	SetRetryOnTimeout(val interface{})
	RetryOnTimeoutInput() interface{}
	RunIf() *string
	SetRunIf(val *string)
	RunIfInput() *string
	RunJobTask() JobTaskRunJobTaskOutputReference
	RunJobTaskInput() *JobTaskRunJobTask
	SparkJarTask() JobTaskSparkJarTaskOutputReference
	SparkJarTaskInput() *JobTaskSparkJarTask
	SparkPythonTask() JobTaskSparkPythonTaskOutputReference
	SparkPythonTaskInput() *JobTaskSparkPythonTask
	SparkSubmitTask() JobTaskSparkSubmitTaskOutputReference
	SparkSubmitTaskInput() *JobTaskSparkSubmitTask
	SqlTask() JobTaskSqlTaskOutputReference
	SqlTaskInput() *JobTaskSqlTask
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
	WebhookNotifications() JobTaskWebhookNotificationsOutputReference
	WebhookNotificationsInput() *JobTaskWebhookNotifications
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
	PutCleanRoomsNotebookTask(value *JobTaskCleanRoomsNotebookTask)
	PutConditionTask(value *JobTaskConditionTask)
	PutDashboardTask(value *JobTaskDashboardTask)
	PutDbtCloudTask(value *JobTaskDbtCloudTask)
	PutDbtPlatformTask(value *JobTaskDbtPlatformTask)
	PutDbtTask(value *JobTaskDbtTask)
	PutDependsOn(value interface{})
	PutEmailNotifications(value *JobTaskEmailNotifications)
	PutForEachTask(value *JobTaskForEachTask)
	PutGenAiComputeTask(value *JobTaskGenAiComputeTask)
	PutHealth(value *JobTaskHealth)
	PutLibrary(value interface{})
	PutNewCluster(value *JobTaskNewCluster)
	PutNotebookTask(value *JobTaskNotebookTask)
	PutNotificationSettings(value *JobTaskNotificationSettings)
	PutPipelineTask(value *JobTaskPipelineTask)
	PutPowerBiTask(value *JobTaskPowerBiTask)
	PutPythonWheelTask(value *JobTaskPythonWheelTask)
	PutRunJobTask(value *JobTaskRunJobTask)
	PutSparkJarTask(value *JobTaskSparkJarTask)
	PutSparkPythonTask(value *JobTaskSparkPythonTask)
	PutSparkSubmitTask(value *JobTaskSparkSubmitTask)
	PutSqlTask(value *JobTaskSqlTask)
	PutWebhookNotifications(value *JobTaskWebhookNotifications)
	ResetCleanRoomsNotebookTask()
	ResetConditionTask()
	ResetDashboardTask()
	ResetDbtCloudTask()
	ResetDbtPlatformTask()
	ResetDbtTask()
	ResetDependsOn()
	ResetDescription()
	ResetDisableAutoOptimization()
	ResetEmailNotifications()
	ResetEnvironmentKey()
	ResetExistingClusterId()
	ResetForEachTask()
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

// The jsii proxy struct for JobTaskOutputReference
type jsiiProxy_JobTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskOutputReference) CleanRoomsNotebookTask() JobTaskCleanRoomsNotebookTaskOutputReference {
	var returns JobTaskCleanRoomsNotebookTaskOutputReference
	_jsii_.Get(
		j,
		"cleanRoomsNotebookTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) CleanRoomsNotebookTaskInput() *JobTaskCleanRoomsNotebookTask {
	var returns *JobTaskCleanRoomsNotebookTask
	_jsii_.Get(
		j,
		"cleanRoomsNotebookTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) ConditionTask() JobTaskConditionTaskOutputReference {
	var returns JobTaskConditionTaskOutputReference
	_jsii_.Get(
		j,
		"conditionTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) ConditionTaskInput() *JobTaskConditionTask {
	var returns *JobTaskConditionTask
	_jsii_.Get(
		j,
		"conditionTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DashboardTask() JobTaskDashboardTaskOutputReference {
	var returns JobTaskDashboardTaskOutputReference
	_jsii_.Get(
		j,
		"dashboardTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DashboardTaskInput() *JobTaskDashboardTask {
	var returns *JobTaskDashboardTask
	_jsii_.Get(
		j,
		"dashboardTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DbtCloudTask() JobTaskDbtCloudTaskOutputReference {
	var returns JobTaskDbtCloudTaskOutputReference
	_jsii_.Get(
		j,
		"dbtCloudTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DbtCloudTaskInput() *JobTaskDbtCloudTask {
	var returns *JobTaskDbtCloudTask
	_jsii_.Get(
		j,
		"dbtCloudTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DbtPlatformTask() JobTaskDbtPlatformTaskOutputReference {
	var returns JobTaskDbtPlatformTaskOutputReference
	_jsii_.Get(
		j,
		"dbtPlatformTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DbtPlatformTaskInput() *JobTaskDbtPlatformTask {
	var returns *JobTaskDbtPlatformTask
	_jsii_.Get(
		j,
		"dbtPlatformTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DbtTask() JobTaskDbtTaskOutputReference {
	var returns JobTaskDbtTaskOutputReference
	_jsii_.Get(
		j,
		"dbtTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DbtTaskInput() *JobTaskDbtTask {
	var returns *JobTaskDbtTask
	_jsii_.Get(
		j,
		"dbtTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DependsOn() JobTaskDependsOnList {
	var returns JobTaskDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DisableAutoOptimization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) DisableAutoOptimizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) EmailNotifications() JobTaskEmailNotificationsOutputReference {
	var returns JobTaskEmailNotificationsOutputReference
	_jsii_.Get(
		j,
		"emailNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) EmailNotificationsInput() *JobTaskEmailNotifications {
	var returns *JobTaskEmailNotifications
	_jsii_.Get(
		j,
		"emailNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) EnvironmentKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) EnvironmentKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) ExistingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) ExistingClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) ForEachTask() JobTaskForEachTaskOutputReference {
	var returns JobTaskForEachTaskOutputReference
	_jsii_.Get(
		j,
		"forEachTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) ForEachTaskInput() *JobTaskForEachTask {
	var returns *JobTaskForEachTask
	_jsii_.Get(
		j,
		"forEachTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) GenAiComputeTask() JobTaskGenAiComputeTaskOutputReference {
	var returns JobTaskGenAiComputeTaskOutputReference
	_jsii_.Get(
		j,
		"genAiComputeTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) GenAiComputeTaskInput() *JobTaskGenAiComputeTask {
	var returns *JobTaskGenAiComputeTask
	_jsii_.Get(
		j,
		"genAiComputeTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) Health() JobTaskHealthOutputReference {
	var returns JobTaskHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) HealthInput() *JobTaskHealth {
	var returns *JobTaskHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) JobClusterKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobClusterKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) JobClusterKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobClusterKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) Library() JobTaskLibraryList {
	var returns JobTaskLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) MinRetryIntervalMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) MinRetryIntervalMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) NewCluster() JobTaskNewClusterOutputReference {
	var returns JobTaskNewClusterOutputReference
	_jsii_.Get(
		j,
		"newCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) NewClusterInput() *JobTaskNewCluster {
	var returns *JobTaskNewCluster
	_jsii_.Get(
		j,
		"newClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) NotebookTask() JobTaskNotebookTaskOutputReference {
	var returns JobTaskNotebookTaskOutputReference
	_jsii_.Get(
		j,
		"notebookTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) NotebookTaskInput() *JobTaskNotebookTask {
	var returns *JobTaskNotebookTask
	_jsii_.Get(
		j,
		"notebookTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) NotificationSettings() JobTaskNotificationSettingsOutputReference {
	var returns JobTaskNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) NotificationSettingsInput() *JobTaskNotificationSettings {
	var returns *JobTaskNotificationSettings
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) PipelineTask() JobTaskPipelineTaskOutputReference {
	var returns JobTaskPipelineTaskOutputReference
	_jsii_.Get(
		j,
		"pipelineTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) PipelineTaskInput() *JobTaskPipelineTask {
	var returns *JobTaskPipelineTask
	_jsii_.Get(
		j,
		"pipelineTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) PowerBiTask() JobTaskPowerBiTaskOutputReference {
	var returns JobTaskPowerBiTaskOutputReference
	_jsii_.Get(
		j,
		"powerBiTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) PowerBiTaskInput() *JobTaskPowerBiTask {
	var returns *JobTaskPowerBiTask
	_jsii_.Get(
		j,
		"powerBiTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) PythonWheelTask() JobTaskPythonWheelTaskOutputReference {
	var returns JobTaskPythonWheelTaskOutputReference
	_jsii_.Get(
		j,
		"pythonWheelTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) PythonWheelTaskInput() *JobTaskPythonWheelTask {
	var returns *JobTaskPythonWheelTask
	_jsii_.Get(
		j,
		"pythonWheelTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) RetryOnTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) RetryOnTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) RunIf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) RunIfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) RunJobTask() JobTaskRunJobTaskOutputReference {
	var returns JobTaskRunJobTaskOutputReference
	_jsii_.Get(
		j,
		"runJobTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) RunJobTaskInput() *JobTaskRunJobTask {
	var returns *JobTaskRunJobTask
	_jsii_.Get(
		j,
		"runJobTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) SparkJarTask() JobTaskSparkJarTaskOutputReference {
	var returns JobTaskSparkJarTaskOutputReference
	_jsii_.Get(
		j,
		"sparkJarTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) SparkJarTaskInput() *JobTaskSparkJarTask {
	var returns *JobTaskSparkJarTask
	_jsii_.Get(
		j,
		"sparkJarTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) SparkPythonTask() JobTaskSparkPythonTaskOutputReference {
	var returns JobTaskSparkPythonTaskOutputReference
	_jsii_.Get(
		j,
		"sparkPythonTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) SparkPythonTaskInput() *JobTaskSparkPythonTask {
	var returns *JobTaskSparkPythonTask
	_jsii_.Get(
		j,
		"sparkPythonTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) SparkSubmitTask() JobTaskSparkSubmitTaskOutputReference {
	var returns JobTaskSparkSubmitTaskOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) SparkSubmitTaskInput() *JobTaskSparkSubmitTask {
	var returns *JobTaskSparkSubmitTask
	_jsii_.Get(
		j,
		"sparkSubmitTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) SqlTask() JobTaskSqlTaskOutputReference {
	var returns JobTaskSqlTaskOutputReference
	_jsii_.Get(
		j,
		"sqlTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) SqlTaskInput() *JobTaskSqlTask {
	var returns *JobTaskSqlTask
	_jsii_.Get(
		j,
		"sqlTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) TaskKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) TaskKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) WebhookNotifications() JobTaskWebhookNotificationsOutputReference {
	var returns JobTaskWebhookNotificationsOutputReference
	_jsii_.Get(
		j,
		"webhookNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) WebhookNotificationsInput() *JobTaskWebhookNotifications {
	var returns *JobTaskWebhookNotifications
	_jsii_.Get(
		j,
		"webhookNotificationsInput",
		&returns,
	)
	return returns
}


func NewJobTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobTaskOutputReference_Override(j JobTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetDisableAutoOptimization(val interface{}) {
	if err := j.validateSetDisableAutoOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoOptimization",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetEnvironmentKey(val *string) {
	if err := j.validateSetEnvironmentKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentKey",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetExistingClusterId(val *string) {
	if err := j.validateSetExistingClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingClusterId",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetJobClusterKey(val *string) {
	if err := j.validateSetJobClusterKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobClusterKey",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetMinRetryIntervalMillis(val *float64) {
	if err := j.validateSetMinRetryIntervalMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRetryIntervalMillis",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetRetryOnTimeout(val interface{}) {
	if err := j.validateSetRetryOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOnTimeout",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetRunIf(val *string) {
	if err := j.validateSetRunIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runIf",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetTaskKey(val *string) {
	if err := j.validateSetTaskKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskKey",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskOutputReference) PutCleanRoomsNotebookTask(value *JobTaskCleanRoomsNotebookTask) {
	if err := j.validatePutCleanRoomsNotebookTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCleanRoomsNotebookTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutConditionTask(value *JobTaskConditionTask) {
	if err := j.validatePutConditionTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putConditionTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutDashboardTask(value *JobTaskDashboardTask) {
	if err := j.validatePutDashboardTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDashboardTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutDbtCloudTask(value *JobTaskDbtCloudTask) {
	if err := j.validatePutDbtCloudTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDbtCloudTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutDbtPlatformTask(value *JobTaskDbtPlatformTask) {
	if err := j.validatePutDbtPlatformTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDbtPlatformTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutDbtTask(value *JobTaskDbtTask) {
	if err := j.validatePutDbtTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDbtTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutDependsOn(value interface{}) {
	if err := j.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutEmailNotifications(value *JobTaskEmailNotifications) {
	if err := j.validatePutEmailNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEmailNotifications",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutForEachTask(value *JobTaskForEachTask) {
	if err := j.validatePutForEachTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putForEachTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutGenAiComputeTask(value *JobTaskGenAiComputeTask) {
	if err := j.validatePutGenAiComputeTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGenAiComputeTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutHealth(value *JobTaskHealth) {
	if err := j.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHealth",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutLibrary(value interface{}) {
	if err := j.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLibrary",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutNewCluster(value *JobTaskNewCluster) {
	if err := j.validatePutNewClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNewCluster",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutNotebookTask(value *JobTaskNotebookTask) {
	if err := j.validatePutNotebookTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNotebookTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutNotificationSettings(value *JobTaskNotificationSettings) {
	if err := j.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutPipelineTask(value *JobTaskPipelineTask) {
	if err := j.validatePutPipelineTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPipelineTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutPowerBiTask(value *JobTaskPowerBiTask) {
	if err := j.validatePutPowerBiTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPowerBiTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutPythonWheelTask(value *JobTaskPythonWheelTask) {
	if err := j.validatePutPythonWheelTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPythonWheelTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutRunJobTask(value *JobTaskRunJobTask) {
	if err := j.validatePutRunJobTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putRunJobTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutSparkJarTask(value *JobTaskSparkJarTask) {
	if err := j.validatePutSparkJarTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkJarTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutSparkPythonTask(value *JobTaskSparkPythonTask) {
	if err := j.validatePutSparkPythonTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkPythonTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutSparkSubmitTask(value *JobTaskSparkSubmitTask) {
	if err := j.validatePutSparkSubmitTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSparkSubmitTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutSqlTask(value *JobTaskSqlTask) {
	if err := j.validatePutSqlTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSqlTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) PutWebhookNotifications(value *JobTaskWebhookNotifications) {
	if err := j.validatePutWebhookNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWebhookNotifications",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetCleanRoomsNotebookTask() {
	_jsii_.InvokeVoid(
		j,
		"resetCleanRoomsNotebookTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetConditionTask() {
	_jsii_.InvokeVoid(
		j,
		"resetConditionTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetDashboardTask() {
	_jsii_.InvokeVoid(
		j,
		"resetDashboardTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetDbtCloudTask() {
	_jsii_.InvokeVoid(
		j,
		"resetDbtCloudTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetDbtPlatformTask() {
	_jsii_.InvokeVoid(
		j,
		"resetDbtPlatformTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetDbtTask() {
	_jsii_.InvokeVoid(
		j,
		"resetDbtTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		j,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		j,
		"resetDescription",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetDisableAutoOptimization() {
	_jsii_.InvokeVoid(
		j,
		"resetDisableAutoOptimization",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetEmailNotifications() {
	_jsii_.InvokeVoid(
		j,
		"resetEmailNotifications",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetEnvironmentKey() {
	_jsii_.InvokeVoid(
		j,
		"resetEnvironmentKey",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetExistingClusterId() {
	_jsii_.InvokeVoid(
		j,
		"resetExistingClusterId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetForEachTask() {
	_jsii_.InvokeVoid(
		j,
		"resetForEachTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetGenAiComputeTask() {
	_jsii_.InvokeVoid(
		j,
		"resetGenAiComputeTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetHealth() {
	_jsii_.InvokeVoid(
		j,
		"resetHealth",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetJobClusterKey() {
	_jsii_.InvokeVoid(
		j,
		"resetJobClusterKey",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetLibrary() {
	_jsii_.InvokeVoid(
		j,
		"resetLibrary",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetMinRetryIntervalMillis() {
	_jsii_.InvokeVoid(
		j,
		"resetMinRetryIntervalMillis",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetNewCluster() {
	_jsii_.InvokeVoid(
		j,
		"resetNewCluster",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetNotebookTask() {
	_jsii_.InvokeVoid(
		j,
		"resetNotebookTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		j,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetPipelineTask() {
	_jsii_.InvokeVoid(
		j,
		"resetPipelineTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetPowerBiTask() {
	_jsii_.InvokeVoid(
		j,
		"resetPowerBiTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetPythonWheelTask() {
	_jsii_.InvokeVoid(
		j,
		"resetPythonWheelTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetRetryOnTimeout() {
	_jsii_.InvokeVoid(
		j,
		"resetRetryOnTimeout",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetRunIf() {
	_jsii_.InvokeVoid(
		j,
		"resetRunIf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetRunJobTask() {
	_jsii_.InvokeVoid(
		j,
		"resetRunJobTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetSparkJarTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkJarTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetSparkPythonTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkPythonTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetSparkSubmitTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkSubmitTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetSqlTask() {
	_jsii_.InvokeVoid(
		j,
		"resetSqlTask",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) ResetWebhookNotifications() {
	_jsii_.InvokeVoid(
		j,
		"resetWebhookNotifications",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

