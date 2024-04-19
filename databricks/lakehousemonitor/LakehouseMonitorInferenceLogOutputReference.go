// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/lakehousemonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LakehouseMonitorInferenceLogOutputReference interface {
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
	Granularities() *[]*string
	SetGranularities(val *[]*string)
	GranularitiesInput() *[]*string
	InternalValue() *LakehouseMonitorInferenceLog
	SetInternalValue(val *LakehouseMonitorInferenceLog)
	LabelCol() *string
	SetLabelCol(val *string)
	LabelColInput() *string
	ModelIdCol() *string
	SetModelIdCol(val *string)
	ModelIdColInput() *string
	PredictionCol() *string
	SetPredictionCol(val *string)
	PredictionColInput() *string
	PredictionProbaCol() *string
	SetPredictionProbaCol(val *string)
	PredictionProbaColInput() *string
	ProblemType() *string
	SetProblemType(val *string)
	ProblemTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimestampCol() *string
	SetTimestampCol(val *string)
	TimestampColInput() *string
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
	ResetLabelCol()
	ResetPredictionProbaCol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LakehouseMonitorInferenceLogOutputReference
type jsiiProxy_LakehouseMonitorInferenceLogOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) Granularities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"granularities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GranularitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"granularitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) InternalValue() *LakehouseMonitorInferenceLog {
	var returns *LakehouseMonitorInferenceLog
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) LabelCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) LabelColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ModelIdCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ModelIdColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) PredictionCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) PredictionColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) PredictionProbaCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionProbaCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) PredictionProbaColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionProbaColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ProblemType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"problemType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ProblemTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"problemTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) TimestampCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) TimestampColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColInput",
		&returns,
	)
	return returns
}


func NewLakehouseMonitorInferenceLogOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LakehouseMonitorInferenceLogOutputReference {
	_init_.Initialize()

	if err := validateNewLakehouseMonitorInferenceLogOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakehouseMonitorInferenceLogOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.lakehouseMonitor.LakehouseMonitorInferenceLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLakehouseMonitorInferenceLogOutputReference_Override(l LakehouseMonitorInferenceLogOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.lakehouseMonitor.LakehouseMonitorInferenceLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetGranularities(val *[]*string) {
	if err := j.validateSetGranularitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularities",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetInternalValue(val *LakehouseMonitorInferenceLog) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetLabelCol(val *string) {
	if err := j.validateSetLabelColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelCol",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetModelIdCol(val *string) {
	if err := j.validateSetModelIdColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelIdCol",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetPredictionCol(val *string) {
	if err := j.validateSetPredictionColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictionCol",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetPredictionProbaCol(val *string) {
	if err := j.validateSetPredictionProbaColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictionProbaCol",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetProblemType(val *string) {
	if err := j.validateSetProblemTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"problemType",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LakehouseMonitorInferenceLogOutputReference)SetTimestampCol(val *string) {
	if err := j.validateSetTimestampColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampCol",
		val,
	)
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ResetLabelCol() {
	_jsii_.InvokeVoid(
		l,
		"resetLabelCol",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ResetPredictionProbaCol() {
	_jsii_.InvokeVoid(
		l,
		"resetPredictionProbaCol",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakehouseMonitorInferenceLogOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

