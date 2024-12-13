// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscurrentconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabrickscurrentconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/current_config databricks_current_config}.
type DataDatabricksCurrentConfig interface {
	cdktf.TerraformDataSource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudType() *string
	SetCloudType(val *string)
	CloudTypeInput() *string
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
	Host() *string
	SetHost(val *string)
	HostInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsAccount() interface{}
	SetIsAccount(val interface{})
	IsAccountInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
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
	ResetAccountId()
	ResetAuthType()
	ResetCloudType()
	ResetHost()
	ResetId()
	ResetIsAccount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataDatabricksCurrentConfig
type jsiiProxy_DataDatabricksCurrentConfig struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) CloudType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) CloudTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) IsAccount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) IsAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/current_config databricks_current_config} Data Source.
func NewDataDatabricksCurrentConfig(scope constructs.Construct, id *string, config *DataDatabricksCurrentConfigConfig) DataDatabricksCurrentConfig {
	_init_.Initialize()

	if err := validateNewDataDatabricksCurrentConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCurrentConfig{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCurrentConfig.DataDatabricksCurrentConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/current_config databricks_current_config} Data Source.
func NewDataDatabricksCurrentConfig_Override(d DataDatabricksCurrentConfig, scope constructs.Construct, id *string, config *DataDatabricksCurrentConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCurrentConfig.DataDatabricksCurrentConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetCloudType(val *string) {
	if err := j.validateSetCloudTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetIsAccount(val interface{}) {
	if err := j.validateSetIsAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAccount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksCurrentConfig resource upon running "cdktf plan <stack-name>".
func DataDatabricksCurrentConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksCurrentConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCurrentConfig.DataDatabricksCurrentConfig",
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
func DataDatabricksCurrentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCurrentConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCurrentConfig.DataDatabricksCurrentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksCurrentConfig_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCurrentConfig_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCurrentConfig.DataDatabricksCurrentConfig",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksCurrentConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCurrentConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCurrentConfig.DataDatabricksCurrentConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksCurrentConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksCurrentConfig.DataDatabricksCurrentConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCurrentConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ResetAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ResetAuthType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ResetCloudType() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ResetHost() {
	_jsii_.InvokeVoid(
		d,
		"resetHost",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ResetIsAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetIsAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCurrentConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

