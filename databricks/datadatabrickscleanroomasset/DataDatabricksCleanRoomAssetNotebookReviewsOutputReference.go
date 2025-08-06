// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomasset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomasset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomAssetNotebookReviewsOutputReference interface {
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
	CreatedAtMillis() *float64
	SetCreatedAtMillis(val *float64)
	CreatedAtMillisInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksCleanRoomAssetNotebookReviews
	SetInternalValue(val *DataDatabricksCleanRoomAssetNotebookReviews)
	ReviewerCollaboratorAlias() *string
	SetReviewerCollaboratorAlias(val *string)
	ReviewerCollaboratorAliasInput() *string
	ReviewState() *string
	SetReviewState(val *string)
	ReviewStateInput() *string
	ReviewSubReason() *string
	SetReviewSubReason(val *string)
	ReviewSubReasonInput() *string
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
	ResetComment()
	ResetCreatedAtMillis()
	ResetReviewerCollaboratorAlias()
	ResetReviewState()
	ResetReviewSubReason()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksCleanRoomAssetNotebookReviewsOutputReference
type jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) CreatedAtMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) CreatedAtMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) InternalValue() *DataDatabricksCleanRoomAssetNotebookReviews {
	var returns *DataDatabricksCleanRoomAssetNotebookReviews
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ReviewerCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewerCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ReviewerCollaboratorAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewerCollaboratorAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ReviewState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ReviewStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ReviewSubReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewSubReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ReviewSubReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewSubReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksCleanRoomAssetNotebookReviewsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksCleanRoomAssetNotebookReviewsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomAssetNotebookReviewsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAssetNotebookReviewsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksCleanRoomAssetNotebookReviewsOutputReference_Override(d DataDatabricksCleanRoomAssetNotebookReviewsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAsset.DataDatabricksCleanRoomAssetNotebookReviewsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetCreatedAtMillis(val *float64) {
	if err := j.validateSetCreatedAtMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAtMillis",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetInternalValue(val *DataDatabricksCleanRoomAssetNotebookReviews) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetReviewerCollaboratorAlias(val *string) {
	if err := j.validateSetReviewerCollaboratorAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewerCollaboratorAlias",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetReviewState(val *string) {
	if err := j.validateSetReviewStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewState",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetReviewSubReason(val *string) {
	if err := j.validateSetReviewSubReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewSubReason",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ResetCreatedAtMillis() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAtMillis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ResetReviewerCollaboratorAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetReviewerCollaboratorAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ResetReviewState() {
	_jsii_.InvokeVoid(
		d,
		"resetReviewState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ResetReviewSubReason() {
	_jsii_.InvokeVoid(
		d,
		"resetReviewSubReason",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAssetNotebookReviewsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

