package share

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v6/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v6/share/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ShareObjectOutputReference interface {
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
	Partition() ShareObjectPartitionList
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
	ResetHistoryDataSharingStatus()
	ResetPartition()
	ResetSharedAs()
	ResetStartVersion()
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ShareObjectOutputReference
type jsiiProxy_ShareObjectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ShareObjectOutputReference) AddedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) AddedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) AddedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) AddedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) CdfEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdfEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) CdfEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdfEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) DataObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataObjectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) DataObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataObjectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) HistoryDataSharingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyDataSharingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) HistoryDataSharingStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyDataSharingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) Partition() ShareObjectPartitionList {
	var returns ShareObjectPartitionList
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) PartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) SharedAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) SharedAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) StartVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) StartVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShareObjectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewShareObjectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ShareObjectOutputReference {
	_init_.Initialize()

	if err := validateNewShareObjectOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ShareObjectOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.share.ShareObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewShareObjectOutputReference_Override(s ShareObjectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.share.ShareObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetAddedAt(val *float64) {
	if err := j.validateSetAddedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addedAt",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetAddedBy(val *string) {
	if err := j.validateSetAddedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addedBy",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetCdfEnabled(val interface{}) {
	if err := j.validateSetCdfEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdfEnabled",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetDataObjectType(val *string) {
	if err := j.validateSetDataObjectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataObjectType",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetHistoryDataSharingStatus(val *string) {
	if err := j.validateSetHistoryDataSharingStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"historyDataSharingStatus",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetSharedAs(val *string) {
	if err := j.validateSetSharedAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAs",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetStartVersion(val *float64) {
	if err := j.validateSetStartVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startVersion",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ShareObjectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ShareObjectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ShareObjectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ShareObjectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ShareObjectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ShareObjectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ShareObjectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ShareObjectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ShareObjectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ShareObjectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ShareObjectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ShareObjectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ShareObjectOutputReference) PutPartition(value interface{}) {
	if err := s.validatePutPartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPartition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetAddedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetAddedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetAddedBy() {
	_jsii_.InvokeVoid(
		s,
		"resetAddedBy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetCdfEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCdfEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetHistoryDataSharingStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetHistoryDataSharingStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetPartition() {
	_jsii_.InvokeVoid(
		s,
		"resetPartition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetSharedAs() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedAs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetStartVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetStartVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ShareObjectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_ShareObjectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

