// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference interface {
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
	ConditionTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskConditionTaskOutputReference
	ConditionTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskConditionTask
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DashboardTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTaskOutputReference
	DashboardTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTask
	DbtTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDbtTaskOutputReference
	DbtTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDbtTask
	DependsOn() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDependsOnList
	DependsOnInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmailNotifications() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskEmailNotificationsOutputReference
	EmailNotificationsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskEmailNotifications
	EnvironmentKey() *string
	SetEnvironmentKey(val *string)
	EnvironmentKeyInput() *string
	ExistingClusterId() *string
	SetExistingClusterId(val *string)
	ExistingClusterIdInput() *string
	// Experimental.
	Fqn() *string
	Health() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealthOutputReference
	HealthInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealth
	InternalValue() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTask
	SetInternalValue(val *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTask)
	JobClusterKey() *string
	SetJobClusterKey(val *string)
	JobClusterKeyInput() *string
	Library() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryList
	LibraryInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MinRetryIntervalMillis() *float64
	SetMinRetryIntervalMillis(val *float64)
	MinRetryIntervalMillisInput() *float64
	NewCluster() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterOutputReference
	NewClusterInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewCluster
	NotebookTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotebookTaskOutputReference
	NotebookTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotebookTask
	NotificationSettings() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference
	NotificationSettingsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings
	PipelineTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPipelineTaskOutputReference
	PipelineTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPipelineTask
	PowerBiTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTaskOutputReference
	PowerBiTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTask
	PythonWheelTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTaskOutputReference
	PythonWheelTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTask
	RetryOnTimeout() interface{}
	SetRetryOnTimeout(val interface{})
	RetryOnTimeoutInput() interface{}
	RunIf() *string
	SetRunIf(val *string)
	RunIfInput() *string
	RunJobTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskRunJobTaskOutputReference
	RunJobTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskRunJobTask
	SparkJarTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTaskOutputReference
	SparkJarTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTask
	SparkPythonTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkPythonTaskOutputReference
	SparkPythonTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkPythonTask
	SparkSubmitTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkSubmitTaskOutputReference
	SparkSubmitTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkSubmitTask
	SqlTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTaskOutputReference
	SqlTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTask
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
	WebhookNotifications() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotificationsOutputReference
	WebhookNotificationsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotifications
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutConditionTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskConditionTask)
	PutDashboardTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTask)
	PutDbtTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDbtTask)
	PutDependsOn(value interface{})
	PutEmailNotifications(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskEmailNotifications)
	PutHealth(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealth)
	PutLibrary(value interface{})
	PutNewCluster(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewCluster)
	PutNotebookTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotebookTask)
	PutNotificationSettings(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings)
	PutPipelineTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPipelineTask)
	PutPowerBiTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTask)
	PutPythonWheelTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTask)
	PutRunJobTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskRunJobTask)
	PutSparkJarTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTask)
	PutSparkPythonTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkPythonTask)
	PutSparkSubmitTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkSubmitTask)
	PutSqlTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTask)
	PutWebhookNotifications(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotifications)
	ResetConditionTask()
	ResetDashboardTask()
	ResetDbtTask()
	ResetDependsOn()
	ResetDescription()
	ResetEmailNotifications()
	ResetEnvironmentKey()
	ResetExistingClusterId()
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
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ConditionTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskConditionTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskConditionTaskOutputReference
	_jsii_.Get(
		j,
		"conditionTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ConditionTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskConditionTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskConditionTask
	_jsii_.Get(
		j,
		"conditionTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) DashboardTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTaskOutputReference
	_jsii_.Get(
		j,
		"dashboardTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) DashboardTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTask
	_jsii_.Get(
		j,
		"dashboardTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) DbtTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDbtTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDbtTaskOutputReference
	_jsii_.Get(
		j,
		"dbtTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) DbtTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDbtTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDbtTask
	_jsii_.Get(
		j,
		"dbtTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) DependsOn() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDependsOnList {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) EmailNotifications() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskEmailNotificationsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskEmailNotificationsOutputReference
	_jsii_.Get(
		j,
		"emailNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) EmailNotificationsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskEmailNotifications {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskEmailNotifications
	_jsii_.Get(
		j,
		"emailNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) EnvironmentKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) EnvironmentKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ExistingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ExistingClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) Health() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealthOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) HealthInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealth {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) InternalValue() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) JobClusterKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobClusterKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) JobClusterKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobClusterKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) Library() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryList {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) MinRetryIntervalMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) MinRetryIntervalMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) NewCluster() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterOutputReference
	_jsii_.Get(
		j,
		"newCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) NewClusterInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewCluster {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewCluster
	_jsii_.Get(
		j,
		"newClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) NotebookTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotebookTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotebookTaskOutputReference
	_jsii_.Get(
		j,
		"notebookTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) NotebookTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotebookTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotebookTask
	_jsii_.Get(
		j,
		"notebookTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) NotificationSettings() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) NotificationSettingsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PipelineTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPipelineTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPipelineTaskOutputReference
	_jsii_.Get(
		j,
		"pipelineTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PipelineTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPipelineTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPipelineTask
	_jsii_.Get(
		j,
		"pipelineTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PowerBiTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTaskOutputReference
	_jsii_.Get(
		j,
		"powerBiTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PowerBiTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTask
	_jsii_.Get(
		j,
		"powerBiTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PythonWheelTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTaskOutputReference
	_jsii_.Get(
		j,
		"pythonWheelTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PythonWheelTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTask
	_jsii_.Get(
		j,
		"pythonWheelTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) RetryOnTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) RetryOnTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) RunIf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) RunIfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) RunJobTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskRunJobTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskRunJobTaskOutputReference
	_jsii_.Get(
		j,
		"runJobTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) RunJobTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskRunJobTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskRunJobTask
	_jsii_.Get(
		j,
		"runJobTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) SparkJarTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTaskOutputReference
	_jsii_.Get(
		j,
		"sparkJarTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) SparkJarTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTask
	_jsii_.Get(
		j,
		"sparkJarTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) SparkPythonTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkPythonTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkPythonTaskOutputReference
	_jsii_.Get(
		j,
		"sparkPythonTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) SparkPythonTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkPythonTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkPythonTask
	_jsii_.Get(
		j,
		"sparkPythonTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) SparkSubmitTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkSubmitTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkSubmitTaskOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) SparkSubmitTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkSubmitTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkSubmitTask
	_jsii_.Get(
		j,
		"sparkSubmitTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) SqlTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTaskOutputReference
	_jsii_.Get(
		j,
		"sqlTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) SqlTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTask
	_jsii_.Get(
		j,
		"sqlTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) TaskKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) TaskKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) WebhookNotifications() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotificationsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotificationsOutputReference
	_jsii_.Get(
		j,
		"webhookNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) WebhookNotificationsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotifications {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotifications
	_jsii_.Get(
		j,
		"webhookNotificationsInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference_Override(d DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetEnvironmentKey(val *string) {
	if err := j.validateSetEnvironmentKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetExistingClusterId(val *string) {
	if err := j.validateSetExistingClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingClusterId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetInternalValue(val *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetJobClusterKey(val *string) {
	if err := j.validateSetJobClusterKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobClusterKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetMinRetryIntervalMillis(val *float64) {
	if err := j.validateSetMinRetryIntervalMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRetryIntervalMillis",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetRetryOnTimeout(val interface{}) {
	if err := j.validateSetRetryOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOnTimeout",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetRunIf(val *string) {
	if err := j.validateSetRunIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runIf",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetTaskKey(val *string) {
	if err := j.validateSetTaskKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutConditionTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskConditionTask) {
	if err := d.validatePutConditionTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditionTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutDashboardTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTask) {
	if err := d.validatePutDashboardTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDashboardTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutDbtTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDbtTask) {
	if err := d.validatePutDbtTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDbtTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutDependsOn(value interface{}) {
	if err := d.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutEmailNotifications(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskEmailNotifications) {
	if err := d.validatePutEmailNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmailNotifications",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutHealth(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealth) {
	if err := d.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHealth",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutLibrary(value interface{}) {
	if err := d.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLibrary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutNewCluster(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewCluster) {
	if err := d.validatePutNewClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNewCluster",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutNotebookTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotebookTask) {
	if err := d.validatePutNotebookTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotebookTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutNotificationSettings(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings) {
	if err := d.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutPipelineTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPipelineTask) {
	if err := d.validatePutPipelineTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPipelineTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutPowerBiTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTask) {
	if err := d.validatePutPowerBiTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPowerBiTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutPythonWheelTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTask) {
	if err := d.validatePutPythonWheelTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPythonWheelTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutRunJobTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskRunJobTask) {
	if err := d.validatePutRunJobTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunJobTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutSparkJarTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTask) {
	if err := d.validatePutSparkJarTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkJarTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutSparkPythonTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkPythonTask) {
	if err := d.validatePutSparkPythonTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkPythonTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutSparkSubmitTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkSubmitTask) {
	if err := d.validatePutSparkSubmitTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkSubmitTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutSqlTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTask) {
	if err := d.validatePutSqlTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) PutWebhookNotifications(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotifications) {
	if err := d.validatePutWebhookNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWebhookNotifications",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetConditionTask() {
	_jsii_.InvokeVoid(
		d,
		"resetConditionTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetDashboardTask() {
	_jsii_.InvokeVoid(
		d,
		"resetDashboardTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetDbtTask() {
	_jsii_.InvokeVoid(
		d,
		"resetDbtTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		d,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetEmailNotifications() {
	_jsii_.InvokeVoid(
		d,
		"resetEmailNotifications",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetEnvironmentKey() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironmentKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetExistingClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetExistingClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetHealth() {
	_jsii_.InvokeVoid(
		d,
		"resetHealth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetJobClusterKey() {
	_jsii_.InvokeVoid(
		d,
		"resetJobClusterKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetLibrary() {
	_jsii_.InvokeVoid(
		d,
		"resetLibrary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetMinRetryIntervalMillis() {
	_jsii_.InvokeVoid(
		d,
		"resetMinRetryIntervalMillis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetNewCluster() {
	_jsii_.InvokeVoid(
		d,
		"resetNewCluster",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetNotebookTask() {
	_jsii_.InvokeVoid(
		d,
		"resetNotebookTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetPipelineTask() {
	_jsii_.InvokeVoid(
		d,
		"resetPipelineTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetPowerBiTask() {
	_jsii_.InvokeVoid(
		d,
		"resetPowerBiTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetPythonWheelTask() {
	_jsii_.InvokeVoid(
		d,
		"resetPythonWheelTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetRetryOnTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetRetryOnTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetRunIf() {
	_jsii_.InvokeVoid(
		d,
		"resetRunIf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetRunJobTask() {
	_jsii_.InvokeVoid(
		d,
		"resetRunJobTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetSparkJarTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkJarTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetSparkPythonTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkPythonTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetSparkSubmitTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkSubmitTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetSqlTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ResetWebhookNotifications() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhookNotifications",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

