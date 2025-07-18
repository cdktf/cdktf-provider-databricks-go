// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/modelservingprovisionedthroughput/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ServedEntityName() *string
	SetServedEntityName(val *string)
	ServedEntityNameInput() *string
	ServedModelName() *string
	SetServedModelName(val *string)
	ServedModelNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficPercentage() *float64
	SetTrafficPercentage(val *float64)
	TrafficPercentageInput() *float64
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
	ResetServedEntityName()
	ResetServedModelName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference
type jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ServedEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servedEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ServedEntityNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servedEntityNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ServedModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servedModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ServedModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servedModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) TrafficPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) TrafficPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentageInput",
		&returns,
	)
	return returns
}


func NewModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServingProvisionedThroughput.ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference_Override(m ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServingProvisionedThroughput.ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference)SetServedEntityName(val *string) {
	if err := j.validateSetServedEntityNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servedEntityName",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference)SetServedModelName(val *string) {
	if err := j.validateSetServedModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servedModelName",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference)SetTrafficPercentage(val *float64) {
	if err := j.validateSetTrafficPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficPercentage",
		val,
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ResetServedEntityName() {
	_jsii_.InvokeVoid(
		m,
		"resetServedEntityName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ResetServedModelName() {
	_jsii_.InvokeVoid(
		m,
		"resetServedModelName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelServingProvisionedThroughputConfigTrafficConfigRoutesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

