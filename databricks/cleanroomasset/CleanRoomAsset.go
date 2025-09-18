// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/cleanroomasset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/clean_room_asset databricks_clean_room_asset}.
type CleanRoomAsset interface {
	cdktf.TerraformResource
	AddedAt() *float64
	AssetType() *string
	SetAssetType(val *string)
	AssetTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CleanRoomName() *string
	SetCleanRoomName(val *string)
	CleanRoomNameInput() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForeignTable() CleanRoomAssetForeignTableOutputReference
	ForeignTableInput() interface{}
	ForeignTableLocalDetails() CleanRoomAssetForeignTableLocalDetailsOutputReference
	ForeignTableLocalDetailsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Notebook() CleanRoomAssetNotebookOutputReference
	NotebookInput() interface{}
	OwnerCollaboratorAlias() *string
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
	Status() *string
	Table() CleanRoomAssetTableOutputReference
	TableInput() interface{}
	TableLocalDetails() CleanRoomAssetTableLocalDetailsOutputReference
	TableLocalDetailsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	View() CleanRoomAssetViewOutputReference
	ViewInput() interface{}
	ViewLocalDetails() CleanRoomAssetViewLocalDetailsOutputReference
	ViewLocalDetailsInput() interface{}
	VolumeLocalDetails() CleanRoomAssetVolumeLocalDetailsOutputReference
	VolumeLocalDetailsInput() interface{}
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	PutForeignTable(value *CleanRoomAssetForeignTable)
	PutForeignTableLocalDetails(value *CleanRoomAssetForeignTableLocalDetails)
	PutNotebook(value *CleanRoomAssetNotebook)
	PutTable(value *CleanRoomAssetTable)
	PutTableLocalDetails(value *CleanRoomAssetTableLocalDetails)
	PutView(value *CleanRoomAssetView)
	PutViewLocalDetails(value *CleanRoomAssetViewLocalDetails)
	PutVolumeLocalDetails(value *CleanRoomAssetVolumeLocalDetails)
	ResetCleanRoomName()
	ResetForeignTable()
	ResetForeignTableLocalDetails()
	ResetNotebook()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTable()
	ResetTableLocalDetails()
	ResetView()
	ResetViewLocalDetails()
	ResetVolumeLocalDetails()
	ResetWorkspaceId()
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

// The jsii proxy struct for CleanRoomAsset
type jsiiProxy_CleanRoomAsset struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CleanRoomAsset) AddedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) AssetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) AssetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) CleanRoomName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) CleanRoomNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ForeignTable() CleanRoomAssetForeignTableOutputReference {
	var returns CleanRoomAssetForeignTableOutputReference
	_jsii_.Get(
		j,
		"foreignTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ForeignTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ForeignTableLocalDetails() CleanRoomAssetForeignTableLocalDetailsOutputReference {
	var returns CleanRoomAssetForeignTableLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"foreignTableLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ForeignTableLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignTableLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Notebook() CleanRoomAssetNotebookOutputReference {
	var returns CleanRoomAssetNotebookOutputReference
	_jsii_.Get(
		j,
		"notebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) NotebookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) OwnerCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) Table() CleanRoomAssetTableOutputReference {
	var returns CleanRoomAssetTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) TableLocalDetails() CleanRoomAssetTableLocalDetailsOutputReference {
	var returns CleanRoomAssetTableLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"tableLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) TableLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) View() CleanRoomAssetViewOutputReference {
	var returns CleanRoomAssetViewOutputReference
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ViewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ViewLocalDetails() CleanRoomAssetViewLocalDetailsOutputReference {
	var returns CleanRoomAssetViewLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"viewLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) ViewLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) VolumeLocalDetails() CleanRoomAssetVolumeLocalDetailsOutputReference {
	var returns CleanRoomAssetVolumeLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"volumeLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) VolumeLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAsset) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/clean_room_asset databricks_clean_room_asset} Resource.
func NewCleanRoomAsset(scope constructs.Construct, id *string, config *CleanRoomAssetConfig) CleanRoomAsset {
	_init_.Initialize()

	if err := validateNewCleanRoomAssetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanRoomAsset{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAsset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/clean_room_asset databricks_clean_room_asset} Resource.
func NewCleanRoomAsset_Override(c CleanRoomAsset, scope constructs.Construct, id *string, config *CleanRoomAssetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAsset",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetAssetType(val *string) {
	if err := j.validateSetAssetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetType",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetCleanRoomName(val *string) {
	if err := j.validateSetCleanRoomNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanRoomName",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAsset)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTF code for importing a CleanRoomAsset resource upon running "cdktf plan <stack-name>".
func CleanRoomAsset_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCleanRoomAsset_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAsset",
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
func CleanRoomAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCleanRoomAsset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CleanRoomAsset_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCleanRoomAsset_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAsset",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CleanRoomAsset_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCleanRoomAsset_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAsset",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CleanRoomAsset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAsset",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CleanRoomAsset) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CleanRoomAsset) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CleanRoomAsset) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CleanRoomAsset) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CleanRoomAsset) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CleanRoomAsset) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CleanRoomAsset) PutForeignTable(value *CleanRoomAssetForeignTable) {
	if err := c.validatePutForeignTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putForeignTable",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) PutForeignTableLocalDetails(value *CleanRoomAssetForeignTableLocalDetails) {
	if err := c.validatePutForeignTableLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putForeignTableLocalDetails",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) PutNotebook(value *CleanRoomAssetNotebook) {
	if err := c.validatePutNotebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNotebook",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) PutTable(value *CleanRoomAssetTable) {
	if err := c.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTable",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) PutTableLocalDetails(value *CleanRoomAssetTableLocalDetails) {
	if err := c.validatePutTableLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTableLocalDetails",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) PutView(value *CleanRoomAssetView) {
	if err := c.validatePutViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putView",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) PutViewLocalDetails(value *CleanRoomAssetViewLocalDetails) {
	if err := c.validatePutViewLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putViewLocalDetails",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) PutVolumeLocalDetails(value *CleanRoomAssetVolumeLocalDetails) {
	if err := c.validatePutVolumeLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolumeLocalDetails",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetCleanRoomName() {
	_jsii_.InvokeVoid(
		c,
		"resetCleanRoomName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetForeignTable() {
	_jsii_.InvokeVoid(
		c,
		"resetForeignTable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetForeignTableLocalDetails() {
	_jsii_.InvokeVoid(
		c,
		"resetForeignTableLocalDetails",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetNotebook() {
	_jsii_.InvokeVoid(
		c,
		"resetNotebook",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetTable() {
	_jsii_.InvokeVoid(
		c,
		"resetTable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetTableLocalDetails() {
	_jsii_.InvokeVoid(
		c,
		"resetTableLocalDetails",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetView() {
	_jsii_.InvokeVoid(
		c,
		"resetView",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetViewLocalDetails() {
	_jsii_.InvokeVoid(
		c,
		"resetViewLocalDetails",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetVolumeLocalDetails() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeLocalDetails",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAsset) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAsset) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

