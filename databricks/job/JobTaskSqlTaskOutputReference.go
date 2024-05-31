// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskSqlTaskOutputReference interface {
	cdktf.ComplexObject
	Alert() JobTaskSqlTaskAlertOutputReference
	AlertInput() *JobTaskSqlTaskAlert
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
	Dashboard() JobTaskSqlTaskDashboardOutputReference
	DashboardInput() *JobTaskSqlTaskDashboard
	File() JobTaskSqlTaskFileOutputReference
	FileInput() *JobTaskSqlTaskFile
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskSqlTask
	SetInternalValue(val *JobTaskSqlTask)
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	Query() JobTaskSqlTaskQueryOutputReference
	QueryInput() *JobTaskSqlTaskQuery
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
	PutAlert(value *JobTaskSqlTaskAlert)
	PutDashboard(value *JobTaskSqlTaskDashboard)
	PutFile(value *JobTaskSqlTaskFile)
	PutQuery(value *JobTaskSqlTaskQuery)
	ResetAlert()
	ResetDashboard()
	ResetFile()
	ResetParameters()
	ResetQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskSqlTaskOutputReference
type jsiiProxy_JobTaskSqlTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) Alert() JobTaskSqlTaskAlertOutputReference {
	var returns JobTaskSqlTaskAlertOutputReference
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) AlertInput() *JobTaskSqlTaskAlert {
	var returns *JobTaskSqlTaskAlert
	_jsii_.Get(
		j,
		"alertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) Dashboard() JobTaskSqlTaskDashboardOutputReference {
	var returns JobTaskSqlTaskDashboardOutputReference
	_jsii_.Get(
		j,
		"dashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) DashboardInput() *JobTaskSqlTaskDashboard {
	var returns *JobTaskSqlTaskDashboard
	_jsii_.Get(
		j,
		"dashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) File() JobTaskSqlTaskFileOutputReference {
	var returns JobTaskSqlTaskFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) FileInput() *JobTaskSqlTaskFile {
	var returns *JobTaskSqlTaskFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) InternalValue() *JobTaskSqlTask {
	var returns *JobTaskSqlTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) Query() JobTaskSqlTaskQueryOutputReference {
	var returns JobTaskSqlTaskQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) QueryInput() *JobTaskSqlTaskQuery {
	var returns *JobTaskSqlTaskQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


func NewJobTaskSqlTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskSqlTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskSqlTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskSqlTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskSqlTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskSqlTaskOutputReference_Override(j JobTaskSqlTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskSqlTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference)SetInternalValue(val *JobTaskSqlTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) PutAlert(value *JobTaskSqlTaskAlert) {
	if err := j.validatePutAlertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAlert",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) PutDashboard(value *JobTaskSqlTaskDashboard) {
	if err := j.validatePutDashboardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDashboard",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) PutFile(value *JobTaskSqlTaskFile) {
	if err := j.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFile",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) PutQuery(value *JobTaskSqlTaskQuery) {
	if err := j.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putQuery",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ResetAlert() {
	_jsii_.InvokeVoid(
		j,
		"resetAlert",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ResetDashboard() {
	_jsii_.InvokeVoid(
		j,
		"resetDashboard",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		j,
		"resetFile",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		j,
		"resetParameters",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		j,
		"resetQuery",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskSqlTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

