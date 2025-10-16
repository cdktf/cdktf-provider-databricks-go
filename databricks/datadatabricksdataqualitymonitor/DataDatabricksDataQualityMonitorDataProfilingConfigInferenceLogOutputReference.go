// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksdataqualitymonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference interface {
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
	LabelColumn() *string
	SetLabelColumn(val *string)
	LabelColumnInput() *string
	ModelIdColumn() *string
	SetModelIdColumn(val *string)
	ModelIdColumnInput() *string
	PredictionColumn() *string
	SetPredictionColumn(val *string)
	PredictionColumnInput() *string
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
	TimestampColumn() *string
	SetTimestampColumn(val *string)
	TimestampColumnInput() *string
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
	ResetLabelColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference
type jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) Granularities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"granularities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GranularitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"granularitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) LabelColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) LabelColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ModelIdColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ModelIdColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) PredictionColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) PredictionColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictionColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ProblemType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"problemType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ProblemTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"problemTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) TimestampColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) TimestampColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDataQualityMonitor.DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference_Override(d DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDataQualityMonitor.DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetGranularities(val *[]*string) {
	if err := j.validateSetGranularitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularities",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetLabelColumn(val *string) {
	if err := j.validateSetLabelColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelColumn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetModelIdColumn(val *string) {
	if err := j.validateSetModelIdColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelIdColumn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetPredictionColumn(val *string) {
	if err := j.validateSetPredictionColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictionColumn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetProblemType(val *string) {
	if err := j.validateSetProblemTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"problemType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference)SetTimestampColumn(val *string) {
	if err := j.validateSetTimestampColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampColumn",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ResetLabelColumn() {
	_jsii_.InvokeVoid(
		d,
		"resetLabelColumn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

