// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onlinetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/onlinetable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OnlineTableSpecOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *OnlineTableSpec
	SetInternalValue(val *OnlineTableSpec)
	PerformFullCopy() interface{}
	SetPerformFullCopy(val interface{})
	PerformFullCopyInput() interface{}
	PipelineId() *string
	PrimaryKeyColumns() *[]*string
	SetPrimaryKeyColumns(val *[]*string)
	PrimaryKeyColumnsInput() *[]*string
	RunContinuously() OnlineTableSpecRunContinuouslyOutputReference
	RunContinuouslyInput() *OnlineTableSpecRunContinuously
	RunTriggered() OnlineTableSpecRunTriggeredOutputReference
	RunTriggeredInput() *OnlineTableSpecRunTriggered
	SourceTableFullName() *string
	SetSourceTableFullName(val *string)
	SourceTableFullNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeseriesKey() *string
	SetTimeseriesKey(val *string)
	TimeseriesKeyInput() *string
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
	PutRunContinuously(value *OnlineTableSpecRunContinuously)
	PutRunTriggered(value *OnlineTableSpecRunTriggered)
	ResetPerformFullCopy()
	ResetPrimaryKeyColumns()
	ResetRunContinuously()
	ResetRunTriggered()
	ResetSourceTableFullName()
	ResetTimeseriesKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OnlineTableSpecOutputReference
type jsiiProxy_OnlineTableSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) InternalValue() *OnlineTableSpec {
	var returns *OnlineTableSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) PerformFullCopy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performFullCopy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) PerformFullCopyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performFullCopyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) PrimaryKeyColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) PrimaryKeyColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) RunContinuously() OnlineTableSpecRunContinuouslyOutputReference {
	var returns OnlineTableSpecRunContinuouslyOutputReference
	_jsii_.Get(
		j,
		"runContinuously",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) RunContinuouslyInput() *OnlineTableSpecRunContinuously {
	var returns *OnlineTableSpecRunContinuously
	_jsii_.Get(
		j,
		"runContinuouslyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) RunTriggered() OnlineTableSpecRunTriggeredOutputReference {
	var returns OnlineTableSpecRunTriggeredOutputReference
	_jsii_.Get(
		j,
		"runTriggered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) RunTriggeredInput() *OnlineTableSpecRunTriggered {
	var returns *OnlineTableSpecRunTriggered
	_jsii_.Get(
		j,
		"runTriggeredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) SourceTableFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) SourceTableFullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) TimeseriesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableSpecOutputReference) TimeseriesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKeyInput",
		&returns,
	)
	return returns
}


func NewOnlineTableSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OnlineTableSpecOutputReference {
	_init_.Initialize()

	if err := validateNewOnlineTableSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OnlineTableSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.onlineTable.OnlineTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOnlineTableSpecOutputReference_Override(o OnlineTableSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.onlineTable.OnlineTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetInternalValue(val *OnlineTableSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetPerformFullCopy(val interface{}) {
	if err := j.validateSetPerformFullCopyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performFullCopy",
		val,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetPrimaryKeyColumns(val *[]*string) {
	if err := j.validateSetPrimaryKeyColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyColumns",
		val,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetSourceTableFullName(val *string) {
	if err := j.validateSetSourceTableFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTableFullName",
		val,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OnlineTableSpecOutputReference)SetTimeseriesKey(val *string) {
	if err := j.validateSetTimeseriesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeseriesKey",
		val,
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) PutRunContinuously(value *OnlineTableSpecRunContinuously) {
	if err := o.validatePutRunContinuouslyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRunContinuously",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) PutRunTriggered(value *OnlineTableSpecRunTriggered) {
	if err := o.validatePutRunTriggeredParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRunTriggered",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) ResetPerformFullCopy() {
	_jsii_.InvokeVoid(
		o,
		"resetPerformFullCopy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) ResetPrimaryKeyColumns() {
	_jsii_.InvokeVoid(
		o,
		"resetPrimaryKeyColumns",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) ResetRunContinuously() {
	_jsii_.InvokeVoid(
		o,
		"resetRunContinuously",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) ResetRunTriggered() {
	_jsii_.InvokeVoid(
		o,
		"resetRunTriggered",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) ResetSourceTableFullName() {
	_jsii_.InvokeVoid(
		o,
		"resetSourceTableFullName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) ResetTimeseriesKey() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeseriesKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

