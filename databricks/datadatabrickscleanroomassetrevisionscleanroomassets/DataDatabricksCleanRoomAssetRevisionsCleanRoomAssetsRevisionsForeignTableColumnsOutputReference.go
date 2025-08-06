// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassetrevisionscleanroomassets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomassetrevisionscleanroomassets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference interface {
	cdktf.ComplexObject
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumns
	SetInternalValue(val *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumns)
	Mask() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsMaskOutputReference
	MaskInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nullable() interface{}
	SetNullable(val interface{})
	NullableInput() interface{}
	PartitionIndex() *float64
	SetPartitionIndex(val *float64)
	PartitionIndexInput() *float64
	Position() *float64
	SetPosition(val *float64)
	PositionInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TypeIntervalType() *string
	SetTypeIntervalType(val *string)
	TypeIntervalTypeInput() *string
	TypeJson() *string
	SetTypeJson(val *string)
	TypeJsonInput() *string
	TypeName() *string
	SetTypeName(val *string)
	TypeNameInput() *string
	TypePrecision() *float64
	SetTypePrecision(val *float64)
	TypePrecisionInput() *float64
	TypeScale() *float64
	SetTypeScale(val *float64)
	TypeScaleInput() *float64
	TypeText() *string
	SetTypeText(val *string)
	TypeTextInput() *string
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
	PutMask(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsMask)
	ResetComment()
	ResetMask()
	ResetName()
	ResetNullable()
	ResetPartitionIndex()
	ResetPosition()
	ResetTypeIntervalType()
	ResetTypeJson()
	ResetTypeName()
	ResetTypePrecision()
	ResetTypeScale()
	ResetTypeText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference
type jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) InternalValue() *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumns {
	var returns *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumns
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) Mask() DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsMaskOutputReference {
	var returns DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsMaskOutputReference
	_jsii_.Get(
		j,
		"mask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) MaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) Nullable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) NullableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) PartitionIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) PartitionIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) Position() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) PositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"positionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeIntervalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeIntervalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeIntervalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeIntervalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypePrecision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typePrecision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypePrecisionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typePrecisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeScale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeScaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) TypeTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeTextInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAssetRevisionsCleanRoomAssets.DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference_Override(d DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAssetRevisionsCleanRoomAssets.DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetInternalValue(val *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumns) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetNullable(val interface{}) {
	if err := j.validateSetNullableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetPartitionIndex(val *float64) {
	if err := j.validateSetPartitionIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetPosition(val *float64) {
	if err := j.validateSetPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"position",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetTypeIntervalType(val *string) {
	if err := j.validateSetTypeIntervalTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeIntervalType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetTypeJson(val *string) {
	if err := j.validateSetTypeJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeJson",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetTypePrecision(val *float64) {
	if err := j.validateSetTypePrecisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typePrecision",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetTypeScale(val *float64) {
	if err := j.validateSetTypeScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeScale",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference)SetTypeText(val *string) {
	if err := j.validateSetTypeTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeText",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) PutMask(value *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsMask) {
	if err := d.validatePutMaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetMask() {
	_jsii_.InvokeVoid(
		d,
		"resetMask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetNullable() {
	_jsii_.InvokeVoid(
		d,
		"resetNullable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetPartitionIndex() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionIndex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetPosition() {
	_jsii_.InvokeVoid(
		d,
		"resetPosition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetTypeIntervalType() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeIntervalType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetTypeJson() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetTypeName() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetTypePrecision() {
	_jsii_.InvokeVoid(
		d,
		"resetTypePrecision",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetTypeScale() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeScale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ResetTypeText() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

