// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/modelserving/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingAiGatewayOutputReference interface {
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
	FallbackConfig() ModelServingAiGatewayFallbackConfigOutputReference
	FallbackConfigInput() *ModelServingAiGatewayFallbackConfig
	// Experimental.
	Fqn() *string
	Guardrails() ModelServingAiGatewayGuardrailsOutputReference
	GuardrailsInput() *ModelServingAiGatewayGuardrails
	InferenceTableConfig() ModelServingAiGatewayInferenceTableConfigOutputReference
	InferenceTableConfigInput() *ModelServingAiGatewayInferenceTableConfig
	InternalValue() *ModelServingAiGateway
	SetInternalValue(val *ModelServingAiGateway)
	RateLimits() ModelServingAiGatewayRateLimitsList
	RateLimitsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageTrackingConfig() ModelServingAiGatewayUsageTrackingConfigOutputReference
	UsageTrackingConfigInput() *ModelServingAiGatewayUsageTrackingConfig
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutFallbackConfig(value *ModelServingAiGatewayFallbackConfig)
	PutGuardrails(value *ModelServingAiGatewayGuardrails)
	PutInferenceTableConfig(value *ModelServingAiGatewayInferenceTableConfig)
	PutRateLimits(value interface{})
	PutUsageTrackingConfig(value *ModelServingAiGatewayUsageTrackingConfig)
	ResetFallbackConfig()
	ResetGuardrails()
	ResetInferenceTableConfig()
	ResetRateLimits()
	ResetUsageTrackingConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingAiGatewayOutputReference
type jsiiProxy_ModelServingAiGatewayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) FallbackConfig() ModelServingAiGatewayFallbackConfigOutputReference {
	var returns ModelServingAiGatewayFallbackConfigOutputReference
	_jsii_.Get(
		j,
		"fallbackConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) FallbackConfigInput() *ModelServingAiGatewayFallbackConfig {
	var returns *ModelServingAiGatewayFallbackConfig
	_jsii_.Get(
		j,
		"fallbackConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) Guardrails() ModelServingAiGatewayGuardrailsOutputReference {
	var returns ModelServingAiGatewayGuardrailsOutputReference
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) GuardrailsInput() *ModelServingAiGatewayGuardrails {
	var returns *ModelServingAiGatewayGuardrails
	_jsii_.Get(
		j,
		"guardrailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) InferenceTableConfig() ModelServingAiGatewayInferenceTableConfigOutputReference {
	var returns ModelServingAiGatewayInferenceTableConfigOutputReference
	_jsii_.Get(
		j,
		"inferenceTableConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) InferenceTableConfigInput() *ModelServingAiGatewayInferenceTableConfig {
	var returns *ModelServingAiGatewayInferenceTableConfig
	_jsii_.Get(
		j,
		"inferenceTableConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) InternalValue() *ModelServingAiGateway {
	var returns *ModelServingAiGateway
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) RateLimits() ModelServingAiGatewayRateLimitsList {
	var returns ModelServingAiGatewayRateLimitsList
	_jsii_.Get(
		j,
		"rateLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) RateLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) UsageTrackingConfig() ModelServingAiGatewayUsageTrackingConfigOutputReference {
	var returns ModelServingAiGatewayUsageTrackingConfigOutputReference
	_jsii_.Get(
		j,
		"usageTrackingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference) UsageTrackingConfigInput() *ModelServingAiGatewayUsageTrackingConfig {
	var returns *ModelServingAiGatewayUsageTrackingConfig
	_jsii_.Get(
		j,
		"usageTrackingConfigInput",
		&returns,
	)
	return returns
}


func NewModelServingAiGatewayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelServingAiGatewayOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingAiGatewayOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingAiGatewayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingAiGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingAiGatewayOutputReference_Override(m ModelServingAiGatewayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingAiGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference)SetInternalValue(val *ModelServingAiGateway) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingAiGatewayOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) PutFallbackConfig(value *ModelServingAiGatewayFallbackConfig) {
	if err := m.validatePutFallbackConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFallbackConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) PutGuardrails(value *ModelServingAiGatewayGuardrails) {
	if err := m.validatePutGuardrailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGuardrails",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) PutInferenceTableConfig(value *ModelServingAiGatewayInferenceTableConfig) {
	if err := m.validatePutInferenceTableConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInferenceTableConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) PutRateLimits(value interface{}) {
	if err := m.validatePutRateLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRateLimits",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) PutUsageTrackingConfig(value *ModelServingAiGatewayUsageTrackingConfig) {
	if err := m.validatePutUsageTrackingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putUsageTrackingConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) ResetFallbackConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetFallbackConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) ResetGuardrails() {
	_jsii_.InvokeVoid(
		m,
		"resetGuardrails",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) ResetInferenceTableConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetInferenceTableConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) ResetRateLimits() {
	_jsii_.InvokeVoid(
		m,
		"resetRateLimits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) ResetUsageTrackingConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetUsageTrackingConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingAiGatewayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

