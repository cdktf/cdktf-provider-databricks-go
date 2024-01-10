// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/pipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineLibraryOutputReference interface {
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
	File() PipelineLibraryFileOutputReference
	FileInput() *PipelineLibraryFile
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Jar() *string
	SetJar(val *string)
	JarInput() *string
	Maven() PipelineLibraryMavenOutputReference
	MavenInput() *PipelineLibraryMaven
	Notebook() PipelineLibraryNotebookOutputReference
	NotebookInput() *PipelineLibraryNotebook
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Whl() *string
	SetWhl(val *string)
	WhlInput() *string
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
	PutFile(value *PipelineLibraryFile)
	PutMaven(value *PipelineLibraryMaven)
	PutNotebook(value *PipelineLibraryNotebook)
	ResetFile()
	ResetJar()
	ResetMaven()
	ResetNotebook()
	ResetWhl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineLibraryOutputReference
type jsiiProxy_PipelineLibraryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipelineLibraryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) File() PipelineLibraryFileOutputReference {
	var returns PipelineLibraryFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) FileInput() *PipelineLibraryFile {
	var returns *PipelineLibraryFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) Jar() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) JarInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) Maven() PipelineLibraryMavenOutputReference {
	var returns PipelineLibraryMavenOutputReference
	_jsii_.Get(
		j,
		"maven",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) MavenInput() *PipelineLibraryMaven {
	var returns *PipelineLibraryMaven
	_jsii_.Get(
		j,
		"mavenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) Notebook() PipelineLibraryNotebookOutputReference {
	var returns PipelineLibraryNotebookOutputReference
	_jsii_.Get(
		j,
		"notebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) NotebookInput() *PipelineLibraryNotebook {
	var returns *PipelineLibraryNotebook
	_jsii_.Get(
		j,
		"notebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) Whl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineLibraryOutputReference) WhlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whlInput",
		&returns,
	)
	return returns
}


func NewPipelineLibraryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PipelineLibraryOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineLibraryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineLibraryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPipelineLibraryOutputReference_Override(p PipelineLibraryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PipelineLibraryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineLibraryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineLibraryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineLibraryOutputReference)SetJar(val *string) {
	if err := j.validateSetJarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jar",
		val,
	)
}

func (j *jsiiProxy_PipelineLibraryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineLibraryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipelineLibraryOutputReference)SetWhl(val *string) {
	if err := j.validateSetWhlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whl",
		val,
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineLibraryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineLibraryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) PutFile(value *PipelineLibraryFile) {
	if err := p.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) PutMaven(value *PipelineLibraryMaven) {
	if err := p.validatePutMavenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMaven",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) PutNotebook(value *PipelineLibraryNotebook) {
	if err := p.validatePutNotebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNotebook",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		p,
		"resetFile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) ResetJar() {
	_jsii_.InvokeVoid(
		p,
		"resetJar",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) ResetMaven() {
	_jsii_.InvokeVoid(
		p,
		"resetMaven",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) ResetNotebook() {
	_jsii_.InvokeVoid(
		p,
		"resetNotebook",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) ResetWhl() {
	_jsii_.InvokeVoid(
		p,
		"resetWhl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineLibraryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineLibraryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

