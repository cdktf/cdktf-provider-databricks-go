// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksservingendpoints/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MicrosoftEntraClientId() *string
	SetMicrosoftEntraClientId(val *string)
	MicrosoftEntraClientIdInput() *string
	MicrosoftEntraClientSecret() *string
	SetMicrosoftEntraClientSecret(val *string)
	MicrosoftEntraClientSecretInput() *string
	MicrosoftEntraClientSecretPlaintext() *string
	SetMicrosoftEntraClientSecretPlaintext(val *string)
	MicrosoftEntraClientSecretPlaintextInput() *string
	MicrosoftEntraTenantId() *string
	SetMicrosoftEntraTenantId(val *string)
	MicrosoftEntraTenantIdInput() *string
	OpenaiApiBase() *string
	SetOpenaiApiBase(val *string)
	OpenaiApiBaseInput() *string
	OpenaiApiKey() *string
	SetOpenaiApiKey(val *string)
	OpenaiApiKeyInput() *string
	OpenaiApiKeyPlaintext() *string
	SetOpenaiApiKeyPlaintext(val *string)
	OpenaiApiKeyPlaintextInput() *string
	OpenaiApiType() *string
	SetOpenaiApiType(val *string)
	OpenaiApiTypeInput() *string
	OpenaiApiVersion() *string
	SetOpenaiApiVersion(val *string)
	OpenaiApiVersionInput() *string
	OpenaiDeploymentName() *string
	SetOpenaiDeploymentName(val *string)
	OpenaiDeploymentNameInput() *string
	OpenaiOrganization() *string
	SetOpenaiOrganization(val *string)
	OpenaiOrganizationInput() *string
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
	ResetMicrosoftEntraClientId()
	ResetMicrosoftEntraClientSecret()
	ResetMicrosoftEntraClientSecretPlaintext()
	ResetMicrosoftEntraTenantId()
	ResetOpenaiApiBase()
	ResetOpenaiApiKey()
	ResetOpenaiApiKeyPlaintext()
	ResetOpenaiApiType()
	ResetOpenaiApiVersion()
	ResetOpenaiDeploymentName()
	ResetOpenaiOrganization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference
type jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) MicrosoftEntraClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftEntraClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) MicrosoftEntraClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftEntraClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) MicrosoftEntraClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftEntraClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) MicrosoftEntraClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftEntraClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) MicrosoftEntraClientSecretPlaintext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftEntraClientSecretPlaintext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) MicrosoftEntraClientSecretPlaintextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftEntraClientSecretPlaintextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) MicrosoftEntraTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftEntraTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) MicrosoftEntraTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftEntraTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiKeyPlaintext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiKeyPlaintext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiKeyPlaintextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiKeyPlaintextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiDeploymentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiDeploymentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiDeploymentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiDeploymentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiOrganization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) OpenaiOrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference_Override(d DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetMicrosoftEntraClientId(val *string) {
	if err := j.validateSetMicrosoftEntraClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftEntraClientId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetMicrosoftEntraClientSecret(val *string) {
	if err := j.validateSetMicrosoftEntraClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftEntraClientSecret",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetMicrosoftEntraClientSecretPlaintext(val *string) {
	if err := j.validateSetMicrosoftEntraClientSecretPlaintextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftEntraClientSecretPlaintext",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetMicrosoftEntraTenantId(val *string) {
	if err := j.validateSetMicrosoftEntraTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftEntraTenantId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiBase(val *string) {
	if err := j.validateSetOpenaiApiBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiBase",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiKey(val *string) {
	if err := j.validateSetOpenaiApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiKeyPlaintext(val *string) {
	if err := j.validateSetOpenaiApiKeyPlaintextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiKeyPlaintext",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiType(val *string) {
	if err := j.validateSetOpenaiApiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiApiVersion(val *string) {
	if err := j.validateSetOpenaiApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiApiVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiDeploymentName(val *string) {
	if err := j.validateSetOpenaiDeploymentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiDeploymentName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetOpenaiOrganization(val *string) {
	if err := j.validateSetOpenaiOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiOrganization",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetMicrosoftEntraClientId() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftEntraClientId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetMicrosoftEntraClientSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftEntraClientSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetMicrosoftEntraClientSecretPlaintext() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftEntraClientSecretPlaintext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetMicrosoftEntraTenantId() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftEntraTenantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiApiBase() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenaiApiBase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiApiKey() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenaiApiKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiApiKeyPlaintext() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenaiApiKeyPlaintext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiApiType() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenaiApiType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiApiVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenaiApiVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiDeploymentName() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenaiDeploymentName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ResetOpenaiOrganization() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenaiOrganization",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelOpenaiConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

