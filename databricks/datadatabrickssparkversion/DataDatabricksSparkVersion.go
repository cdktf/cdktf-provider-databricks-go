package datadatabrickssparkversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/datadatabrickssparkversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/data-sources/spark_version databricks_spark_version}.
type DataDatabricksSparkVersion interface {
	cdktf.TerraformDataSource
	Beta() interface{}
	SetBeta(val interface{})
	BetaInput() interface{}
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Genomics() interface{}
	SetGenomics(val interface{})
	GenomicsInput() interface{}
	Gpu() interface{}
	SetGpu(val interface{})
	GpuInput() interface{}
	Graviton() interface{}
	SetGraviton(val interface{})
	GravitonInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Latest() interface{}
	SetLatest(val interface{})
	LatestInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LongTermSupport() interface{}
	SetLongTermSupport(val interface{})
	LongTermSupportInput() interface{}
	Ml() interface{}
	SetMl(val interface{})
	MlInput() interface{}
	// The tree node.
	Node() constructs.Node
	Photon() interface{}
	SetPhoton(val interface{})
	PhotonInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Scala() *string
	SetScala(val *string)
	ScalaInput() *string
	SparkVersion() *string
	SetSparkVersion(val *string)
	SparkVersionInput() *string
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
	ResetBeta()
	ResetGenomics()
	ResetGpu()
	ResetGraviton()
	ResetId()
	ResetLatest()
	ResetLongTermSupport()
	ResetMl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhoton()
	ResetScala()
	ResetSparkVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataDatabricksSparkVersion
type jsiiProxy_DataDatabricksSparkVersion struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Beta() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"beta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) BetaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"betaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Genomics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genomics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) GenomicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genomicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Gpu() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) GpuInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Graviton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"graviton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) GravitonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gravitonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Latest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) LatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) LongTermSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"longTermSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) LongTermSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"longTermSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Ml() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) MlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Photon() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"photon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) PhotonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"photonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) Scala() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scala",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) ScalaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSparkVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/data-sources/spark_version databricks_spark_version} Data Source.
func NewDataDatabricksSparkVersion(scope constructs.Construct, id *string, config *DataDatabricksSparkVersionConfig) DataDatabricksSparkVersion {
	_init_.Initialize()

	if err := validateNewDataDatabricksSparkVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksSparkVersion{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSparkVersion.DataDatabricksSparkVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/data-sources/spark_version databricks_spark_version} Data Source.
func NewDataDatabricksSparkVersion_Override(d DataDatabricksSparkVersion, scope constructs.Construct, id *string, config *DataDatabricksSparkVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSparkVersion.DataDatabricksSparkVersion",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetBeta(val interface{}) {
	if err := j.validateSetBetaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beta",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetGenomics(val interface{}) {
	if err := j.validateSetGenomicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"genomics",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetGpu(val interface{}) {
	if err := j.validateSetGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpu",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetGraviton(val interface{}) {
	if err := j.validateSetGravitonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graviton",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetLatest(val interface{}) {
	if err := j.validateSetLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latest",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetLongTermSupport(val interface{}) {
	if err := j.validateSetLongTermSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longTermSupport",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetMl(val interface{}) {
	if err := j.validateSetMlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ml",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetPhoton(val interface{}) {
	if err := j.validateSetPhotonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"photon",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetScala(val *string) {
	if err := j.validateSetScalaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scala",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSparkVersion)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
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
func DataDatabricksSparkVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksSparkVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksSparkVersion.DataDatabricksSparkVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksSparkVersion_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksSparkVersion_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksSparkVersion.DataDatabricksSparkVersion",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksSparkVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksSparkVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksSparkVersion.DataDatabricksSparkVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksSparkVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksSparkVersion.DataDatabricksSparkVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksSparkVersion) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSparkVersion) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetBeta() {
	_jsii_.InvokeVoid(
		d,
		"resetBeta",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetGenomics() {
	_jsii_.InvokeVoid(
		d,
		"resetGenomics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetGpu() {
	_jsii_.InvokeVoid(
		d,
		"resetGpu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetGraviton() {
	_jsii_.InvokeVoid(
		d,
		"resetGraviton",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetLatest() {
	_jsii_.InvokeVoid(
		d,
		"resetLatest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetLongTermSupport() {
	_jsii_.InvokeVoid(
		d,
		"resetLongTermSupport",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetMl() {
	_jsii_.InvokeVoid(
		d,
		"resetMl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetPhoton() {
	_jsii_.InvokeVoid(
		d,
		"resetPhoton",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetScala() {
	_jsii_.InvokeVoid(
		d,
		"resetScala",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ResetSparkVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSparkVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSparkVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

