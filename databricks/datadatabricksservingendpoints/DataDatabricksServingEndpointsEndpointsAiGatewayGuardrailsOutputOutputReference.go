// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksservingendpoints/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference interface {
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
	InvalidKeywords() *[]*string
	SetInvalidKeywords(val *[]*string)
	InvalidKeywordsInput() *[]*string
	Pii() DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputPiiList
	PiiInput() interface{}
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutPii(value interface{})
	ResetInvalidKeywords()
	ResetPii()
	ResetSafety()
	ResetValidTopics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference
type jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) InvalidKeywords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidKeywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) InvalidKeywordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidKeywordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) Pii() DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputPiiList {
	var returns DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputPiiList
	_jsii_.Get(
		j,
		"pii",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) PiiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"piiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) Safety() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"safety",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) SafetyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"safetyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ValidTopics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ValidTopicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validTopicsInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference_Override(d DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference)SetInvalidKeywords(val *[]*string) {
	if err := j.validateSetInvalidKeywordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invalidKeywords",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference)SetSafety(val interface{}) {
	if err := j.validateSetSafetyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"safety",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference)SetValidTopics(val *[]*string) {
	if err := j.validateSetValidTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validTopics",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) PutPii(value interface{}) {
	if err := d.validatePutPiiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPii",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ResetInvalidKeywords() {
	_jsii_.InvokeVoid(
		d,
		"resetInvalidKeywords",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ResetPii() {
	_jsii_.InvokeVoid(
		d,
		"resetPii",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ResetSafety() {
	_jsii_.InvokeVoid(
		d,
		"resetSafety",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ResetValidTopics() {
	_jsii_.InvokeVoid(
		d,
		"resetValidTopics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

