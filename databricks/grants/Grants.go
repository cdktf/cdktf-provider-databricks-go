// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grants

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/grants/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/grants databricks_grants}.
type Grants interface {
	cdktf.TerraformResource
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
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
	Credential() *string
	SetCredential(val *string)
	CredentialInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExternalLocation() *string
	SetExternalLocation(val *string)
	ExternalLocationInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForeignConnection() *string
	SetForeignConnection(val *string)
	ForeignConnectionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Function() *string
	SetFunction(val *string)
	FunctionInput() *string
	Grant() GrantsGrantList
	GrantInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metastore() *string
	SetMetastore(val *string)
	MetastoreInput() *string
	Model() *string
	SetModel(val *string)
	ModelInput() *string
	// The tree node.
	Node() constructs.Node
	Pipeline() *string
	SetPipeline(val *string)
	PipelineInput() *string
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
	Recipient() *string
	SetRecipient(val *string)
	RecipientInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Share() *string
	SetShare(val *string)
	ShareInput() *string
	StorageCredential() *string
	SetStorageCredential(val *string)
	StorageCredentialInput() *string
	Table() *string
	SetTable(val *string)
	TableInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Volume() *string
	SetVolume(val *string)
	VolumeInput() *string
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
	PutGrant(value interface{})
	ResetCatalog()
	ResetCredential()
	ResetExternalLocation()
	ResetForeignConnection()
	ResetFunction()
	ResetId()
	ResetMetastore()
	ResetModel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPipeline()
	ResetRecipient()
	ResetSchema()
	ResetShare()
	ResetStorageCredential()
	ResetTable()
	ResetVolume()
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

// The jsii proxy struct for Grants
type jsiiProxy_Grants struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Grants) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Credential() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) CredentialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) ExternalLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) ExternalLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) ForeignConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foreignConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) ForeignConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foreignConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Function() *string {
	var returns *string
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) FunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Grant() GrantsGrantList {
	var returns GrantsGrantList
	_jsii_.Get(
		j,
		"grant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) GrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Metastore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) MetastoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Model() *string {
	var returns *string
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) ModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Pipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) PipelineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Recipient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) RecipientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Share() *string {
	var returns *string
	_jsii_.Get(
		j,
		"share",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) ShareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) StorageCredential() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) StorageCredentialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) Volume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grants) VolumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/grants databricks_grants} Resource.
func NewGrants(scope constructs.Construct, id *string, config *GrantsConfig) Grants {
	_init_.Initialize()

	if err := validateNewGrantsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Grants{}

	_jsii_.Create(
		"@cdktf/provider-databricks.grants.Grants",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/grants databricks_grants} Resource.
func NewGrants_Override(g Grants, scope constructs.Construct, id *string, config *GrantsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.grants.Grants",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_Grants)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_Grants)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Grants)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Grants)SetCredential(val *string) {
	if err := j.validateSetCredentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credential",
		val,
	)
}

func (j *jsiiProxy_Grants)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Grants)SetExternalLocation(val *string) {
	if err := j.validateSetExternalLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalLocation",
		val,
	)
}

func (j *jsiiProxy_Grants)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Grants)SetForeignConnection(val *string) {
	if err := j.validateSetForeignConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foreignConnection",
		val,
	)
}

func (j *jsiiProxy_Grants)SetFunction(val *string) {
	if err := j.validateSetFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"function",
		val,
	)
}

func (j *jsiiProxy_Grants)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Grants)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Grants)SetMetastore(val *string) {
	if err := j.validateSetMetastoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastore",
		val,
	)
}

func (j *jsiiProxy_Grants)SetModel(val *string) {
	if err := j.validateSetModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"model",
		val,
	)
}

func (j *jsiiProxy_Grants)SetPipeline(val *string) {
	if err := j.validateSetPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipeline",
		val,
	)
}

func (j *jsiiProxy_Grants)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Grants)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Grants)SetRecipient(val *string) {
	if err := j.validateSetRecipientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipient",
		val,
	)
}

func (j *jsiiProxy_Grants)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_Grants)SetShare(val *string) {
	if err := j.validateSetShareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"share",
		val,
	)
}

func (j *jsiiProxy_Grants)SetStorageCredential(val *string) {
	if err := j.validateSetStorageCredentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCredential",
		val,
	)
}

func (j *jsiiProxy_Grants)SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_Grants)SetVolume(val *string) {
	if err := j.validateSetVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volume",
		val,
	)
}

// Generates CDKTF code for importing a Grants resource upon running "cdktf plan <stack-name>".
func Grants_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGrants_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.grants.Grants",
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
func Grants_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrants_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.grants.Grants",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Grants_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrants_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.grants.Grants",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Grants_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrants_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.grants.Grants",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Grants_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.grants.Grants",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_Grants) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_Grants) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_Grants) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_Grants) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_Grants) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_Grants) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_Grants) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_Grants) PutGrant(value interface{}) {
	if err := g.validatePutGrantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrant",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Grants) ResetCatalog() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetCredential() {
	_jsii_.InvokeVoid(
		g,
		"resetCredential",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetExternalLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetForeignConnection() {
	_jsii_.InvokeVoid(
		g,
		"resetForeignConnection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetFunction() {
	_jsii_.InvokeVoid(
		g,
		"resetFunction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetMetastore() {
	_jsii_.InvokeVoid(
		g,
		"resetMetastore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetModel() {
	_jsii_.InvokeVoid(
		g,
		"resetModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetPipeline() {
	_jsii_.InvokeVoid(
		g,
		"resetPipeline",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetRecipient() {
	_jsii_.InvokeVoid(
		g,
		"resetRecipient",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetShare() {
	_jsii_.InvokeVoid(
		g,
		"resetShare",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetStorageCredential() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageCredential",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetTable() {
	_jsii_.InvokeVoid(
		g,
		"resetTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) ResetVolume() {
	_jsii_.InvokeVoid(
		g,
		"resetVolume",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grants) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grants) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

