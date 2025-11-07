// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/databaseinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/database_instance databricks_database_instance}.
type DatabaseInstance interface {
	cdktf.TerraformResource
	Capacity() *string
	SetCapacity(val *string)
	CapacityInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChildInstanceRefs() DatabaseInstanceChildInstanceRefsList
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
	CreationTime() *string
	Creator() *string
	CustomTags() DatabaseInstanceCustomTagsList
	CustomTagsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveCapacity() *string
	EffectiveCustomTags() DatabaseInstanceEffectiveCustomTagsList
	EffectiveEnablePgNativeLogin() cdktf.IResolvable
	EffectiveEnableReadableSecondaries() cdktf.IResolvable
	EffectiveNodeCount() *float64
	EffectiveRetentionWindowInDays() *float64
	EffectiveStopped() cdktf.IResolvable
	EffectiveUsagePolicyId() *string
	EnablePgNativeLogin() interface{}
	SetEnablePgNativeLogin(val interface{})
	EnablePgNativeLoginInput() interface{}
	EnableReadableSecondaries() interface{}
	SetEnableReadableSecondaries(val interface{})
	EnableReadableSecondariesInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	ParentInstanceRef() DatabaseInstanceParentInstanceRefOutputReference
	ParentInstanceRefInput() interface{}
	PgVersion() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PurgeOnDelete() interface{}
	SetPurgeOnDelete(val interface{})
	PurgeOnDeleteInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReadOnlyDns() *string
	ReadWriteDns() *string
	RetentionWindowInDays() *float64
	SetRetentionWindowInDays(val *float64)
	RetentionWindowInDaysInput() *float64
	State() *string
	Stopped() interface{}
	SetStopped(val interface{})
	StoppedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Uid() *string
	UsagePolicyId() *string
	SetUsagePolicyId(val *string)
	UsagePolicyIdInput() *string
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
	PutCustomTags(value interface{})
	PutParentInstanceRef(value *DatabaseInstanceParentInstanceRef)
	ResetCapacity()
	ResetCustomTags()
	ResetEnablePgNativeLogin()
	ResetEnableReadableSecondaries()
	ResetNodeCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentInstanceRef()
	ResetPurgeOnDelete()
	ResetRetentionWindowInDays()
	ResetStopped()
	ResetUsagePolicyId()
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

// The jsii proxy struct for DatabaseInstance
type jsiiProxy_DatabaseInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseInstance) Capacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) CapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) ChildInstanceRefs() DatabaseInstanceChildInstanceRefsList {
	var returns DatabaseInstanceChildInstanceRefsList
	_jsii_.Get(
		j,
		"childInstanceRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) CustomTags() DatabaseInstanceCustomTagsList {
	var returns DatabaseInstanceCustomTagsList
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) CustomTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EffectiveCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EffectiveCustomTags() DatabaseInstanceEffectiveCustomTagsList {
	var returns DatabaseInstanceEffectiveCustomTagsList
	_jsii_.Get(
		j,
		"effectiveCustomTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EffectiveEnablePgNativeLogin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveEnablePgNativeLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EffectiveEnableReadableSecondaries() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveEnableReadableSecondaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EffectiveNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EffectiveRetentionWindowInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveRetentionWindowInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EffectiveStopped() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveStopped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EffectiveUsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveUsagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EnablePgNativeLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePgNativeLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EnablePgNativeLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePgNativeLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EnableReadableSecondaries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableReadableSecondaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EnableReadableSecondariesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableReadableSecondariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) ParentInstanceRef() DatabaseInstanceParentInstanceRefOutputReference {
	var returns DatabaseInstanceParentInstanceRefOutputReference
	_jsii_.Get(
		j,
		"parentInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) ParentInstanceRefInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parentInstanceRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) PgVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) PurgeOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) PurgeOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) ReadOnlyDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readOnlyDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) ReadWriteDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readWriteDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) RetentionWindowInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) RetentionWindowInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Stopped() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) StoppedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) UsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) UsagePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/database_instance databricks_database_instance} Resource.
func NewDatabaseInstance(scope constructs.Construct, id *string, config *DatabaseInstanceConfig) DatabaseInstance {
	_init_.Initialize()

	if err := validateNewDatabaseInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseInstance{}

	_jsii_.Create(
		"@cdktf/provider-databricks.databaseInstance.DatabaseInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/database_instance databricks_database_instance} Resource.
func NewDatabaseInstance_Override(d DatabaseInstance, scope constructs.Construct, id *string, config *DatabaseInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.databaseInstance.DatabaseInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetCapacity(val *string) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetEnablePgNativeLogin(val interface{}) {
	if err := j.validateSetEnablePgNativeLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePgNativeLogin",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetEnableReadableSecondaries(val interface{}) {
	if err := j.validateSetEnableReadableSecondariesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableReadableSecondaries",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetPurgeOnDelete(val interface{}) {
	if err := j.validateSetPurgeOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purgeOnDelete",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetRetentionWindowInDays(val *float64) {
	if err := j.validateSetRetentionWindowInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionWindowInDays",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetStopped(val interface{}) {
	if err := j.validateSetStoppedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopped",
		val,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetUsagePolicyId(val *string) {
	if err := j.validateSetUsagePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usagePolicyId",
		val,
	)
}

// Generates CDKTF code for importing a DatabaseInstance resource upon running "cdktf plan <stack-name>".
func DatabaseInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabaseInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.databaseInstance.DatabaseInstance",
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
func DatabaseInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.databaseInstance.DatabaseInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.databaseInstance.DatabaseInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.databaseInstance.DatabaseInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.databaseInstance.DatabaseInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseInstance) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabaseInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabaseInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseInstance) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabaseInstance) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseInstance) PutCustomTags(value interface{}) {
	if err := d.validatePutCustomTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseInstance) PutParentInstanceRef(value *DatabaseInstanceParentInstanceRef) {
	if err := d.validatePutParentInstanceRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParentInstanceRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetCapacity() {
	_jsii_.InvokeVoid(
		d,
		"resetCapacity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetEnablePgNativeLogin() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePgNativeLogin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetEnableReadableSecondaries() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableReadableSecondaries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetNodeCount() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetParentInstanceRef() {
	_jsii_.InvokeVoid(
		d,
		"resetParentInstanceRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetPurgeOnDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetPurgeOnDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetRetentionWindowInDays() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionWindowInDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetStopped() {
	_jsii_.InvokeVoid(
		d,
		"resetStopped",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ResetUsagePolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetUsagePolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

