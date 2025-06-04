// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/accountnetworkpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference interface {
	cdktf.ComplexObject
	AzureStorageAccount() *string
	SetAzureStorageAccount(val *string)
	AzureStorageAccountInput() *string
	AzureStorageService() *string
	SetAzureStorageService(val *string)
	AzureStorageServiceInput() *string
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	StorageDestinationType() *string
	SetStorageDestinationType(val *string)
	StorageDestinationTypeInput() *string
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
	ResetAzureStorageAccount()
	ResetAzureStorageService()
	ResetBucketName()
	ResetRegion()
	ResetStorageDestinationType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference
type jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) AzureStorageAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) AzureStorageAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) AzureStorageService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) AzureStorageServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) StorageDestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageDestinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) StorageDestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageDestinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference {
	_init_.Initialize()

	if err := validateNewAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference_Override(a AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetAzureStorageAccount(val *string) {
	if err := j.validateSetAzureStorageAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageAccount",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetAzureStorageService(val *string) {
	if err := j.validateSetAzureStorageServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageService",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetStorageDestinationType(val *string) {
	if err := j.validateSetStorageDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageDestinationType",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetAzureStorageAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureStorageAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetAzureStorageService() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureStorageService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetStorageDestinationType() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageDestinationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

