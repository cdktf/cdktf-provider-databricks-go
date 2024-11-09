// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/sqlendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_endpoint databricks_sql_endpoint}.
type SqlEndpoint interface {
	cdktf.TerraformResource
	AutoStopMins() *float64
	SetAutoStopMins(val *float64)
	AutoStopMinsInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Channel() SqlEndpointChannelOutputReference
	ChannelInput() *SqlEndpointChannel
	ClusterSize() *string
	SetClusterSize(val *string)
	ClusterSizeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatorName() *string
	DataSourceId() *string
	SetDataSourceId(val *string)
	DataSourceIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnablePhoton() interface{}
	SetEnablePhoton(val interface{})
	EnablePhotonInput() interface{}
	EnableServerlessCompute() interface{}
	SetEnableServerlessCompute(val interface{})
	EnableServerlessComputeInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() SqlEndpointHealthList
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	JdbcUrl() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxNumClusters() *float64
	SetMaxNumClusters(val *float64)
	MaxNumClustersInput() *float64
	MinNumClusters() *float64
	SetMinNumClusters(val *float64)
	MinNumClustersInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NumActiveSessions() *float64
	NumClusters() *float64
	OdbcParams() SqlEndpointOdbcParamsList
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
	SpotInstancePolicy() *string
	SetSpotInstancePolicy(val *string)
	SpotInstancePolicyInput() *string
	State() *string
	Tags() SqlEndpointTagsOutputReference
	TagsInput() *SqlEndpointTags
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SqlEndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	WarehouseType() *string
	SetWarehouseType(val *string)
	WarehouseTypeInput() *string
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
	PutChannel(value *SqlEndpointChannel)
	PutTags(value *SqlEndpointTags)
	PutTimeouts(value *SqlEndpointTimeouts)
	ResetAutoStopMins()
	ResetChannel()
	ResetDataSourceId()
	ResetEnablePhoton()
	ResetEnableServerlessCompute()
	ResetId()
	ResetInstanceProfileArn()
	ResetMaxNumClusters()
	ResetMinNumClusters()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSpotInstancePolicy()
	ResetTags()
	ResetTimeouts()
	ResetWarehouseType()
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

// The jsii proxy struct for SqlEndpoint
type jsiiProxy_SqlEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SqlEndpoint) AutoStopMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoStopMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) AutoStopMinsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoStopMinsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Channel() SqlEndpointChannelOutputReference {
	var returns SqlEndpointChannelOutputReference
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) ChannelInput() *SqlEndpointChannel {
	var returns *SqlEndpointChannel
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) ClusterSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) ClusterSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) CreatorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) DataSourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) EnablePhoton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePhoton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) EnablePhotonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePhotonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) EnableServerlessCompute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServerlessCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) EnableServerlessComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServerlessComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Health() SqlEndpointHealthList {
	var returns SqlEndpointHealthList
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) JdbcUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jdbcUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) MaxNumClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) MaxNumClustersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) MinNumClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) MinNumClustersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) NumActiveSessions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numActiveSessions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) NumClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) OdbcParams() SqlEndpointOdbcParamsList {
	var returns SqlEndpointOdbcParamsList
	_jsii_.Get(
		j,
		"odbcParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) SpotInstancePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) SpotInstancePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Tags() SqlEndpointTagsOutputReference {
	var returns SqlEndpointTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) TagsInput() *SqlEndpointTags {
	var returns *SqlEndpointTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) Timeouts() SqlEndpointTimeoutsOutputReference {
	var returns SqlEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) WarehouseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlEndpoint) WarehouseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_endpoint databricks_sql_endpoint} Resource.
func NewSqlEndpoint(scope constructs.Construct, id *string, config *SqlEndpointConfig) SqlEndpoint {
	_init_.Initialize()

	if err := validateNewSqlEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-databricks.sqlEndpoint.SqlEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_endpoint databricks_sql_endpoint} Resource.
func NewSqlEndpoint_Override(s SqlEndpoint, scope constructs.Construct, id *string, config *SqlEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.sqlEndpoint.SqlEndpoint",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetAutoStopMins(val *float64) {
	if err := j.validateSetAutoStopMinsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStopMins",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetClusterSize(val *string) {
	if err := j.validateSetClusterSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSize",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetDataSourceId(val *string) {
	if err := j.validateSetDataSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetEnablePhoton(val interface{}) {
	if err := j.validateSetEnablePhotonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePhoton",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetEnableServerlessCompute(val interface{}) {
	if err := j.validateSetEnableServerlessComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServerlessCompute",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetMaxNumClusters(val *float64) {
	if err := j.validateSetMaxNumClustersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNumClusters",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetMinNumClusters(val *float64) {
	if err := j.validateSetMinNumClustersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumClusters",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetSpotInstancePolicy(val *string) {
	if err := j.validateSetSpotInstancePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotInstancePolicy",
		val,
	)
}

func (j *jsiiProxy_SqlEndpoint)SetWarehouseType(val *string) {
	if err := j.validateSetWarehouseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseType",
		val,
	)
}

// Generates CDKTF code for importing a SqlEndpoint resource upon running "cdktf plan <stack-name>".
func SqlEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSqlEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.sqlEndpoint.SqlEndpoint",
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
func SqlEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.sqlEndpoint.SqlEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqlEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.sqlEndpoint.SqlEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqlEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.sqlEndpoint.SqlEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SqlEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.sqlEndpoint.SqlEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SqlEndpoint) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SqlEndpoint) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SqlEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SqlEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SqlEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SqlEndpoint) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SqlEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SqlEndpoint) PutChannel(value *SqlEndpointChannel) {
	if err := s.validatePutChannelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putChannel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlEndpoint) PutTags(value *SqlEndpointTags) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlEndpoint) PutTimeouts(value *SqlEndpointTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetAutoStopMins() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoStopMins",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetDataSourceId() {
	_jsii_.InvokeVoid(
		s,
		"resetDataSourceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetEnablePhoton() {
	_jsii_.InvokeVoid(
		s,
		"resetEnablePhoton",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetEnableServerlessCompute() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableServerlessCompute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetMaxNumClusters() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxNumClusters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetMinNumClusters() {
	_jsii_.InvokeVoid(
		s,
		"resetMinNumClusters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetSpotInstancePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotInstancePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) ResetWarehouseType() {
	_jsii_.InvokeVoid(
		s,
		"resetWarehouseType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

