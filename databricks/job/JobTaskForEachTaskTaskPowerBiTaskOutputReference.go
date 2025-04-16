// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskPowerBiTaskOutputReference interface {
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
	ConnectionResourceName() *string
	SetConnectionResourceName(val *string)
	ConnectionResourceNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskForEachTaskTaskPowerBiTask
	SetInternalValue(val *JobTaskForEachTaskTaskPowerBiTask)
	PowerBiModel() JobTaskForEachTaskTaskPowerBiTaskPowerBiModelOutputReference
	PowerBiModelInput() *JobTaskForEachTaskTaskPowerBiTaskPowerBiModel
	RefreshAfterUpdate() interface{}
	SetRefreshAfterUpdate(val interface{})
	RefreshAfterUpdateInput() interface{}
	Tables() JobTaskForEachTaskTaskPowerBiTaskTablesList
	TablesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WarehouseId() *string
	SetWarehouseId(val *string)
	WarehouseIdInput() *string
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
	PutPowerBiModel(value *JobTaskForEachTaskTaskPowerBiTaskPowerBiModel)
	PutTables(value interface{})
	ResetConnectionResourceName()
	ResetPowerBiModel()
	ResetRefreshAfterUpdate()
	ResetTables()
	ResetWarehouseId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskPowerBiTaskOutputReference
type jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ConnectionResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ConnectionResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) InternalValue() *JobTaskForEachTaskTaskPowerBiTask {
	var returns *JobTaskForEachTaskTaskPowerBiTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) PowerBiModel() JobTaskForEachTaskTaskPowerBiTaskPowerBiModelOutputReference {
	var returns JobTaskForEachTaskTaskPowerBiTaskPowerBiModelOutputReference
	_jsii_.Get(
		j,
		"powerBiModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) PowerBiModelInput() *JobTaskForEachTaskTaskPowerBiTaskPowerBiModel {
	var returns *JobTaskForEachTaskTaskPowerBiTaskPowerBiModel
	_jsii_.Get(
		j,
		"powerBiModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) RefreshAfterUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshAfterUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) RefreshAfterUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshAfterUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) Tables() JobTaskForEachTaskTaskPowerBiTaskTablesList {
	var returns JobTaskForEachTaskTaskPowerBiTaskTablesList
	_jsii_.Get(
		j,
		"tables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) TablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskPowerBiTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskPowerBiTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskPowerBiTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskPowerBiTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskPowerBiTaskOutputReference_Override(j JobTaskForEachTaskTaskPowerBiTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskPowerBiTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference)SetConnectionResourceName(val *string) {
	if err := j.validateSetConnectionResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionResourceName",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskPowerBiTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference)SetRefreshAfterUpdate(val interface{}) {
	if err := j.validateSetRefreshAfterUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshAfterUpdate",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) PutPowerBiModel(value *JobTaskForEachTaskTaskPowerBiTaskPowerBiModel) {
	if err := j.validatePutPowerBiModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPowerBiModel",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) PutTables(value interface{}) {
	if err := j.validatePutTablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTables",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ResetConnectionResourceName() {
	_jsii_.InvokeVoid(
		j,
		"resetConnectionResourceName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ResetPowerBiModel() {
	_jsii_.InvokeVoid(
		j,
		"resetPowerBiModel",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ResetRefreshAfterUpdate() {
	_jsii_.InvokeVoid(
		j,
		"resetRefreshAfterUpdate",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ResetTables() {
	_jsii_.InvokeVoid(
		j,
		"resetTables",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		j,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskPowerBiTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

