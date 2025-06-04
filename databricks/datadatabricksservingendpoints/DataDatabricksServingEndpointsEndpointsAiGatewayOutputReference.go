// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksservingendpoints/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference interface {
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
	FallbackConfig() DataDatabricksServingEndpointsEndpointsAiGatewayFallbackConfigList
	FallbackConfigInput() interface{}
	// Experimental.
	Fqn() *string
	Guardrails() DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsList
	GuardrailsInput() interface{}
	InferenceTableConfig() DataDatabricksServingEndpointsEndpointsAiGatewayInferenceTableConfigList
	InferenceTableConfigInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RateLimits() DataDatabricksServingEndpointsEndpointsAiGatewayRateLimitsList
	RateLimitsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageTrackingConfig() DataDatabricksServingEndpointsEndpointsAiGatewayUsageTrackingConfigList
	UsageTrackingConfigInput() interface{}
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
	PutFallbackConfig(value interface{})
	PutGuardrails(value interface{})
	PutInferenceTableConfig(value interface{})
	PutRateLimits(value interface{})
	PutUsageTrackingConfig(value interface{})
	ResetFallbackConfig()
	ResetGuardrails()
	ResetInferenceTableConfig()
	ResetRateLimits()
	ResetUsageTrackingConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference
type jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) FallbackConfig() DataDatabricksServingEndpointsEndpointsAiGatewayFallbackConfigList {
	var returns DataDatabricksServingEndpointsEndpointsAiGatewayFallbackConfigList
	_jsii_.Get(
		j,
		"fallbackConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) FallbackConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) Guardrails() DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsList {
	var returns DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsList
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GuardrailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guardrailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) InferenceTableConfig() DataDatabricksServingEndpointsEndpointsAiGatewayInferenceTableConfigList {
	var returns DataDatabricksServingEndpointsEndpointsAiGatewayInferenceTableConfigList
	_jsii_.Get(
		j,
		"inferenceTableConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) InferenceTableConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceTableConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) RateLimits() DataDatabricksServingEndpointsEndpointsAiGatewayRateLimitsList {
	var returns DataDatabricksServingEndpointsEndpointsAiGatewayRateLimitsList
	_jsii_.Get(
		j,
		"rateLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) RateLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) UsageTrackingConfig() DataDatabricksServingEndpointsEndpointsAiGatewayUsageTrackingConfigList {
	var returns DataDatabricksServingEndpointsEndpointsAiGatewayUsageTrackingConfigList
	_jsii_.Get(
		j,
		"usageTrackingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) UsageTrackingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usageTrackingConfigInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksServingEndpointsEndpointsAiGatewayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksServingEndpointsEndpointsAiGatewayOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksServingEndpointsEndpointsAiGatewayOutputReference_Override(d DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) PutFallbackConfig(value interface{}) {
	if err := d.validatePutFallbackConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFallbackConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) PutGuardrails(value interface{}) {
	if err := d.validatePutGuardrailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGuardrails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) PutInferenceTableConfig(value interface{}) {
	if err := d.validatePutInferenceTableConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInferenceTableConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) PutRateLimits(value interface{}) {
	if err := d.validatePutRateLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRateLimits",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) PutUsageTrackingConfig(value interface{}) {
	if err := d.validatePutUsageTrackingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUsageTrackingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ResetFallbackConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetFallbackConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ResetGuardrails() {
	_jsii_.InvokeVoid(
		d,
		"resetGuardrails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ResetInferenceTableConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceTableConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ResetRateLimits() {
	_jsii_.InvokeVoid(
		d,
		"resetRateLimits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ResetUsageTrackingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetUsageTrackingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

