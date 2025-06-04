// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/modelserving/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference interface {
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
	DatabricksApiToken() *string
	SetDatabricksApiToken(val *string)
	DatabricksApiTokenInput() *string
	DatabricksApiTokenPlaintext() *string
	SetDatabricksApiTokenPlaintext(val *string)
	DatabricksApiTokenPlaintextInput() *string
	DatabricksWorkspaceUrl() *string
	SetDatabricksWorkspaceUrl(val *string)
	DatabricksWorkspaceUrlInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig
	SetInternalValue(val *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig)
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
	ResetDatabricksApiToken()
	ResetDatabricksApiTokenPlaintext()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference
type jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) DatabricksApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksApiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) DatabricksApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksApiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) DatabricksApiTokenPlaintext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksApiTokenPlaintext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) DatabricksApiTokenPlaintextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksApiTokenPlaintextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) DatabricksWorkspaceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksWorkspaceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) DatabricksWorkspaceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksWorkspaceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) InternalValue() *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference_Override(m ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)SetDatabricksApiToken(val *string) {
	if err := j.validateSetDatabricksApiTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databricksApiToken",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)SetDatabricksApiTokenPlaintext(val *string) {
	if err := j.validateSetDatabricksApiTokenPlaintextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databricksApiTokenPlaintext",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)SetDatabricksWorkspaceUrl(val *string) {
	if err := j.validateSetDatabricksWorkspaceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databricksWorkspaceUrl",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)SetInternalValue(val *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) ResetDatabricksApiToken() {
	_jsii_.InvokeVoid(
		m,
		"resetDatabricksApiToken",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) ResetDatabricksApiTokenPlaintext() {
	_jsii_.InvokeVoid(
		m,
		"resetDatabricksApiTokenPlaintext",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

