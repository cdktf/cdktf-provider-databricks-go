// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomassets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomAssetsAssetsOutputReference interface {
	cdktf.ComplexObject
	AddedAt() *float64
	AssetType() *string
	SetAssetType(val *string)
	AssetTypeInput() *string
	CleanRoomName() *string
	SetCleanRoomName(val *string)
	CleanRoomNameInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ForeignTable() DataDatabricksCleanRoomAssetsAssetsForeignTableOutputReference
	ForeignTableInput() interface{}
	ForeignTableLocalDetails() DataDatabricksCleanRoomAssetsAssetsForeignTableLocalDetailsOutputReference
	ForeignTableLocalDetailsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksCleanRoomAssetsAssets
	SetInternalValue(val *DataDatabricksCleanRoomAssetsAssets)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Notebook() DataDatabricksCleanRoomAssetsAssetsNotebookOutputReference
	NotebookInput() interface{}
	OwnerCollaboratorAlias() *string
	Status() *string
	Table() DataDatabricksCleanRoomAssetsAssetsTableOutputReference
	TableInput() interface{}
	TableLocalDetails() DataDatabricksCleanRoomAssetsAssetsTableLocalDetailsOutputReference
	TableLocalDetailsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	View() DataDatabricksCleanRoomAssetsAssetsViewOutputReference
	ViewInput() interface{}
	ViewLocalDetails() DataDatabricksCleanRoomAssetsAssetsViewLocalDetailsOutputReference
	ViewLocalDetailsInput() interface{}
	VolumeLocalDetails() DataDatabricksCleanRoomAssetsAssetsVolumeLocalDetailsOutputReference
	VolumeLocalDetailsInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutForeignTable(value *DataDatabricksCleanRoomAssetsAssetsForeignTable)
	PutForeignTableLocalDetails(value *DataDatabricksCleanRoomAssetsAssetsForeignTableLocalDetails)
	PutNotebook(value *DataDatabricksCleanRoomAssetsAssetsNotebook)
	PutTable(value *DataDatabricksCleanRoomAssetsAssetsTable)
	PutTableLocalDetails(value *DataDatabricksCleanRoomAssetsAssetsTableLocalDetails)
	PutView(value *DataDatabricksCleanRoomAssetsAssetsView)
	PutViewLocalDetails(value *DataDatabricksCleanRoomAssetsAssetsViewLocalDetails)
	PutVolumeLocalDetails(value *DataDatabricksCleanRoomAssetsAssetsVolumeLocalDetails)
	ResetCleanRoomName()
	ResetForeignTable()
	ResetForeignTableLocalDetails()
	ResetNotebook()
	ResetTable()
	ResetTableLocalDetails()
	ResetView()
	ResetViewLocalDetails()
	ResetVolumeLocalDetails()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksCleanRoomAssetsAssetsOutputReference
type jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) AddedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) AssetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) AssetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) CleanRoomName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) CleanRoomNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ForeignTable() DataDatabricksCleanRoomAssetsAssetsForeignTableOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsForeignTableOutputReference
	_jsii_.Get(
		j,
		"foreignTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ForeignTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ForeignTableLocalDetails() DataDatabricksCleanRoomAssetsAssetsForeignTableLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsForeignTableLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"foreignTableLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ForeignTableLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignTableLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) InternalValue() *DataDatabricksCleanRoomAssetsAssets {
	var returns *DataDatabricksCleanRoomAssetsAssets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) Notebook() DataDatabricksCleanRoomAssetsAssetsNotebookOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsNotebookOutputReference
	_jsii_.Get(
		j,
		"notebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) NotebookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) OwnerCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) Table() DataDatabricksCleanRoomAssetsAssetsTableOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) TableLocalDetails() DataDatabricksCleanRoomAssetsAssetsTableLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsTableLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"tableLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) TableLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) View() DataDatabricksCleanRoomAssetsAssetsViewOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsViewOutputReference
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ViewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ViewLocalDetails() DataDatabricksCleanRoomAssetsAssetsViewLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsViewLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"viewLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ViewLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) VolumeLocalDetails() DataDatabricksCleanRoomAssetsAssetsVolumeLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsVolumeLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"volumeLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) VolumeLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeLocalDetailsInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksCleanRoomAssetsAssetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksCleanRoomAssetsAssetsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomAssetsAssetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAssets.DataDatabricksCleanRoomAssetsAssetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksCleanRoomAssetsAssetsOutputReference_Override(d DataDatabricksCleanRoomAssetsAssetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAssets.DataDatabricksCleanRoomAssetsAssetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference)SetAssetType(val *string) {
	if err := j.validateSetAssetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference)SetCleanRoomName(val *string) {
	if err := j.validateSetCleanRoomNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanRoomName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference)SetInternalValue(val *DataDatabricksCleanRoomAssetsAssets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) PutForeignTable(value *DataDatabricksCleanRoomAssetsAssetsForeignTable) {
	if err := d.validatePutForeignTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForeignTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) PutForeignTableLocalDetails(value *DataDatabricksCleanRoomAssetsAssetsForeignTableLocalDetails) {
	if err := d.validatePutForeignTableLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForeignTableLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) PutNotebook(value *DataDatabricksCleanRoomAssetsAssetsNotebook) {
	if err := d.validatePutNotebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotebook",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) PutTable(value *DataDatabricksCleanRoomAssetsAssetsTable) {
	if err := d.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) PutTableLocalDetails(value *DataDatabricksCleanRoomAssetsAssetsTableLocalDetails) {
	if err := d.validatePutTableLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTableLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) PutView(value *DataDatabricksCleanRoomAssetsAssetsView) {
	if err := d.validatePutViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putView",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) PutViewLocalDetails(value *DataDatabricksCleanRoomAssetsAssetsViewLocalDetails) {
	if err := d.validatePutViewLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putViewLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) PutVolumeLocalDetails(value *DataDatabricksCleanRoomAssetsAssetsVolumeLocalDetails) {
	if err := d.validatePutVolumeLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVolumeLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetCleanRoomName() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanRoomName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetForeignTable() {
	_jsii_.InvokeVoid(
		d,
		"resetForeignTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetForeignTableLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetForeignTableLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetNotebook() {
	_jsii_.InvokeVoid(
		d,
		"resetNotebook",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetTableLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetTableLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetView() {
	_jsii_.InvokeVoid(
		d,
		"resetView",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetViewLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetViewLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ResetVolumeLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

