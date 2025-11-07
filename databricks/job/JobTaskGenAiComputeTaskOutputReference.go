// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskGenAiComputeTaskOutputReference interface {
	cdktf.ComplexObject
	Command() *string
	SetCommand(val *string)
	CommandInput() *string
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
	Compute() JobTaskGenAiComputeTaskComputeOutputReference
	ComputeInput() *JobTaskGenAiComputeTaskCompute
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DlRuntimeImage() *string
	SetDlRuntimeImage(val *string)
	DlRuntimeImageInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskGenAiComputeTask
	SetInternalValue(val *JobTaskGenAiComputeTask)
	MlflowExperimentName() *string
	SetMlflowExperimentName(val *string)
	MlflowExperimentNameInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrainingScriptPath() *string
	SetTrainingScriptPath(val *string)
	TrainingScriptPathInput() *string
	YamlParameters() *string
	SetYamlParameters(val *string)
	YamlParametersFilePath() *string
	SetYamlParametersFilePath(val *string)
	YamlParametersFilePathInput() *string
	YamlParametersInput() *string
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
	PutCompute(value *JobTaskGenAiComputeTaskCompute)
	ResetCommand()
	ResetCompute()
	ResetMlflowExperimentName()
	ResetSource()
	ResetTrainingScriptPath()
	ResetYamlParameters()
	ResetYamlParametersFilePath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskGenAiComputeTaskOutputReference
type jsiiProxy_JobTaskGenAiComputeTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) Command() *string {
	var returns *string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) CommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) Compute() JobTaskGenAiComputeTaskComputeOutputReference {
	var returns JobTaskGenAiComputeTaskComputeOutputReference
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ComputeInput() *JobTaskGenAiComputeTaskCompute {
	var returns *JobTaskGenAiComputeTaskCompute
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) DlRuntimeImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dlRuntimeImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) DlRuntimeImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dlRuntimeImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) InternalValue() *JobTaskGenAiComputeTask {
	var returns *JobTaskGenAiComputeTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) MlflowExperimentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowExperimentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) MlflowExperimentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowExperimentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) TrainingScriptPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingScriptPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) TrainingScriptPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingScriptPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) YamlParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yamlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) YamlParametersFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yamlParametersFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) YamlParametersFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yamlParametersFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) YamlParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yamlParametersInput",
		&returns,
	)
	return returns
}


func NewJobTaskGenAiComputeTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskGenAiComputeTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskGenAiComputeTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskGenAiComputeTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskGenAiComputeTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskGenAiComputeTaskOutputReference_Override(j JobTaskGenAiComputeTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskGenAiComputeTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetCommand(val *string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetDlRuntimeImage(val *string) {
	if err := j.validateSetDlRuntimeImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dlRuntimeImage",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetInternalValue(val *JobTaskGenAiComputeTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetMlflowExperimentName(val *string) {
	if err := j.validateSetMlflowExperimentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlflowExperimentName",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetTrainingScriptPath(val *string) {
	if err := j.validateSetTrainingScriptPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingScriptPath",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetYamlParameters(val *string) {
	if err := j.validateSetYamlParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yamlParameters",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference)SetYamlParametersFilePath(val *string) {
	if err := j.validateSetYamlParametersFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yamlParametersFilePath",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) PutCompute(value *JobTaskGenAiComputeTaskCompute) {
	if err := j.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCompute",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		j,
		"resetCommand",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ResetCompute() {
	_jsii_.InvokeVoid(
		j,
		"resetCompute",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ResetMlflowExperimentName() {
	_jsii_.InvokeVoid(
		j,
		"resetMlflowExperimentName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		j,
		"resetSource",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ResetTrainingScriptPath() {
	_jsii_.InvokeVoid(
		j,
		"resetTrainingScriptPath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ResetYamlParameters() {
	_jsii_.InvokeVoid(
		j,
		"resetYamlParameters",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ResetYamlParametersFilePath() {
	_jsii_.InvokeVoid(
		j,
		"resetYamlParametersFilePath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

