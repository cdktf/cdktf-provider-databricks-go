// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/qualitymonitorpluginframework/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QualityMonitorPluginframeworkInferenceLogOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
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

// The jsii proxy struct for QualityMonitorPluginframeworkInferenceLogOutputReference
type jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) Granularities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"granularities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GranularitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"granularitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) LabelCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) LabelColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ModelIdCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ModelIdColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) PredictionCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) PredictionColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) PredictionProbaCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionProbaCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) PredictionProbaColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionProbaColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ProblemType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"problemType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ProblemTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"problemTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) TimestampCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) TimestampColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColInput",
		&returns,
	)
	return returns
}


func NewQualityMonitorPluginframeworkInferenceLogOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QualityMonitorPluginframeworkInferenceLogOutputReference {
	_init_.Initialize()

	if err := validateNewQualityMonitorPluginframeworkInferenceLogOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframeworkInferenceLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQualityMonitorPluginframeworkInferenceLogOutputReference_Override(q QualityMonitorPluginframeworkInferenceLogOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframeworkInferenceLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetGranularities(val *[]*string) {
	if err := j.validateSetGranularitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularities",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetLabelCol(val *string) {
	if err := j.validateSetLabelColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelCol",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetModelIdCol(val *string) {
	if err := j.validateSetModelIdColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelIdCol",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetPredictionCol(val *string) {
	if err := j.validateSetPredictionColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictionCol",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetPredictionProbaCol(val *string) {
	if err := j.validateSetPredictionProbaColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictionProbaCol",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetProblemType(val *string) {
	if err := j.validateSetProblemTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"problemType",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference)SetTimestampCol(val *string) {
	if err := j.validateSetTimestampColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampCol",
		val,
	)
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ResetLabelCol() {
	_jsii_.InvokeVoid(
		q,
		"resetLabelCol",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ResetPredictionProbaCol() {
	_jsii_.InvokeVoid(
		q,
		"resetPredictionProbaCol",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframeworkInferenceLogOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

