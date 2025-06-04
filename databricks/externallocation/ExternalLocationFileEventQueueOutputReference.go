// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externallocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/externallocation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExternalLocationFileEventQueueOutputReference interface {
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
	InternalValue() *ExternalLocationFileEventQueue
	SetInternalValue(val *ExternalLocationFileEventQueue)
	ManagedAqs() ExternalLocationFileEventQueueManagedAqsOutputReference
	ManagedAqsInput() *ExternalLocationFileEventQueueManagedAqs
	ManagedPubsub() ExternalLocationFileEventQueueManagedPubsubOutputReference
	ManagedPubsubInput() *ExternalLocationFileEventQueueManagedPubsub
	ManagedSqs() ExternalLocationFileEventQueueManagedSqsOutputReference
	ManagedSqsInput() *ExternalLocationFileEventQueueManagedSqs
	ProvidedAqs() ExternalLocationFileEventQueueProvidedAqsOutputReference
	ProvidedAqsInput() *ExternalLocationFileEventQueueProvidedAqs
	ProvidedPubsub() ExternalLocationFileEventQueueProvidedPubsubOutputReference
	ProvidedPubsubInput() *ExternalLocationFileEventQueueProvidedPubsub
	ProvidedSqs() ExternalLocationFileEventQueueProvidedSqsOutputReference
	ProvidedSqsInput() *ExternalLocationFileEventQueueProvidedSqs
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
	PutManagedAqs(value *ExternalLocationFileEventQueueManagedAqs)
	PutManagedPubsub(value *ExternalLocationFileEventQueueManagedPubsub)
	PutManagedSqs(value *ExternalLocationFileEventQueueManagedSqs)
	PutProvidedAqs(value *ExternalLocationFileEventQueueProvidedAqs)
	PutProvidedPubsub(value *ExternalLocationFileEventQueueProvidedPubsub)
	PutProvidedSqs(value *ExternalLocationFileEventQueueProvidedSqs)
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

// The jsii proxy struct for ExternalLocationFileEventQueueOutputReference
type jsiiProxy_ExternalLocationFileEventQueueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) InternalValue() *ExternalLocationFileEventQueue {
	var returns *ExternalLocationFileEventQueue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ManagedAqs() ExternalLocationFileEventQueueManagedAqsOutputReference {
	var returns ExternalLocationFileEventQueueManagedAqsOutputReference
	_jsii_.Get(
		j,
		"managedAqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ManagedAqsInput() *ExternalLocationFileEventQueueManagedAqs {
	var returns *ExternalLocationFileEventQueueManagedAqs
	_jsii_.Get(
		j,
		"managedAqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ManagedPubsub() ExternalLocationFileEventQueueManagedPubsubOutputReference {
	var returns ExternalLocationFileEventQueueManagedPubsubOutputReference
	_jsii_.Get(
		j,
		"managedPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ManagedPubsubInput() *ExternalLocationFileEventQueueManagedPubsub {
	var returns *ExternalLocationFileEventQueueManagedPubsub
	_jsii_.Get(
		j,
		"managedPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ManagedSqs() ExternalLocationFileEventQueueManagedSqsOutputReference {
	var returns ExternalLocationFileEventQueueManagedSqsOutputReference
	_jsii_.Get(
		j,
		"managedSqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ManagedSqsInput() *ExternalLocationFileEventQueueManagedSqs {
	var returns *ExternalLocationFileEventQueueManagedSqs
	_jsii_.Get(
		j,
		"managedSqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ProvidedAqs() ExternalLocationFileEventQueueProvidedAqsOutputReference {
	var returns ExternalLocationFileEventQueueProvidedAqsOutputReference
	_jsii_.Get(
		j,
		"providedAqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ProvidedAqsInput() *ExternalLocationFileEventQueueProvidedAqs {
	var returns *ExternalLocationFileEventQueueProvidedAqs
	_jsii_.Get(
		j,
		"providedAqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ProvidedPubsub() ExternalLocationFileEventQueueProvidedPubsubOutputReference {
	var returns ExternalLocationFileEventQueueProvidedPubsubOutputReference
	_jsii_.Get(
		j,
		"providedPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ProvidedPubsubInput() *ExternalLocationFileEventQueueProvidedPubsub {
	var returns *ExternalLocationFileEventQueueProvidedPubsub
	_jsii_.Get(
		j,
		"providedPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ProvidedSqs() ExternalLocationFileEventQueueProvidedSqsOutputReference {
	var returns ExternalLocationFileEventQueueProvidedSqsOutputReference
	_jsii_.Get(
		j,
		"providedSqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ProvidedSqsInput() *ExternalLocationFileEventQueueProvidedSqs {
	var returns *ExternalLocationFileEventQueueProvidedSqs
	_jsii_.Get(
		j,
		"providedSqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExternalLocationFileEventQueueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExternalLocationFileEventQueueOutputReference {
	_init_.Initialize()

	if err := validateNewExternalLocationFileEventQueueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalLocationFileEventQueueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.externalLocation.ExternalLocationFileEventQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExternalLocationFileEventQueueOutputReference_Override(e ExternalLocationFileEventQueueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.externalLocation.ExternalLocationFileEventQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference)SetInternalValue(val *ExternalLocationFileEventQueue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExternalLocationFileEventQueueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) PutManagedAqs(value *ExternalLocationFileEventQueueManagedAqs) {
	if err := e.validatePutManagedAqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedAqs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) PutManagedPubsub(value *ExternalLocationFileEventQueueManagedPubsub) {
	if err := e.validatePutManagedPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedPubsub",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) PutManagedSqs(value *ExternalLocationFileEventQueueManagedSqs) {
	if err := e.validatePutManagedSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedSqs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) PutProvidedAqs(value *ExternalLocationFileEventQueueProvidedAqs) {
	if err := e.validatePutProvidedAqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProvidedAqs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) PutProvidedPubsub(value *ExternalLocationFileEventQueueProvidedPubsub) {
	if err := e.validatePutProvidedPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProvidedPubsub",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) PutProvidedSqs(value *ExternalLocationFileEventQueueProvidedSqs) {
	if err := e.validatePutProvidedSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProvidedSqs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ResetManagedAqs() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedAqs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ResetManagedPubsub() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedPubsub",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ResetManagedSqs() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedSqs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ResetProvidedAqs() {
	_jsii_.InvokeVoid(
		e,
		"resetProvidedAqs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ResetProvidedPubsub() {
	_jsii_.InvokeVoid(
		e,
		"resetProvidedPubsub",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ResetProvidedSqs() {
	_jsii_.InvokeVoid(
		e,
		"resetProvidedSqs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationFileEventQueueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

