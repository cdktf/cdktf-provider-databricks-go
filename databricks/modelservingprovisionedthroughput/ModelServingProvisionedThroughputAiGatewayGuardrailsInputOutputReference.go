// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/modelservingprovisionedthroughput/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference interface {
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
	InternalValue() *ModelServingProvisionedThroughputAiGatewayGuardrailsInput
	SetInternalValue(val *ModelServingProvisionedThroughputAiGatewayGuardrailsInput)
	InvalidKeywords() *[]*string
	SetInvalidKeywords(val *[]*string)
	InvalidKeywordsInput() *[]*string
	Pii() ModelServingProvisionedThroughputAiGatewayGuardrailsInputPiiOutputReference
	PiiInput() *ModelServingProvisionedThroughputAiGatewayGuardrailsInputPii
	Safety() interface{}
	SetSafety(val interface{})
	SafetyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidTopics() *[]*string
	SetValidTopics(val *[]*string)
	ValidTopicsInput() *[]*string
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
	PutPii(value *ModelServingProvisionedThroughputAiGatewayGuardrailsInputPii)
	ResetInvalidKeywords()
	ResetPii()
	ResetSafety()
	ResetValidTopics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference
type jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) InternalValue() *ModelServingProvisionedThroughputAiGatewayGuardrailsInput {
	var returns *ModelServingProvisionedThroughputAiGatewayGuardrailsInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) InvalidKeywords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidKeywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) InvalidKeywordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidKeywordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) Pii() ModelServingProvisionedThroughputAiGatewayGuardrailsInputPiiOutputReference {
	var returns ModelServingProvisionedThroughputAiGatewayGuardrailsInputPiiOutputReference
	_jsii_.Get(
		j,
		"pii",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) PiiInput() *ModelServingProvisionedThroughputAiGatewayGuardrailsInputPii {
	var returns *ModelServingProvisionedThroughputAiGatewayGuardrailsInputPii
	_jsii_.Get(
		j,
		"piiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) Safety() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"safety",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) SafetyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"safetyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ValidTopics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ValidTopicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validTopicsInput",
		&returns,
	)
	return returns
}


func NewModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServingProvisionedThroughput.ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference_Override(m ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServingProvisionedThroughput.ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference)SetInternalValue(val *ModelServingProvisionedThroughputAiGatewayGuardrailsInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference)SetInvalidKeywords(val *[]*string) {
	if err := j.validateSetInvalidKeywordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invalidKeywords",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference)SetSafety(val interface{}) {
	if err := j.validateSetSafetyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"safety",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference)SetValidTopics(val *[]*string) {
	if err := j.validateSetValidTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validTopics",
		val,
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) PutPii(value *ModelServingProvisionedThroughputAiGatewayGuardrailsInputPii) {
	if err := m.validatePutPiiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPii",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ResetInvalidKeywords() {
	_jsii_.InvokeVoid(
		m,
		"resetInvalidKeywords",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ResetPii() {
	_jsii_.InvokeVoid(
		m,
		"resetPii",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ResetSafety() {
	_jsii_.InvokeVoid(
		m,
		"resetSafety",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ResetValidTopics() {
	_jsii_.InvokeVoid(
		m,
		"resetValidTopics",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputAiGatewayGuardrailsInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

