// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipalfederationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/serviceprincipalfederationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicePrincipalFederationPolicyOidcPolicyOutputReference interface {
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

// The jsii proxy struct for ServicePrincipalFederationPolicyOidcPolicyOutputReference
type jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) Audiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) AudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) JwksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) JwksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) JwksUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) JwksUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) SubjectClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) SubjectClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServicePrincipalFederationPolicyOidcPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServicePrincipalFederationPolicyOidcPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewServicePrincipalFederationPolicyOidcPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.servicePrincipalFederationPolicy.ServicePrincipalFederationPolicyOidcPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicePrincipalFederationPolicyOidcPolicyOutputReference_Override(s ServicePrincipalFederationPolicyOidcPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.servicePrincipalFederationPolicy.ServicePrincipalFederationPolicyOidcPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetAudiences(val *[]*string) {
	if err := j.validateSetAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audiences",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetJwksJson(val *string) {
	if err := j.validateSetJwksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksJson",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetJwksUri(val *string) {
	if err := j.validateSetJwksUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksUri",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetSubjectClaim(val *string) {
	if err := j.validateSetSubjectClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectClaim",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetAudiences() {
	_jsii_.InvokeVoid(
		s,
		"resetAudiences",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		s,
		"resetIssuer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetJwksJson() {
	_jsii_.InvokeVoid(
		s,
		"resetJwksJson",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetJwksUri() {
	_jsii_.InvokeVoid(
		s,
		"resetJwksUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetSubject() {
	_jsii_.InvokeVoid(
		s,
		"resetSubject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetSubjectClaim() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectClaim",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFederationPolicyOidcPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

