// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/pipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference interface {
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
	ExcludeColumns() *[]*string
	SetExcludeColumns(val *[]*string)
	ExcludeColumnsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeColumns() *[]*string
	SetIncludeColumns(val *[]*string)
	IncludeColumnsInput() *[]*string
	InternalValue() *PipelineIngestionDefinitionObjectsTableTableConfiguration
	SetInternalValue(val *PipelineIngestionDefinitionObjectsTableTableConfiguration)
	PrimaryKeys() *[]*string
	SetPrimaryKeys(val *[]*string)
	PrimaryKeysInput() *[]*string
	SalesforceIncludeFormulaFields() interface{}
	SetSalesforceIncludeFormulaFields(val interface{})
	SalesforceIncludeFormulaFieldsInput() interface{}
	ScdType() *string
	SetScdType(val *string)
	ScdTypeInput() *string
	SequenceBy() *[]*string
	SetSequenceBy(val *[]*string)
	SequenceByInput() *[]*string
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
	ResetExcludeColumns()
	ResetIncludeColumns()
	ResetPrimaryKeys()
	ResetSalesforceIncludeFormulaFields()
	ResetScdType()
	ResetSequenceBy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ExcludeColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ExcludeColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) IncludeColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) IncludeColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsTableTableConfiguration {
	var returns *PipelineIngestionDefinitionObjectsTableTableConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) PrimaryKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) PrimaryKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SalesforceIncludeFormulaFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceIncludeFormulaFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SalesforceIncludeFormulaFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceIncludeFormulaFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ScdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ScdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SequenceBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sequenceBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SequenceByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sequenceByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsTableTableConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference_Override(p PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetExcludeColumns(val *[]*string) {
	if err := j.validateSetExcludeColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeColumns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetIncludeColumns(val *[]*string) {
	if err := j.validateSetIncludeColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeColumns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsTableTableConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetPrimaryKeys(val *[]*string) {
	if err := j.validateSetPrimaryKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeys",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetSalesforceIncludeFormulaFields(val interface{}) {
	if err := j.validateSetSalesforceIncludeFormulaFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"salesforceIncludeFormulaFields",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetScdType(val *string) {
	if err := j.validateSetScdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scdType",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetSequenceBy(val *[]*string) {
	if err := j.validateSetSequenceByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sequenceBy",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetExcludeColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludeColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetIncludeColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetPrimaryKeys() {
	_jsii_.InvokeVoid(
		p,
		"resetPrimaryKeys",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetSalesforceIncludeFormulaFields() {
	_jsii_.InvokeVoid(
		p,
		"resetSalesforceIncludeFormulaFields",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetScdType() {
	_jsii_.InvokeVoid(
		p,
		"resetScdType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetSequenceBy() {
	_jsii_.InvokeVoid(
		p,
		"resetSequenceBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

