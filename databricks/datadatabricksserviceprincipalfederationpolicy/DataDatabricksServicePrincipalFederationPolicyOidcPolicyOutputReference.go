// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipalfederationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksserviceprincipalfederationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference interface {
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
	InternalValue() *DataDatabricksServicePrincipalFederationPolicyOidcPolicy
	SetInternalValue(val *DataDatabricksServicePrincipalFederationPolicyOidcPolicy)
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

// The jsii proxy struct for DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference
type jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) Audiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) AudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) InternalValue() *DataDatabricksServicePrincipalFederationPolicyOidcPolicy {
	var returns *DataDatabricksServicePrincipalFederationPolicyOidcPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) JwksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) JwksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) JwksUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) JwksUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) SubjectClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) SubjectClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServicePrincipalFederationPolicy.DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference_Override(d DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServicePrincipalFederationPolicy.DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetAudiences(val *[]*string) {
	if err := j.validateSetAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audiences",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetInternalValue(val *DataDatabricksServicePrincipalFederationPolicyOidcPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetJwksJson(val *string) {
	if err := j.validateSetJwksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksJson",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetJwksUri(val *string) {
	if err := j.validateSetJwksUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksUri",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetSubjectClaim(val *string) {
	if err := j.validateSetSubjectClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectClaim",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetAudiences() {
	_jsii_.InvokeVoid(
		d,
		"resetAudiences",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		d,
		"resetIssuer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetJwksJson() {
	_jsii_.InvokeVoid(
		d,
		"resetJwksJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetJwksUri() {
	_jsii_.InvokeVoid(
		d,
		"resetJwksUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetSubject() {
	_jsii_.InvokeVoid(
		d,
		"resetSubject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ResetSubjectClaim() {
	_jsii_.InvokeVoid(
		d,
		"resetSubjectClaim",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksServicePrincipalFederationPolicyOidcPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

