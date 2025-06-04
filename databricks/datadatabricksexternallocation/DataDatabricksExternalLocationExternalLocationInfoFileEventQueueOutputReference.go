// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksexternallocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksexternallocation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference interface {
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
	InternalValue() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueue
	SetInternalValue(val *DataDatabricksExternalLocationExternalLocationInfoFileEventQueue)
	ManagedAqs() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqsOutputReference
	ManagedAqsInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqs
	ManagedPubsub() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsubOutputReference
	ManagedPubsubInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsub
	ManagedSqs() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqsOutputReference
	ManagedSqsInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqs
	ProvidedAqs() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqsOutputReference
	ProvidedAqsInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqs
	ProvidedPubsub() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsubOutputReference
	ProvidedPubsubInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsub
	ProvidedSqs() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqsOutputReference
	ProvidedSqsInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqs
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
	PutManagedAqs(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqs)
	PutManagedPubsub(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsub)
	PutManagedSqs(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqs)
	PutProvidedAqs(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqs)
	PutProvidedPubsub(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsub)
	PutProvidedSqs(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqs)
	ResetManagedAqs()
	ResetManagedPubsub()
	ResetManagedSqs()
	ResetProvidedAqs()
	ResetProvidedPubsub()
	ResetProvidedSqs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference
type jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) InternalValue() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueue {
	var returns *DataDatabricksExternalLocationExternalLocationInfoFileEventQueue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ManagedAqs() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqsOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqsOutputReference
	_jsii_.Get(
		j,
		"managedAqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ManagedAqsInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqs
	_jsii_.Get(
		j,
		"managedAqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ManagedPubsub() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsubOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsubOutputReference
	_jsii_.Get(
		j,
		"managedPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ManagedPubsubInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsub {
	var returns *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsub
	_jsii_.Get(
		j,
		"managedPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ManagedSqs() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqsOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqsOutputReference
	_jsii_.Get(
		j,
		"managedSqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ManagedSqsInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqs
	_jsii_.Get(
		j,
		"managedSqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ProvidedAqs() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqsOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqsOutputReference
	_jsii_.Get(
		j,
		"providedAqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ProvidedAqsInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqs
	_jsii_.Get(
		j,
		"providedAqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ProvidedPubsub() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsubOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsubOutputReference
	_jsii_.Get(
		j,
		"providedPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ProvidedPubsubInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsub {
	var returns *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsub
	_jsii_.Get(
		j,
		"providedPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ProvidedSqs() DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqsOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqsOutputReference
	_jsii_.Get(
		j,
		"providedSqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ProvidedSqsInput() *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqs
	_jsii_.Get(
		j,
		"providedSqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksExternalLocation.DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference_Override(d DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksExternalLocation.DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference)SetInternalValue(val *DataDatabricksExternalLocationExternalLocationInfoFileEventQueue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) PutManagedAqs(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqs) {
	if err := d.validatePutManagedAqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManagedAqs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) PutManagedPubsub(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsub) {
	if err := d.validatePutManagedPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManagedPubsub",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) PutManagedSqs(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqs) {
	if err := d.validatePutManagedSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManagedSqs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) PutProvidedAqs(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqs) {
	if err := d.validatePutProvidedAqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvidedAqs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) PutProvidedPubsub(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsub) {
	if err := d.validatePutProvidedPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvidedPubsub",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) PutProvidedSqs(value *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqs) {
	if err := d.validatePutProvidedSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvidedSqs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ResetManagedAqs() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedAqs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ResetManagedPubsub() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedPubsub",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ResetManagedSqs() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedSqs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ResetProvidedAqs() {
	_jsii_.InvokeVoid(
		d,
		"resetProvidedAqs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ResetProvidedPubsub() {
	_jsii_.InvokeVoid(
		d,
		"resetProvidedPubsub",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ResetProvidedSqs() {
	_jsii_.InvokeVoid(
		d,
		"resetProvidedSqs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoFileEventQueueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

