package metastoredataaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v9/metastoredataaccess/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/metastore_data_access databricks_metastore_data_access}.
type MetastoreDataAccess interface {
	cdktf.TerraformResource
	AwsIamRole() MetastoreDataAccessAwsIamRoleOutputReference
	AwsIamRoleInput() *MetastoreDataAccessAwsIamRole
	AzureManagedIdentity() MetastoreDataAccessAzureManagedIdentityOutputReference
	AzureManagedIdentityInput() *MetastoreDataAccessAzureManagedIdentity
	AzureServicePrincipal() MetastoreDataAccessAzureServicePrincipalOutputReference
	AzureServicePrincipalInput() *MetastoreDataAccessAzureServicePrincipal
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigurationType() *string
	SetConfigurationType(val *string)
	ConfigurationTypeInput() *string
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
	DatabricksGcpServiceAccount() MetastoreDataAccessDatabricksGcpServiceAccountOutputReference
	DatabricksGcpServiceAccountInput() *MetastoreDataAccessDatabricksGcpServiceAccount
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
	GcpServiceAccountKey() MetastoreDataAccessGcpServiceAccountKeyOutputReference
	GcpServiceAccountKeyInput() *MetastoreDataAccessGcpServiceAccountKey
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsDefault() interface{}
	SetIsDefault(val interface{})
	IsDefaultInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
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
	PutAwsIamRole(value *MetastoreDataAccessAwsIamRole)
	PutAzureManagedIdentity(value *MetastoreDataAccessAzureManagedIdentity)
	PutAzureServicePrincipal(value *MetastoreDataAccessAzureServicePrincipal)
	PutDatabricksGcpServiceAccount(value *MetastoreDataAccessDatabricksGcpServiceAccount)
	PutGcpServiceAccountKey(value *MetastoreDataAccessGcpServiceAccountKey)
	ResetAwsIamRole()
	ResetAzureManagedIdentity()
	ResetAzureServicePrincipal()
	ResetConfigurationType()
	ResetDatabricksGcpServiceAccount()
	ResetGcpServiceAccountKey()
	ResetId()
	ResetIsDefault()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MetastoreDataAccess
type jsiiProxy_MetastoreDataAccess struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MetastoreDataAccess) AwsIamRole() MetastoreDataAccessAwsIamRoleOutputReference {
	var returns MetastoreDataAccessAwsIamRoleOutputReference
	_jsii_.Get(
		j,
		"awsIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) AwsIamRoleInput() *MetastoreDataAccessAwsIamRole {
	var returns *MetastoreDataAccessAwsIamRole
	_jsii_.Get(
		j,
		"awsIamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) AzureManagedIdentity() MetastoreDataAccessAzureManagedIdentityOutputReference {
	var returns MetastoreDataAccessAzureManagedIdentityOutputReference
	_jsii_.Get(
		j,
		"azureManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) AzureManagedIdentityInput() *MetastoreDataAccessAzureManagedIdentity {
	var returns *MetastoreDataAccessAzureManagedIdentity
	_jsii_.Get(
		j,
		"azureManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) AzureServicePrincipal() MetastoreDataAccessAzureServicePrincipalOutputReference {
	var returns MetastoreDataAccessAzureServicePrincipalOutputReference
	_jsii_.Get(
		j,
		"azureServicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) AzureServicePrincipalInput() *MetastoreDataAccessAzureServicePrincipal {
	var returns *MetastoreDataAccessAzureServicePrincipal
	_jsii_.Get(
		j,
		"azureServicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) ConfigurationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) ConfigurationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) DatabricksGcpServiceAccount() MetastoreDataAccessDatabricksGcpServiceAccountOutputReference {
	var returns MetastoreDataAccessDatabricksGcpServiceAccountOutputReference
	_jsii_.Get(
		j,
		"databricksGcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) DatabricksGcpServiceAccountInput() *MetastoreDataAccessDatabricksGcpServiceAccount {
	var returns *MetastoreDataAccessDatabricksGcpServiceAccount
	_jsii_.Get(
		j,
		"databricksGcpServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) GcpServiceAccountKey() MetastoreDataAccessGcpServiceAccountKeyOutputReference {
	var returns MetastoreDataAccessGcpServiceAccountKeyOutputReference
	_jsii_.Get(
		j,
		"gcpServiceAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) GcpServiceAccountKeyInput() *MetastoreDataAccessGcpServiceAccountKey {
	var returns *MetastoreDataAccessGcpServiceAccountKey
	_jsii_.Get(
		j,
		"gcpServiceAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) IsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) IsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccess) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/metastore_data_access databricks_metastore_data_access} Resource.
func NewMetastoreDataAccess(scope constructs.Construct, id *string, config *MetastoreDataAccessConfig) MetastoreDataAccess {
	_init_.Initialize()

	if err := validateNewMetastoreDataAccessParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MetastoreDataAccess{}

	_jsii_.Create(
		"@cdktf/provider-databricks.metastoreDataAccess.MetastoreDataAccess",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/metastore_data_access databricks_metastore_data_access} Resource.
func NewMetastoreDataAccess_Override(m MetastoreDataAccess, scope constructs.Construct, id *string, config *MetastoreDataAccessConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.metastoreDataAccess.MetastoreDataAccess",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetConfigurationType(val *string) {
	if err := j.validateSetConfigurationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationType",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetIsDefault(val interface{}) {
	if err := j.validateSetIsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDefault",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccess)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func MetastoreDataAccess_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetastoreDataAccess_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.metastoreDataAccess.MetastoreDataAccess",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MetastoreDataAccess_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetastoreDataAccess_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.metastoreDataAccess.MetastoreDataAccess",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MetastoreDataAccess_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetastoreDataAccess_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.metastoreDataAccess.MetastoreDataAccess",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MetastoreDataAccess_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.metastoreDataAccess.MetastoreDataAccess",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MetastoreDataAccess) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MetastoreDataAccess) PutAwsIamRole(value *MetastoreDataAccessAwsIamRole) {
	if err := m.validatePutAwsIamRoleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAwsIamRole",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MetastoreDataAccess) PutAzureManagedIdentity(value *MetastoreDataAccessAzureManagedIdentity) {
	if err := m.validatePutAzureManagedIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAzureManagedIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MetastoreDataAccess) PutAzureServicePrincipal(value *MetastoreDataAccessAzureServicePrincipal) {
	if err := m.validatePutAzureServicePrincipalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAzureServicePrincipal",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MetastoreDataAccess) PutDatabricksGcpServiceAccount(value *MetastoreDataAccessDatabricksGcpServiceAccount) {
	if err := m.validatePutDatabricksGcpServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDatabricksGcpServiceAccount",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MetastoreDataAccess) PutGcpServiceAccountKey(value *MetastoreDataAccessGcpServiceAccountKey) {
	if err := m.validatePutGcpServiceAccountKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGcpServiceAccountKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetAwsIamRole() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsIamRole",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetAzureManagedIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetAzureManagedIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetAzureServicePrincipal() {
	_jsii_.InvokeVoid(
		m,
		"resetAzureServicePrincipal",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetConfigurationType() {
	_jsii_.InvokeVoid(
		m,
		"resetConfigurationType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetDatabricksGcpServiceAccount() {
	_jsii_.InvokeVoid(
		m,
		"resetDatabricksGcpServiceAccount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetGcpServiceAccountKey() {
	_jsii_.InvokeVoid(
		m,
		"resetGcpServiceAccountKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetIsDefault() {
	_jsii_.InvokeVoid(
		m,
		"resetIsDefault",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccess) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccess) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

