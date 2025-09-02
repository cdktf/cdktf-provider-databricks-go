// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomasset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomasset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset databricks_clean_room_asset}.
type DataDatabricksCleanRoomAsset interface {
	cdktf.TerraformDataSource
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
	ForeignTable() DataDatabricksCleanRoomAssetForeignTableOutputReference
	ForeignTableInput() interface{}
	ForeignTableLocalDetails() DataDatabricksCleanRoomAssetForeignTableLocalDetailsOutputReference
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
	Notebook() DataDatabricksCleanRoomAssetNotebookOutputReference
	NotebookInput() interface{}
	OwnerCollaboratorAlias() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Status() *string
	Table() DataDatabricksCleanRoomAssetTableOutputReference
	TableInput() interface{}
	TableLocalDetails() DataDatabricksCleanRoomAssetTableLocalDetailsOutputReference
	TableLocalDetailsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	View() DataDatabricksCleanRoomAssetViewOutputReference
	ViewInput() interface{}
	ViewLocalDetails() DataDatabricksCleanRoomAssetViewLocalDetailsOutputReference
	ViewLocalDetailsInput() interface{}
	VolumeLocalDetails() DataDatabricksCleanRoomAssetVolumeLocalDetailsOutputReference
	VolumeLocalDetailsInput() interface{}
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	PutForeignTable(value *DataDatabricksCleanRoomAssetForeignTable)
	PutForeignTableLocalDetails(value *DataDatabricksCleanRoomAssetForeignTableLocalDetails)
	PutNotebook(value *DataDatabricksCleanRoomAssetNotebook)
	PutTable(value *DataDatabricksCleanRoomAssetTable)
	PutTableLocalDetails(value *DataDatabricksCleanRoomAssetTableLocalDetails)
	PutView(value *DataDatabricksCleanRoomAssetView)
	PutViewLocalDetails(value *DataDatabricksCleanRoomAssetViewLocalDetails)
	PutVolumeLocalDetails(value *DataDatabricksCleanRoomAssetVolumeLocalDetails)
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

// The jsii proxy struct for DataDatabricksCleanRoomAsset
type jsiiProxy_DataDatabricksCleanRoomAsset struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) AddedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) AssetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) AssetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) CleanRoomName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) CleanRoomNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ForeignTable() DataDatabricksCleanRoomAssetForeignTableOutputReference {
	var returns DataDatabricksCleanRoomAssetForeignTableOutputReference
	_jsii_.Get(
		j,
		"foreignTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ForeignTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ForeignTableLocalDetails() DataDatabricksCleanRoomAssetForeignTableLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetForeignTableLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"foreignTableLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ForeignTableLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignTableLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Notebook() DataDatabricksCleanRoomAssetNotebookOutputReference {
	var returns DataDatabricksCleanRoomAssetNotebookOutputReference
	_jsii_.Get(
		j,
		"notebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) NotebookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) OwnerCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) Table() DataDatabricksCleanRoomAssetTableOutputReference {
	var returns DataDatabricksCleanRoomAssetTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) TableLocalDetails() DataDatabricksCleanRoomAssetTableLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetTableLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"tableLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) TableLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) View() DataDatabricksCleanRoomAssetViewOutputReference {
	var returns DataDatabricksCleanRoomAssetViewOutputReference
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ViewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ViewLocalDetails() DataDatabricksCleanRoomAssetViewLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetViewLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"viewLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) ViewLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) VolumeLocalDetails() DataDatabricksCleanRoomAssetVolumeLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetVolumeLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"volumeLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) VolumeLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset databricks_clean_room_asset} Data Source.
func NewDataDatabricksCleanRoomAsset(scope constructs.Construct, id *string, config *DataDatabricksCleanRoomAssetConfig) DataDatabricksCleanRoomAsset {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomAssetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomAsset{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAsset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset databricks_clean_room_asset} Data Source.
func NewDataDatabricksCleanRoomAsset_Override(d DataDatabricksCleanRoomAsset, scope constructs.Construct, id *string, config *DataDatabricksCleanRoomAssetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAsset",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetAssetType(val *string) {
	if err := j.validateSetAssetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetCleanRoomName(val *string) {
	if err := j.validateSetCleanRoomNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanRoomName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAsset)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksCleanRoomAsset resource upon running "cdktf plan <stack-name>".
func DataDatabricksCleanRoomAsset_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksCleanRoomAsset_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAsset",
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
func DataDatabricksCleanRoomAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCleanRoomAsset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksCleanRoomAsset_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCleanRoomAsset_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAsset",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksCleanRoomAsset_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCleanRoomAsset_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAsset",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksCleanRoomAsset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAsset",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) PutForeignTable(value *DataDatabricksCleanRoomAssetForeignTable) {
	if err := d.validatePutForeignTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForeignTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) PutForeignTableLocalDetails(value *DataDatabricksCleanRoomAssetForeignTableLocalDetails) {
	if err := d.validatePutForeignTableLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForeignTableLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) PutNotebook(value *DataDatabricksCleanRoomAssetNotebook) {
	if err := d.validatePutNotebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotebook",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) PutTable(value *DataDatabricksCleanRoomAssetTable) {
	if err := d.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) PutTableLocalDetails(value *DataDatabricksCleanRoomAssetTableLocalDetails) {
	if err := d.validatePutTableLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTableLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) PutView(value *DataDatabricksCleanRoomAssetView) {
	if err := d.validatePutViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putView",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) PutViewLocalDetails(value *DataDatabricksCleanRoomAssetViewLocalDetails) {
	if err := d.validatePutViewLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putViewLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) PutVolumeLocalDetails(value *DataDatabricksCleanRoomAssetVolumeLocalDetails) {
	if err := d.validatePutVolumeLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVolumeLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetCleanRoomName() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanRoomName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetForeignTable() {
	_jsii_.InvokeVoid(
		d,
		"resetForeignTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetForeignTableLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetForeignTableLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetNotebook() {
	_jsii_.InvokeVoid(
		d,
		"resetNotebook",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetTableLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetTableLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetView() {
	_jsii_.InvokeVoid(
		d,
		"resetView",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetViewLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetViewLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetVolumeLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAsset) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

