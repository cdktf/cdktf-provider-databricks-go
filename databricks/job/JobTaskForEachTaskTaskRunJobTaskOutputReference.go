// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskRunJobTaskOutputReference interface {
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
	DbtCommands() *[]*string
	SetDbtCommands(val *[]*string)
	DbtCommandsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskForEachTaskTaskRunJobTask
	SetInternalValue(val *JobTaskForEachTaskTaskRunJobTask)
	JarParams() *[]*string
	SetJarParams(val *[]*string)
	JarParamsInput() *[]*string
	JobId() *float64
	SetJobId(val *float64)
	JobIdInput() *float64
	JobParameters() *map[string]*string
	SetJobParameters(val *map[string]*string)
	JobParametersInput() *map[string]*string
	NotebookParams() *map[string]*string
	SetNotebookParams(val *map[string]*string)
	NotebookParamsInput() *map[string]*string
	PipelineParams() JobTaskForEachTaskTaskRunJobTaskPipelineParamsOutputReference
	PipelineParamsInput() *JobTaskForEachTaskTaskRunJobTaskPipelineParams
	PythonNamedParams() *map[string]*string
	SetPythonNamedParams(val *map[string]*string)
	PythonNamedParamsInput() *map[string]*string
	PythonParams() *[]*string
	SetPythonParams(val *[]*string)
	PythonParamsInput() *[]*string
	SparkSubmitParams() *[]*string
	SetSparkSubmitParams(val *[]*string)
	SparkSubmitParamsInput() *[]*string
	SqlParams() *map[string]*string
	SetSqlParams(val *map[string]*string)
	SqlParamsInput() *map[string]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutPipelineParams(value *JobTaskForEachTaskTaskRunJobTaskPipelineParams)
	ResetDbtCommands()
	ResetJarParams()
	ResetJobParameters()
	ResetNotebookParams()
	ResetPipelineParams()
	ResetPythonNamedParams()
	ResetPythonParams()
	ResetSparkSubmitParams()
	ResetSqlParams()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskRunJobTaskOutputReference
type jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) DbtCommands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbtCommands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) DbtCommandsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbtCommandsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) InternalValue() *JobTaskForEachTaskTaskRunJobTask {
	var returns *JobTaskForEachTaskTaskRunJobTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) JarParams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) JarParamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) JobId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) JobIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) JobParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"jobParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) JobParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"jobParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) NotebookParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"notebookParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) NotebookParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"notebookParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) PipelineParams() JobTaskForEachTaskTaskRunJobTaskPipelineParamsOutputReference {
	var returns JobTaskForEachTaskTaskRunJobTaskPipelineParamsOutputReference
	_jsii_.Get(
		j,
		"pipelineParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) PipelineParamsInput() *JobTaskForEachTaskTaskRunJobTaskPipelineParams {
	var returns *JobTaskForEachTaskTaskRunJobTaskPipelineParams
	_jsii_.Get(
		j,
		"pipelineParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) PythonNamedParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pythonNamedParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) PythonNamedParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pythonNamedParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) PythonParams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pythonParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) PythonParamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pythonParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) SparkSubmitParams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sparkSubmitParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) SparkSubmitParamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sparkSubmitParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) SqlParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sqlParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) SqlParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sqlParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskRunJobTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskRunJobTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskRunJobTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskRunJobTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskRunJobTaskOutputReference_Override(j JobTaskForEachTaskTaskRunJobTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskRunJobTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetDbtCommands(val *[]*string) {
	if err := j.validateSetDbtCommandsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbtCommands",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskRunJobTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetJarParams(val *[]*string) {
	if err := j.validateSetJarParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarParams",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetJobId(val *float64) {
	if err := j.validateSetJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetJobParameters(val *map[string]*string) {
	if err := j.validateSetJobParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobParameters",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetNotebookParams(val *map[string]*string) {
	if err := j.validateSetNotebookParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookParams",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetPythonNamedParams(val *map[string]*string) {
	if err := j.validateSetPythonNamedParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonNamedParams",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetPythonParams(val *[]*string) {
	if err := j.validateSetPythonParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonParams",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetSparkSubmitParams(val *[]*string) {
	if err := j.validateSetSparkSubmitParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkSubmitParams",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetSqlParams(val *map[string]*string) {
	if err := j.validateSetSqlParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlParams",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) PutPipelineParams(value *JobTaskForEachTaskTaskRunJobTaskPipelineParams) {
	if err := j.validatePutPipelineParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPipelineParams",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetDbtCommands() {
	_jsii_.InvokeVoid(
		j,
		"resetDbtCommands",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetJarParams() {
	_jsii_.InvokeVoid(
		j,
		"resetJarParams",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetJobParameters() {
	_jsii_.InvokeVoid(
		j,
		"resetJobParameters",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetNotebookParams() {
	_jsii_.InvokeVoid(
		j,
		"resetNotebookParams",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetPipelineParams() {
	_jsii_.InvokeVoid(
		j,
		"resetPipelineParams",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetPythonNamedParams() {
	_jsii_.InvokeVoid(
		j,
		"resetPythonNamedParams",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetPythonParams() {
	_jsii_.InvokeVoid(
		j,
		"resetPythonParams",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetSparkSubmitParams() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkSubmitParams",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ResetSqlParams() {
	_jsii_.InvokeVoid(
		j,
		"resetSqlParams",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskRunJobTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

