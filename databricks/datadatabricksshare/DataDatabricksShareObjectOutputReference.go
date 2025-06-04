// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksshare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksshare/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksShareObjectOutputReference interface {
	cdktf.ComplexObject
	AddedAt() *float64
	SetAddedAt(val *float64)
	AddedAtInput() *float64
	AddedBy() *string
	SetAddedBy(val *string)
	AddedByInput() *string
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
	Partition() DataDatabricksShareObjectPartitionList
	PartitionInput() interface{}
	SharedAs() *string
	SetSharedAs(val *string)
	SharedAsInput() *string
	StartVersion() *float64
	SetStartVersion(val *float64)
	StartVersionInput() *float64
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	ResetAddedAt()
	ResetAddedBy()
	ResetCdfEnabled()
	ResetComment()
	ResetContent()
	ResetHistoryDataSharingStatus()
	ResetPartition()
	ResetSharedAs()
	ResetStartVersion()
	ResetStatus()
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

// The jsii proxy struct for DataDatabricksShareObjectOutputReference
type jsiiProxy_DataDatabricksShareObjectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) AddedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) AddedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) AddedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) AddedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) CdfEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdfEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) CdfEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdfEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) DataObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataObjectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) DataObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataObjectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) HistoryDataSharingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyDataSharingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) HistoryDataSharingStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyDataSharingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) Partition() DataDatabricksShareObjectPartitionList {
	var returns DataDatabricksShareObjectPartitionList
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) PartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) SharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) SharedAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) StartVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) StartVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) StringSharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringSharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) StringSharedAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringSharedAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksShareObjectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksShareObjectOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksShareObjectOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksShareObjectOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksShare.DataDatabricksShareObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksShareObjectOutputReference_Override(d DataDatabricksShareObjectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksShare.DataDatabricksShareObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetAddedAt(val *float64) {
	if err := j.validateSetAddedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetAddedBy(val *string) {
	if err := j.validateSetAddedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addedBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetCdfEnabled(val interface{}) {
	if err := j.validateSetCdfEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdfEnabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetDataObjectType(val *string) {
	if err := j.validateSetDataObjectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataObjectType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetHistoryDataSharingStatus(val *string) {
	if err := j.validateSetHistoryDataSharingStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"historyDataSharingStatus",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetSharedAs(val *string) {
	if err := j.validateSetSharedAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAs",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetStartVersion(val *float64) {
	if err := j.validateSetStartVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetStringSharedAs(val *string) {
	if err := j.validateSetStringSharedAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringSharedAs",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksShareObjectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) PutPartition(value interface{}) {
	if err := d.validatePutPartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPartition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetAddedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetAddedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetAddedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetAddedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetCdfEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetCdfEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		d,
		"resetContent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetHistoryDataSharingStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetHistoryDataSharingStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetPartition() {
	_jsii_.InvokeVoid(
		d,
		"resetPartition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetSharedAs() {
	_jsii_.InvokeVoid(
		d,
		"resetSharedAs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetStartVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetStartVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ResetStringSharedAs() {
	_jsii_.InvokeVoid(
		d,
		"resetStringSharedAs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksShareObjectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

