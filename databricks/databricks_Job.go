// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/databricks/r/job databricks_job}.
type Job interface {
	cdktf.TerraformResource
	AlwaysRunning() interface{}
	SetAlwaysRunning(val interface{})
	AlwaysRunningInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EmailNotifications() JobEmailNotificationsOutputReference
	EmailNotificationsInput() *JobEmailNotifications
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
	// Experimental.
	RawOverrides() interface{}
	RetryOnTimeout() interface{}
	SetRetryOnTimeout(val interface{})
	RetryOnTimeoutInput() interface{}
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
	Url() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutEmailNotifications(value *JobEmailNotifications)
	PutGitSource(value *JobGitSource)
	PutJobCluster(value interface{})
	PutLibrary(value interface{})
	PutNewCluster(value *JobNewCluster)
	PutNotebookTask(value *JobNotebookTask)
	PutPipelineTask(value *JobPipelineTask)
	PutPythonWheelTask(value *JobPythonWheelTask)
	PutSchedule(value *JobSchedule)
	PutSparkJarTask(value *JobSparkJarTask)
	PutSparkPythonTask(value *JobSparkPythonTask)
	PutSparkSubmitTask(value *JobSparkSubmitTask)
	PutTask(value interface{})
	PutTimeouts(value *JobTimeouts)
	ResetAlwaysRunning()
	ResetEmailNotifications()
	ResetExistingClusterId()
	ResetFormat()
	ResetGitSource()
	ResetId()
	ResetJobCluster()
	ResetLibrary()
	ResetMaxConcurrentRuns()
	ResetMaxRetries()
	ResetMinRetryIntervalMillis()
	ResetName()
	ResetNewCluster()
	ResetNotebookTask()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPipelineTask()
	ResetPythonWheelTask()
	ResetRetryOnTimeout()
	ResetSchedule()
	ResetSparkJarTask()
	ResetSparkPythonTask()
	ResetSparkSubmitTask()
	ResetTags()
	ResetTask()
	ResetTimeouts()
	ResetTimeoutSeconds()
	SynthesizeAttributes() *map[string]interface{}
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

func (j *jsiiProxy_Job) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
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

func (j *jsiiProxy_Job) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/databricks/r/job databricks_job} Resource.
func NewJob(scope constructs.Construct, id *string, config *JobConfig) Job {
	_init_.Initialize()

	j := jsiiProxy_Job{}

	_jsii_.Create(
		"@cdktf/provider-databricks.Job",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/databricks/r/job databricks_job} Resource.
func NewJob_Override(j Job, scope constructs.Construct, id *string, config *JobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.Job",
		[]interface{}{scope, id, config},
		j,
	)
}

func (j *jsiiProxy_Job) SetAlwaysRunning(val interface{}) {
	_jsii_.Set(
		j,
		"alwaysRunning",
		val,
	)
}

func (j *jsiiProxy_Job) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Job) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Job) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Job) SetExistingClusterId(val *string) {
	_jsii_.Set(
		j,
		"existingClusterId",
		val,
	)
}

func (j *jsiiProxy_Job) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Job) SetFormat(val *string) {
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_Job) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Job) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Job) SetMaxConcurrentRuns(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrentRuns",
		val,
	)
}

func (j *jsiiProxy_Job) SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_Job) SetMinRetryIntervalMillis(val *float64) {
	_jsii_.Set(
		j,
		"minRetryIntervalMillis",
		val,
	)
}

func (j *jsiiProxy_Job) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Job) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Job) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Job) SetRetryOnTimeout(val interface{}) {
	_jsii_.Set(
		j,
		"retryOnTimeout",
		val,
	)
}

func (j *jsiiProxy_Job) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Job) SetTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.Job",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Job_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.Job",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (j *jsiiProxy_Job) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		j,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (j *jsiiProxy_Job) PutEmailNotifications(value *JobEmailNotifications) {
	_jsii_.InvokeVoid(
		j,
		"putEmailNotifications",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutGitSource(value *JobGitSource) {
	_jsii_.InvokeVoid(
		j,
		"putGitSource",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutJobCluster(value interface{}) {
	_jsii_.InvokeVoid(
		j,
		"putJobCluster",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutLibrary(value interface{}) {
	_jsii_.InvokeVoid(
		j,
		"putLibrary",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutNewCluster(value *JobNewCluster) {
	_jsii_.InvokeVoid(
		j,
		"putNewCluster",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutNotebookTask(value *JobNotebookTask) {
	_jsii_.InvokeVoid(
		j,
		"putNotebookTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutPipelineTask(value *JobPipelineTask) {
	_jsii_.InvokeVoid(
		j,
		"putPipelineTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutPythonWheelTask(value *JobPythonWheelTask) {
	_jsii_.InvokeVoid(
		j,
		"putPythonWheelTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutSchedule(value *JobSchedule) {
	_jsii_.InvokeVoid(
		j,
		"putSchedule",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutSparkJarTask(value *JobSparkJarTask) {
	_jsii_.InvokeVoid(
		j,
		"putSparkJarTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutSparkPythonTask(value *JobSparkPythonTask) {
	_jsii_.InvokeVoid(
		j,
		"putSparkPythonTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutSparkSubmitTask(value *JobSparkSubmitTask) {
	_jsii_.InvokeVoid(
		j,
		"putSparkSubmitTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutTask(value interface{}) {
	_jsii_.InvokeVoid(
		j,
		"putTask",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutTimeouts(value *JobTimeouts) {
	_jsii_.InvokeVoid(
		j,
		"putTimeouts",
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

func (j *jsiiProxy_Job) ResetEmailNotifications() {
	_jsii_.InvokeVoid(
		j,
		"resetEmailNotifications",
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

func (j *jsiiProxy_Job) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		j,
		"resetOverrideLogicalId",
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

func (j *jsiiProxy_Job) ResetRetryOnTimeout() {
	_jsii_.InvokeVoid(
		j,
		"resetRetryOnTimeout",
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

