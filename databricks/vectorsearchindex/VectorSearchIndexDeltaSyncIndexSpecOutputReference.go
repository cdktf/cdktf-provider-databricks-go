// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/vectorsearchindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VectorSearchIndexDeltaSyncIndexSpecOutputReference interface {
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
	EmbeddingSourceColumns() VectorSearchIndexDeltaSyncIndexSpecEmbeddingSourceColumnsList
	EmbeddingSourceColumnsInput() interface{}
	EmbeddingVectorColumns() VectorSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList
	EmbeddingVectorColumnsInput() interface{}
	EmbeddingWritebackTable() *string
	SetEmbeddingWritebackTable(val *string)
	EmbeddingWritebackTableInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *VectorSearchIndexDeltaSyncIndexSpec
	SetInternalValue(val *VectorSearchIndexDeltaSyncIndexSpec)
	PipelineId() *string
	PipelineType() *string
	SetPipelineType(val *string)
	PipelineTypeInput() *string
	SourceTable() *string
	SetSourceTable(val *string)
	SourceTableInput() *string
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
	PutEmbeddingSourceColumns(value interface{})
	PutEmbeddingVectorColumns(value interface{})
	ResetEmbeddingSourceColumns()
	ResetEmbeddingVectorColumns()
	ResetEmbeddingWritebackTable()
	ResetPipelineType()
	ResetSourceTable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VectorSearchIndexDeltaSyncIndexSpecOutputReference
type jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingSourceColumns() VectorSearchIndexDeltaSyncIndexSpecEmbeddingSourceColumnsList {
	var returns VectorSearchIndexDeltaSyncIndexSpecEmbeddingSourceColumnsList
	_jsii_.Get(
		j,
		"embeddingSourceColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingSourceColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingSourceColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingVectorColumns() VectorSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList {
	var returns VectorSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList
	_jsii_.Get(
		j,
		"embeddingVectorColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingVectorColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingVectorColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingWritebackTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingWritebackTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingWritebackTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingWritebackTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) InternalValue() *VectorSearchIndexDeltaSyncIndexSpec {
	var returns *VectorSearchIndexDeltaSyncIndexSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) PipelineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) PipelineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) SourceTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) SourceTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVectorSearchIndexDeltaSyncIndexSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VectorSearchIndexDeltaSyncIndexSpecOutputReference {
	_init_.Initialize()

	if err := validateNewVectorSearchIndexDeltaSyncIndexSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.vectorSearchIndex.VectorSearchIndexDeltaSyncIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVectorSearchIndexDeltaSyncIndexSpecOutputReference_Override(v VectorSearchIndexDeltaSyncIndexSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.vectorSearchIndex.VectorSearchIndexDeltaSyncIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference)SetEmbeddingWritebackTable(val *string) {
	if err := j.validateSetEmbeddingWritebackTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingWritebackTable",
		val,
	)
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference)SetInternalValue(val *VectorSearchIndexDeltaSyncIndexSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference)SetPipelineType(val *string) {
	if err := j.validateSetPipelineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineType",
		val,
	)
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference)SetSourceTable(val *string) {
	if err := j.validateSetSourceTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTable",
		val,
	)
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) PutEmbeddingSourceColumns(value interface{}) {
	if err := v.validatePutEmbeddingSourceColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEmbeddingSourceColumns",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) PutEmbeddingVectorColumns(value interface{}) {
	if err := v.validatePutEmbeddingVectorColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEmbeddingVectorColumns",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ResetEmbeddingSourceColumns() {
	_jsii_.InvokeVoid(
		v,
		"resetEmbeddingSourceColumns",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ResetEmbeddingVectorColumns() {
	_jsii_.InvokeVoid(
		v,
		"resetEmbeddingVectorColumns",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ResetEmbeddingWritebackTable() {
	_jsii_.InvokeVoid(
		v,
		"resetEmbeddingWritebackTable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ResetPipelineType() {
	_jsii_.InvokeVoid(
		v,
		"resetPipelineType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ResetSourceTable() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceTable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorSearchIndexDeltaSyncIndexSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

