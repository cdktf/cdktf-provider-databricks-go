package storagecredential

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/storagecredential/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.19.0/docs/resources/storage_credential databricks_storage_credential}.
type StorageCredential interface {
	cdktf.TerraformResource
	AwsIamRole() StorageCredentialAwsIamRoleOutputReference
	AwsIamRoleInput() *StorageCredentialAwsIamRole
	AzureManagedIdentity() StorageCredentialAzureManagedIdentityOutputReference
	AzureManagedIdentityInput() *StorageCredentialAzureManagedIdentity
	AzureServicePrincipal() StorageCredentialAzureServicePrincipalOutputReference
	AzureServicePrincipalInput() *StorageCredentialAzureServicePrincipal
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	DatabricksGcpServiceAccount() StorageCredentialDatabricksGcpServiceAccountOutputReference
	DatabricksGcpServiceAccountInput() *StorageCredentialDatabricksGcpServiceAccount
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
	GcpServiceAccountKey() StorageCredentialGcpServiceAccountKeyOutputReference
	GcpServiceAccountKeyInput() *StorageCredentialGcpServiceAccountKey
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
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
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
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
	PutAwsIamRole(value *StorageCredentialAwsIamRole)
	PutAzureManagedIdentity(value *StorageCredentialAzureManagedIdentity)
	PutAzureServicePrincipal(value *StorageCredentialAzureServicePrincipal)
	PutDatabricksGcpServiceAccount(value *StorageCredentialDatabricksGcpServiceAccount)
	PutGcpServiceAccountKey(value *StorageCredentialGcpServiceAccountKey)
	ResetAwsIamRole()
	ResetAzureManagedIdentity()
	ResetAzureServicePrincipal()
	ResetComment()
	ResetDatabricksGcpServiceAccount()
	ResetGcpServiceAccountKey()
	ResetId()
	ResetMetastoreId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetReadOnly()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StorageCredential
type jsiiProxy_StorageCredential struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageCredential) AwsIamRole() StorageCredentialAwsIamRoleOutputReference {
	var returns StorageCredentialAwsIamRoleOutputReference
	_jsii_.Get(
		j,
		"awsIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) AwsIamRoleInput() *StorageCredentialAwsIamRole {
	var returns *StorageCredentialAwsIamRole
	_jsii_.Get(
		j,
		"awsIamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) AzureManagedIdentity() StorageCredentialAzureManagedIdentityOutputReference {
	var returns StorageCredentialAzureManagedIdentityOutputReference
	_jsii_.Get(
		j,
		"azureManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) AzureManagedIdentityInput() *StorageCredentialAzureManagedIdentity {
	var returns *StorageCredentialAzureManagedIdentity
	_jsii_.Get(
		j,
		"azureManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) AzureServicePrincipal() StorageCredentialAzureServicePrincipalOutputReference {
	var returns StorageCredentialAzureServicePrincipalOutputReference
	_jsii_.Get(
		j,
		"azureServicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) AzureServicePrincipalInput() *StorageCredentialAzureServicePrincipal {
	var returns *StorageCredentialAzureServicePrincipal
	_jsii_.Get(
		j,
		"azureServicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) DatabricksGcpServiceAccount() StorageCredentialDatabricksGcpServiceAccountOutputReference {
	var returns StorageCredentialDatabricksGcpServiceAccountOutputReference
	_jsii_.Get(
		j,
		"databricksGcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) DatabricksGcpServiceAccountInput() *StorageCredentialDatabricksGcpServiceAccount {
	var returns *StorageCredentialDatabricksGcpServiceAccount
	_jsii_.Get(
		j,
		"databricksGcpServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) GcpServiceAccountKey() StorageCredentialGcpServiceAccountKeyOutputReference {
	var returns StorageCredentialGcpServiceAccountKeyOutputReference
	_jsii_.Get(
		j,
		"gcpServiceAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) GcpServiceAccountKeyInput() *StorageCredentialGcpServiceAccountKey {
	var returns *StorageCredentialGcpServiceAccountKey
	_jsii_.Get(
		j,
		"gcpServiceAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageCredential) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.19.0/docs/resources/storage_credential databricks_storage_credential} Resource.
func NewStorageCredential(scope constructs.Construct, id *string, config *StorageCredentialConfig) StorageCredential {
	_init_.Initialize()

	if err := validateNewStorageCredentialParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageCredential{}

	_jsii_.Create(
		"@cdktf/provider-databricks.storageCredential.StorageCredential",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.19.0/docs/resources/storage_credential databricks_storage_credential} Resource.
func NewStorageCredential_Override(s StorageCredential, scope constructs.Construct, id *string, config *StorageCredentialConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.storageCredential.StorageCredential",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageCredential)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageCredential)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
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
func StorageCredential_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageCredential_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.storageCredential.StorageCredential",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageCredential_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageCredential_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.storageCredential.StorageCredential",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageCredential_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageCredential_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.storageCredential.StorageCredential",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageCredential_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.storageCredential.StorageCredential",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageCredential) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageCredential) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageCredential) PutAwsIamRole(value *StorageCredentialAwsIamRole) {
	if err := s.validatePutAwsIamRoleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsIamRole",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageCredential) PutAzureManagedIdentity(value *StorageCredentialAzureManagedIdentity) {
	if err := s.validatePutAzureManagedIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureManagedIdentity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageCredential) PutAzureServicePrincipal(value *StorageCredentialAzureServicePrincipal) {
	if err := s.validatePutAzureServicePrincipalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureServicePrincipal",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageCredential) PutDatabricksGcpServiceAccount(value *StorageCredentialDatabricksGcpServiceAccount) {
	if err := s.validatePutDatabricksGcpServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatabricksGcpServiceAccount",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageCredential) PutGcpServiceAccountKey(value *StorageCredentialGcpServiceAccountKey) {
	if err := s.validatePutGcpServiceAccountKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcpServiceAccountKey",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageCredential) ResetAwsIamRole() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsIamRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetAzureManagedIdentity() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureManagedIdentity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetAzureServicePrincipal() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureServicePrincipal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetDatabricksGcpServiceAccount() {
	_jsii_.InvokeVoid(
		s,
		"resetDatabricksGcpServiceAccount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetGcpServiceAccountKey() {
	_jsii_.InvokeVoid(
		s,
		"resetGcpServiceAccountKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		s,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetOwner() {
	_jsii_.InvokeVoid(
		s,
		"resetOwner",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) ResetReadOnly() {
	_jsii_.InvokeVoid(
		s,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageCredential) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageCredential) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

