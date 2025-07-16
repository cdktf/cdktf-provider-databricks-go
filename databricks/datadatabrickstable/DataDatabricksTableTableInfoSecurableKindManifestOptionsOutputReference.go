// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickstable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference interface {
	cdktf.ComplexObject
	AllowedValues() *[]*string
	SetAllowedValues(val *[]*string)
	AllowedValuesInput() *[]*string
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
	DefaultValue() *string
	SetDefaultValue(val *string)
	DefaultValueInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	Hint() *string
	SetHint(val *string)
	HintInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsCopiable() interface{}
	SetIsCopiable(val interface{})
	IsCopiableInput() interface{}
	IsCreatable() interface{}
	SetIsCreatable(val interface{})
	IsCreatableInput() interface{}
	IsHidden() interface{}
	SetIsHidden(val interface{})
	IsHiddenInput() interface{}
	IsLoggable() interface{}
	SetIsLoggable(val interface{})
	IsLoggableInput() interface{}
	IsRequired() interface{}
	SetIsRequired(val interface{})
	IsRequiredInput() interface{}
	IsSecret() interface{}
	SetIsSecret(val interface{})
	IsSecretInput() interface{}
	IsUpdatable() interface{}
	SetIsUpdatable(val interface{})
	IsUpdatableInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	OauthStage() *string
	SetOauthStage(val *string)
	OauthStageInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAllowedValues()
	ResetDefaultValue()
	ResetDescription()
	ResetHint()
	ResetIsCopiable()
	ResetIsCreatable()
	ResetIsHidden()
	ResetIsLoggable()
	ResetIsRequired()
	ResetIsSecret()
	ResetIsUpdatable()
	ResetName()
	ResetOauthStage()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference
type jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) AllowedValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) AllowedValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) Hint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) HintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsCopiable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCopiable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsCopiableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCopiableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsCreatable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCreatable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsCreatableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCreatableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsHidden() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHidden",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsHiddenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHiddenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsLoggable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLoggable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsLoggableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLoggableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsUpdatable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isUpdatable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) IsUpdatableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isUpdatableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) OauthStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) OauthStageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksTable.DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference_Override(d DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksTable.DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetAllowedValues(val *[]*string) {
	if err := j.validateSetAllowedValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedValues",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetHint(val *string) {
	if err := j.validateSetHintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hint",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetIsCopiable(val interface{}) {
	if err := j.validateSetIsCopiableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCopiable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetIsCreatable(val interface{}) {
	if err := j.validateSetIsCreatableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCreatable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetIsHidden(val interface{}) {
	if err := j.validateSetIsHiddenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isHidden",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetIsLoggable(val interface{}) {
	if err := j.validateSetIsLoggableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLoggable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetIsRequired(val interface{}) {
	if err := j.validateSetIsRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRequired",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetIsSecret(val interface{}) {
	if err := j.validateSetIsSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecret",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetIsUpdatable(val interface{}) {
	if err := j.validateSetIsUpdatableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isUpdatable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetOauthStage(val *string) {
	if err := j.validateSetOauthStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthStage",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetAllowedValues() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetHint() {
	_jsii_.InvokeVoid(
		d,
		"resetHint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetIsCopiable() {
	_jsii_.InvokeVoid(
		d,
		"resetIsCopiable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetIsCreatable() {
	_jsii_.InvokeVoid(
		d,
		"resetIsCreatable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetIsHidden() {
	_jsii_.InvokeVoid(
		d,
		"resetIsHidden",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetIsLoggable() {
	_jsii_.InvokeVoid(
		d,
		"resetIsLoggable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetIsRequired() {
	_jsii_.InvokeVoid(
		d,
		"resetIsRequired",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetIsSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetIsSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetIsUpdatable() {
	_jsii_.InvokeVoid(
		d,
		"resetIsUpdatable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetOauthStage() {
	_jsii_.InvokeVoid(
		d,
		"resetOauthStage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoSecurableKindManifestOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

