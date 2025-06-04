// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssharepluginframework

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/datadatabrickssharepluginframework/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksSharePluginframeworkObjectOutputReference interface {
	cdktf.ComplexObject
	AddedAt() *float64
	AddedBy() *string
	CdfEnabled() interface{}
	SetCdfEnabled(val interface{})
	CdfEnabledInput() interface{}
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
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataObjectType() *string
	SetDataObjectType(val *string)
	DataObjectTypeInput() *string
	EffectiveCdfEnabled() cdktf.IResolvable
	EffectiveHistoryDataSharingStatus() *string
	EffectiveSharedAs() *string
	EffectiveStartVersion() *float64
	// Experimental.
	Fqn() *string
	HistoryDataSharingStatus() *string
	SetHistoryDataSharingStatus(val *string)
	HistoryDataSharingStatusInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Partition() DataDatabricksSharePluginframeworkObjectPartitionList
	PartitionInput() interface{}
	SharedAs() *string
	SetSharedAs(val *string)
	SharedAsInput() *string
	StartVersion() *float64
	SetStartVersion(val *float64)
	StartVersionInput() *float64
	Status() *string
	StringSharedAs() *string
	SetStringSharedAs(val *string)
	StringSharedAsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutPartition(value interface{})
	ResetCdfEnabled()
	ResetComment()
	ResetContent()
	ResetDataObjectType()
	ResetHistoryDataSharingStatus()
	ResetPartition()
	ResetSharedAs()
	ResetStartVersion()
	ResetStringSharedAs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksSharePluginframeworkObjectOutputReference
type jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) AddedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) AddedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) CdfEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdfEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) CdfEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdfEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) DataObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataObjectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) DataObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataObjectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) EffectiveCdfEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveCdfEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) EffectiveHistoryDataSharingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveHistoryDataSharingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) EffectiveSharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveSharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) EffectiveStartVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveStartVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) HistoryDataSharingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyDataSharingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) HistoryDataSharingStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyDataSharingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) Partition() DataDatabricksSharePluginframeworkObjectPartitionList {
	var returns DataDatabricksSharePluginframeworkObjectPartitionList
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) PartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) SharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) SharedAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) StartVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) StartVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) StringSharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringSharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) StringSharedAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringSharedAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksSharePluginframeworkObjectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksSharePluginframeworkObjectOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksSharePluginframeworkObjectOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSharePluginframework.DataDatabricksSharePluginframeworkObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksSharePluginframeworkObjectOutputReference_Override(d DataDatabricksSharePluginframeworkObjectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSharePluginframework.DataDatabricksSharePluginframeworkObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetCdfEnabled(val interface{}) {
	if err := j.validateSetCdfEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdfEnabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetDataObjectType(val *string) {
	if err := j.validateSetDataObjectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataObjectType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetHistoryDataSharingStatus(val *string) {
	if err := j.validateSetHistoryDataSharingStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"historyDataSharingStatus",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetSharedAs(val *string) {
	if err := j.validateSetSharedAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAs",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetStartVersion(val *float64) {
	if err := j.validateSetStartVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetStringSharedAs(val *string) {
	if err := j.validateSetStringSharedAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringSharedAs",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) PutPartition(value interface{}) {
	if err := d.validatePutPartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPartition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetCdfEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetCdfEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		d,
		"resetContent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetDataObjectType() {
	_jsii_.InvokeVoid(
		d,
		"resetDataObjectType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetHistoryDataSharingStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetHistoryDataSharingStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetPartition() {
	_jsii_.InvokeVoid(
		d,
		"resetPartition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetSharedAs() {
	_jsii_.InvokeVoid(
		d,
		"resetSharedAs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetStartVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetStartVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ResetStringSharedAs() {
	_jsii_.InvokeVoid(
		d,
		"resetStringSharedAs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksSharePluginframeworkObjectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

