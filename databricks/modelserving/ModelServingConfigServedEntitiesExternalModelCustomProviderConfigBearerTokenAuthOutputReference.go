// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/modelserving/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference interface {
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
	InternalValue() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth
	SetInternalValue(val *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	TokenPlaintext() *string
	SetTokenPlaintext(val *string)
	TokenPlaintextInput() *string
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
	ResetToken()
	ResetTokenPlaintext()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference
type jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) InternalValue() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth {
	var returns *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) TokenPlaintext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenPlaintext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) TokenPlaintextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenPlaintextInput",
		&returns,
	)
	return returns
}


func NewModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference_Override(m ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference)SetInternalValue(val *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference)SetTokenPlaintext(val *string) {
	if err := j.validateSetTokenPlaintextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPlaintext",
		val,
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) ResetToken() {
	_jsii_.InvokeVoid(
		m,
		"resetToken",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) ResetTokenPlaintext() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenPlaintext",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

