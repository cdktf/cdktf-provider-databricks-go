// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksdatabaseinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/database_instance databricks_database_instance}.
type DataDatabricksDatabaseInstance interface {
	cdktf.TerraformDataSource
	Capacity() *string
	SetCapacity(val *string)
	CapacityInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChildInstanceRefs() DataDatabricksDatabaseInstanceChildInstanceRefsList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
	Creator() *string
	CustomTags() DataDatabricksDatabaseInstanceCustomTagsList
	CustomTagsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveCapacity() *string
	EffectiveCustomTags() DataDatabricksDatabaseInstanceEffectiveCustomTagsList
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
	ParentInstanceRef() DataDatabricksDatabaseInstanceParentInstanceRefOutputReference
	ParentInstanceRefInput() interface{}
	PgVersion() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
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
	PutCustomTags(value interface{})
	PutParentInstanceRef(value *DataDatabricksDatabaseInstanceParentInstanceRef)
	ResetCapacity()
	ResetCustomTags()
	ResetEnablePgNativeLogin()
	ResetEnableReadableSecondaries()
	ResetNodeCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentInstanceRef()
	ResetRetentionWindowInDays()
	ResetStopped()
	ResetUsagePolicyId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataDatabricksDatabaseInstance
type jsiiProxy_DataDatabricksDatabaseInstance struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Capacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) CapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) ChildInstanceRefs() DataDatabricksDatabaseInstanceChildInstanceRefsList {
	var returns DataDatabricksDatabaseInstanceChildInstanceRefsList
	_jsii_.Get(
		j,
		"childInstanceRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) CustomTags() DataDatabricksDatabaseInstanceCustomTagsList {
	var returns DataDatabricksDatabaseInstanceCustomTagsList
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) CustomTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EffectiveCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EffectiveCustomTags() DataDatabricksDatabaseInstanceEffectiveCustomTagsList {
	var returns DataDatabricksDatabaseInstanceEffectiveCustomTagsList
	_jsii_.Get(
		j,
		"effectiveCustomTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EffectiveEnablePgNativeLogin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveEnablePgNativeLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EffectiveEnableReadableSecondaries() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveEnableReadableSecondaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EffectiveNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EffectiveRetentionWindowInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveRetentionWindowInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EffectiveStopped() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveStopped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EffectiveUsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveUsagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EnablePgNativeLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePgNativeLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EnablePgNativeLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePgNativeLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EnableReadableSecondaries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableReadableSecondaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) EnableReadableSecondariesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableReadableSecondariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) ParentInstanceRef() DataDatabricksDatabaseInstanceParentInstanceRefOutputReference {
	var returns DataDatabricksDatabaseInstanceParentInstanceRefOutputReference
	_jsii_.Get(
		j,
		"parentInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) ParentInstanceRefInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parentInstanceRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) PgVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) ReadOnlyDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readOnlyDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) ReadWriteDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readWriteDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) RetentionWindowInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) RetentionWindowInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Stopped() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) StoppedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) UsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance) UsagePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/database_instance databricks_database_instance} Data Source.
func NewDataDatabricksDatabaseInstance(scope constructs.Construct, id *string, config *DataDatabricksDatabaseInstanceConfig) DataDatabricksDatabaseInstance {
	_init_.Initialize()

	if err := validateNewDataDatabricksDatabaseInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDatabaseInstance{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstance.DataDatabricksDatabaseInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/database_instance databricks_database_instance} Data Source.
func NewDataDatabricksDatabaseInstance_Override(d DataDatabricksDatabaseInstance, scope constructs.Construct, id *string, config *DataDatabricksDatabaseInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstance.DataDatabricksDatabaseInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetCapacity(val *string) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetEnablePgNativeLogin(val interface{}) {
	if err := j.validateSetEnablePgNativeLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePgNativeLogin",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetEnableReadableSecondaries(val interface{}) {
	if err := j.validateSetEnableReadableSecondariesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableReadableSecondaries",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetRetentionWindowInDays(val *float64) {
	if err := j.validateSetRetentionWindowInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionWindowInDays",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetStopped(val interface{}) {
	if err := j.validateSetStoppedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopped",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstance)SetUsagePolicyId(val *string) {
	if err := j.validateSetUsagePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usagePolicyId",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksDatabaseInstance resource upon running "cdktf plan <stack-name>".
func DataDatabricksDatabaseInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksDatabaseInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstance.DataDatabricksDatabaseInstance",
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
func DataDatabricksDatabaseInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksDatabaseInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstance.DataDatabricksDatabaseInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksDatabaseInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksDatabaseInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstance.DataDatabricksDatabaseInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksDatabaseInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksDatabaseInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstance.DataDatabricksDatabaseInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksDatabaseInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstance.DataDatabricksDatabaseInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) PutCustomTags(value interface{}) {
	if err := d.validatePutCustomTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) PutParentInstanceRef(value *DataDatabricksDatabaseInstanceParentInstanceRef) {
	if err := d.validatePutParentInstanceRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParentInstanceRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetCapacity() {
	_jsii_.InvokeVoid(
		d,
		"resetCapacity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetEnablePgNativeLogin() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePgNativeLogin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetEnableReadableSecondaries() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableReadableSecondaries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetNodeCount() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetParentInstanceRef() {
	_jsii_.InvokeVoid(
		d,
		"resetParentInstanceRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetRetentionWindowInDays() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionWindowInDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetStopped() {
	_jsii_.InvokeVoid(
		d,
		"resetStopped",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ResetUsagePolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetUsagePolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

