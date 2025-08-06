// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomassets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference interface {
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
	InternalValue() *DataDatabricksCleanRoomAssetsAssetsTableColumns
	SetInternalValue(val *DataDatabricksCleanRoomAssetsAssetsTableColumns)
	Mask() DataDatabricksCleanRoomAssetsAssetsTableColumnsMaskOutputReference
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
	PutMask(value *DataDatabricksCleanRoomAssetsAssetsTableColumnsMask)
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

// The jsii proxy struct for DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference
type jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) InternalValue() *DataDatabricksCleanRoomAssetsAssetsTableColumns {
	var returns *DataDatabricksCleanRoomAssetsAssetsTableColumns
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) Mask() DataDatabricksCleanRoomAssetsAssetsTableColumnsMaskOutputReference {
	var returns DataDatabricksCleanRoomAssetsAssetsTableColumnsMaskOutputReference
	_jsii_.Get(
		j,
		"mask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) MaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) Nullable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) NullableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) PartitionIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) PartitionIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) Position() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) PositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"positionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeIntervalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeIntervalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeIntervalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeIntervalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypePrecision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typePrecision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypePrecisionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typePrecisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeScale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeScaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) TypeTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeTextInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAssets.DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference_Override(d DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAssets.DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetInternalValue(val *DataDatabricksCleanRoomAssetsAssetsTableColumns) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetNullable(val interface{}) {
	if err := j.validateSetNullableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetPartitionIndex(val *float64) {
	if err := j.validateSetPartitionIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetPosition(val *float64) {
	if err := j.validateSetPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"position",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetTypeIntervalType(val *string) {
	if err := j.validateSetTypeIntervalTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeIntervalType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetTypeJson(val *string) {
	if err := j.validateSetTypeJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeJson",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetTypePrecision(val *float64) {
	if err := j.validateSetTypePrecisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typePrecision",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetTypeScale(val *float64) {
	if err := j.validateSetTypeScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeScale",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference)SetTypeText(val *string) {
	if err := j.validateSetTypeTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeText",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) PutMask(value *DataDatabricksCleanRoomAssetsAssetsTableColumnsMask) {
	if err := d.validatePutMaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMask",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetMask() {
	_jsii_.InvokeVoid(
		d,
		"resetMask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetNullable() {
	_jsii_.InvokeVoid(
		d,
		"resetNullable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetPartitionIndex() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionIndex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetPosition() {
	_jsii_.InvokeVoid(
		d,
		"resetPosition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetTypeIntervalType() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeIntervalType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetTypeJson() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetTypeName() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetTypePrecision() {
	_jsii_.InvokeVoid(
		d,
		"resetTypePrecision",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetTypeScale() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeScale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ResetTypeText() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetsAssetsTableColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

