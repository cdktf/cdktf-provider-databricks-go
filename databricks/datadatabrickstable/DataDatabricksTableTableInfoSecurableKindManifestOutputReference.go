// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickstable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksTableTableInfoSecurableKindManifestOutputReference interface {
	cdktf.ComplexObject
	AssignablePrivileges() *[]*string
	SetAssignablePrivileges(val *[]*string)
	AssignablePrivilegesInput() *[]*string
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	CapabilitiesInput() *[]*string
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
	InternalValue() *DataDatabricksTableTableInfoSecurableKindManifest
	SetInternalValue(val *DataDatabricksTableTableInfoSecurableKindManifest)
	Options() DataDatabricksTableTableInfoSecurableKindManifestOptionsList
	OptionsInput() interface{}
	SecurableKind() *string
	SetSecurableKind(val *string)
	SecurableKindInput() *string
	SecurableType() *string
	SetSecurableType(val *string)
	SecurableTypeInput() *string
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
	PutOptions(value interface{})
	ResetAssignablePrivileges()
	ResetCapabilities()
	ResetOptions()
	ResetSecurableKind()
	ResetSecurableType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksTableTableInfoSecurableKindManifestOutputReference
type jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) AssignablePrivileges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assignablePrivileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) AssignablePrivilegesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assignablePrivilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) CapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) InternalValue() *DataDatabricksTableTableInfoSecurableKindManifest {
	var returns *DataDatabricksTableTableInfoSecurableKindManifest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) Options() DataDatabricksTableTableInfoSecurableKindManifestOptionsList {
	var returns DataDatabricksTableTableInfoSecurableKindManifestOptionsList
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) OptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) SecurableKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) SecurableKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) SecurableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) SecurableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksTableTableInfoSecurableKindManifestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksTableTableInfoSecurableKindManifestOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksTableTableInfoSecurableKindManifestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksTable.DataDatabricksTableTableInfoSecurableKindManifestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksTableTableInfoSecurableKindManifestOutputReference_Override(d DataDatabricksTableTableInfoSecurableKindManifestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksTable.DataDatabricksTableTableInfoSecurableKindManifestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetAssignablePrivileges(val *[]*string) {
	if err := j.validateSetAssignablePrivilegesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignablePrivileges",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetCapabilities(val *[]*string) {
	if err := j.validateSetCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetInternalValue(val *DataDatabricksTableTableInfoSecurableKindManifest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetSecurableKind(val *string) {
	if err := j.validateSetSecurableKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securableKind",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetSecurableType(val *string) {
	if err := j.validateSetSecurableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securableType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) PutOptions(value interface{}) {
	if err := d.validatePutOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ResetAssignablePrivileges() {
	_jsii_.InvokeVoid(
		d,
		"resetAssignablePrivileges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		d,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ResetOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ResetSecurableKind() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurableKind",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ResetSecurableType() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurableType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

