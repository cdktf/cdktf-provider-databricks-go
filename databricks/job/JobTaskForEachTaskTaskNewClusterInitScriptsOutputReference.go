// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference interface {
	cdktf.ComplexObject
	Abfss() JobTaskForEachTaskTaskNewClusterInitScriptsAbfssOutputReference
	AbfssInput() *JobTaskForEachTaskTaskNewClusterInitScriptsAbfss
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
	Dbfs() JobTaskForEachTaskTaskNewClusterInitScriptsDbfsOutputReference
	DbfsInput() *JobTaskForEachTaskTaskNewClusterInitScriptsDbfs
	File() JobTaskForEachTaskTaskNewClusterInitScriptsFileOutputReference
	FileInput() *JobTaskForEachTaskTaskNewClusterInitScriptsFile
	// Experimental.
	Fqn() *string
	Gcs() JobTaskForEachTaskTaskNewClusterInitScriptsGcsOutputReference
	GcsInput() *JobTaskForEachTaskTaskNewClusterInitScriptsGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3() JobTaskForEachTaskTaskNewClusterInitScriptsS3OutputReference
	S3Input() *JobTaskForEachTaskTaskNewClusterInitScriptsS3
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() JobTaskForEachTaskTaskNewClusterInitScriptsVolumesOutputReference
	VolumesInput() *JobTaskForEachTaskTaskNewClusterInitScriptsVolumes
	Workspace() JobTaskForEachTaskTaskNewClusterInitScriptsWorkspaceOutputReference
	WorkspaceInput() *JobTaskForEachTaskTaskNewClusterInitScriptsWorkspace
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
	PutAbfss(value *JobTaskForEachTaskTaskNewClusterInitScriptsAbfss)
	PutDbfs(value *JobTaskForEachTaskTaskNewClusterInitScriptsDbfs)
	PutFile(value *JobTaskForEachTaskTaskNewClusterInitScriptsFile)
	PutGcs(value *JobTaskForEachTaskTaskNewClusterInitScriptsGcs)
	PutS3(value *JobTaskForEachTaskTaskNewClusterInitScriptsS3)
	PutVolumes(value *JobTaskForEachTaskTaskNewClusterInitScriptsVolumes)
	PutWorkspace(value *JobTaskForEachTaskTaskNewClusterInitScriptsWorkspace)
	ResetAbfss()
	ResetDbfs()
	ResetFile()
	ResetGcs()
	ResetS3()
	ResetVolumes()
	ResetWorkspace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference
type jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Abfss() JobTaskForEachTaskTaskNewClusterInitScriptsAbfssOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterInitScriptsAbfssOutputReference
	_jsii_.Get(
		j,
		"abfss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) AbfssInput() *JobTaskForEachTaskTaskNewClusterInitScriptsAbfss {
	var returns *JobTaskForEachTaskTaskNewClusterInitScriptsAbfss
	_jsii_.Get(
		j,
		"abfssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Dbfs() JobTaskForEachTaskTaskNewClusterInitScriptsDbfsOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterInitScriptsDbfsOutputReference
	_jsii_.Get(
		j,
		"dbfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) DbfsInput() *JobTaskForEachTaskTaskNewClusterInitScriptsDbfs {
	var returns *JobTaskForEachTaskTaskNewClusterInitScriptsDbfs
	_jsii_.Get(
		j,
		"dbfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) File() JobTaskForEachTaskTaskNewClusterInitScriptsFileOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterInitScriptsFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) FileInput() *JobTaskForEachTaskTaskNewClusterInitScriptsFile {
	var returns *JobTaskForEachTaskTaskNewClusterInitScriptsFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Gcs() JobTaskForEachTaskTaskNewClusterInitScriptsGcsOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterInitScriptsGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GcsInput() *JobTaskForEachTaskTaskNewClusterInitScriptsGcs {
	var returns *JobTaskForEachTaskTaskNewClusterInitScriptsGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) S3() JobTaskForEachTaskTaskNewClusterInitScriptsS3OutputReference {
	var returns JobTaskForEachTaskTaskNewClusterInitScriptsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) S3Input() *JobTaskForEachTaskTaskNewClusterInitScriptsS3 {
	var returns *JobTaskForEachTaskTaskNewClusterInitScriptsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Volumes() JobTaskForEachTaskTaskNewClusterInitScriptsVolumesOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterInitScriptsVolumesOutputReference
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) VolumesInput() *JobTaskForEachTaskTaskNewClusterInitScriptsVolumes {
	var returns *JobTaskForEachTaskTaskNewClusterInitScriptsVolumes
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Workspace() JobTaskForEachTaskTaskNewClusterInitScriptsWorkspaceOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterInitScriptsWorkspaceOutputReference
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) WorkspaceInput() *JobTaskForEachTaskTaskNewClusterInitScriptsWorkspace {
	var returns *JobTaskForEachTaskTaskNewClusterInitScriptsWorkspace
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskNewClusterInitScriptsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskNewClusterInitScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskNewClusterInitScriptsOutputReference_Override(j JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutAbfss(value *JobTaskForEachTaskTaskNewClusterInitScriptsAbfss) {
	if err := j.validatePutAbfssParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAbfss",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutDbfs(value *JobTaskForEachTaskTaskNewClusterInitScriptsDbfs) {
	if err := j.validatePutDbfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDbfs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutFile(value *JobTaskForEachTaskTaskNewClusterInitScriptsFile) {
	if err := j.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFile",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutGcs(value *JobTaskForEachTaskTaskNewClusterInitScriptsGcs) {
	if err := j.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGcs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutS3(value *JobTaskForEachTaskTaskNewClusterInitScriptsS3) {
	if err := j.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putS3",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutVolumes(value *JobTaskForEachTaskTaskNewClusterInitScriptsVolumes) {
	if err := j.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVolumes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutWorkspace(value *JobTaskForEachTaskTaskNewClusterInitScriptsWorkspace) {
	if err := j.validatePutWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWorkspace",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetAbfss() {
	_jsii_.InvokeVoid(
		j,
		"resetAbfss",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetDbfs() {
	_jsii_.InvokeVoid(
		j,
		"resetDbfs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		j,
		"resetFile",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		j,
		"resetGcs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		j,
		"resetS3",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		j,
		"resetVolumes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetWorkspace() {
	_jsii_.InvokeVoid(
		j,
		"resetWorkspace",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

