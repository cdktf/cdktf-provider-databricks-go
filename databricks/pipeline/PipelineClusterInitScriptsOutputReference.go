// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/pipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineClusterInitScriptsOutputReference interface {
	cdktf.ComplexObject
	Abfss() PipelineClusterInitScriptsAbfssOutputReference
	AbfssInput() *PipelineClusterInitScriptsAbfss
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
	Dbfs() PipelineClusterInitScriptsDbfsOutputReference
	DbfsInput() *PipelineClusterInitScriptsDbfs
	File() PipelineClusterInitScriptsFileOutputReference
	FileInput() *PipelineClusterInitScriptsFile
	// Experimental.
	Fqn() *string
	Gcs() PipelineClusterInitScriptsGcsOutputReference
	GcsInput() *PipelineClusterInitScriptsGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3() PipelineClusterInitScriptsS3OutputReference
	S3Input() *PipelineClusterInitScriptsS3
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() PipelineClusterInitScriptsVolumesOutputReference
	VolumesInput() *PipelineClusterInitScriptsVolumes
	Workspace() PipelineClusterInitScriptsWorkspaceOutputReference
	WorkspaceInput() *PipelineClusterInitScriptsWorkspace
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
	PutAbfss(value *PipelineClusterInitScriptsAbfss)
	PutDbfs(value *PipelineClusterInitScriptsDbfs)
	PutFile(value *PipelineClusterInitScriptsFile)
	PutGcs(value *PipelineClusterInitScriptsGcs)
	PutS3(value *PipelineClusterInitScriptsS3)
	PutVolumes(value *PipelineClusterInitScriptsVolumes)
	PutWorkspace(value *PipelineClusterInitScriptsWorkspace)
	ResetAbfss()
	ResetDbfs()
	ResetFile()
	ResetGcs()
	ResetS3()
	ResetVolumes()
	ResetWorkspace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineClusterInitScriptsOutputReference
type jsiiProxy_PipelineClusterInitScriptsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) Abfss() PipelineClusterInitScriptsAbfssOutputReference {
	var returns PipelineClusterInitScriptsAbfssOutputReference
	_jsii_.Get(
		j,
		"abfss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) AbfssInput() *PipelineClusterInitScriptsAbfss {
	var returns *PipelineClusterInitScriptsAbfss
	_jsii_.Get(
		j,
		"abfssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) Dbfs() PipelineClusterInitScriptsDbfsOutputReference {
	var returns PipelineClusterInitScriptsDbfsOutputReference
	_jsii_.Get(
		j,
		"dbfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) DbfsInput() *PipelineClusterInitScriptsDbfs {
	var returns *PipelineClusterInitScriptsDbfs
	_jsii_.Get(
		j,
		"dbfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) File() PipelineClusterInitScriptsFileOutputReference {
	var returns PipelineClusterInitScriptsFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) FileInput() *PipelineClusterInitScriptsFile {
	var returns *PipelineClusterInitScriptsFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) Gcs() PipelineClusterInitScriptsGcsOutputReference {
	var returns PipelineClusterInitScriptsGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) GcsInput() *PipelineClusterInitScriptsGcs {
	var returns *PipelineClusterInitScriptsGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) S3() PipelineClusterInitScriptsS3OutputReference {
	var returns PipelineClusterInitScriptsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) S3Input() *PipelineClusterInitScriptsS3 {
	var returns *PipelineClusterInitScriptsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) Volumes() PipelineClusterInitScriptsVolumesOutputReference {
	var returns PipelineClusterInitScriptsVolumesOutputReference
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) VolumesInput() *PipelineClusterInitScriptsVolumes {
	var returns *PipelineClusterInitScriptsVolumes
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) Workspace() PipelineClusterInitScriptsWorkspaceOutputReference {
	var returns PipelineClusterInitScriptsWorkspaceOutputReference
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference) WorkspaceInput() *PipelineClusterInitScriptsWorkspace {
	var returns *PipelineClusterInitScriptsWorkspace
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}


func NewPipelineClusterInitScriptsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PipelineClusterInitScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineClusterInitScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineClusterInitScriptsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPipelineClusterInitScriptsOutputReference_Override(p PipelineClusterInitScriptsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterInitScriptsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) PutAbfss(value *PipelineClusterInitScriptsAbfss) {
	if err := p.validatePutAbfssParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAbfss",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) PutDbfs(value *PipelineClusterInitScriptsDbfs) {
	if err := p.validatePutDbfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDbfs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) PutFile(value *PipelineClusterInitScriptsFile) {
	if err := p.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) PutGcs(value *PipelineClusterInitScriptsGcs) {
	if err := p.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGcs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) PutS3(value *PipelineClusterInitScriptsS3) {
	if err := p.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putS3",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) PutVolumes(value *PipelineClusterInitScriptsVolumes) {
	if err := p.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVolumes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) PutWorkspace(value *PipelineClusterInitScriptsWorkspace) {
	if err := p.validatePutWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWorkspace",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ResetAbfss() {
	_jsii_.InvokeVoid(
		p,
		"resetAbfss",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ResetDbfs() {
	_jsii_.InvokeVoid(
		p,
		"resetDbfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		p,
		"resetFile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		p,
		"resetGcs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		p,
		"resetS3",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ResetWorkspace() {
	_jsii_.InvokeVoid(
		p,
		"resetWorkspace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterInitScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

