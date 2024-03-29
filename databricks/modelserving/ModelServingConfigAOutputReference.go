// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/modelserving/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingConfigAOutputReference interface {
	cdktf.ComplexObject
	AutoCaptureConfig() ModelServingConfigAutoCaptureConfigOutputReference
	AutoCaptureConfigInput() *ModelServingConfigAutoCaptureConfig
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
	InternalValue() *ModelServingConfigA
	SetInternalValue(val *ModelServingConfigA)
	ServedEntities() ModelServingConfigServedEntitiesList
	ServedEntitiesInput() interface{}
	ServedModels() ModelServingConfigServedModelsList
	ServedModelsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficConfig() ModelServingConfigTrafficConfigOutputReference
	TrafficConfigInput() *ModelServingConfigTrafficConfig
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
	PutAutoCaptureConfig(value *ModelServingConfigAutoCaptureConfig)
	PutServedEntities(value interface{})
	PutServedModels(value interface{})
	PutTrafficConfig(value *ModelServingConfigTrafficConfig)
	ResetAutoCaptureConfig()
	ResetServedEntities()
	ResetServedModels()
	ResetTrafficConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingConfigAOutputReference
type jsiiProxy_ModelServingConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) AutoCaptureConfig() ModelServingConfigAutoCaptureConfigOutputReference {
	var returns ModelServingConfigAutoCaptureConfigOutputReference
	_jsii_.Get(
		j,
		"autoCaptureConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) AutoCaptureConfigInput() *ModelServingConfigAutoCaptureConfig {
	var returns *ModelServingConfigAutoCaptureConfig
	_jsii_.Get(
		j,
		"autoCaptureConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) InternalValue() *ModelServingConfigA {
	var returns *ModelServingConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) ServedEntities() ModelServingConfigServedEntitiesList {
	var returns ModelServingConfigServedEntitiesList
	_jsii_.Get(
		j,
		"servedEntities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) ServedEntitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servedEntitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) ServedModels() ModelServingConfigServedModelsList {
	var returns ModelServingConfigServedModelsList
	_jsii_.Get(
		j,
		"servedModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) ServedModelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servedModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) TrafficConfig() ModelServingConfigTrafficConfigOutputReference {
	var returns ModelServingConfigTrafficConfigOutputReference
	_jsii_.Get(
		j,
		"trafficConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigAOutputReference) TrafficConfigInput() *ModelServingConfigTrafficConfig {
	var returns *ModelServingConfigTrafficConfig
	_jsii_.Get(
		j,
		"trafficConfigInput",
		&returns,
	)
	return returns
}


func NewModelServingConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelServingConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingConfigAOutputReference_Override(m ModelServingConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigAOutputReference)SetInternalValue(val *ModelServingConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) PutAutoCaptureConfig(value *ModelServingConfigAutoCaptureConfig) {
	if err := m.validatePutAutoCaptureConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAutoCaptureConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) PutServedEntities(value interface{}) {
	if err := m.validatePutServedEntitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putServedEntities",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) PutServedModels(value interface{}) {
	if err := m.validatePutServedModelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putServedModels",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) PutTrafficConfig(value *ModelServingConfigTrafficConfig) {
	if err := m.validatePutTrafficConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTrafficConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) ResetAutoCaptureConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoCaptureConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) ResetServedEntities() {
	_jsii_.InvokeVoid(
		m,
		"resetServedEntities",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) ResetServedModels() {
	_jsii_.InvokeVoid(
		m,
		"resetServedModels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) ResetTrafficConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetTrafficConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelServingConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

