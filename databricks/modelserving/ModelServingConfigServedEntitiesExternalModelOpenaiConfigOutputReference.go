// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/modelserving/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference interface {
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
	InternalValue() *ModelServingConfigServedEntitiesExternalModelOpenaiConfig
	SetInternalValue(val *ModelServingConfigServedEntitiesExternalModelOpenaiConfig)
	OpenaiApiBase() *string
	SetOpenaiApiBase(val *string)
	OpenaiApiBaseInput() *string
	OpenaiApiKey() *string
	SetOpenaiApiKey(val *string)
	OpenaiApiKeyInput() *string
	OpenaiApiType() *string
	SetOpenaiApiType(val *string)
	OpenaiApiTypeInput() *string
	OpenaiApiVersion() *string
	SetOpenaiApiVersion(val *string)
	OpenaiApiVersionInput() *string
	OpenaiDeploymentName() *string
	SetOpenaiDeploymentName(val *string)
	OpenaiDeploymentNameInput() *string
	OpenaiOrganization() *string
	SetOpenaiOrganization(val *string)
	OpenaiOrganizationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetOpenaiApiBase()
	ResetOpenaiApiType()
	ResetOpenaiApiVersion()
	ResetOpenaiDeploymentName()
	ResetOpenaiOrganization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference
type jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) InternalValue() *ModelServingConfigServedEntitiesExternalModelOpenaiConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelOpenaiConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiDeploymentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiDeploymentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiDeploymentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiDeploymentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiOrganization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiOrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference_Override(m ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetInternalValue(val *ModelServingConfigServedEntitiesExternalModelOpenaiConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiBase(val *string) {
	if err := j.validateSetOpenaiApiBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiBase",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiKey(val *string) {
	if err := j.validateSetOpenaiApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiKey",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiType(val *string) {
	if err := j.validateSetOpenaiApiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiType",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiVersion(val *string) {
	if err := j.validateSetOpenaiApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiVersion",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiDeploymentName(val *string) {
	if err := j.validateSetOpenaiDeploymentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiDeploymentName",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiOrganization(val *string) {
	if err := j.validateSetOpenaiOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiOrganization",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiApiBase() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenaiApiBase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiApiType() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenaiApiType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiApiVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenaiApiVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiDeploymentName() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenaiDeploymentName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiOrganization() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenaiOrganization",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

