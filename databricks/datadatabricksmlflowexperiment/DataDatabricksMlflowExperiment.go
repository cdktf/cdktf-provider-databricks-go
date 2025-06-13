// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmlflowexperiment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksmlflowexperiment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/mlflow_experiment databricks_mlflow_experiment}.
type DataDatabricksMlflowExperiment interface {
	cdktf.TerraformDataSource
	ArtifactLocation() *string
	SetArtifactLocation(val *string)
	ArtifactLocationInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *float64
	SetCreationTime(val *float64)
	CreationTimeInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	LastUpdateTime() *float64
	SetLastUpdateTime(val *float64)
	LastUpdateTimeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleStage() *string
	SetLifecycleStage(val *string)
	LifecycleStageInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Tags() DataDatabricksMlflowExperimentTagsList
	TagsInput() interface{}
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
	PutTags(value interface{})
	ResetArtifactLocation()
	ResetCreationTime()
	ResetExperimentId()
	ResetId()
	ResetLastUpdateTime()
	ResetLifecycleStage()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
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

// The jsii proxy struct for DataDatabricksMlflowExperiment
type jsiiProxy_DataDatabricksMlflowExperiment struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) ArtifactLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) ArtifactLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) CreationTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) ExperimentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) ExperimentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) LastUpdateTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastUpdateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) LastUpdateTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastUpdateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) LifecycleStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) LifecycleStageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) Tags() DataDatabricksMlflowExperimentTagsList {
	var returns DataDatabricksMlflowExperimentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/mlflow_experiment databricks_mlflow_experiment} Data Source.
func NewDataDatabricksMlflowExperiment(scope constructs.Construct, id *string, config *DataDatabricksMlflowExperimentConfig) DataDatabricksMlflowExperiment {
	_init_.Initialize()

	if err := validateNewDataDatabricksMlflowExperimentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksMlflowExperiment{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMlflowExperiment.DataDatabricksMlflowExperiment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/mlflow_experiment databricks_mlflow_experiment} Data Source.
func NewDataDatabricksMlflowExperiment_Override(d DataDatabricksMlflowExperiment, scope constructs.Construct, id *string, config *DataDatabricksMlflowExperimentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMlflowExperiment.DataDatabricksMlflowExperiment",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetArtifactLocation(val *string) {
	if err := j.validateSetArtifactLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactLocation",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetCreationTime(val *float64) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetExperimentId(val *string) {
	if err := j.validateSetExperimentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experimentId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetLastUpdateTime(val *float64) {
	if err := j.validateSetLastUpdateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastUpdateTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetLifecycleStage(val *string) {
	if err := j.validateSetLifecycleStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleStage",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowExperiment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksMlflowExperiment resource upon running "cdktf plan <stack-name>".
func DataDatabricksMlflowExperiment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksMlflowExperiment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksMlflowExperiment.DataDatabricksMlflowExperiment",
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
func DataDatabricksMlflowExperiment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksMlflowExperiment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksMlflowExperiment.DataDatabricksMlflowExperiment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksMlflowExperiment_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksMlflowExperiment_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksMlflowExperiment.DataDatabricksMlflowExperiment",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksMlflowExperiment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksMlflowExperiment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksMlflowExperiment.DataDatabricksMlflowExperiment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksMlflowExperiment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksMlflowExperiment.DataDatabricksMlflowExperiment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMlflowExperiment) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetArtifactLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetArtifactLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetCreationTime() {
	_jsii_.InvokeVoid(
		d,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetExperimentId() {
	_jsii_.InvokeVoid(
		d,
		"resetExperimentId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetLastUpdateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLastUpdateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetLifecycleStage() {
	_jsii_.InvokeVoid(
		d,
		"resetLifecycleStage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMlflowExperiment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

