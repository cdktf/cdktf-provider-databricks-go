// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTriggerOutputReference interface {
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
	FileArrival() JobTriggerFileArrivalOutputReference
	FileArrivalInput() *JobTriggerFileArrival
	// Experimental.
	Fqn() *string
	InternalValue() *JobTrigger
	SetInternalValue(val *JobTrigger)
	PauseStatus() *string
	SetPauseStatus(val *string)
	PauseStatusInput() *string
	Periodic() JobTriggerPeriodicOutputReference
	PeriodicInput() *JobTriggerPeriodic
	Table() JobTriggerTableOutputReference
	TableInput() *JobTriggerTable
	TableUpdate() JobTriggerTableUpdateOutputReference
	TableUpdateInput() *JobTriggerTableUpdate
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
	PutFileArrival(value *JobTriggerFileArrival)
	PutPeriodic(value *JobTriggerPeriodic)
	PutTable(value *JobTriggerTable)
	PutTableUpdate(value *JobTriggerTableUpdate)
	ResetFileArrival()
	ResetPauseStatus()
	ResetPeriodic()
	ResetTable()
	ResetTableUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTriggerOutputReference
type jsiiProxy_JobTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) FileArrival() JobTriggerFileArrivalOutputReference {
	var returns JobTriggerFileArrivalOutputReference
	_jsii_.Get(
		j,
		"fileArrival",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) FileArrivalInput() *JobTriggerFileArrival {
	var returns *JobTriggerFileArrival
	_jsii_.Get(
		j,
		"fileArrivalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) InternalValue() *JobTrigger {
	var returns *JobTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) PauseStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) PauseStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) Periodic() JobTriggerPeriodicOutputReference {
	var returns JobTriggerPeriodicOutputReference
	_jsii_.Get(
		j,
		"periodic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) PeriodicInput() *JobTriggerPeriodic {
	var returns *JobTriggerPeriodic
	_jsii_.Get(
		j,
		"periodicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) Table() JobTriggerTableOutputReference {
	var returns JobTriggerTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) TableInput() *JobTriggerTable {
	var returns *JobTriggerTable
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) TableUpdate() JobTriggerTableUpdateOutputReference {
	var returns JobTriggerTableUpdateOutputReference
	_jsii_.Get(
		j,
		"tableUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) TableUpdateInput() *JobTriggerTableUpdate {
	var returns *JobTriggerTableUpdate
	_jsii_.Get(
		j,
		"tableUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTriggerOutputReference {
	_init_.Initialize()

	if err := validateNewJobTriggerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTriggerOutputReference_Override(j JobTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTriggerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTriggerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTriggerOutputReference)SetInternalValue(val *JobTrigger) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTriggerOutputReference)SetPauseStatus(val *string) {
	if err := j.validateSetPauseStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pauseStatus",
		val,
	)
}

func (j *jsiiProxy_JobTriggerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTriggerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) PutFileArrival(value *JobTriggerFileArrival) {
	if err := j.validatePutFileArrivalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFileArrival",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) PutPeriodic(value *JobTriggerPeriodic) {
	if err := j.validatePutPeriodicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPeriodic",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) PutTable(value *JobTriggerTable) {
	if err := j.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTable",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) PutTableUpdate(value *JobTriggerTableUpdate) {
	if err := j.validatePutTableUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTableUpdate",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) ResetFileArrival() {
	_jsii_.InvokeVoid(
		j,
		"resetFileArrival",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) ResetPauseStatus() {
	_jsii_.InvokeVoid(
		j,
		"resetPauseStatus",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) ResetPeriodic() {
	_jsii_.InvokeVoid(
		j,
		"resetPeriodic",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		j,
		"resetTable",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) ResetTableUpdate() {
	_jsii_.InvokeVoid(
		j,
		"resetTableUpdate",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

