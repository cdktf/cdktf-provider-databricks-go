// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/databricks/r/azure_adls_gen1_mount databricks_azure_adls_gen1_mount}.
type AzureAdlsGen1Mount interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecretKey() *string
	SetClientSecretKey(val *string)
	ClientSecretKeyInput() *string
	ClientSecretScope() *string
	SetClientSecretScope(val *string)
	ClientSecretScopeInput() *string
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	Directory() *string
	SetDirectory(val *string)
	DirectoryInput() *string
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
	MountName() *string
	SetMountName(val *string)
	MountNameInput() *string
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
	Source() *string
	SparkConfPrefix() *string
	SetSparkConfPrefix(val *string)
	SparkConfPrefixInput() *string
	StorageResourceName() *string
	SetStorageResourceName(val *string)
	StorageResourceNameInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetClusterId()
	ResetDirectory()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSparkConfPrefix()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AzureAdlsGen1Mount
type jsiiProxy_AzureAdlsGen1Mount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AzureAdlsGen1Mount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ClientSecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ClientSecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ClientSecretScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ClientSecretScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Directory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) DirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) MountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SparkConfPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkConfPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SparkConfPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkConfPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) StorageResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) StorageResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAdlsGen1Mount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/databricks/r/azure_adls_gen1_mount databricks_azure_adls_gen1_mount} Resource.
func NewAzureAdlsGen1Mount(scope constructs.Construct, id *string, config *AzureAdlsGen1MountConfig) AzureAdlsGen1Mount {
	_init_.Initialize()

	j := jsiiProxy_AzureAdlsGen1Mount{}

	_jsii_.Create(
		"@cdktf/provider-databricks.AzureAdlsGen1Mount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/databricks/r/azure_adls_gen1_mount databricks_azure_adls_gen1_mount} Resource.
func NewAzureAdlsGen1Mount_Override(a AzureAdlsGen1Mount, scope constructs.Construct, id *string, config *AzureAdlsGen1MountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.AzureAdlsGen1Mount",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetClientSecretKey(val *string) {
	_jsii_.Set(
		j,
		"clientSecretKey",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetClientSecretScope(val *string) {
	_jsii_.Set(
		j,
		"clientSecretScope",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetDirectory(val *string) {
	_jsii_.Set(
		j,
		"directory",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetMountName(val *string) {
	_jsii_.Set(
		j,
		"mountName",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetSparkConfPrefix(val *string) {
	_jsii_.Set(
		j,
		"sparkConfPrefix",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetStorageResourceName(val *string) {
	_jsii_.Set(
		j,
		"storageResourceName",
		val,
	)
}

func (j *jsiiProxy_AzureAdlsGen1Mount) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
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
func AzureAdlsGen1Mount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.AzureAdlsGen1Mount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AzureAdlsGen1Mount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.AzureAdlsGen1Mount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzureAdlsGen1Mount) ResetClusterId() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAdlsGen1Mount) ResetDirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAdlsGen1Mount) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAdlsGen1Mount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAdlsGen1Mount) ResetSparkConfPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetSparkConfPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAdlsGen1Mount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAdlsGen1Mount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
