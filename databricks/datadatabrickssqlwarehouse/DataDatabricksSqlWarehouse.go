// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssqlwarehouse

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/datadatabrickssqlwarehouse/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/sql_warehouse databricks_sql_warehouse}.
type DataDatabricksSqlWarehouse interface {
	cdktf.TerraformDataSource
	AutoStopMins() *float64
	SetAutoStopMins(val *float64)
	AutoStopMinsInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Channel() DataDatabricksSqlWarehouseChannelOutputReference
	ChannelInput() *DataDatabricksSqlWarehouseChannel
	ClusterSize() *string
	SetClusterSize(val *string)
	ClusterSizeInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	JdbcUrl() *string
	SetJdbcUrl(val *string)
	JdbcUrlInput() *string
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
	NumClusters() *float64
	SetNumClusters(val *float64)
	NumClustersInput() *float64
	OdbcParams() DataDatabricksSqlWarehouseOdbcParamsOutputReference
	OdbcParamsInput() *DataDatabricksSqlWarehouseOdbcParams
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SpotInstancePolicy() *string
	SetSpotInstancePolicy(val *string)
	SpotInstancePolicyInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	Tags() DataDatabricksSqlWarehouseTagsOutputReference
	TagsInput() *DataDatabricksSqlWarehouseTags
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
	PutChannel(value *DataDatabricksSqlWarehouseChannel)
	PutOdbcParams(value *DataDatabricksSqlWarehouseOdbcParams)
	PutTags(value *DataDatabricksSqlWarehouseTags)
	ResetAutoStopMins()
	ResetChannel()
	ResetClusterSize()
	ResetDataSourceId()
	ResetEnablePhoton()
	ResetEnableServerlessCompute()
	ResetId()
	ResetInstanceProfileArn()
	ResetJdbcUrl()
	ResetMaxNumClusters()
	ResetMinNumClusters()
	ResetName()
	ResetNumClusters()
	ResetOdbcParams()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSpotInstancePolicy()
	ResetState()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataDatabricksSqlWarehouse
type jsiiProxy_DataDatabricksSqlWarehouse struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) AutoStopMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoStopMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) AutoStopMinsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoStopMinsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Channel() DataDatabricksSqlWarehouseChannelOutputReference {
	var returns DataDatabricksSqlWarehouseChannelOutputReference
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) ChannelInput() *DataDatabricksSqlWarehouseChannel {
	var returns *DataDatabricksSqlWarehouseChannel
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) ClusterSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) ClusterSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) DataSourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) EnablePhoton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePhoton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) EnablePhotonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePhotonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) EnableServerlessCompute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServerlessCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) EnableServerlessComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServerlessComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) JdbcUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jdbcUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) JdbcUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jdbcUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) MaxNumClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) MaxNumClustersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) MinNumClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) MinNumClustersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) NumClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) NumClustersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) OdbcParams() DataDatabricksSqlWarehouseOdbcParamsOutputReference {
	var returns DataDatabricksSqlWarehouseOdbcParamsOutputReference
	_jsii_.Get(
		j,
		"odbcParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) OdbcParamsInput() *DataDatabricksSqlWarehouseOdbcParams {
	var returns *DataDatabricksSqlWarehouseOdbcParams
	_jsii_.Get(
		j,
		"odbcParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) SpotInstancePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) SpotInstancePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) Tags() DataDatabricksSqlWarehouseTagsOutputReference {
	var returns DataDatabricksSqlWarehouseTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) TagsInput() *DataDatabricksSqlWarehouseTags {
	var returns *DataDatabricksSqlWarehouseTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/sql_warehouse databricks_sql_warehouse} Data Source.
func NewDataDatabricksSqlWarehouse(scope constructs.Construct, id *string, config *DataDatabricksSqlWarehouseConfig) DataDatabricksSqlWarehouse {
	_init_.Initialize()

	if err := validateNewDataDatabricksSqlWarehouseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksSqlWarehouse{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSqlWarehouse.DataDatabricksSqlWarehouse",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/data-sources/sql_warehouse databricks_sql_warehouse} Data Source.
func NewDataDatabricksSqlWarehouse_Override(d DataDatabricksSqlWarehouse, scope constructs.Construct, id *string, config *DataDatabricksSqlWarehouseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSqlWarehouse.DataDatabricksSqlWarehouse",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetAutoStopMins(val *float64) {
	if err := j.validateSetAutoStopMinsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStopMins",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetClusterSize(val *string) {
	if err := j.validateSetClusterSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSize",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetDataSourceId(val *string) {
	if err := j.validateSetDataSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetEnablePhoton(val interface{}) {
	if err := j.validateSetEnablePhotonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePhoton",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetEnableServerlessCompute(val interface{}) {
	if err := j.validateSetEnableServerlessComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServerlessCompute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetJdbcUrl(val *string) {
	if err := j.validateSetJdbcUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcUrl",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetMaxNumClusters(val *float64) {
	if err := j.validateSetMaxNumClustersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNumClusters",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetMinNumClusters(val *float64) {
	if err := j.validateSetMinNumClustersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumClusters",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetNumClusters(val *float64) {
	if err := j.validateSetNumClustersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numClusters",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetSpotInstancePolicy(val *string) {
	if err := j.validateSetSpotInstancePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotInstancePolicy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSqlWarehouse)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksSqlWarehouse resource upon running "cdktf plan <stack-name>".
func DataDatabricksSqlWarehouse_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksSqlWarehouse_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksSqlWarehouse.DataDatabricksSqlWarehouse",
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
func DataDatabricksSqlWarehouse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksSqlWarehouse_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksSqlWarehouse.DataDatabricksSqlWarehouse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksSqlWarehouse_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksSqlWarehouse_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksSqlWarehouse.DataDatabricksSqlWarehouse",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksSqlWarehouse_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksSqlWarehouse_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksSqlWarehouse.DataDatabricksSqlWarehouse",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksSqlWarehouse_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksSqlWarehouse.DataDatabricksSqlWarehouse",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSqlWarehouse) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) PutChannel(value *DataDatabricksSqlWarehouseChannel) {
	if err := d.validatePutChannelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChannel",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) PutOdbcParams(value *DataDatabricksSqlWarehouseOdbcParams) {
	if err := d.validatePutOdbcParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOdbcParams",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) PutTags(value *DataDatabricksSqlWarehouseTags) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetAutoStopMins() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoStopMins",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetChannel() {
	_jsii_.InvokeVoid(
		d,
		"resetChannel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetClusterSize() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetDataSourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetEnablePhoton() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePhoton",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetEnableServerlessCompute() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableServerlessCompute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetJdbcUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetJdbcUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetMaxNumClusters() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxNumClusters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetMinNumClusters() {
	_jsii_.InvokeVoid(
		d,
		"resetMinNumClusters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetNumClusters() {
	_jsii_.InvokeVoid(
		d,
		"resetNumClusters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetOdbcParams() {
	_jsii_.InvokeVoid(
		d,
		"resetOdbcParams",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetSpotInstancePolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetSpotInstancePolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetState() {
	_jsii_.InvokeVoid(
		d,
		"resetState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSqlWarehouse) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

