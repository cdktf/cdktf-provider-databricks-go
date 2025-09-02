// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountfederationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/accountfederationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountFederationPolicyOidcPolicyOutputReference interface {
	cdktf.ComplexObject
	Audiences() *[]*string
	SetAudiences(val *[]*string)
	AudiencesInput() *[]*string
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
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	JwksJson() *string
	SetJwksJson(val *string)
	JwksJsonInput() *string
	JwksUri() *string
	SetJwksUri(val *string)
	JwksUriInput() *string
	Subject() *string
	SetSubject(val *string)
	SubjectClaim() *string
	SetSubjectClaim(val *string)
	SubjectClaimInput() *string
	SubjectInput() *string
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
	ResetAudiences()
	ResetIssuer()
	ResetJwksJson()
	ResetJwksUri()
	ResetSubject()
	ResetSubjectClaim()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountFederationPolicyOidcPolicyOutputReference
type jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) Audiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) AudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) JwksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) JwksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) JwksUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) JwksUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) SubjectClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) SubjectClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountFederationPolicyOidcPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccountFederationPolicyOidcPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAccountFederationPolicyOidcPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.accountFederationPolicy.AccountFederationPolicyOidcPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountFederationPolicyOidcPolicyOutputReference_Override(a AccountFederationPolicyOidcPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.accountFederationPolicy.AccountFederationPolicyOidcPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetAudiences(val *[]*string) {
	if err := j.validateSetAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audiences",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetJwksJson(val *string) {
	if err := j.validateSetJwksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksJson",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetJwksUri(val *string) {
	if err := j.validateSetJwksUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksUri",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetSubjectClaim(val *string) {
	if err := j.validateSetSubjectClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectClaim",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ResetAudiences() {
	_jsii_.InvokeVoid(
		a,
		"resetAudiences",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ResetJwksJson() {
	_jsii_.InvokeVoid(
		a,
		"resetJwksJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ResetJwksUri() {
	_jsii_.InvokeVoid(
		a,
		"resetJwksUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ResetSubject() {
	_jsii_.InvokeVoid(
		a,
		"resetSubject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ResetSubjectClaim() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectClaim",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountFederationPolicyOidcPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

