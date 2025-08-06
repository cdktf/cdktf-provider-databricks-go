// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassetrevisionscleanroomassets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomassetrevisionscleanroomassets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference interface {
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
	ForeignTable() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableOutputReference
	ForeignTableInput() interface{}
	ForeignTableLocalDetails() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableLocalDetailsOutputReference
	ForeignTableLocalDetailsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisions
	SetInternalValue(val *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisions)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Notebook() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsNotebookOutputReference
	NotebookInput() interface{}
	OwnerCollaboratorAlias() *string
	Status() *string
	Table() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableOutputReference
	TableInput() interface{}
	TableLocalDetails() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableLocalDetailsOutputReference
	TableLocalDetailsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	View() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewOutputReference
	ViewInput() interface{}
	ViewLocalDetails() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewLocalDetailsOutputReference
	ViewLocalDetailsInput() interface{}
	VolumeLocalDetails() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsVolumeLocalDetailsOutputReference
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
	PutForeignTable(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTable)
	PutForeignTableLocalDetails(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableLocalDetails)
	PutNotebook(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsNotebook)
	PutTable(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTable)
	PutTableLocalDetails(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableLocalDetails)
	PutView(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsView)
	PutViewLocalDetails(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewLocalDetails)
	PutVolumeLocalDetails(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsVolumeLocalDetails)
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

// The jsii proxy struct for DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference
type jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) AddedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) AssetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) AssetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) CleanRoomName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) CleanRoomNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ForeignTable() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableOutputReference
	_jsii_.Get(
		j,
		"foreignTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ForeignTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ForeignTableLocalDetails() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"foreignTableLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ForeignTableLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignTableLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) InternalValue() *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisions {
	var returns *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) Notebook() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsNotebookOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsNotebookOutputReference
	_jsii_.Get(
		j,
		"notebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) NotebookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) OwnerCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) Table() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) TableLocalDetails() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"tableLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) TableLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) View() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewOutputReference
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ViewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ViewLocalDetails() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"viewLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ViewLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewLocalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) VolumeLocalDetails() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsVolumeLocalDetailsOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsVolumeLocalDetailsOutputReference
	_jsii_.Get(
		j,
		"volumeLocalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) VolumeLocalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeLocalDetailsInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAssetRevisionsCleanRoomAssets.DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference_Override(d DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAssetRevisionsCleanRoomAssets.DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference)SetAssetType(val *string) {
	if err := j.validateSetAssetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference)SetCleanRoomName(val *string) {
	if err := j.validateSetCleanRoomNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanRoomName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference)SetInternalValue(val *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) PutForeignTable(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTable) {
	if err := d.validatePutForeignTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForeignTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) PutForeignTableLocalDetails(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableLocalDetails) {
	if err := d.validatePutForeignTableLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForeignTableLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) PutNotebook(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsNotebook) {
	if err := d.validatePutNotebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotebook",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) PutTable(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTable) {
	if err := d.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) PutTableLocalDetails(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableLocalDetails) {
	if err := d.validatePutTableLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTableLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) PutView(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsView) {
	if err := d.validatePutViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putView",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) PutViewLocalDetails(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewLocalDetails) {
	if err := d.validatePutViewLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putViewLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) PutVolumeLocalDetails(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsVolumeLocalDetails) {
	if err := d.validatePutVolumeLocalDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVolumeLocalDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetCleanRoomName() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanRoomName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetForeignTable() {
	_jsii_.InvokeVoid(
		d,
		"resetForeignTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetForeignTableLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetForeignTableLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetNotebook() {
	_jsii_.InvokeVoid(
		d,
		"resetNotebook",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetTableLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetTableLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetView() {
	_jsii_.InvokeVoid(
		d,
		"resetView",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetViewLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetViewLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ResetVolumeLocalDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeLocalDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

