// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/cleanroomasset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanRoomAssetNotebookReviewsOutputReference interface {
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
	InternalValue() *CleanRoomAssetNotebookReviews
	SetInternalValue(val *CleanRoomAssetNotebookReviews)
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

// The jsii proxy struct for CleanRoomAssetNotebookReviewsOutputReference
type jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) CreatedAtMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) CreatedAtMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) InternalValue() *CleanRoomAssetNotebookReviews {
	var returns *CleanRoomAssetNotebookReviews
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ReviewerCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewerCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ReviewerCollaboratorAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewerCollaboratorAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ReviewState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ReviewStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ReviewSubReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewSubReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ReviewSubReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewSubReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanRoomAssetNotebookReviewsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CleanRoomAssetNotebookReviewsOutputReference {
	_init_.Initialize()

	if err := validateNewCleanRoomAssetNotebookReviewsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAssetNotebookReviewsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCleanRoomAssetNotebookReviewsOutputReference_Override(c CleanRoomAssetNotebookReviewsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomAsset.CleanRoomAssetNotebookReviewsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetCreatedAtMillis(val *float64) {
	if err := j.validateSetCreatedAtMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAtMillis",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetInternalValue(val *CleanRoomAssetNotebookReviews) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetReviewerCollaboratorAlias(val *string) {
	if err := j.validateSetReviewerCollaboratorAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewerCollaboratorAlias",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetReviewState(val *string) {
	if err := j.validateSetReviewStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewState",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetReviewSubReason(val *string) {
	if err := j.validateSetReviewSubReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewSubReason",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ResetCreatedAtMillis() {
	_jsii_.InvokeVoid(
		c,
		"resetCreatedAtMillis",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ResetReviewerCollaboratorAlias() {
	_jsii_.InvokeVoid(
		c,
		"resetReviewerCollaboratorAlias",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ResetReviewState() {
	_jsii_.InvokeVoid(
		c,
		"resetReviewState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ResetReviewSubReason() {
	_jsii_.InvokeVoid(
		c,
		"resetReviewSubReason",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanRoomAssetNotebookReviewsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

