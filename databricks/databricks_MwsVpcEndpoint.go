// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/databricks/r/mws_vpc_endpoint databricks_mws_vpc_endpoint}.
type MwsVpcEndpoint interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	AwsEndpointServiceId() *string
	SetAwsEndpointServiceId(val *string)
	AwsEndpointServiceIdInput() *string
	AwsVpcEndpointId() *string
	SetAwsVpcEndpointId(val *string)
	AwsVpcEndpointIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseCase() *string
	SetUseCase(val *string)
	UseCaseInput() *string
	VpcEndpointId() *string
	SetVpcEndpointId(val *string)
	VpcEndpointIdInput() *string
	VpcEndpointName() *string
	SetVpcEndpointName(val *string)
	VpcEndpointNameInput() *string
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
	ResetAwsAccountId()
	ResetAwsEndpointServiceId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetState()
	ResetUseCase()
	ResetVpcEndpointId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MwsVpcEndpoint
type jsiiProxy_MwsVpcEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MwsVpcEndpoint) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) AwsEndpointServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsEndpointServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) AwsEndpointServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsEndpointServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) AwsVpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsVpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) AwsVpcEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsVpcEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) UseCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) UseCaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) VpcEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) VpcEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpoint) VpcEndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/databricks/r/mws_vpc_endpoint databricks_mws_vpc_endpoint} Resource.
func NewMwsVpcEndpoint(scope constructs.Construct, id *string, config *MwsVpcEndpointConfig) MwsVpcEndpoint {
	_init_.Initialize()

	j := jsiiProxy_MwsVpcEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-databricks.MwsVpcEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/databricks/r/mws_vpc_endpoint databricks_mws_vpc_endpoint} Resource.
func NewMwsVpcEndpoint_Override(m MwsVpcEndpoint, scope constructs.Construct, id *string, config *MwsVpcEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.MwsVpcEndpoint",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetAwsEndpointServiceId(val *string) {
	_jsii_.Set(
		j,
		"awsEndpointServiceId",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetAwsVpcEndpointId(val *string) {
	_jsii_.Set(
		j,
		"awsVpcEndpointId",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetUseCase(val *string) {
	_jsii_.Set(
		j,
		"useCase",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetVpcEndpointId(val *string) {
	_jsii_.Set(
		j,
		"vpcEndpointId",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpoint) SetVpcEndpointName(val *string) {
	_jsii_.Set(
		j,
		"vpcEndpointName",
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
func MwsVpcEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.MwsVpcEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MwsVpcEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.MwsVpcEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) ResetAccountId() {
	_jsii_.InvokeVoid(
		m,
		"resetAccountId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) ResetAwsEndpointServiceId() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsEndpointServiceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) ResetState() {
	_jsii_.InvokeVoid(
		m,
		"resetState",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) ResetUseCase() {
	_jsii_.InvokeVoid(
		m,
		"resetUseCase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) ResetVpcEndpointId() {
	_jsii_.InvokeVoid(
		m,
		"resetVpcEndpointId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
