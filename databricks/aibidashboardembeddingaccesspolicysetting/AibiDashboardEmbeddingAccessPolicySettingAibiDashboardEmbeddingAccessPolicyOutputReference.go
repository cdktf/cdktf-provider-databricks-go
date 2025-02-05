// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aibidashboardembeddingaccesspolicysetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/aibidashboardembeddingaccesspolicysetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference interface {
	cdktf.ComplexObject
	AccessPolicyType() *string
	SetAccessPolicyType(val *string)
	AccessPolicyTypeInput() *string
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
	InternalValue() *AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicy
	SetInternalValue(val *AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicy)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference
type jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) AccessPolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) AccessPolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) InternalValue() *AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicy {
	var returns *AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.aibiDashboardEmbeddingAccessPolicySetting.AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference_Override(a AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.aibiDashboardEmbeddingAccessPolicySetting.AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference)SetAccessPolicyType(val *string) {
	if err := j.validateSetAccessPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPolicyType",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference)SetInternalValue(val *AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

