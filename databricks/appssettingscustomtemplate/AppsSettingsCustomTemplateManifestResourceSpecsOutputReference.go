// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appssettingscustomtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/appssettingscustomtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsSettingsCustomTemplateManifestResourceSpecsOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobSpec() AppsSettingsCustomTemplateManifestResourceSpecsJobSpecOutputReference
	JobSpecInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	SecretSpec() AppsSettingsCustomTemplateManifestResourceSpecsSecretSpecOutputReference
	SecretSpecInput() interface{}
	ServingEndpointSpec() AppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpecOutputReference
	ServingEndpointSpecInput() interface{}
	SqlWarehouseSpec() AppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpecOutputReference
	SqlWarehouseSpecInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UcSecurableSpec() AppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpecOutputReference
	UcSecurableSpecInput() interface{}
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
	PutJobSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsJobSpec)
	PutSecretSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsSecretSpec)
	PutServingEndpointSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpec)
	PutSqlWarehouseSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpec)
	PutUcSecurableSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpec)
	ResetDescription()
	ResetJobSpec()
	ResetSecretSpec()
	ResetServingEndpointSpec()
	ResetSqlWarehouseSpec()
	ResetUcSecurableSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsSettingsCustomTemplateManifestResourceSpecsOutputReference
type jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) JobSpec() AppsSettingsCustomTemplateManifestResourceSpecsJobSpecOutputReference {
	var returns AppsSettingsCustomTemplateManifestResourceSpecsJobSpecOutputReference
	_jsii_.Get(
		j,
		"jobSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) JobSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) SecretSpec() AppsSettingsCustomTemplateManifestResourceSpecsSecretSpecOutputReference {
	var returns AppsSettingsCustomTemplateManifestResourceSpecsSecretSpecOutputReference
	_jsii_.Get(
		j,
		"secretSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) SecretSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ServingEndpointSpec() AppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpecOutputReference {
	var returns AppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpecOutputReference
	_jsii_.Get(
		j,
		"servingEndpointSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ServingEndpointSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingEndpointSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) SqlWarehouseSpec() AppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpecOutputReference {
	var returns AppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpecOutputReference
	_jsii_.Get(
		j,
		"sqlWarehouseSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) SqlWarehouseSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlWarehouseSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) UcSecurableSpec() AppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpecOutputReference {
	var returns AppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpecOutputReference
	_jsii_.Get(
		j,
		"ucSecurableSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) UcSecurableSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucSecurableSpecInput",
		&returns,
	)
	return returns
}


func NewAppsSettingsCustomTemplateManifestResourceSpecsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppsSettingsCustomTemplateManifestResourceSpecsOutputReference {
	_init_.Initialize()

	if err := validateNewAppsSettingsCustomTemplateManifestResourceSpecsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.appsSettingsCustomTemplate.AppsSettingsCustomTemplateManifestResourceSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppsSettingsCustomTemplateManifestResourceSpecsOutputReference_Override(a AppsSettingsCustomTemplateManifestResourceSpecsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.appsSettingsCustomTemplate.AppsSettingsCustomTemplateManifestResourceSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutJobSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsJobSpec) {
	if err := a.validatePutJobSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJobSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutSecretSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsSecretSpec) {
	if err := a.validatePutSecretSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecretSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutServingEndpointSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpec) {
	if err := a.validatePutServingEndpointSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServingEndpointSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutSqlWarehouseSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpec) {
	if err := a.validatePutSqlWarehouseSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqlWarehouseSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutUcSecurableSpec(value *AppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpec) {
	if err := a.validatePutUcSecurableSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUcSecurableSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetJobSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetJobSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetSecretSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetServingEndpointSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetServingEndpointSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetSqlWarehouseSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetSqlWarehouseSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetUcSecurableSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetUcSecurableSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

