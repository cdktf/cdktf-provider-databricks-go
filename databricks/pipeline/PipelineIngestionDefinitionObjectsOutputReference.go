// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/pipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineIngestionDefinitionObjectsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Report() PipelineIngestionDefinitionObjectsReportOutputReference
	ReportInput() *PipelineIngestionDefinitionObjectsReport
	Schema() PipelineIngestionDefinitionObjectsSchemaOutputReference
	SchemaInput() *PipelineIngestionDefinitionObjectsSchema
	Table() PipelineIngestionDefinitionObjectsTableOutputReference
	TableInput() *PipelineIngestionDefinitionObjectsTable
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
	PutReport(value *PipelineIngestionDefinitionObjectsReport)
	PutSchema(value *PipelineIngestionDefinitionObjectsSchema)
	PutTable(value *PipelineIngestionDefinitionObjectsTable)
	ResetReport()
	ResetSchema()
	ResetTable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) Report() PipelineIngestionDefinitionObjectsReportOutputReference {
	var returns PipelineIngestionDefinitionObjectsReportOutputReference
	_jsii_.Get(
		j,
		"report",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) ReportInput() *PipelineIngestionDefinitionObjectsReport {
	var returns *PipelineIngestionDefinitionObjectsReport
	_jsii_.Get(
		j,
		"reportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) Schema() PipelineIngestionDefinitionObjectsSchemaOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) SchemaInput() *PipelineIngestionDefinitionObjectsSchema {
	var returns *PipelineIngestionDefinitionObjectsSchema
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) Table() PipelineIngestionDefinitionObjectsTableOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) TableInput() *PipelineIngestionDefinitionObjectsTable {
	var returns *PipelineIngestionDefinitionObjectsTable
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PipelineIngestionDefinitionObjectsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsOutputReference_Override(p PipelineIngestionDefinitionObjectsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) PutReport(value *PipelineIngestionDefinitionObjectsReport) {
	if err := p.validatePutReportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putReport",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) PutSchema(value *PipelineIngestionDefinitionObjectsSchema) {
	if err := p.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSchema",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) PutTable(value *PipelineIngestionDefinitionObjectsTable) {
	if err := p.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTable",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) ResetReport() {
	_jsii_.InvokeVoid(
		p,
		"resetReport",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		p,
		"resetSchema",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		p,
		"resetTable",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

