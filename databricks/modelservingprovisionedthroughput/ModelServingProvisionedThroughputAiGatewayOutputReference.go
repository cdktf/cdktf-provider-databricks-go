// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/modelservingprovisionedthroughput/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingProvisionedThroughputAiGatewayOutputReference interface {
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
	FallbackConfig() ModelServingProvisionedThroughputAiGatewayFallbackConfigOutputReference
	FallbackConfigInput() *ModelServingProvisionedThroughputAiGatewayFallbackConfig
	// Experimental.
	Fqn() *string
	Guardrails() ModelServingProvisionedThroughputAiGatewayGuardrailsOutputReference
	GuardrailsInput() *ModelServingProvisionedThroughputAiGatewayGuardrails
	InferenceTableConfig() ModelServingProvisionedThroughputAiGatewayInferenceTableConfigOutputReference
	InferenceTableConfigInput() *ModelServingProvisionedThroughputAiGatewayInferenceTableConfig
	InternalValue() *ModelServingProvisionedThroughputAiGateway
	SetInternalValue(val *ModelServingProvisionedThroughputAiGateway)
	RateLimits() ModelServingProvisionedThroughputAiGatewayRateLimitsList
	RateLimitsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageTrackingConfig() ModelServingProvisionedThroughputAiGatewayUsageTrackingConfigOutputReference
	UsageTrackingConfigInput() *ModelServingProvisionedThroughputAiGatewayUsageTrackingConfig
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
	PutFallbackConfig(value *ModelServingProvisionedThroughputAiGatewayFallbackConfig)
	PutGuardrails(value *ModelServingProvisionedThroughputAiGatewayGuardrails)
	PutInferenceTableConfig(value *ModelServingProvisionedThroughputAiGatewayInferenceTableConfig)
	PutRateLimits(value interface{})
	PutUsageTrackingConfig(value *ModelServingProvisionedThroughputAiGatewayUsageTrackingConfig)
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

// The jsii proxy struct for ModelServingProvisionedThroughputAiGatewayOutputReference
type jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) FallbackConfig() ModelServingProvisionedThroughputAiGatewayFallbackConfigOutputReference {
	var returns ModelServingProvisionedThroughputAiGatewayFallbackConfigOutputReference
	_jsii_.Get(
		j,
		"fallbackConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) FallbackConfigInput() *ModelServingProvisionedThroughputAiGatewayFallbackConfig {
	var returns *ModelServingProvisionedThroughputAiGatewayFallbackConfig
	_jsii_.Get(
		j,
		"fallbackConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) Guardrails() ModelServingProvisionedThroughputAiGatewayGuardrailsOutputReference {
	var returns ModelServingProvisionedThroughputAiGatewayGuardrailsOutputReference
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GuardrailsInput() *ModelServingProvisionedThroughputAiGatewayGuardrails {
	var returns *ModelServingProvisionedThroughputAiGatewayGuardrails
	_jsii_.Get(
		j,
		"guardrailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) InferenceTableConfig() ModelServingProvisionedThroughputAiGatewayInferenceTableConfigOutputReference {
	var returns ModelServingProvisionedThroughputAiGatewayInferenceTableConfigOutputReference
	_jsii_.Get(
		j,
		"inferenceTableConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) InferenceTableConfigInput() *ModelServingProvisionedThroughputAiGatewayInferenceTableConfig {
	var returns *ModelServingProvisionedThroughputAiGatewayInferenceTableConfig
	_jsii_.Get(
		j,
		"inferenceTableConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) InternalValue() *ModelServingProvisionedThroughputAiGateway {
	var returns *ModelServingProvisionedThroughputAiGateway
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) RateLimits() ModelServingProvisionedThroughputAiGatewayRateLimitsList {
	var returns ModelServingProvisionedThroughputAiGatewayRateLimitsList
	_jsii_.Get(
		j,
		"rateLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) RateLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) UsageTrackingConfig() ModelServingProvisionedThroughputAiGatewayUsageTrackingConfigOutputReference {
	var returns ModelServingProvisionedThroughputAiGatewayUsageTrackingConfigOutputReference
	_jsii_.Get(
		j,
		"usageTrackingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) UsageTrackingConfigInput() *ModelServingProvisionedThroughputAiGatewayUsageTrackingConfig {
	var returns *ModelServingProvisionedThroughputAiGatewayUsageTrackingConfig
	_jsii_.Get(
		j,
		"usageTrackingConfigInput",
		&returns,
	)
	return returns
}


func NewModelServingProvisionedThroughputAiGatewayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelServingProvisionedThroughputAiGatewayOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingProvisionedThroughputAiGatewayOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServingProvisionedThroughput.ModelServingProvisionedThroughputAiGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingProvisionedThroughputAiGatewayOutputReference_Override(m ModelServingProvisionedThroughputAiGatewayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServingProvisionedThroughput.ModelServingProvisionedThroughputAiGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference)SetInternalValue(val *ModelServingProvisionedThroughputAiGateway) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) PutFallbackConfig(value *ModelServingProvisionedThroughputAiGatewayFallbackConfig) {
	if err := m.validatePutFallbackConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFallbackConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) PutGuardrails(value *ModelServingProvisionedThroughputAiGatewayGuardrails) {
	if err := m.validatePutGuardrailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGuardrails",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) PutInferenceTableConfig(value *ModelServingProvisionedThroughputAiGatewayInferenceTableConfig) {
	if err := m.validatePutInferenceTableConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInferenceTableConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) PutRateLimits(value interface{}) {
	if err := m.validatePutRateLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRateLimits",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) PutUsageTrackingConfig(value *ModelServingProvisionedThroughputAiGatewayUsageTrackingConfig) {
	if err := m.validatePutUsageTrackingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putUsageTrackingConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ResetFallbackConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetFallbackConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ResetGuardrails() {
	_jsii_.InvokeVoid(
		m,
		"resetGuardrails",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ResetInferenceTableConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetInferenceTableConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ResetRateLimits() {
	_jsii_.InvokeVoid(
		m,
		"resetRateLimits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ResetUsageTrackingConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetUsageTrackingConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

