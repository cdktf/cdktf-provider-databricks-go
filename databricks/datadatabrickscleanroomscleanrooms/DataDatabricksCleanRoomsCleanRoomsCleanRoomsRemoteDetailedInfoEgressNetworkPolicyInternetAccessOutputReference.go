// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomscleanrooms/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference interface {
	cdktf.ComplexObject
	AllowedInternetDestinations() DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinationsList
	AllowedInternetDestinationsInput() interface{}
	AllowedStorageDestinations() DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinationsList
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
	LogOnlyMode() DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyModeOutputReference
	LogOnlyModeInput() interface{}
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
	PutLogOnlyMode(value *DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyMode)
	ResetAllowedInternetDestinations()
	ResetAllowedStorageDestinations()
	ResetLogOnlyMode()
	ResetRestrictionMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference
type jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) AllowedInternetDestinations() DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinationsList {
	var returns DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinationsList
	_jsii_.Get(
		j,
		"allowedInternetDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) AllowedInternetDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedInternetDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) AllowedStorageDestinations() DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinationsList {
	var returns DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinationsList
	_jsii_.Get(
		j,
		"allowedStorageDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) AllowedStorageDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedStorageDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) LogOnlyMode() DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyModeOutputReference {
	var returns DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyModeOutputReference
	_jsii_.Get(
		j,
		"logOnlyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) LogOnlyModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logOnlyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) RestrictionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) RestrictionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomsCleanRooms.DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference_Override(d DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomsCleanRooms.DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetRestrictionMode(val *string) {
	if err := j.validateSetRestrictionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictionMode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) PutAllowedInternetDestinations(value interface{}) {
	if err := d.validatePutAllowedInternetDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllowedInternetDestinations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) PutAllowedStorageDestinations(value interface{}) {
	if err := d.validatePutAllowedStorageDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllowedStorageDestinations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) PutLogOnlyMode(value *DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyMode) {
	if err := d.validatePutLogOnlyModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogOnlyMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ResetAllowedInternetDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedInternetDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ResetAllowedStorageDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedStorageDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ResetLogOnlyMode() {
	_jsii_.InvokeVoid(
		d,
		"resetLogOnlyMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ResetRestrictionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetRestrictionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

