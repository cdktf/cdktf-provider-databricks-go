// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksaccountnetworkpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference interface {
	cdktf.ComplexObject
	AllowedInternetDestinations() DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedInternetDestinationsList
	AllowedInternetDestinationsInput() interface{}
	AllowedStorageDestinations() DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsList
	AllowedStorageDestinationsInput() interface{}
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
	PolicyEnforcement() DataDatabricksAccountNetworkPolicyEgressNetworkAccessPolicyEnforcementOutputReference
	PolicyEnforcementInput() interface{}
	RestrictionMode() *string
	SetRestrictionMode(val *string)
	RestrictionModeInput() *string
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
	PutAllowedInternetDestinations(value interface{})
	PutAllowedStorageDestinations(value interface{})
	PutPolicyEnforcement(value *DataDatabricksAccountNetworkPolicyEgressNetworkAccessPolicyEnforcement)
	ResetAllowedInternetDestinations()
	ResetAllowedStorageDestinations()
	ResetPolicyEnforcement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference
type jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) AllowedInternetDestinations() DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedInternetDestinationsList {
	var returns DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedInternetDestinationsList
	_jsii_.Get(
		j,
		"allowedInternetDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) AllowedInternetDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedInternetDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) AllowedStorageDestinations() DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsList {
	var returns DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsList
	_jsii_.Get(
		j,
		"allowedStorageDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) AllowedStorageDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedStorageDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) PolicyEnforcement() DataDatabricksAccountNetworkPolicyEgressNetworkAccessPolicyEnforcementOutputReference {
	var returns DataDatabricksAccountNetworkPolicyEgressNetworkAccessPolicyEnforcementOutputReference
	_jsii_.Get(
		j,
		"policyEnforcement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) PolicyEnforcementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyEnforcementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) RestrictionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) RestrictionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference_Override(d DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference)SetRestrictionMode(val *string) {
	if err := j.validateSetRestrictionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictionMode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) PutAllowedInternetDestinations(value interface{}) {
	if err := d.validatePutAllowedInternetDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllowedInternetDestinations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) PutAllowedStorageDestinations(value interface{}) {
	if err := d.validatePutAllowedStorageDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllowedStorageDestinations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) PutPolicyEnforcement(value *DataDatabricksAccountNetworkPolicyEgressNetworkAccessPolicyEnforcement) {
	if err := d.validatePutPolicyEnforcementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPolicyEnforcement",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) ResetAllowedInternetDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedInternetDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) ResetAllowedStorageDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedStorageDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) ResetPolicyEnforcement() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyEnforcement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

