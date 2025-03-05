// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/instancepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/instance_pool databricks_instance_pool}.
type InstancePool interface {
	cdktf.TerraformResource
	AwsAttributes() InstancePoolAwsAttributesOutputReference
	AwsAttributesInput() *InstancePoolAwsAttributes
	AzureAttributes() InstancePoolAzureAttributesOutputReference
	AzureAttributesInput() *InstancePoolAzureAttributes
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CustomTags() *map[string]*string
	SetCustomTags(val *map[string]*string)
	CustomTagsInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskSpec() InstancePoolDiskSpecOutputReference
	DiskSpecInput() *InstancePoolDiskSpec
	EnableElasticDisk() interface{}
	SetEnableElasticDisk(val interface{})
	EnableElasticDiskInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpAttributes() InstancePoolGcpAttributesOutputReference
	GcpAttributesInput() *InstancePoolGcpAttributes
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdleInstanceAutoterminationMinutes() *float64
	SetIdleInstanceAutoterminationMinutes(val *float64)
	IdleInstanceAutoterminationMinutesInput() *float64
	InstancePoolFleetAttributes() InstancePoolInstancePoolFleetAttributesOutputReference
	InstancePoolFleetAttributesInput() *InstancePoolInstancePoolFleetAttributes
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InstancePoolName() *string
	SetInstancePoolName(val *string)
	InstancePoolNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	MinIdleInstances() *float64
	SetMinIdleInstances(val *float64)
	MinIdleInstancesInput() *float64
	// The tree node.
	Node() constructs.Node
	NodeTypeId() *string
	SetNodeTypeId(val *string)
	NodeTypeIdInput() *string
	PreloadedDockerImage() InstancePoolPreloadedDockerImageList
	PreloadedDockerImageInput() interface{}
	PreloadedSparkVersions() *[]*string
	SetPreloadedSparkVersions(val *[]*string)
	PreloadedSparkVersionsInput() *[]*string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAwsAttributes(value *InstancePoolAwsAttributes)
	PutAzureAttributes(value *InstancePoolAzureAttributes)
	PutDiskSpec(value *InstancePoolDiskSpec)
	PutGcpAttributes(value *InstancePoolGcpAttributes)
	PutInstancePoolFleetAttributes(value *InstancePoolInstancePoolFleetAttributes)
	PutPreloadedDockerImage(value interface{})
	ResetAwsAttributes()
	ResetAzureAttributes()
	ResetCustomTags()
	ResetDiskSpec()
	ResetEnableElasticDisk()
	ResetGcpAttributes()
	ResetId()
	ResetInstancePoolFleetAttributes()
	ResetInstancePoolId()
	ResetMaxCapacity()
	ResetMinIdleInstances()
	ResetNodeTypeId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreloadedDockerImage()
	ResetPreloadedSparkVersions()
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

// The jsii proxy struct for InstancePool
type jsiiProxy_InstancePool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_InstancePool) AwsAttributes() InstancePoolAwsAttributesOutputReference {
	var returns InstancePoolAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) AwsAttributesInput() *InstancePoolAwsAttributes {
	var returns *InstancePoolAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) AzureAttributes() InstancePoolAzureAttributesOutputReference {
	var returns InstancePoolAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) AzureAttributesInput() *InstancePoolAzureAttributes {
	var returns *InstancePoolAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) DiskSpec() InstancePoolDiskSpecOutputReference {
	var returns InstancePoolDiskSpecOutputReference
	_jsii_.Get(
		j,
		"diskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) DiskSpecInput() *InstancePoolDiskSpec {
	var returns *InstancePoolDiskSpec
	_jsii_.Get(
		j,
		"diskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) GcpAttributes() InstancePoolGcpAttributesOutputReference {
	var returns InstancePoolGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) GcpAttributesInput() *InstancePoolGcpAttributes {
	var returns *InstancePoolGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) IdleInstanceAutoterminationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInstanceAutoterminationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) IdleInstanceAutoterminationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInstanceAutoterminationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) InstancePoolFleetAttributes() InstancePoolInstancePoolFleetAttributesOutputReference {
	var returns InstancePoolInstancePoolFleetAttributesOutputReference
	_jsii_.Get(
		j,
		"instancePoolFleetAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) InstancePoolFleetAttributesInput() *InstancePoolInstancePoolFleetAttributes {
	var returns *InstancePoolInstancePoolFleetAttributes
	_jsii_.Get(
		j,
		"instancePoolFleetAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) InstancePoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) InstancePoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) MinIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) MinIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) PreloadedDockerImage() InstancePoolPreloadedDockerImageList {
	var returns InstancePoolPreloadedDockerImageList
	_jsii_.Get(
		j,
		"preloadedDockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) PreloadedDockerImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preloadedDockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) PreloadedSparkVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preloadedSparkVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) PreloadedSparkVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preloadedSparkVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/instance_pool databricks_instance_pool} Resource.
func NewInstancePool(scope constructs.Construct, id *string, config *InstancePoolConfig) InstancePool {
	_init_.Initialize()

	if err := validateNewInstancePoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstancePool{}

	_jsii_.Create(
		"@cdktf/provider-databricks.instancePool.InstancePool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/instance_pool databricks_instance_pool} Resource.
func NewInstancePool_Override(i InstancePool, scope constructs.Construct, id *string, config *InstancePoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.instancePool.InstancePool",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_InstancePool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetIdleInstanceAutoterminationMinutes(val *float64) {
	if err := j.validateSetIdleInstanceAutoterminationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleInstanceAutoterminationMinutes",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetInstancePoolName(val *string) {
	if err := j.validateSetInstancePoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolName",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetMinIdleInstances(val *float64) {
	if err := j.validateSetMinIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleInstances",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetPreloadedSparkVersions(val *[]*string) {
	if err := j.validateSetPreloadedSparkVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preloadedSparkVersions",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_InstancePool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a InstancePool resource upon running "cdktf plan <stack-name>".
func InstancePool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateInstancePool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.instancePool.InstancePool",
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
func InstancePool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInstancePool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.instancePool.InstancePool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InstancePool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInstancePool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.instancePool.InstancePool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InstancePool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInstancePool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.instancePool.InstancePool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func InstancePool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.instancePool.InstancePool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InstancePool) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_InstancePool) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_InstancePool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_InstancePool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InstancePool) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_InstancePool) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InstancePool) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_InstancePool) PutAwsAttributes(value *InstancePoolAwsAttributes) {
	if err := i.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePool) PutAzureAttributes(value *InstancePoolAzureAttributes) {
	if err := i.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePool) PutDiskSpec(value *InstancePoolDiskSpec) {
	if err := i.validatePutDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDiskSpec",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePool) PutGcpAttributes(value *InstancePoolGcpAttributes) {
	if err := i.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePool) PutInstancePoolFleetAttributes(value *InstancePoolInstancePoolFleetAttributes) {
	if err := i.validatePutInstancePoolFleetAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putInstancePoolFleetAttributes",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePool) PutPreloadedDockerImage(value interface{}) {
	if err := i.validatePutPreloadedDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPreloadedDockerImage",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePool) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetCustomTags() {
	_jsii_.InvokeVoid(
		i,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetDiskSpec() {
	_jsii_.InvokeVoid(
		i,
		"resetDiskSpec",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		i,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetInstancePoolFleetAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetInstancePoolFleetAttributes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		i,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetMinIdleInstances() {
	_jsii_.InvokeVoid(
		i,
		"resetMinIdleInstances",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		i,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetPreloadedDockerImage() {
	_jsii_.InvokeVoid(
		i,
		"resetPreloadedDockerImage",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) ResetPreloadedSparkVersions() {
	_jsii_.InvokeVoid(
		i,
		"resetPreloadedSparkVersions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

