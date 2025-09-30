// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksalertsv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAlertsV2AlertsEvaluationOutputReference interface {
	cdktf.ComplexObject
	ComparisonOperator() *string
	SetComparisonOperator(val *string)
	ComparisonOperatorInput() *string
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
	EmptyResultState() *string
	SetEmptyResultState(val *string)
	EmptyResultStateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastEvaluatedAt() *string
	Notification() DataDatabricksAlertsV2AlertsEvaluationNotificationOutputReference
	NotificationInput() interface{}
	Source() DataDatabricksAlertsV2AlertsEvaluationSourceOutputReference
	SourceInput() interface{}
	State() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() DataDatabricksAlertsV2AlertsEvaluationThresholdOutputReference
	ThresholdInput() interface{}
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
	PutNotification(value *DataDatabricksAlertsV2AlertsEvaluationNotification)
	PutSource(value *DataDatabricksAlertsV2AlertsEvaluationSource)
	PutThreshold(value *DataDatabricksAlertsV2AlertsEvaluationThreshold)
	ResetComparisonOperator()
	ResetEmptyResultState()
	ResetNotification()
	ResetSource()
	ResetThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAlertsV2AlertsEvaluationOutputReference
type jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ComparisonOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) EmptyResultState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyResultState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) EmptyResultStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyResultStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) LastEvaluatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastEvaluatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) Notification() DataDatabricksAlertsV2AlertsEvaluationNotificationOutputReference {
	var returns DataDatabricksAlertsV2AlertsEvaluationNotificationOutputReference
	_jsii_.Get(
		j,
		"notification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) NotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) Source() DataDatabricksAlertsV2AlertsEvaluationSourceOutputReference {
	var returns DataDatabricksAlertsV2AlertsEvaluationSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) Threshold() DataDatabricksAlertsV2AlertsEvaluationThresholdOutputReference {
	var returns DataDatabricksAlertsV2AlertsEvaluationThresholdOutputReference
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ThresholdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAlertsV2AlertsEvaluationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksAlertsV2AlertsEvaluationOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAlertsV2AlertsEvaluationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAlertsV2.DataDatabricksAlertsV2AlertsEvaluationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAlertsV2AlertsEvaluationOutputReference_Override(d DataDatabricksAlertsV2AlertsEvaluationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAlertsV2.DataDatabricksAlertsV2AlertsEvaluationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference)SetEmptyResultState(val *string) {
	if err := j.validateSetEmptyResultStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyResultState",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) PutNotification(value *DataDatabricksAlertsV2AlertsEvaluationNotification) {
	if err := d.validatePutNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) PutSource(value *DataDatabricksAlertsV2AlertsEvaluationSource) {
	if err := d.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) PutThreshold(value *DataDatabricksAlertsV2AlertsEvaluationThreshold) {
	if err := d.validatePutThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putThreshold",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ResetComparisonOperator() {
	_jsii_.InvokeVoid(
		d,
		"resetComparisonOperator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ResetEmptyResultState() {
	_jsii_.InvokeVoid(
		d,
		"resetEmptyResultState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ResetNotification() {
	_jsii_.InvokeVoid(
		d,
		"resetNotification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		d,
		"resetSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsEvaluationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

