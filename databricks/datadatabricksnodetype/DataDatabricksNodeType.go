// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksnodetype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksnodetype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/node_type databricks_node_type}.
type DataDatabricksNodeType interface {
	cdktf.TerraformDataSource
	Arm() interface{}
	SetArm(val interface{})
	ArmInput() interface{}
	Category() *string
	SetCategory(val *string)
	CategoryInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Fleet() interface{}
	SetFleet(val interface{})
	FleetInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GbPerCore() *float64
	SetGbPerCore(val *float64)
	GbPerCoreInput() *float64
	Graviton() interface{}
	SetGraviton(val interface{})
	GravitonInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsIoCacheEnabled() interface{}
	SetIsIoCacheEnabled(val interface{})
	IsIoCacheEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalDisk() interface{}
	SetLocalDisk(val interface{})
	LocalDiskInput() interface{}
	LocalDiskMinSize() *float64
	SetLocalDiskMinSize(val *float64)
	LocalDiskMinSizeInput() *float64
	MinCores() *float64
	SetMinCores(val *float64)
	MinCoresInput() *float64
	MinGpus() *float64
	SetMinGpus(val *float64)
	MinGpusInput() *float64
	MinMemoryGb() *float64
	SetMinMemoryGb(val *float64)
	MinMemoryGbInput() *float64
	// The tree node.
	Node() constructs.Node
	PhotonDriverCapable() interface{}
	SetPhotonDriverCapable(val interface{})
	PhotonDriverCapableInput() interface{}
	PhotonWorkerCapable() interface{}
	SetPhotonWorkerCapable(val interface{})
	PhotonWorkerCapableInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SupportPortForwarding() interface{}
	SetSupportPortForwarding(val interface{})
	SupportPortForwardingInput() interface{}
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
	ResetArm()
	ResetCategory()
	ResetFleet()
	ResetGbPerCore()
	ResetGraviton()
	ResetId()
	ResetIsIoCacheEnabled()
	ResetLocalDisk()
	ResetLocalDiskMinSize()
	ResetMinCores()
	ResetMinGpus()
	ResetMinMemoryGb()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhotonDriverCapable()
	ResetPhotonWorkerCapable()
	ResetSupportPortForwarding()
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

// The jsii proxy struct for DataDatabricksNodeType
type jsiiProxy_DataDatabricksNodeType struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksNodeType) Arm() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) ArmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"armInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) CategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Fleet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) FleetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) GbPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gbPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) GbPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gbPerCoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Graviton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"graviton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) GravitonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gravitonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) IsIoCacheEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIoCacheEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) IsIoCacheEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIoCacheEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) LocalDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) LocalDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) LocalDiskMinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localDiskMinSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) LocalDiskMinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localDiskMinSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) MinCores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) MinCoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) MinGpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minGpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) MinGpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minGpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) MinMemoryGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) MinMemoryGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) PhotonDriverCapable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"photonDriverCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) PhotonDriverCapableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"photonDriverCapableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) PhotonWorkerCapable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"photonWorkerCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) PhotonWorkerCapableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"photonWorkerCapableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) SupportPortForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportPortForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) SupportPortForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportPortForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksNodeType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/node_type databricks_node_type} Data Source.
func NewDataDatabricksNodeType(scope constructs.Construct, id *string, config *DataDatabricksNodeTypeConfig) DataDatabricksNodeType {
	_init_.Initialize()

	if err := validateNewDataDatabricksNodeTypeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksNodeType{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/node_type databricks_node_type} Data Source.
func NewDataDatabricksNodeType_Override(d DataDatabricksNodeType, scope constructs.Construct, id *string, config *DataDatabricksNodeTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeType",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetArm(val interface{}) {
	if err := j.validateSetArmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arm",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetCategory(val *string) {
	if err := j.validateSetCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetFleet(val interface{}) {
	if err := j.validateSetFleetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetGbPerCore(val *float64) {
	if err := j.validateSetGbPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gbPerCore",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetGraviton(val interface{}) {
	if err := j.validateSetGravitonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graviton",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetIsIoCacheEnabled(val interface{}) {
	if err := j.validateSetIsIoCacheEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIoCacheEnabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetLocalDisk(val interface{}) {
	if err := j.validateSetLocalDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localDisk",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetLocalDiskMinSize(val *float64) {
	if err := j.validateSetLocalDiskMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localDiskMinSize",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetMinCores(val *float64) {
	if err := j.validateSetMinCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCores",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetMinGpus(val *float64) {
	if err := j.validateSetMinGpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minGpus",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetMinMemoryGb(val *float64) {
	if err := j.validateSetMinMemoryGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMemoryGb",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetPhotonDriverCapable(val interface{}) {
	if err := j.validateSetPhotonDriverCapableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"photonDriverCapable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetPhotonWorkerCapable(val interface{}) {
	if err := j.validateSetPhotonWorkerCapableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"photonWorkerCapable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksNodeType)SetSupportPortForwarding(val interface{}) {
	if err := j.validateSetSupportPortForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportPortForwarding",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksNodeType resource upon running "cdktf plan <stack-name>".
func DataDatabricksNodeType_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksNodeType_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeType",
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
func DataDatabricksNodeType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksNodeType_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksNodeType_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksNodeType_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeType",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksNodeType_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksNodeType_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeType",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksNodeType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeType",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksNodeType) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksNodeType) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksNodeType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksNodeType) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksNodeType) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksNodeType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksNodeType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksNodeType) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksNodeType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksNodeType) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksNodeType) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetArm() {
	_jsii_.InvokeVoid(
		d,
		"resetArm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetCategory() {
	_jsii_.InvokeVoid(
		d,
		"resetCategory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetFleet() {
	_jsii_.InvokeVoid(
		d,
		"resetFleet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetGbPerCore() {
	_jsii_.InvokeVoid(
		d,
		"resetGbPerCore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetGraviton() {
	_jsii_.InvokeVoid(
		d,
		"resetGraviton",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetIsIoCacheEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIsIoCacheEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetLocalDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetLocalDiskMinSize() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalDiskMinSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetMinCores() {
	_jsii_.InvokeVoid(
		d,
		"resetMinCores",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetMinGpus() {
	_jsii_.InvokeVoid(
		d,
		"resetMinGpus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetMinMemoryGb() {
	_jsii_.InvokeVoid(
		d,
		"resetMinMemoryGb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetPhotonDriverCapable() {
	_jsii_.InvokeVoid(
		d,
		"resetPhotonDriverCapable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetPhotonWorkerCapable() {
	_jsii_.InvokeVoid(
		d,
		"resetPhotonWorkerCapable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) ResetSupportPortForwarding() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportPortForwarding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksNodeType) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksNodeType) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksNodeType) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksNodeType) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksNodeType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksNodeType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

