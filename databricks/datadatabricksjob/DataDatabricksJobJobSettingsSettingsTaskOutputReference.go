// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsTaskOutputReference interface {
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
	ComputeKey() *string
	SetComputeKey(val *string)
	ComputeKeyInput() *string
	ConditionTask() DataDatabricksJobJobSettingsSettingsTaskConditionTaskOutputReference
	ConditionTaskInput() *DataDatabricksJobJobSettingsSettingsTaskConditionTask
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DbtTask() DataDatabricksJobJobSettingsSettingsTaskDbtTaskOutputReference
	DbtTaskInput() *DataDatabricksJobJobSettingsSettingsTaskDbtTask
	DependsOn() DataDatabricksJobJobSettingsSettingsTaskDependsOnList
	DependsOnInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmailNotifications() DataDatabricksJobJobSettingsSettingsTaskEmailNotificationsOutputReference
	EmailNotificationsInput() *DataDatabricksJobJobSettingsSettingsTaskEmailNotifications
	ExistingClusterId() *string
	SetExistingClusterId(val *string)
	ExistingClusterIdInput() *string
	ForEachTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskOutputReference
	ForEachTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTask
	// Experimental.
	Fqn() *string
	Health() DataDatabricksJobJobSettingsSettingsTaskHealthOutputReference
	HealthInput() *DataDatabricksJobJobSettingsSettingsTaskHealth
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobClusterKey() *string
	SetJobClusterKey(val *string)
	JobClusterKeyInput() *string
	Library() DataDatabricksJobJobSettingsSettingsTaskLibraryList
	LibraryInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MinRetryIntervalMillis() *float64
	SetMinRetryIntervalMillis(val *float64)
	MinRetryIntervalMillisInput() *float64
	NewCluster() DataDatabricksJobJobSettingsSettingsTaskNewClusterOutputReference
	NewClusterInput() *DataDatabricksJobJobSettingsSettingsTaskNewCluster
	NotebookTask() DataDatabricksJobJobSettingsSettingsTaskNotebookTaskOutputReference
	NotebookTaskInput() *DataDatabricksJobJobSettingsSettingsTaskNotebookTask
	NotificationSettings() DataDatabricksJobJobSettingsSettingsTaskNotificationSettingsOutputReference
	NotificationSettingsInput() *DataDatabricksJobJobSettingsSettingsTaskNotificationSettings
	PipelineTask() DataDatabricksJobJobSettingsSettingsTaskPipelineTaskOutputReference
	PipelineTaskInput() *DataDatabricksJobJobSettingsSettingsTaskPipelineTask
	PythonWheelTask() DataDatabricksJobJobSettingsSettingsTaskPythonWheelTaskOutputReference
	PythonWheelTaskInput() *DataDatabricksJobJobSettingsSettingsTaskPythonWheelTask
	RetryOnTimeout() interface{}
	SetRetryOnTimeout(val interface{})
	RetryOnTimeoutInput() interface{}
	RunIf() *string
	SetRunIf(val *string)
	RunIfInput() *string
	RunJobTask() DataDatabricksJobJobSettingsSettingsTaskRunJobTaskOutputReference
	RunJobTaskInput() *DataDatabricksJobJobSettingsSettingsTaskRunJobTask
	SparkJarTask() DataDatabricksJobJobSettingsSettingsTaskSparkJarTaskOutputReference
	SparkJarTaskInput() *DataDatabricksJobJobSettingsSettingsTaskSparkJarTask
	SparkPythonTask() DataDatabricksJobJobSettingsSettingsTaskSparkPythonTaskOutputReference
	SparkPythonTaskInput() *DataDatabricksJobJobSettingsSettingsTaskSparkPythonTask
	SparkSubmitTask() DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTaskOutputReference
	SparkSubmitTaskInput() *DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTask
	SqlTask() DataDatabricksJobJobSettingsSettingsTaskSqlTaskOutputReference
	SqlTaskInput() *DataDatabricksJobJobSettingsSettingsTaskSqlTask
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
	WebhookNotifications() DataDatabricksJobJobSettingsSettingsTaskWebhookNotificationsOutputReference
	WebhookNotificationsInput() *DataDatabricksJobJobSettingsSettingsTaskWebhookNotifications
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
	PutConditionTask(value *DataDatabricksJobJobSettingsSettingsTaskConditionTask)
	PutDbtTask(value *DataDatabricksJobJobSettingsSettingsTaskDbtTask)
	PutDependsOn(value interface{})
	PutEmailNotifications(value *DataDatabricksJobJobSettingsSettingsTaskEmailNotifications)
	PutForEachTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTask)
	PutHealth(value *DataDatabricksJobJobSettingsSettingsTaskHealth)
	PutLibrary(value interface{})
	PutNewCluster(value *DataDatabricksJobJobSettingsSettingsTaskNewCluster)
	PutNotebookTask(value *DataDatabricksJobJobSettingsSettingsTaskNotebookTask)
	PutNotificationSettings(value *DataDatabricksJobJobSettingsSettingsTaskNotificationSettings)
	PutPipelineTask(value *DataDatabricksJobJobSettingsSettingsTaskPipelineTask)
	PutPythonWheelTask(value *DataDatabricksJobJobSettingsSettingsTaskPythonWheelTask)
	PutRunJobTask(value *DataDatabricksJobJobSettingsSettingsTaskRunJobTask)
	PutSparkJarTask(value *DataDatabricksJobJobSettingsSettingsTaskSparkJarTask)
	PutSparkPythonTask(value *DataDatabricksJobJobSettingsSettingsTaskSparkPythonTask)
	PutSparkSubmitTask(value *DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTask)
	PutSqlTask(value *DataDatabricksJobJobSettingsSettingsTaskSqlTask)
	PutWebhookNotifications(value *DataDatabricksJobJobSettingsSettingsTaskWebhookNotifications)
	ResetComputeKey()
	ResetConditionTask()
	ResetDbtTask()
	ResetDependsOn()
	ResetDescription()
	ResetEmailNotifications()
	ResetExistingClusterId()
	ResetForEachTask()
	ResetHealth()
	ResetJobClusterKey()
	ResetLibrary()
	ResetMaxRetries()
	ResetMinRetryIntervalMillis()
	ResetNewCluster()
	ResetNotebookTask()
	ResetNotificationSettings()
	ResetPipelineTask()
	ResetPythonWheelTask()
	ResetRetryOnTimeout()
	ResetRunIf()
	ResetRunJobTask()
	ResetSparkJarTask()
	ResetSparkPythonTask()
	ResetSparkSubmitTask()
	ResetSqlTask()
	ResetTaskKey()
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

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsTaskOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ComputeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ComputeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ConditionTask() DataDatabricksJobJobSettingsSettingsTaskConditionTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskConditionTaskOutputReference
	_jsii_.Get(
		j,
		"conditionTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ConditionTaskInput() *DataDatabricksJobJobSettingsSettingsTaskConditionTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskConditionTask
	_jsii_.Get(
		j,
		"conditionTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) DbtTask() DataDatabricksJobJobSettingsSettingsTaskDbtTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskDbtTaskOutputReference
	_jsii_.Get(
		j,
		"dbtTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) DbtTaskInput() *DataDatabricksJobJobSettingsSettingsTaskDbtTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskDbtTask
	_jsii_.Get(
		j,
		"dbtTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) DependsOn() DataDatabricksJobJobSettingsSettingsTaskDependsOnList {
	var returns DataDatabricksJobJobSettingsSettingsTaskDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) EmailNotifications() DataDatabricksJobJobSettingsSettingsTaskEmailNotificationsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskEmailNotificationsOutputReference
	_jsii_.Get(
		j,
		"emailNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) EmailNotificationsInput() *DataDatabricksJobJobSettingsSettingsTaskEmailNotifications {
	var returns *DataDatabricksJobJobSettingsSettingsTaskEmailNotifications
	_jsii_.Get(
		j,
		"emailNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ExistingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ExistingClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ForEachTask() DataDatabricksJobJobSettingsSettingsTaskForEachTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskOutputReference
	_jsii_.Get(
		j,
		"forEachTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ForEachTaskInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTask
	_jsii_.Get(
		j,
		"forEachTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) Health() DataDatabricksJobJobSettingsSettingsTaskHealthOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) HealthInput() *DataDatabricksJobJobSettingsSettingsTaskHealth {
	var returns *DataDatabricksJobJobSettingsSettingsTaskHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) JobClusterKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobClusterKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) JobClusterKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobClusterKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) Library() DataDatabricksJobJobSettingsSettingsTaskLibraryList {
	var returns DataDatabricksJobJobSettingsSettingsTaskLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) MinRetryIntervalMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) MinRetryIntervalMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetryIntervalMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) NewCluster() DataDatabricksJobJobSettingsSettingsTaskNewClusterOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskNewClusterOutputReference
	_jsii_.Get(
		j,
		"newCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) NewClusterInput() *DataDatabricksJobJobSettingsSettingsTaskNewCluster {
	var returns *DataDatabricksJobJobSettingsSettingsTaskNewCluster
	_jsii_.Get(
		j,
		"newClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) NotebookTask() DataDatabricksJobJobSettingsSettingsTaskNotebookTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskNotebookTaskOutputReference
	_jsii_.Get(
		j,
		"notebookTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) NotebookTaskInput() *DataDatabricksJobJobSettingsSettingsTaskNotebookTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskNotebookTask
	_jsii_.Get(
		j,
		"notebookTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) NotificationSettings() DataDatabricksJobJobSettingsSettingsTaskNotificationSettingsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) NotificationSettingsInput() *DataDatabricksJobJobSettingsSettingsTaskNotificationSettings {
	var returns *DataDatabricksJobJobSettingsSettingsTaskNotificationSettings
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PipelineTask() DataDatabricksJobJobSettingsSettingsTaskPipelineTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskPipelineTaskOutputReference
	_jsii_.Get(
		j,
		"pipelineTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PipelineTaskInput() *DataDatabricksJobJobSettingsSettingsTaskPipelineTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskPipelineTask
	_jsii_.Get(
		j,
		"pipelineTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PythonWheelTask() DataDatabricksJobJobSettingsSettingsTaskPythonWheelTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskPythonWheelTaskOutputReference
	_jsii_.Get(
		j,
		"pythonWheelTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PythonWheelTaskInput() *DataDatabricksJobJobSettingsSettingsTaskPythonWheelTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskPythonWheelTask
	_jsii_.Get(
		j,
		"pythonWheelTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) RetryOnTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) RetryOnTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) RunIf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) RunIfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) RunJobTask() DataDatabricksJobJobSettingsSettingsTaskRunJobTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskRunJobTaskOutputReference
	_jsii_.Get(
		j,
		"runJobTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) RunJobTaskInput() *DataDatabricksJobJobSettingsSettingsTaskRunJobTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskRunJobTask
	_jsii_.Get(
		j,
		"runJobTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) SparkJarTask() DataDatabricksJobJobSettingsSettingsTaskSparkJarTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskSparkJarTaskOutputReference
	_jsii_.Get(
		j,
		"sparkJarTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) SparkJarTaskInput() *DataDatabricksJobJobSettingsSettingsTaskSparkJarTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskSparkJarTask
	_jsii_.Get(
		j,
		"sparkJarTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) SparkPythonTask() DataDatabricksJobJobSettingsSettingsTaskSparkPythonTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskSparkPythonTaskOutputReference
	_jsii_.Get(
		j,
		"sparkPythonTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) SparkPythonTaskInput() *DataDatabricksJobJobSettingsSettingsTaskSparkPythonTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskSparkPythonTask
	_jsii_.Get(
		j,
		"sparkPythonTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) SparkSubmitTask() DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTaskOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) SparkSubmitTaskInput() *DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTask
	_jsii_.Get(
		j,
		"sparkSubmitTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) SqlTask() DataDatabricksJobJobSettingsSettingsTaskSqlTaskOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskSqlTaskOutputReference
	_jsii_.Get(
		j,
		"sqlTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) SqlTaskInput() *DataDatabricksJobJobSettingsSettingsTaskSqlTask {
	var returns *DataDatabricksJobJobSettingsSettingsTaskSqlTask
	_jsii_.Get(
		j,
		"sqlTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) TaskKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) TaskKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) WebhookNotifications() DataDatabricksJobJobSettingsSettingsTaskWebhookNotificationsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskWebhookNotificationsOutputReference
	_jsii_.Get(
		j,
		"webhookNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) WebhookNotificationsInput() *DataDatabricksJobJobSettingsSettingsTaskWebhookNotifications {
	var returns *DataDatabricksJobJobSettingsSettingsTaskWebhookNotifications
	_jsii_.Get(
		j,
		"webhookNotificationsInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksJobJobSettingsSettingsTaskOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsTaskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsTaskOutputReference_Override(d DataDatabricksJobJobSettingsSettingsTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetComputeKey(val *string) {
	if err := j.validateSetComputeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetExistingClusterId(val *string) {
	if err := j.validateSetExistingClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingClusterId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetJobClusterKey(val *string) {
	if err := j.validateSetJobClusterKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobClusterKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetMinRetryIntervalMillis(val *float64) {
	if err := j.validateSetMinRetryIntervalMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRetryIntervalMillis",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetRetryOnTimeout(val interface{}) {
	if err := j.validateSetRetryOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOnTimeout",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetRunIf(val *string) {
	if err := j.validateSetRunIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runIf",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetTaskKey(val *string) {
	if err := j.validateSetTaskKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutConditionTask(value *DataDatabricksJobJobSettingsSettingsTaskConditionTask) {
	if err := d.validatePutConditionTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditionTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutDbtTask(value *DataDatabricksJobJobSettingsSettingsTaskDbtTask) {
	if err := d.validatePutDbtTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDbtTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutDependsOn(value interface{}) {
	if err := d.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutEmailNotifications(value *DataDatabricksJobJobSettingsSettingsTaskEmailNotifications) {
	if err := d.validatePutEmailNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmailNotifications",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutForEachTask(value *DataDatabricksJobJobSettingsSettingsTaskForEachTask) {
	if err := d.validatePutForEachTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForEachTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutHealth(value *DataDatabricksJobJobSettingsSettingsTaskHealth) {
	if err := d.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHealth",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutLibrary(value interface{}) {
	if err := d.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLibrary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutNewCluster(value *DataDatabricksJobJobSettingsSettingsTaskNewCluster) {
	if err := d.validatePutNewClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNewCluster",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutNotebookTask(value *DataDatabricksJobJobSettingsSettingsTaskNotebookTask) {
	if err := d.validatePutNotebookTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotebookTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutNotificationSettings(value *DataDatabricksJobJobSettingsSettingsTaskNotificationSettings) {
	if err := d.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutPipelineTask(value *DataDatabricksJobJobSettingsSettingsTaskPipelineTask) {
	if err := d.validatePutPipelineTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPipelineTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutPythonWheelTask(value *DataDatabricksJobJobSettingsSettingsTaskPythonWheelTask) {
	if err := d.validatePutPythonWheelTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPythonWheelTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutRunJobTask(value *DataDatabricksJobJobSettingsSettingsTaskRunJobTask) {
	if err := d.validatePutRunJobTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunJobTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutSparkJarTask(value *DataDatabricksJobJobSettingsSettingsTaskSparkJarTask) {
	if err := d.validatePutSparkJarTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkJarTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutSparkPythonTask(value *DataDatabricksJobJobSettingsSettingsTaskSparkPythonTask) {
	if err := d.validatePutSparkPythonTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkPythonTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutSparkSubmitTask(value *DataDatabricksJobJobSettingsSettingsTaskSparkSubmitTask) {
	if err := d.validatePutSparkSubmitTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkSubmitTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutSqlTask(value *DataDatabricksJobJobSettingsSettingsTaskSqlTask) {
	if err := d.validatePutSqlTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlTask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) PutWebhookNotifications(value *DataDatabricksJobJobSettingsSettingsTaskWebhookNotifications) {
	if err := d.validatePutWebhookNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWebhookNotifications",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetComputeKey() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetConditionTask() {
	_jsii_.InvokeVoid(
		d,
		"resetConditionTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetDbtTask() {
	_jsii_.InvokeVoid(
		d,
		"resetDbtTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		d,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetEmailNotifications() {
	_jsii_.InvokeVoid(
		d,
		"resetEmailNotifications",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetExistingClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetExistingClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetForEachTask() {
	_jsii_.InvokeVoid(
		d,
		"resetForEachTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetHealth() {
	_jsii_.InvokeVoid(
		d,
		"resetHealth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetJobClusterKey() {
	_jsii_.InvokeVoid(
		d,
		"resetJobClusterKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetLibrary() {
	_jsii_.InvokeVoid(
		d,
		"resetLibrary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetMinRetryIntervalMillis() {
	_jsii_.InvokeVoid(
		d,
		"resetMinRetryIntervalMillis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetNewCluster() {
	_jsii_.InvokeVoid(
		d,
		"resetNewCluster",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetNotebookTask() {
	_jsii_.InvokeVoid(
		d,
		"resetNotebookTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetPipelineTask() {
	_jsii_.InvokeVoid(
		d,
		"resetPipelineTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetPythonWheelTask() {
	_jsii_.InvokeVoid(
		d,
		"resetPythonWheelTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetRetryOnTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetRetryOnTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetRunIf() {
	_jsii_.InvokeVoid(
		d,
		"resetRunIf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetRunJobTask() {
	_jsii_.InvokeVoid(
		d,
		"resetRunJobTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetSparkJarTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkJarTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetSparkPythonTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkPythonTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetSparkSubmitTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkSubmitTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetSqlTask() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetTaskKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTaskKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ResetWebhookNotifications() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhookNotifications",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

