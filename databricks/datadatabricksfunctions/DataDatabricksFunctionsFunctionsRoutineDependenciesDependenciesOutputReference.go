// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksfunctions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference interface {
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
	Connection() DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnectionOutputReference
	ConnectionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Credential() DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredentialOutputReference
	CredentialInput() interface{}
	// Experimental.
	Fqn() *string
	Function() DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunctionOutputReference
	FunctionInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Table() DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTableOutputReference
	TableInput() interface{}
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
	PutConnection(value *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnection)
	PutCredential(value *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredential)
	PutFunction(value *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunction)
	PutTable(value *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTable)
	ResetConnection()
	ResetCredential()
	ResetFunction()
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

// The jsii proxy struct for DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference
type jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) Connection() DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnectionOutputReference {
	var returns DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) Credential() DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredentialOutputReference {
	var returns DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredentialOutputReference
	_jsii_.Get(
		j,
		"credential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) CredentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) Function() DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunctionOutputReference {
	var returns DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunctionOutputReference
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) FunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) Table() DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTableOutputReference {
	var returns DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksFunctions.DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference_Override(d DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksFunctions.DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) PutConnection(value *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnection) {
	if err := d.validatePutConnectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) PutCredential(value *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredential) {
	if err := d.validatePutCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCredential",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) PutFunction(value *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunction) {
	if err := d.validatePutFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFunction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) PutTable(value *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTable) {
	if err := d.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ResetCredential() {
	_jsii_.InvokeVoid(
		d,
		"resetCredential",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ResetFunction() {
	_jsii_.InvokeVoid(
		d,
		"resetFunction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

