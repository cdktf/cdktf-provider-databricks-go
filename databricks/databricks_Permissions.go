// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/databricks/r/permissions databricks_permissions}.
type Permissions interface {
	cdktf.TerraformResource
	AccessControl() PermissionsAccessControlList
	AccessControlInput() interface{}
	Authorization() *string
	SetAuthorization(val *string)
	AuthorizationInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterPolicyId() *string
	SetClusterPolicyId(val *string)
	ClusterPolicyIdInput() *string
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
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	DirectoryPath() *string
	SetDirectoryPath(val *string)
	DirectoryPathInput() *string
	ExperimentId() *string
	SetExperimentId(val *string)
	ExperimentIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	JobId() *string
	SetJobId(val *string)
	JobIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotebookId() *string
	SetNotebookId(val *string)
	NotebookIdInput() *string
	NotebookPath() *string
	SetNotebookPath(val *string)
	NotebookPathInput() *string
	ObjectType() *string
	SetObjectType(val *string)
	ObjectTypeInput() *string
	PipelineId() *string
	SetPipelineId(val *string)
	PipelineIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RegisteredModelId() *string
	SetRegisteredModelId(val *string)
	RegisteredModelIdInput() *string
	RepoId() *string
	SetRepoId(val *string)
	RepoIdInput() *string
	RepoPath() *string
	SetRepoPath(val *string)
	RepoPathInput() *string
	SqlAlertId() *string
	SetSqlAlertId(val *string)
	SqlAlertIdInput() *string
	SqlDashboardId() *string
	SetSqlDashboardId(val *string)
	SqlDashboardIdInput() *string
	SqlEndpointId() *string
	SetSqlEndpointId(val *string)
	SqlEndpointIdInput() *string
	SqlQueryId() *string
	SetSqlQueryId(val *string)
	SqlQueryIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAccessControl(value interface{})
	ResetAuthorization()
	ResetClusterId()
	ResetClusterPolicyId()
	ResetDirectoryId()
	ResetDirectoryPath()
	ResetExperimentId()
	ResetId()
	ResetInstancePoolId()
	ResetJobId()
	ResetNotebookId()
	ResetNotebookPath()
	ResetObjectType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPipelineId()
	ResetRegisteredModelId()
	ResetRepoId()
	ResetRepoPath()
	ResetSqlAlertId()
	ResetSqlDashboardId()
	ResetSqlEndpointId()
	ResetSqlQueryId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Permissions
type jsiiProxy_Permissions struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Permissions) AccessControl() PermissionsAccessControlList {
	var returns PermissionsAccessControlList
	_jsii_.Get(
		j,
		"accessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) AccessControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Authorization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) AuthorizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ClusterPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ClusterPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) DirectoryPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) DirectoryPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ExperimentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ExperimentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) JobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) NotebookId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) NotebookIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) NotebookPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) NotebookPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) ObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) PipelineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) RegisteredModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) RegisteredModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredModelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) RepoId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) RepoIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) RepoPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) RepoPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) SqlAlertId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlAlertId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) SqlAlertIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlAlertIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) SqlDashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlDashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) SqlDashboardIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlDashboardIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) SqlEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) SqlEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) SqlQueryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlQueryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) SqlQueryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlQueryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Permissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/databricks/r/permissions databricks_permissions} Resource.
func NewPermissions(scope constructs.Construct, id *string, config *PermissionsConfig) Permissions {
	_init_.Initialize()

	j := jsiiProxy_Permissions{}

	_jsii_.Create(
		"@cdktf/provider-databricks.Permissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/databricks/r/permissions databricks_permissions} Resource.
func NewPermissions_Override(p Permissions, scope constructs.Construct, id *string, config *PermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.Permissions",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_Permissions) SetAuthorization(val *string) {
	_jsii_.Set(
		j,
		"authorization",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetClusterPolicyId(val *string) {
	_jsii_.Set(
		j,
		"clusterPolicyId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetDirectoryPath(val *string) {
	_jsii_.Set(
		j,
		"directoryPath",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetExperimentId(val *string) {
	_jsii_.Set(
		j,
		"experimentId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetInstancePoolId(val *string) {
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetJobId(val *string) {
	_jsii_.Set(
		j,
		"jobId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetNotebookId(val *string) {
	_jsii_.Set(
		j,
		"notebookId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetNotebookPath(val *string) {
	_jsii_.Set(
		j,
		"notebookPath",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetObjectType(val *string) {
	_jsii_.Set(
		j,
		"objectType",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetPipelineId(val *string) {
	_jsii_.Set(
		j,
		"pipelineId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetRegisteredModelId(val *string) {
	_jsii_.Set(
		j,
		"registeredModelId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetRepoId(val *string) {
	_jsii_.Set(
		j,
		"repoId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetRepoPath(val *string) {
	_jsii_.Set(
		j,
		"repoPath",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetSqlAlertId(val *string) {
	_jsii_.Set(
		j,
		"sqlAlertId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetSqlDashboardId(val *string) {
	_jsii_.Set(
		j,
		"sqlDashboardId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetSqlEndpointId(val *string) {
	_jsii_.Set(
		j,
		"sqlEndpointId",
		val,
	)
}

func (j *jsiiProxy_Permissions) SetSqlQueryId(val *string) {
	_jsii_.Set(
		j,
		"sqlQueryId",
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
func Permissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.Permissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Permissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.Permissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Permissions) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_Permissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_Permissions) PutAccessControl(value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"putAccessControl",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Permissions) ResetAuthorization() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthorization",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetClusterId() {
	_jsii_.InvokeVoid(
		p,
		"resetClusterId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetClusterPolicyId() {
	_jsii_.InvokeVoid(
		p,
		"resetClusterPolicyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		p,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetDirectoryPath() {
	_jsii_.InvokeVoid(
		p,
		"resetDirectoryPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetExperimentId() {
	_jsii_.InvokeVoid(
		p,
		"resetExperimentId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		p,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetJobId() {
	_jsii_.InvokeVoid(
		p,
		"resetJobId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetNotebookId() {
	_jsii_.InvokeVoid(
		p,
		"resetNotebookId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetNotebookPath() {
	_jsii_.InvokeVoid(
		p,
		"resetNotebookPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetObjectType() {
	_jsii_.InvokeVoid(
		p,
		"resetObjectType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetPipelineId() {
	_jsii_.InvokeVoid(
		p,
		"resetPipelineId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetRegisteredModelId() {
	_jsii_.InvokeVoid(
		p,
		"resetRegisteredModelId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetRepoId() {
	_jsii_.InvokeVoid(
		p,
		"resetRepoId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetRepoPath() {
	_jsii_.InvokeVoid(
		p,
		"resetRepoPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetSqlAlertId() {
	_jsii_.InvokeVoid(
		p,
		"resetSqlAlertId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetSqlDashboardId() {
	_jsii_.InvokeVoid(
		p,
		"resetSqlDashboardId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetSqlEndpointId() {
	_jsii_.InvokeVoid(
		p,
		"resetSqlEndpointId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) ResetSqlQueryId() {
	_jsii_.InvokeVoid(
		p,
		"resetSqlQueryId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Permissions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Permissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

