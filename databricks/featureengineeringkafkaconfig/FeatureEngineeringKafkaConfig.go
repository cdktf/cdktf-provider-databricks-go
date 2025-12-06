// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/featureengineeringkafkaconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config databricks_feature_engineering_kafka_config}.
type FeatureEngineeringKafkaConfig interface {
	cdktf.TerraformResource
	AuthConfig() FeatureEngineeringKafkaConfigAuthConfigOutputReference
	AuthConfigInput() interface{}
	BootstrapServers() *string
	SetBootstrapServers(val *string)
	BootstrapServersInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExtraOptions() *map[string]*string
	SetExtraOptions(val *map[string]*string)
	ExtraOptionsInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	KeySchema() FeatureEngineeringKafkaConfigKeySchemaOutputReference
	KeySchemaInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
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
	SubscriptionMode() FeatureEngineeringKafkaConfigSubscriptionModeOutputReference
	SubscriptionModeInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValueSchema() FeatureEngineeringKafkaConfigValueSchemaOutputReference
	ValueSchemaInput() interface{}
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
	PutAuthConfig(value *FeatureEngineeringKafkaConfigAuthConfig)
	PutKeySchema(value *FeatureEngineeringKafkaConfigKeySchema)
	PutSubscriptionMode(value *FeatureEngineeringKafkaConfigSubscriptionMode)
	PutValueSchema(value *FeatureEngineeringKafkaConfigValueSchema)
	ResetExtraOptions()
	ResetKeySchema()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetValueSchema()
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

// The jsii proxy struct for FeatureEngineeringKafkaConfig
type jsiiProxy_FeatureEngineeringKafkaConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) AuthConfig() FeatureEngineeringKafkaConfigAuthConfigOutputReference {
	var returns FeatureEngineeringKafkaConfigAuthConfigOutputReference
	_jsii_.Get(
		j,
		"authConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) AuthConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) BootstrapServers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) BootstrapServersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) ExtraOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) ExtraOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) KeySchema() FeatureEngineeringKafkaConfigKeySchemaOutputReference {
	var returns FeatureEngineeringKafkaConfigKeySchemaOutputReference
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) KeySchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) SubscriptionMode() FeatureEngineeringKafkaConfigSubscriptionModeOutputReference {
	var returns FeatureEngineeringKafkaConfigSubscriptionModeOutputReference
	_jsii_.Get(
		j,
		"subscriptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) SubscriptionModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) ValueSchema() FeatureEngineeringKafkaConfigValueSchemaOutputReference {
	var returns FeatureEngineeringKafkaConfigValueSchemaOutputReference
	_jsii_.Get(
		j,
		"valueSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig) ValueSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueSchemaInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config databricks_feature_engineering_kafka_config} Resource.
func NewFeatureEngineeringKafkaConfig(scope constructs.Construct, id *string, config *FeatureEngineeringKafkaConfigConfig) FeatureEngineeringKafkaConfig {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringKafkaConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringKafkaConfig{}

	_jsii_.Create(
		"@cdktf/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config databricks_feature_engineering_kafka_config} Resource.
func NewFeatureEngineeringKafkaConfig_Override(f FeatureEngineeringKafkaConfig, scope constructs.Construct, id *string, config *FeatureEngineeringKafkaConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfig",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetBootstrapServers(val *string) {
	if err := j.validateSetBootstrapServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapServers",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetExtraOptions(val *map[string]*string) {
	if err := j.validateSetExtraOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraOptions",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a FeatureEngineeringKafkaConfig resource upon running "cdktf plan <stack-name>".
func FeatureEngineeringKafkaConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFeatureEngineeringKafkaConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfig",
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
func FeatureEngineeringKafkaConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringKafkaConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureEngineeringKafkaConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringKafkaConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureEngineeringKafkaConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringKafkaConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FeatureEngineeringKafkaConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) PutAuthConfig(value *FeatureEngineeringKafkaConfigAuthConfig) {
	if err := f.validatePutAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAuthConfig",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) PutKeySchema(value *FeatureEngineeringKafkaConfigKeySchema) {
	if err := f.validatePutKeySchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putKeySchema",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) PutSubscriptionMode(value *FeatureEngineeringKafkaConfigSubscriptionMode) {
	if err := f.validatePutSubscriptionModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSubscriptionMode",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) PutValueSchema(value *FeatureEngineeringKafkaConfigValueSchema) {
	if err := f.validatePutValueSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putValueSchema",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ResetExtraOptions() {
	_jsii_.InvokeVoid(
		f,
		"resetExtraOptions",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ResetKeySchema() {
	_jsii_.InvokeVoid(
		f,
		"resetKeySchema",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ResetValueSchema() {
	_jsii_.InvokeVoid(
		f,
		"resetValueSchema",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

