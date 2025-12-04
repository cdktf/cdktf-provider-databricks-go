// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksserviceprincipals/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksServicePrincipalsServicePrincipalsOutputReference interface {
	cdktf.ComplexObject
	AclPrincipalId() *string
	SetAclPrincipalId(val *string)
	AclPrincipalIdInput() *string
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	// Experimental.
	Fqn() *string
	Home() *string
	SetHome(val *string)
	HomeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Repos() *string
	SetRepos(val *string)
	ReposInput() *string
	ScimId() *string
	SetScimId(val *string)
	ScimIdInput() *string
	SpId() *string
	SetSpId(val *string)
	SpIdInput() *string
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
	ResetAclPrincipalId()
	ResetActive()
	ResetApplicationId()
	ResetDisplayName()
	ResetExternalId()
	ResetHome()
	ResetId()
	ResetRepos()
	ResetScimId()
	ResetSpId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksServicePrincipalsServicePrincipalsOutputReference
type jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) AclPrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclPrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) AclPrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclPrincipalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) Home() *string {
	var returns *string
	_jsii_.Get(
		j,
		"home",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) HomeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) Repos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ReposInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reposInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ScimId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scimId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ScimIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scimIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) SpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) SpIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksServicePrincipalsServicePrincipalsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksServicePrincipalsServicePrincipalsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksServicePrincipalsServicePrincipalsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServicePrincipals.DataDatabricksServicePrincipalsServicePrincipalsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksServicePrincipalsServicePrincipalsOutputReference_Override(d DataDatabricksServicePrincipalsServicePrincipalsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServicePrincipals.DataDatabricksServicePrincipalsServicePrincipalsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetAclPrincipalId(val *string) {
	if err := j.validateSetAclPrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aclPrincipalId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetActive(val interface{}) {
	if err := j.validateSetActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetHome(val *string) {
	if err := j.validateSetHomeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"home",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetRepos(val *string) {
	if err := j.validateSetReposParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repos",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetScimId(val *string) {
	if err := j.validateSetScimIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scimId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetSpId(val *string) {
	if err := j.validateSetSpIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetAclPrincipalId() {
	_jsii_.InvokeVoid(
		d,
		"resetAclPrincipalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetActive() {
	_jsii_.InvokeVoid(
		d,
		"resetActive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetApplicationId() {
	_jsii_.InvokeVoid(
		d,
		"resetApplicationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetHome() {
	_jsii_.InvokeVoid(
		d,
		"resetHome",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetRepos() {
	_jsii_.InvokeVoid(
		d,
		"resetRepos",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetScimId() {
	_jsii_.InvokeVoid(
		d,
		"resetScimId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ResetSpId() {
	_jsii_.InvokeVoid(
		d,
		"resetSpId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalsServicePrincipalsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

