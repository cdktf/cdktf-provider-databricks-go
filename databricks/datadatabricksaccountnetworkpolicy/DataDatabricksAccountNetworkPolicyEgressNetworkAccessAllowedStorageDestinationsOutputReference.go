// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksaccountnetworkpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference interface {
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

// The jsii proxy struct for DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference
type jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) AzureStorageAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) AzureStorageAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) AzureStorageService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) AzureStorageServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) StorageDestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageDestinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) StorageDestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageDestinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference_Override(d DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetAzureStorageAccount(val *string) {
	if err := j.validateSetAzureStorageAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageAccount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetAzureStorageService(val *string) {
	if err := j.validateSetAzureStorageServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageService",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetStorageDestinationType(val *string) {
	if err := j.validateSetStorageDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageDestinationType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetAzureStorageAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureStorageAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetAzureStorageService() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureStorageService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ResetStorageDestinationType() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageDestinationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

