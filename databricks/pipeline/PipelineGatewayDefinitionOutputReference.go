// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/pipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineGatewayDefinitionOutputReference interface {
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
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	ConnectionName() *string
	SetConnectionName(val *string)
	ConnectionNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GatewayStorageCatalog() *string
	SetGatewayStorageCatalog(val *string)
	GatewayStorageCatalogInput() *string
	GatewayStorageName() *string
	SetGatewayStorageName(val *string)
	GatewayStorageNameInput() *string
	GatewayStorageSchema() *string
	SetGatewayStorageSchema(val *string)
	GatewayStorageSchemaInput() *string
	InternalValue() *PipelineGatewayDefinition
	SetInternalValue(val *PipelineGatewayDefinition)
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
	ResetConnectionId()
	ResetGatewayStorageName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineGatewayDefinitionOutputReference
type jsiiProxy_PipelineGatewayDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) ConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) GatewayStorageCatalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayStorageCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) GatewayStorageCatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayStorageCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) GatewayStorageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayStorageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) GatewayStorageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayStorageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) GatewayStorageSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayStorageSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) GatewayStorageSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayStorageSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) InternalValue() *PipelineGatewayDefinition {
	var returns *PipelineGatewayDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineGatewayDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipelineGatewayDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineGatewayDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineGatewayDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineGatewayDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineGatewayDefinitionOutputReference_Override(p PipelineGatewayDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineGatewayDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetConnectionName(val *string) {
	if err := j.validateSetConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionName",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetGatewayStorageCatalog(val *string) {
	if err := j.validateSetGatewayStorageCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayStorageCatalog",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetGatewayStorageName(val *string) {
	if err := j.validateSetGatewayStorageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayStorageName",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetGatewayStorageSchema(val *string) {
	if err := j.validateSetGatewayStorageSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayStorageSchema",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetInternalValue(val *PipelineGatewayDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineGatewayDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) ResetConnectionId() {
	_jsii_.InvokeVoid(
		p,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) ResetGatewayStorageName() {
	_jsii_.InvokeVoid(
		p,
		"resetGatewayStorageName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineGatewayDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

