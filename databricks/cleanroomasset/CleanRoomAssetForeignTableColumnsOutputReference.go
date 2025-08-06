// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/cleanroomasset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanRoomAssetForeignTableColumnsOutputReference interface {
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
	InternalValue() *CleanRoomAssetForeignTableColumns
	SetInternalValue(val *CleanRoomAssetForeignTableColumns)
	Mask() CleanRoomAssetForeignTableColumnsMaskOutputReference
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
	PutMask(value *CleanRoomAssetForeignTableColumnsMask)
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

// The jsii proxy struct for CleanRoomAssetForeignTableColumnsOutputReference
type jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) InternalValue() *CleanRoomAssetForeignTableColumns {
	var returns *CleanRoomAssetForeignTableColumns
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) Mask() CleanRoomAssetForeignTableColumnsMaskOutputReference {
	var returns CleanRoomAssetForeignTableColumnsMaskOutputReference
	_jsii_.Get(
		j,
		"mask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) MaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) Nullable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) NullableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) PartitionIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) PartitionIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) Position() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) PositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"positionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeIntervalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeIntervalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeIntervalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeIntervalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypePrecision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typePrecision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypePrecisionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typePrecisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeScale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeScaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) TypeTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeTextInput",
		&returns,
	)
	return returns
}


func NewCleanRoomAssetForeignTableColumnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CleanRoomAssetForeignTableColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewCleanRoomAssetForeignTableColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAssetForeignTableColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCleanRoomAssetForeignTableColumnsOutputReference_Override(c CleanRoomAssetForeignTableColumnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAssetForeignTableColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetInternalValue(val *CleanRoomAssetForeignTableColumns) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetNullable(val interface{}) {
	if err := j.validateSetNullableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullable",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetPartitionIndex(val *float64) {
	if err := j.validateSetPartitionIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionIndex",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetPosition(val *float64) {
	if err := j.validateSetPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"position",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetTypeIntervalType(val *string) {
	if err := j.validateSetTypeIntervalTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeIntervalType",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetTypeJson(val *string) {
	if err := j.validateSetTypeJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeJson",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetTypePrecision(val *float64) {
	if err := j.validateSetTypePrecisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typePrecision",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetTypeScale(val *float64) {
	if err := j.validateSetTypeScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeScale",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference)SetTypeText(val *string) {
	if err := j.validateSetTypeTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeText",
		val,
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) PutMask(value *CleanRoomAssetForeignTableColumnsMask) {
	if err := c.validatePutMaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMask",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetMask() {
	_jsii_.InvokeVoid(
		c,
		"resetMask",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetNullable() {
	_jsii_.InvokeVoid(
		c,
		"resetNullable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetPartitionIndex() {
	_jsii_.InvokeVoid(
		c,
		"resetPartitionIndex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetPosition() {
	_jsii_.InvokeVoid(
		c,
		"resetPosition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetTypeIntervalType() {
	_jsii_.InvokeVoid(
		c,
		"resetTypeIntervalType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetTypeJson() {
	_jsii_.InvokeVoid(
		c,
		"resetTypeJson",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetTypeName() {
	_jsii_.InvokeVoid(
		c,
		"resetTypeName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetTypePrecision() {
	_jsii_.InvokeVoid(
		c,
		"resetTypePrecision",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetTypeScale() {
	_jsii_.InvokeVoid(
		c,
		"resetTypeScale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ResetTypeText() {
	_jsii_.InvokeVoid(
		c,
		"resetTypeText",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAssetForeignTableColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

