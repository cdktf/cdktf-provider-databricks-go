// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharepluginframework

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/sharepluginframework/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SharePluginframeworkObjectOutputReference interface {
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
	EffectiveCdfEnabled() interface{}
	SetEffectiveCdfEnabled(val interface{})
	EffectiveCdfEnabledInput() interface{}
	EffectiveHistoryDataSharingStatus() *string
	SetEffectiveHistoryDataSharingStatus(val *string)
	EffectiveHistoryDataSharingStatusInput() *string
	EffectiveSharedAs() *string
	SetEffectiveSharedAs(val *string)
	EffectiveSharedAsInput() *string
	EffectiveStartVersion() *float64
	SetEffectiveStartVersion(val *float64)
	EffectiveStartVersionInput() *float64
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
	Partition() SharePluginframeworkObjectPartitionList
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
	ResetEffectiveCdfEnabled()
	ResetEffectiveHistoryDataSharingStatus()
	ResetEffectiveSharedAs()
	ResetEffectiveStartVersion()
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

// The jsii proxy struct for SharePluginframeworkObjectOutputReference
type jsiiProxy_SharePluginframeworkObjectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) AddedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) AddedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) AddedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) AddedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) CdfEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdfEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) CdfEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdfEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) DataObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataObjectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) DataObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataObjectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) EffectiveCdfEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveCdfEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) EffectiveCdfEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveCdfEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) EffectiveHistoryDataSharingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveHistoryDataSharingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) EffectiveHistoryDataSharingStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveHistoryDataSharingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) EffectiveSharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveSharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) EffectiveSharedAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveSharedAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) EffectiveStartVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveStartVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) EffectiveStartVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveStartVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) HistoryDataSharingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyDataSharingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) HistoryDataSharingStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyDataSharingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) Partition() SharePluginframeworkObjectPartitionList {
	var returns SharePluginframeworkObjectPartitionList
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) PartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) SharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) SharedAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) StartVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) StartVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) StringSharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringSharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) StringSharedAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringSharedAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSharePluginframeworkObjectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SharePluginframeworkObjectOutputReference {
	_init_.Initialize()

	if err := validateNewSharePluginframeworkObjectOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SharePluginframeworkObjectOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.sharePluginframework.SharePluginframeworkObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSharePluginframeworkObjectOutputReference_Override(s SharePluginframeworkObjectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.sharePluginframework.SharePluginframeworkObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetAddedAt(val *float64) {
	if err := j.validateSetAddedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addedAt",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetAddedBy(val *string) {
	if err := j.validateSetAddedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addedBy",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetCdfEnabled(val interface{}) {
	if err := j.validateSetCdfEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdfEnabled",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetDataObjectType(val *string) {
	if err := j.validateSetDataObjectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataObjectType",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetEffectiveCdfEnabled(val interface{}) {
	if err := j.validateSetEffectiveCdfEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveCdfEnabled",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetEffectiveHistoryDataSharingStatus(val *string) {
	if err := j.validateSetEffectiveHistoryDataSharingStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveHistoryDataSharingStatus",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetEffectiveSharedAs(val *string) {
	if err := j.validateSetEffectiveSharedAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveSharedAs",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetEffectiveStartVersion(val *float64) {
	if err := j.validateSetEffectiveStartVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveStartVersion",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetHistoryDataSharingStatus(val *string) {
	if err := j.validateSetHistoryDataSharingStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"historyDataSharingStatus",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetSharedAs(val *string) {
	if err := j.validateSetSharedAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAs",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetStartVersion(val *float64) {
	if err := j.validateSetStartVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startVersion",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetStringSharedAs(val *string) {
	if err := j.validateSetStringSharedAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringSharedAs",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SharePluginframeworkObjectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) PutPartition(value interface{}) {
	if err := s.validatePutPartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPartition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetAddedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetAddedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetAddedBy() {
	_jsii_.InvokeVoid(
		s,
		"resetAddedBy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetCdfEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCdfEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		s,
		"resetContent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetEffectiveCdfEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEffectiveCdfEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetEffectiveHistoryDataSharingStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetEffectiveHistoryDataSharingStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetEffectiveSharedAs() {
	_jsii_.InvokeVoid(
		s,
		"resetEffectiveSharedAs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetEffectiveStartVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetEffectiveStartVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetHistoryDataSharingStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetHistoryDataSharingStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetPartition() {
	_jsii_.InvokeVoid(
		s,
		"resetPartition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetSharedAs() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedAs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetStartVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetStartVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ResetStringSharedAs() {
	_jsii_.InvokeVoid(
		s,
		"resetStringSharedAs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharePluginframeworkObjectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

